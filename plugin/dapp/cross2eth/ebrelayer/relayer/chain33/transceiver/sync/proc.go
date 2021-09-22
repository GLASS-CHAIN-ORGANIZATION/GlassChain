package sync

import (
	"fmt"
	"math"
	"sync/atomic"

	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/types"
	"github.com/33cn/plugin/plugin/dapp/cross2eth/ebrelayer/utils"
	"github.com/pkg/errors"
)

// SeqType
const (
	SeqTypeAdd = int32(1)
	SeqTypeDel = int32(2)
)

var (
	syncLastHeight   = []byte("syncLastHeight:")
	evmTxLogPrefix   = []byte("evmTxLogPrefix:")
	lastSequences    = []byte("lastSequences:")
	seqOperationType = []string{"SeqTypeAdd", "SeqTypeDel"}
)

var evmTxLogsCh chan *types.EVMTxLogsInBlks
var resultCh chan error

func init() {
	evmTxLogsCh = make(chan *types.EVMTxLogsInBlks)
	resultCh = make(chan error)
}

func evmTxLogKey4Height(height int64) []byte {
	return append(evmTxLogPrefix, []byte(fmt.Sprintf("%012d", height))...)
}

// pushTxReceipts push block to backend
func pushTxReceipts(evmTxLogsInBlks *types.EVMTxLogsInBlks) error {
	evmTxLogsCh <- evmTxLogsInBlks
	err := <-resultCh
	return err
}

//EVMTxLogs ...
type EVMTxLogs struct {
	db     dbm.DB
	seqNum int64 
	height int64 
	quit   chan struct{}
}

//NewSyncTxReceipts ...
func NewSyncTxReceipts(db dbm.DB) *EVMTxLogs {
	sync := &EVMTxLogs{
		db: db,
	}
	sync.seqNum, _ = sync.loadBlockLastSequence()
	sync.height, _ = sync.LoadLastBlockHeight()
	sync.quit = make(chan struct{})
	sync.initSyncReceiptDataBase()

	return sync
}

func (syncTx *EVMTxLogs) initSyncReceiptDataBase() {
	txLogs0, _ := syncTx.GetTxLogs(0)
	if nil != txLogs0 {
		return
	}
	logsPerBlock := &types.EVMTxLogPerBlk{
		Height: 0,
	}
	syncTx.setTxLogsPerBlock(logsPerBlock)
}

//Stop ...
func (syncTx *EVMTxLogs) Stop() {
	close(syncTx.quit)
}

// SaveAndSyncTxs2Relayer save block to db
func (syncTx *EVMTxLogs) SaveAndSyncTxs2Relayer() {
	for {
		select {
		case evmTxLogs := <-evmTxLogsCh:
			log.Info("to deal request", "seq", evmTxLogs.Logs4EVMPerBlk[0].SeqNum, "count", len(evmTxLogs.Logs4EVMPerBlk))
			syncTx.dealEVMTxLogs(evmTxLogs)
		case <-syncTx.quit:
			return
		}
	}
}

func (syncTx *EVMTxLogs) dealEVMTxLogs(evmTxLogsInBlks *types.EVMTxLogsInBlks) {
	count, start, evmTxLogsParsed, err := parseEvmTxLogsInBlks(evmTxLogsInBlks)
	if err != nil {
		resultCh <- err
	}

	if start < syncTx.seqNum {
		log.Error("dealEVMTxLogs err: the tx and receipt pushed is old", "start", start, "current_seq", syncTx.seqNum)
		resultCh <- errors.New("The tx and receipt pushed is old")
		return
	}
	var height int64
	for i := 0; i < count; i++ {
		txsPerBlock := evmTxLogsParsed[i]
		if txsPerBlock.AddDelType == SeqTypeAdd {
			syncTx.setTxLogsPerBlock(txsPerBlock)
			syncTx.setBlockLastSequence(txsPerBlock.SeqNum)
			syncTx.setBlockHeight(txsPerBlock.Height)
			height = txsPerBlock.Height
		} else {
			syncTx.delTxReceipts(txsPerBlock.Height)
			syncTx.setBlockLastSequence(txsPerBlock.SeqNum)
			height = txsPerBlock.Height - 1
			syncTx.setBlockHeight(height)
		}
	}

	resultCh <- nil
	log.Debug("dealEVMTxLogs", "seqStart", start, "count", count, "maxBlockHeight", height)
}

func (syncTx *EVMTxLogs) loadBlockLastSequence() (int64, error) {
	return utils.LoadInt64FromDB(lastSequences, syncTx.db)
}

//LoadLastBlockHeight ...
func (syncTx *EVMTxLogs) LoadLastBlockHeight() (int64, error) {
	return utils.LoadInt64FromDB(syncLastHeight, syncTx.db)
}

func (syncTx *EVMTxLogs) setBlockLastSequence(newSequence int64) {
	Sequencebytes := types.Encode(&types.Int64{Data: newSequence})
	if err := syncTx.db.Set(lastSequences, Sequencebytes); nil != err {
		panic("setBlockLastSequence failed due to cause:" + err.Error())
	}
	syncTx.updateSequence(newSequence)
}

func (syncTx *EVMTxLogs) setBlockHeight(height int64) {
	bytes := types.Encode(&types.Int64{Data: height})
	_ = syncTx.db.Set(syncLastHeight, bytes)
	atomic.StoreInt64(&syncTx.height, height)
}

func (syncTx *EVMTxLogs) updateSequence(newSequence int64) {
	atomic.StoreInt64(&syncTx.seqNum, newSequence)
}

func (syncTx *EVMTxLogs) setTxLogsPerBlock(txLogs *types.EVMTxLogPerBlk) {
	key := evmTxLogKey4Height(txLogs.Height)
	value := types.Encode(txLogs)
	if err := syncTx.db.Set(key, value); nil != err {
		panic("setTxLogsPerBlock failed due to:" + err.Error())
	}
}

//GetTxReceipts ...
func (syncTx *EVMTxLogs) GetTxLogs(height int64) (*types.TxReceipts4SubscribePerBlk, error) {
	key := evmTxLogKey4Height(height)
	value, err := syncTx.db.Get(key)
	if err != nil {
		return nil, err
	}
	detail := &types.TxReceipts4SubscribePerBlk{}
	err = types.Decode(value, detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

//GetNextValidTxReceipts ...
func (syncTx *EVMTxLogs) GetNextValidEvmTxLogs(height int64) (*types.EVMTxLogPerBlk, error) {
	key := evmTxLogKey4Height(height)
	helper := dbm.NewListHelper(syncTx.db)
	evmTxLogs := helper.List(evmTxLogPrefix, key, 1, dbm.ListASC)
	if nil == evmTxLogs {
		return nil, nil
	}
	detail := &types.EVMTxLogPerBlk{}
	err := types.Decode(evmTxLogs[0], detail)
	if err != nil {
		return nil, err
	}
	return detail, nil
}

func (syncTx *EVMTxLogs) delTxReceipts(height int64) {
	key := evmTxLogKey4Height(height)
	_ = syncTx.db.Set(key, nil)
}

func parseEvmTxLogsInBlks(evmTxLogs *types.EVMTxLogsInBlks) (count int, start int64, txsWithReceipt []*types.EVMTxLogPerBlk, err error) {
	count = len(evmTxLogs.Logs4EVMPerBlk)
	txsWithReceipt = make([]*types.EVMTxLogPerBlk, 0)
	start = math.MaxInt64
	for i := 0; i < count; i++ {
		if evmTxLogs.Logs4EVMPerBlk[i].AddDelType != SeqTypeAdd && evmTxLogs.Logs4EVMPerBlk[i].AddDelType != SeqTypeDel {
			log.Error("parseEvmTxLogsInBlks seq op not support", "seq", evmTxLogs.Logs4EVMPerBlk[i].SeqNum,
				"height", evmTxLogs.Logs4EVMPerBlk[i].Height, "seqOp", evmTxLogs.Logs4EVMPerBlk[i].AddDelType)
			continue
		}
		txsWithReceipt = append(txsWithReceipt, evmTxLogs.Logs4EVMPerBlk[i])
		if evmTxLogs.Logs4EVMPerBlk[i].SeqNum < start {
			start = evmTxLogs.Logs4EVMPerBlk[i].SeqNum
		}
		log.Debug("parseEvmTxLogsInBlks get one block's tx with receipts", "seq", evmTxLogs.Logs4EVMPerBlk[i].SeqNum,
			"height", evmTxLogs.Logs4EVMPerBlk[i].Height, "seqOpType", seqOperationType[evmTxLogs.Logs4EVMPerBlk[i].AddDelType-1])

	}
	if len(txsWithReceipt) != count {
		err = errors.New("duplicate block's tx logs")
		return
	}
	return
}
