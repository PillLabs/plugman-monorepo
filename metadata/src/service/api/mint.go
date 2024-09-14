package api

import (
	"fmt"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/middleware"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"github.com/PillLabs/plugman-monorepo/metadata/src/utils"
	"strconv"
)

type GetSignatureRequest struct {
	Address  string `json:"address"`
	MintType uint   `json:"mint_type"`
	Count    uint   `json:"count"`
}

type GetSignatureResponse struct {
	Count      string   `json:"count"`
	Nonce      string   `json:"nonce"`
	Timestamp  string   `json:"timestamp"`
	TraitValue []string `json:"traitValue"`
	Signature  string   `json:"signature"`
	MintType   int      `json:"mintType"`
}

func GetSignature(request *GetSignatureRequest) (*GetSignatureResponse, int, error) {
	var (
		nonce       model.Nonce
		nonceNum    uint
		metadata    *[]*model.Metadata
		traitValue  [][32]byte
		traitValueS []string
		err         error
	)
	err = premise.RuntimeVars.Dao.AllocateNonce(request.Address, request.MintType, request.Count, &nonce)
	if err != nil {
		if err == common.ErrInsufficientPresaleQuota {
			return nil, common.INSUFFICIENT_PRESALE_QUOTA, nil
		} else if err == common.ErrInvalidMintType {
			return nil, common.INVALID_MINT_TYPE, nil
		}
		return nil, common.INTERNAL_ERROR, fmt.Errorf("allocate nonce error: %s", err)
	}
	switch request.MintType {
	case common.MINT_TYPE_OG:
		nonceNum = nonce.OGNonce
	case common.MINT_TYPE_WL:
		nonceNum = nonce.WLNonce
	case common.MINT_TYPE_PUBLIC_SALE:
		nonceNum = nonce.PublicNonce
	}
	metadata, err = premise.RuntimeVars.Dao.AllocateMetadata(request.Address, request.MintType, request.Count, nonceNum)
	if err != nil {
		if err == common.ErrNoUnusedMetadata {
			return nil, common.NO_UNUSED_METADATA, nil
		}
		return nil, common.INTERNAL_ERROR, fmt.Errorf("allocate metadata error: %s", err)
	}

	for _, m := range *metadata {
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.Body))
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.BodyColor))
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.Background))
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.Front))
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.Side))
		traitValue = append(traitValue, utils.MustConvertStringToByte32(m.Top))
	}

	sig, err := middleware.PlugmanMintSign(
		request.Address,
		premise.RuntimeVars.PrivateKey,
		request.Count,
		nonceNum,
		request.MintType,
		traitValue,
		(*metadata)[0].Timestamp,
	)
	if err != nil {
		return nil, common.INTERNAL_ERROR, fmt.Errorf("generate signature error: %s", err)
	}

	for _, t := range traitValue {
		traitValueS = append(traitValueS, string(t[:]))
	}

	resp := GetSignatureResponse{
		Count:      strconv.Itoa(int(request.Count)),
		Nonce:      strconv.Itoa(int(nonceNum)),
		Timestamp:  strconv.FormatUint((*metadata)[0].Timestamp, 10),
		TraitValue: traitValueS,
		Signature:  sig,
		MintType:   int(request.MintType),
	}
	return &resp, common.SUCCESS, nil
}
