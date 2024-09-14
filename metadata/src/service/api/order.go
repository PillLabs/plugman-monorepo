package api

import (
	"errors"
	"fmt"
	"github.com/PillLabs/plugman-monorepo/metadata/src/chain"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"github.com/PillLabs/plugman-monorepo/metadata/src/stats"
	common2 "github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm"
)

func PostOrder(request model.PostOrderRequest) (int, error) {
	rpcUrl := conf.GConfig.GetString(fmt.Sprintf("connectedEvmChain.%s.chainRpc", request.ChainId))
	chainName := conf.GConfig.GetString(fmt.Sprintf("connectedEvmChain.%s.chainName", request.ChainId))
	tss := conf.GConfig.GetString(fmt.Sprintf("connectedEvmChain.%s.tssAddress", request.ChainId))

	if rpcUrl == "" || chainName == "" {
		return common.NOT_SUPPORTED_CHAIN_ID, nil
	}

	rpc, err := chain.NewEthChain(rpcUrl, chainName)
	if err != nil {
		return common.INTERNAL_ERROR, err
	}
	if !common2.IsHexAddress(request.FromAddress) {
		return common.INVALID_PARAM, nil
	}

	if request.Nonce < 0 || request.MintType < common.MINT_TYPE_OG || request.MintType > common.MINT_TYPE_PUBLIC_SALE {
		log.Info("[api - PostOrder] nonce %d or mint type %d in request is invalid",
			request.Nonce, request.MintType)
		return common.INVALID_PARAM, nil
	}

	tx, found := rpc.GetTransactionWithRetries(common2.HexToHash(request.TxHash))
	if !found {
		return common.TRANSACTION_NOT_FOUND, nil
	}

	if tx.To().String() != tss {
		log.Info("[api - PostOrder] transfer to address %s is not TSS address", tx.To().String())
		return common.INVALID_PARAM, nil
	}

	sender := chain.GetSenderFromTx(tx)
	if sender.String() != request.FromAddress {
		log.Info("[api - PostOrder] sender address %s is not match with from address %s",
			sender.String(), request.FromAddress)
		return common.INVALID_PARAM, nil
	}

	var metadataList []model.Metadata
	err = premise.RuntimeVars.Dao.GetMetadataByAddressAndMintTypeAndNonce(
		request.FromAddress, uint(request.Nonce), uint(request.MintType), &metadataList)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[api - PostOrder] failed to fetch metadata list from db: %v", err)
		return common.INTERNAL_ERROR, err
	}
	if len(metadataList) == 0 {
		log.Info("[api - PostOrder] metadata list for address %s nonce %d is empty",
			request.FromAddress, request.Nonce)
		return common.INVALID_PARAM, nil
	}

	status, err := chain.CheckCrossChainTransactionStatus(request.TxHash)
	if err != nil {
		stats.ThirdPartyClientError.WithLabelValues("zeta_cross_rpc").Inc()
		log.Info("[api - PostOrder] check cross chain transaction status failed: %v", err)
		return common.INVALID_PARAM, nil
	}

	tssTransfer := model.TssTransfer{
		Model:       gorm.Model{},
		ChainId:     request.ChainId,
		FromAddress: request.FromAddress,
		Amount:      tx.Value().String(),
		TxHash:      tx.Hash().String(),
		MintType:    request.MintType,
		Nonce:       request.Nonce,
		OrderStatus: status,
	}
	err = premise.RuntimeVars.Dao.CreateTssTransfer(&tssTransfer)
	if err != nil {
		log.Error("[api - PostOrder] failed to create tssTransfer: %v", err)
		return common.INTERNAL_ERROR, err
	}
	return common.SUCCESS, nil
}

func GetOrders(address string) (model.GetOrderResponse, int, error) {
	var (
		transfers []model.TssTransfer
		orders    []model.Order
	)
	err := premise.RuntimeVars.Dao.GetTssTransfer(address, &transfers)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[api - GetOrders] failed to get tssTransfer: %s", err)
		stats.DbError.Inc()
		return model.GetOrderResponse{}, common.INTERNAL_ERROR, err
	}
	if len(transfers) == 0 {
		return model.GetOrderResponse{}, common.ORDER_NOT_FOUND, nil
	}

	for _, transfer := range transfers {
		var (
			metadataList []model.Metadata
			tokenIds     []string
		)
		err = premise.RuntimeVars.Dao.GetMetadataByAddressAndMintTypeAndNonce(
			address, uint(transfer.Nonce), uint(transfer.MintType), &metadataList)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("[api - GetOrders] failed to fetch metadata list from db: %v", err)
			stats.DbError.Inc()
			return model.GetOrderResponse{}, common.INTERNAL_ERROR, err
		}
		if len(metadataList) == 0 {
			log.Error("[api - GetOrders] metadata list for address %s nonce %d is empty",
				address, transfer.Nonce)
		}
		for _, metadata := range metadataList {
			if metadata.TokenId != "" {
				tokenIds = append(tokenIds, metadata.TokenId)
			}
		}
		order := model.Order{
			TssTransfer: transfer,
			TokenId:     tokenIds,
		}
		orders = append(orders, order)
	}
	return model.GetOrderResponse{Orders: orders}, common.SUCCESS, nil
}
