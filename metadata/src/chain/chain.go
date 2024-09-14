package chain

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/stats"
	"github.com/ethereum/go-ethereum"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Block struct {
	BaseFee *big.Int `json:"baseFeePerGas"`
	Time    uint64   `json:"timestamp"`
}

type EthChain struct {
	Client           *ethclient.Client
	EIP1559Supported bool
	ChainName        string
	rpcUrl           string
}

func NewEthChain(rpcUrl, chainName string) (*EthChain, error) {
	c, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	_, supportedEIP1559, err := BlockByNumber(context.TODO(), c, rpcUrl, nil)
	if err != nil {
		return nil, err
	}
	return &EthChain{
		Client:           c,
		EIP1559Supported: supportedEIP1559,
		ChainName:        chainName,
		rpcUrl:           rpcUrl,
	}, nil
}

func BlockByNumber(ctx context.Context, client *ethclient.Client, rpcUrl string, number *big.Int) (uint64, bool, error) {
	if number == nil {
		blockNum, err := client.BlockNumber(ctx)
		if err != nil {
			return 0, false, err
		}
		number = big.NewInt(int64(blockNum))
	}
	body := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"" + hexutil.EncodeBig(number) + "\", false], \"id\":\"1\"}"

	req, err := http.NewRequestWithContext(ctx, "POST", rpcUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return 0, false, err
	}
	httpClient := &http.Client{}
	req.Header.Set("Content-Type", "application/json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, false, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, false, err
	}

	var result Block
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return 0, false, err
	}
	if result.BaseFee == nil {
		return result.Time, false, nil
	}
	return result.Time, true, nil
}

func (e *EthChain) GetBlockByNumberWithRetries(blockId uint64) *types.Block {
	for {
		block, err := func() (*types.Block, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.BlockByNumber(ctx, big.NewInt(int64(blockId)))
		}()
		if err != nil {
			log.Error("[EthChain - %s - GetBlockByNumberWithRetries] Error getting block by number: %s", e.ChainName, err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return block
	}
}

func (e *EthChain) GetLatestBlockWithRetries() uint64 {
	for {
		blockNumber, err := func() (uint64, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.BlockNumber(ctx)
		}()
		if err != nil {
			log.Error("[EthChain - %s- GetLatestBlockWithRetries] get latest block error: %s", e.ChainName, err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return blockNumber
	}
}

func (e *EthChain) GetBlockTimeWithRetries(blockNumber uint64) uint64 {
	for {
		blockTime, err := func() (uint64, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			number := big.NewInt(int64(blockNumber))
			t, _, err := BlockByNumber(ctx, e.Client, e.rpcUrl, number)
			if err != nil {
				return 0, err
			}
			return t, nil
		}()
		if err != nil {
			log.Error("[EthChain - %s- GetBlockTimeWithRetries] get block %v time error: %v", e.ChainName, blockNumber, err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return blockTime
	}
}

func (e *EthChain) GetBalanceOfAddressWithRetries(address common2.Address) *big.Int {
	for {
		balance, err := func() (*big.Int, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.BalanceAt(ctx, address, nil)
		}()
		if err != nil {
			log.Error("[EthChain - %s- GetBalanceOfAddressWithRetries] get balance for %v error: %v",
				e.ChainName, address.String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return balance
	}
}

func (e *EthChain) SendTxWithRetries(tx *types.Transaction) error {
	for {
		err := func() error {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.SendTransaction(ctx, tx)
		}()
		if err != nil {
			log.Error("[EthChain - %s - SendTxWithRetries] send raw transaction %s error: %s",
				e.ChainName, tx.Hash().String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			if strings.Contains(err.Error(), "replacement transaction underpriced") {
				return common.ErrReplaceUnderPriced
			}
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return nil
	}
}

func (e *EthChain) GetReceiptWithRetries(hash common2.Hash) (*types.Receipt, bool) {
	for {
		receipt, err := func() (*types.Receipt, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.TransactionReceipt(ctx, hash)
		}()
		if err != nil && err != ethereum.NotFound {
			log.Error("[EthChain - %s - GetReceiptWithRetries] get receipt for transaction %s error: %s",
				e.ChainName, hash.String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		if err == ethereum.NotFound {
			return nil, false
		} else {
			return receipt, true
		}
	}
}

func (e *EthChain) GetNonceWithRetries(address common2.Address) uint64 {
	for {
		nonce, err := func() (uint64, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.NonceAt(ctx, address, nil)
		}()
		if err != nil {
			log.Error("[EthChain - %s - GetNonceWithRetries] get nonce for address %s error: %s",
				e.ChainName, address.String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return nonce
	}
}

func (e *EthChain) GetPendingNonceWithRetries(address common2.Address) uint64 {
	for {
		nonce, err := func() (uint64, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.PendingNonceAt(ctx, address)
		}()
		if err != nil {
			log.Error("[EthChain - %s - GetPendingNonceWithRetries] get nonce for address %s error: %s",
				e.ChainName, address.String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		return nonce
	}
}

func (e *EthChain) GetTransactionWithRetries(hash common2.Hash) (*types.Transaction, bool) {
	for {
		tx, _, err := func() (*types.Transaction, bool, error) {
			rpcTimeout := conf.GConfig.GetDuration("chain.rpcTimeout")
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second*rpcTimeout)
			defer cancel()
			return e.Client.TransactionByHash(ctx, hash)
		}()

		if err != nil && !errors.Is(err, ethereum.NotFound) {
			log.Warn("[EthChain - %s - GetReceiptWithRetries] get transaction %s error: %s",
				e.ChainName, hash.String(), err)
			stats.EthRpcClientError.WithLabelValues(e.ChainName).Inc()
			time.Sleep(conf.GConfig.GetDuration("chain.retryInterval") * time.Second)
			continue
		}
		if errors.Is(err, ethereum.NotFound) {
			return nil, false
		} else {
			return tx, true
		}
	}
}

func GetSenderFromTx(tx *types.Transaction) common2.Address {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		log.Fatal("[chain - GetSenderFromTx] get sender error: %s", err)
	}
	return from
}
