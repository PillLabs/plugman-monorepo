package listener

import (
	"errors"
	"github.com/PillLabs/plugman-monorepo/metadata/src/abi/plugman"
	"github.com/PillLabs/plugman-monorepo/metadata/src/abi/shuttleMachine"
	"github.com/PillLabs/plugman-monorepo/metadata/src/chain"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common"
	"github.com/PillLabs/plugman-monorepo/metadata/src/common/log"
	"github.com/PillLabs/plugman-monorepo/metadata/src/conf"
	"github.com/PillLabs/plugman-monorepo/metadata/src/dao"
	"github.com/PillLabs/plugman-monorepo/metadata/src/model"
	"github.com/PillLabs/plugman-monorepo/metadata/src/premise"
	"github.com/PillLabs/plugman-monorepo/metadata/src/stats"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

type ZetaListener struct {
	db                   *dao.PlugmanDao
	chainId              string
	filter               *plugman.PlugmanFilterer
	shuttleMachineFilter *shuttleMachine.ShuttleMachineFilterer
	rpcClient            *chain.EthChain
	gracefulStop         chan *struct{}
}

func NewZetaListener() *ZetaListener {
	var (
		client *ethclient.Client
		err    error
	)

	client, err = ethclient.Dial(conf.GConfig.GetString("chain.chainRpc"))
	if err != nil {
		log.Fatal("[ZetaListener - NewListener] new eth client error: %s", err)
	}

	plugmanAddress := conf.GConfig.GetString("chain.plugmanContract")
	if ok := common2.IsHexAddress(plugmanAddress); !ok {
		log.Fatal("[ZetaListener - NewListener] invalid verifying contract address: %v", plugmanAddress)
	}

	shuttleMachineAddress := conf.GConfig.GetString("chain.polygonShuttleMachineContract")
	if ok := common2.IsHexAddress(shuttleMachineAddress); !ok {
		log.Fatal("[ZetaListener - NewListener] invalid shuttle machine address: %v", shuttleMachineAddress)
	}

	filter, err := plugman.NewPlugmanFilterer(common2.HexToAddress(plugmanAddress), client)
	if err != nil {
		log.Fatal("[ZetaListener - NewListener] new plugman filter client error: %s", err)
	}

	shuttleMachineFilter, err := shuttleMachine.NewShuttleMachineFilterer(common2.HexToAddress(shuttleMachineAddress), client)
	if err != nil {
		log.Fatal("[ZetaListener - NewListener] new shuttle machine filter client error: %s", err)
	}

	r, err := chain.NewEthChain(conf.GConfig.GetString("chain.chainRpc"), conf.GConfig.GetString("chain.chainName"))
	if err != nil {
		log.Fatal("[ZetaListener - NewListener] new chain client error: %s", err)
	}
	return &ZetaListener{
		db:                   premise.RuntimeVars.Dao,
		chainId:              conf.GConfig.GetString("chain.chainId"),
		filter:               filter,
		shuttleMachineFilter: shuttleMachineFilter,
		rpcClient:            r,
		gracefulStop:         make(chan *struct{}),
	}
}

func (l *ZetaListener) Listen(closed chan *struct{}) {
	defer close(closed)

	timeout := conf.GConfig.GetDuration("chain.scanInterval") * time.Second
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	startId, err := l.getStartBlock()
	if err != nil {
		log.Fatal("[ZetaListener - Listen] get start id error: %v", err)
	}

	log.Debug("[ZetaListener - Listen] get start id: %v", startId)

	endId := l.getEndBlock(startId)

	for {
		stats.EthScanningBlockNumber.WithLabelValues(l.chainId).Set(float64(startId))
		select {
		case <-l.gracefulStop:
			log.Info("[ZetaListener - Listen] receive quit signal from main, scanner is quitting...")
			return
		case <-timer.C:
			wg := sync.WaitGroup{}
			wg.Add(3)
			go func() {
				l.scanPlugman(startId, endId)
				wg.Done()
			}()
			go func() {
				l.scanShuttleMachine(startId, endId)
				wg.Done()
			}()
			go func() {
				l.checkCrossChainTransactionStatus()
				wg.Done()
			}()
			wg.Wait()
			l.punchIn(&endId)
			startId = endId + 1
			endId = l.getEndBlock(startId)
			timer.Reset(timeout)
		}
	}
}

func (l *ZetaListener) scanPlugman(startId, endId uint64) {
	defaultBackoff := time.Second * conf.GConfig.GetDuration("chain.backoffTime")

	opt := &bind.FilterOpts{
		Start:   startId,
		End:     &endId,
		Context: nil,
	}

	filter := func(opt *bind.FilterOpts) *plugman.PlugmanMintIterator {
		for {
			iter, err := l.filter.FilterMint(opt, nil)
			if err != nil {
				log.Error("[ZetaListener - Listen] create filter error: %v", err)
				stats.EthRpcClientError.WithLabelValues(l.rpcClient.ChainName).Inc()
				time.Sleep(defaultBackoff)
				continue
			}
			return iter
		}
	}

	iter := filter(opt)
	if iter == nil {
		return
	}

	for iter.Next() {
		err := l.processMintEvent(iter.Event)
		if err != nil {
			log.Error("[ZetaListener - scanPlugman] process mint event for tx hash %v error: %v",
				iter.Event.Raw.TxHash.String(), err)
		}
	}
}

func (l *ZetaListener) scanShuttleMachine(startId, endId uint64) {
	defaultBackoff := time.Second * conf.GConfig.GetDuration("chain.backoffTime")

	opt := &bind.FilterOpts{
		Start:   startId,
		End:     &endId,
		Context: nil,
	}

	filter := func(opt *bind.FilterOpts) *shuttleMachine.ShuttleMachineCrossChainMintIterator {
		for {
			iter, err := l.shuttleMachineFilter.FilterCrossChainMint(opt, nil)
			if err != nil {
				log.Error("[ZetaListener - scanShuttleMachine] create filter error: %v", err)
				stats.EthRpcClientError.WithLabelValues(l.rpcClient.ChainName).Inc()
				time.Sleep(defaultBackoff)
				continue
			}
			return iter
		}
	}

	iter := filter(opt)
	if iter == nil {
		return
	}

	for iter.Next() {
		err := l.processCrossMintEvent(iter.Event)
		if err != nil {
			log.Error("[ZetaListener - scanPlugman] process mint event for tx hash %v error: %v",
				iter.Event.Raw.TxHash.String(), err)
		}
	}
}

func (l *ZetaListener) GracefulStop() {
	close(l.gracefulStop)
}

func (l *ZetaListener) processMintEvent(event *plugman.PlugmanMint) error {
	address := event.To.String()
	nonce := uint(event.Nonce.Uint64())
	firstTokenId := uint(event.FirstTokenId.Uint64())
	mintType := uint(event.MintType)
	txHash := event.Raw.TxHash.String()
	err := l.db.UpdateMetadataTx(address, txHash, nonce, firstTokenId, mintType)
	if err != nil {
		return err
	}
	return nil
}

func (l *ZetaListener) processCrossMintEvent(event *shuttleMachine.ShuttleMachineCrossChainMint) error {
	log.Info("[ZetaListener - processCrossMintEvent] process event for tx hash %s",
		event.Raw.TxHash.String())
	log.Info("[ZetaListener - processCrossMintEvent] nft receiver %s, eth sender %s, "+
		"amount %s, nonce %s, mint type %d",
		event.To.String(), event.Sender.String(), event.Count.String(), event.Nonce.String(), event.MintType)
	return nil
}

func (l *ZetaListener) checkCrossChainTransactionStatus() {
	var transfers []model.TssTransfer
	err := l.db.GetUnfinishedTssTransfer(&transfers)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("[ZetaListener - checkCrossChainTransactionStatus] get transfers error: %v", err)
		stats.DbError.Inc()
		return
	}
	length := len(transfers)
	if length == 0 {
		return
	}
	log.Debug("[ZetaListener - checkCrossChainTransactionStatus] start track transfer status for %d transfers", length)
	wg := sync.WaitGroup{}
	wg.Add(length)
	for _, transfer := range transfers {
		go func(txHash string) {
			l.updateTssTransferStatus(txHash)
			wg.Done()
		}(transfer.TxHash)
	}
	wg.Wait()
}

func (l *ZetaListener) updateTssTransferStatus(txHash string) {
	status, err := chain.CheckCrossChainTransactionStatus(txHash)
	if err != nil {
		stats.ThirdPartyClientError.WithLabelValues("zeta_cross_rpc").Inc()
		log.Error("[ZetaListener - updateTssTransferStatus] check transaction for %s status error: %v", txHash, err)
		return
	}
	if status != common.ORDER_INIT {
		err = l.db.UpdateTssTransferStatus(txHash, status)
		if err != nil {
			stats.DbError.Inc()
			log.Error("[ZetaListener - updateTssTransferStatus] update transaction for %s status error: %v", txHash, err)
		} else {
			log.Debug("[ZetaListener - updateTssTransferStatus] update transaction for %s status success", txHash)
		}
	}
}

func (l *ZetaListener) punchIn(blockNumber *uint64) {
	var number uint64
	if blockNumber == nil {
		number = l.rpcClient.GetLatestBlockWithRetries()
	} else {
		number = *blockNumber
	}

	block := model.Block{
		ChainName: conf.GConfig.GetString("chain.chainName"),
		ChainId:   l.chainId,
		Number:    strconv.FormatUint(number, 10),
	}
	log.Debug("[EthListener - %v - punchIn] scanning block %v", conf.GConfig.GetString("chain.chainName"), number)
	err := l.db.SetBlockByChain(&block)
	if err != nil {
		log.Error("[EthListener - %v - punchIn] set block to db error: %v", conf.GConfig.GetString("chain.chainName"), err)
	}
}

func (l *ZetaListener) getStartBlock() (uint64, error) {
	var (
		err     error
		block   *model.Block
		blockId uint64
	)
	block, err = l.db.GetBlockByChain(conf.GConfig.GetString("chain.chainName"))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return conf.GConfig.GetUint64("chain.startBlockId"), nil
	} else {
		blockId, err = strconv.ParseUint(block.Number, 10, 64)
		if err != nil {
			return 0, err
		}
	}
	return blockId, nil
}

func (l *ZetaListener) getEndBlock(startId uint64) uint64 {
	candidateEndId := startId + conf.GConfig.GetUint64("chain.filterBlockRange")

	for {
		latestId := l.rpcClient.GetLatestBlockWithRetries()
		if latestId <= startId {
			time.Sleep(5 * time.Second)
			continue
		}
		if latestId > candidateEndId {
			return candidateEndId
		}
		return latestId
	}
}
