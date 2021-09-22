// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package para

import (
	"context"
	"time"

	"strings"

	"sync/atomic"
	"unsafe"

	"sync"

	"strconv"

	"bytes"

	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/types"
	paracross "github.com/33cn/plugin/plugin/dapp/paracross/types"
	pt "github.com/33cn/plugin/plugin/dapp/paracross/types"
	"github.com/pkg/errors"
)

const (
	consensusInterval = 10 //about 1 new block interval
	minerInterval     = 10 

	waitBlocks4CommitMsg int32  = 5 
	waitConsensStopTimes uint32 = 3 //3*10s
)

type paraSelfConsEnable struct {
	startHeight int64
	endHeight   int64
}

type commitMsgClient struct {
	paraClient           *client
	waitMainBlocks       int32  
	waitConsensStopTimes uint32 
	resetCh              chan interface{}
	sendMsgCh            chan *types.Transaction
	minerSwitch          int32
	currentTx            unsafe.Pointer
	chainHeight          int64
	sendingHeight        int64
	consensHeight        int64
	consensDoneHeight    int64
	selfConsensError     int32 
	authAccount          string
	authAccountIn        bool
	isRollBack           int32
	checkTxCommitTimes   int32
	txFeeRate            int64
	selfConsEnableList   []*paraSelfConsEnable 
	privateKey           crypto.PrivKey
	quit                 chan struct{}
	mutex                sync.Mutex
}

type commitCheckParams struct {
	consensStopTimes uint32
}

func newCommitMsgCli(para *client, cfg *subConfig) *commitMsgClient {
	cli := &commitMsgClient{
		paraClient:           para,
		authAccount:          cfg.AuthAccount,
		waitMainBlocks:       waitBlocks4CommitMsg,
		waitConsensStopTimes: waitConsensStopTimes,
		consensHeight:        -2,
		sendingHeight:        -1,
		consensDoneHeight:    -1,
		resetCh:              make(chan interface{}, 1),
		quit:                 make(chan struct{}),
	}
	if cfg.WaitBlocks4CommitMsg > 0 {
		cli.waitMainBlocks = cfg.WaitBlocks4CommitMsg
	}

	if cfg.WaitConsensStopTimes > 0 {
		cli.waitConsensStopTimes = cfg.WaitConsensStopTimes
	}

	
	
	if cfg.ParaConsensStartHeight > 0 {
		cli.consensDoneHeight = cfg.ParaConsensStartHeight - 1
	}
	return cli
}

func (client *commitMsgClient) handler() {
	var readTick <-chan time.Time
	checkParams := &commitCheckParams{}

	client.paraClient.wg.Add(1)
	go client.getMainConsensusInfo()

	if client.authAccount != "" {
		client.paraClient.wg.Add(1)
		client.sendMsgCh = make(chan *types.Transaction, 1)
		go client.sendCommitMsg()

		ticker := time.NewTicker(time.Second * time.Duration(minerInterval))
		readTick = ticker.C
		defer ticker.Stop()
	}

out:
	for {
		select {
		case <-client.resetCh:
			client.resetSend()
			client.createCommitTx()
		case <-readTick:
			client.procChecks(checkParams)
			client.createCommitTx()

		case <-client.quit:
			break out
		}
	}

	client.paraClient.wg.Done()
}

func (client *commitMsgClient) updateChainHeightNotify(height int64, isDel bool) {
	if isDel {
		atomic.StoreInt32(&client.isRollBack, 1)
	} else {
		atomic.StoreInt32(&client.isRollBack, 0)
	}

	atomic.StoreInt64(&client.chainHeight, height)

	client.checkRollback(height)
	client.createCommitTx()
}

func (client *commitMsgClient) setInitChainHeight(height int64) {
	atomic.StoreInt64(&client.chainHeight, height)
}

func (client *commitMsgClient) resetNotify() {
	client.resetCh <- 1
}

func (client *commitMsgClient) commitTxCheckNotify(block *types.ParaTxDetail) {
	if client.checkCommitTxSuccess(block) {
		client.createCommitTx()
	}
}

func (client *commitMsgClient) resetSendEnv() {
	client.sendingHeight = -1
	client.setCurrentTx(nil)
}
func (client *commitMsgClient) resetSend() {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.resetSendEnv()
}

func (client *commitMsgClient) getConsensusHeight() int64 {
	status, err := client.getSelfConsensus()
	if err != nil {
		return atomic.LoadInt64(&client.consensHeight)
	}

	return status.Height
}

func (client *commitMsgClient) createCommitTx() {
	tx := client.getCommitTx()
	if tx == nil {
		return
	}
	//bls sign, send to p2p
	if client.paraClient.subCfg.BlsSign {
		//send to p2p pubsub
		plog.Info("para commitMs send to p2p", "hash", common.ToHex(tx.Hash()))
		act := &pt.ParaP2PSubMsg{Ty: P2pSubCommitTx, Value: &pt.ParaP2PSubMsg_CommitTx{CommitTx: tx}}
		client.paraClient.SendPubP2PMsg(paraBlsSignTopic, types.Encode(act))
		return
	}
	client.pushCommitTx(tx)
}

func (client *commitMsgClient) getCommitTx() *types.Transaction {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	consensHeight := client.getConsensusHeight()
	if consensHeight == -1 && consensHeight < client.consensDoneHeight {
		consensHeight = client.consensDoneHeight
	}

	chainHeight := atomic.LoadInt64(&client.chainHeight)
	sendingHeight := client.sendingHeight
	if sendingHeight < consensHeight {
		sendingHeight = consensHeight
	}

	isSync := client.isSync()
	plog.Info("para commitMsg---status", "chainHeight", chainHeight, "sendingHeight", sendingHeight,
		"consensHeight", consensHeight, "isSendingTx", client.isSendingCommitMsg(), "sync", isSync)

	if !isSync {
		return nil
	}

	if sendingHeight > consensHeight || consensHeight > chainHeight || sendingHeight >= chainHeight {
		return nil
	}

	//　sendingHeight <= consensHeight <= chainHeight && sendingHeight < chainHeight
	signTx, count := client.getSendingTx(sendingHeight, chainHeight)
	if signTx == nil {
		return nil
	}
	client.sendingHeight = sendingHeight + count
	return signTx

}

func (client *commitMsgClient) pushCommitTx(signTx *types.Transaction) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.checkTxCommitTimes = 0
	client.setCurrentTx(signTx)
	client.sendMsgCh <- signTx
}

func (client *commitMsgClient) sendCommitActions(acts []*pt.ParacrossCommitAction) {
	txs, _, err := client.createCommitMsgTxs(acts)
	if err != nil {
		return
	}
	plog.Debug("paracommitmsg sendCommitActions", "txhash", common.ToHex(txs.Hash()))
	for i, msg := range acts {
		plog.Debug("paracommitmsg sendCommitActions", "idx", i, "height", msg.Status.Height, "mainheight", msg.Status.MainBlockHeight,
			"blockhash", common.HashHex(msg.Status.BlockHash), "mainHash", common.HashHex(msg.Status.MainBlockHash),
			"addrsmap", common.ToHex(msg.Bls.AddrsMap), "sign", common.ToHex(msg.Bls.Sign))
	}
	client.pushCommitTx(txs)
}

func (client *commitMsgClient) checkTxIn(block *types.ParaTxDetail, tx *types.Transaction) bool {

	if types.IsParaExecName(string(tx.Execer)) {
		for _, tx := range block.TxDetails {
			if bytes.HasSuffix(tx.Tx.Execer, []byte(pt.ParaX)) && tx.Receipt.Ty == types.ExecOk {
				return true
			}
		}
		return false
	}

	receipt, _ := client.paraClient.QueryTxOnMainByHash(tx.Hash())
	if receipt != nil && receipt.Receipt.Ty == types.ExecOk {
		return true
	}
	return false
}

func (client *commitMsgClient) checkCommitTxSuccess(block *types.ParaTxDetail) bool {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	curTx := client.getCurrentTx()
	if curTx == nil {
		return false
	}

	if block.Type != types.AddBlock {
		if client.checkTxCommitTimes > 0 {
			client.checkTxCommitTimes--
		}
		return false
	}

	if client.checkTxIn(block, curTx) {
		client.setCurrentTx(nil)
		return true
	}

	return client.reSendCommitTx(curTx)
}

func (client *commitMsgClient) reSendCommitTx(tx *types.Transaction) bool {
	client.checkTxCommitTimes++
	if client.checkTxCommitTimes < client.waitMainBlocks {
		return false
	}
	client.checkTxCommitTimes = 0
	client.resetSendEnv()
	return true
}

func (client *commitMsgClient) checkConsensusStop(checks *commitCheckParams) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	consensHeight := client.getConsensusHeight()
	if client.sendingHeight > consensHeight {
		checks.consensStopTimes++
		if checks.consensStopTimes > client.waitConsensStopTimes {
			plog.Debug("para checkConsensusStop", "times", checks.consensStopTimes, "consens", consensHeight, "send", client.sendingHeight)
			checks.consensStopTimes = 0
			client.resetSendEnv()
		}
	}
}

func (client *commitMsgClient) checkAuthAccountIn() {
	nodeStr, err := client.getNodeGroupAddrs()
	if err != nil {
		return
	}
	authExist := strings.Contains(nodeStr, client.authAccount)

	if !client.authAccountIn && authExist {
		client.resetSend()
	}

	client.authAccountIn = authExist
}

func (client *commitMsgClient) procChecks(checks *commitCheckParams) {
	client.checkConsensusStop(checks)
	client.checkAuthAccountIn()
}

func (client *commitMsgClient) isSync() bool {
	height := atomic.LoadInt64(&client.chainHeight)
	if height <= 0 {
		plog.Info("para is not Sync", "chainHeight", height)
		return false
	}

	height = atomic.LoadInt64(&client.consensHeight)
	if height == -2 {
		plog.Info("para is not Sync", "consensHeight", height)
		return false
	}

	if atomic.LoadInt32(&client.selfConsensError) != 0 {
		plog.Info("para is not Sync", "selfConsensError", atomic.LoadInt32(&client.selfConsensError))
		return false
	}

	if !client.authAccountIn {
		plog.Info("para is not Sync", "authAccountIn", client.authAccountIn)
		return false
	}

	if atomic.LoadInt32(&client.minerSwitch) != 1 {
		plog.Info("para is not Sync", "isMiner", atomic.LoadInt32(&client.minerSwitch))
		return false
	}

	if atomic.LoadInt32(&client.isRollBack) == 1 {
		plog.Info("para is not Sync", "isRollBack", atomic.LoadInt32(&client.isRollBack))
		return false
	}

	if !client.paraClient.isCaughtUp() {
		plog.Info("para is not Sync", "caughtUp", client.paraClient.isCaughtUp())
		return false
	}

	if !client.paraClient.blockSyncClient.syncHasCaughtUp() {
		plog.Info("para is not Sync", "syncCaughtUp", client.paraClient.blockSyncClient.syncHasCaughtUp())
		return false
	}

	return true

}

func (client *commitMsgClient) getSendingTx(startHeight, endHeight int64) (*types.Transaction, int64) {
	count := endHeight - startHeight
	if count > int64(types.MaxTxGroupSize) {
		count = int64(types.MaxTxGroupSize)
	}
	status, err := client.getNodeStatus(startHeight+1, startHeight+count)
	if err != nil {
		plog.Error("para commit msg read tick", "err", err.Error())
		return nil, 0
	}
	if len(status) == 0 {
		return nil, 0
	}

	var commits []*pt.ParacrossCommitAction
	for _, stat := range status {
		commits = append(commits, &pt.ParacrossCommitAction{Status: stat})
	}

	if client.paraClient.subCfg.BlsSign {
		err = client.paraClient.blsSignCli.blsSign(commits)
		if err != nil {
			plog.Error("paracommitmsg bls sign", "err", err)
			return nil, 0
		}
	}

	signTx, count, err := client.createCommitMsgTxs(commits)
	if err != nil || signTx == nil {
		return nil, 0
	}

	sendingMsgs := status[:count]
	plog.Debug("paracommitmsg sending", "txhash", common.ToHex(signTx.Hash()), "exec", string(signTx.Execer))
	for i, msg := range sendingMsgs {
		plog.Debug("paracommitmsg sending", "idx", i, "height", msg.Height, "mainheight", msg.MainBlockHeight,
			"blockhash", common.HashHex(msg.BlockHash), "mainHash", common.HashHex(msg.MainBlockHash),
			"from", client.authAccount)
	}

	return signTx, count
}

func (client *commitMsgClient) createCommitMsgTxs(notifications []*pt.ParacrossCommitAction) (*types.Transaction, int64, error) {
	txs, count, err := client.batchCalcTxGroup(notifications, atomic.LoadInt64(&client.txFeeRate))
	if err != nil {
		txs, err = client.singleCalcTx((notifications)[0], atomic.LoadInt64(&client.txFeeRate))
		if err != nil {
			plog.Error("single calc tx", "height", notifications[0].Status.Height)

			return nil, 0, err
		}
		return txs, 1, nil
	}
	return txs, int64(count), nil
}

func (client *commitMsgClient) getTxsGroup(txsArr *types.Transactions) (*types.Transaction, error) {
	if len(txsArr.Txs) < 2 {
		tx := txsArr.Txs[0]
		tx.Sign(types.SECP256K1, client.privateKey)
		return tx, nil
	}
	cfg := client.paraClient.GetAPI().GetConfig()
	group, err := types.CreateTxGroup(txsArr.Txs, cfg.GetMinTxFeeRate())
	if err != nil {
		plog.Error("para CreateTxGroup", "err", err.Error())
		return nil, err
	}
	err = group.Check(cfg, 0, cfg.GetMinTxFeeRate(), cfg.GetMaxTxFee())
	if err != nil {
		plog.Error("para CheckTxGroup", "err", err.Error())
		return nil, err
	}
	for i := range group.Txs {
		group.SignN(i, int32(types.SECP256K1), client.privateKey)
	}

	newtx := group.Tx()
	return newtx, nil
}

func (client *commitMsgClient) getExecName(commitHeight int64) string {
	cfg := client.paraClient.GetAPI().GetConfig()
	if cfg.IsDappFork(commitHeight, pt.ParaX, pt.ForkParaFullMinerHeight) {
		return paracross.GetExecName(cfg)
	}

	if cfg.IsDappFork(commitHeight, pt.ParaX, pt.ForkParaSelfConsStages) {
		return paracross.GetExecName(cfg)
	}

	execName := pt.ParaX
	if client.isSelfConsEnable(commitHeight) {
		execName = paracross.GetExecName(cfg)
	}
	return execName

}

func (client *commitMsgClient) batchCalcTxGroup(notifications []*pt.ParacrossCommitAction, feeRate int64) (*types.Transaction, int, error) {
	var rawTxs types.Transactions
	cfg := client.paraClient.GetAPI().GetConfig()
	for i, notify := range notifications {
		if i >= int(types.MaxTxGroupSize) {
			break
		}
		execName := client.getExecName(notify.Status.Height)
		tx, err := paracross.CreateRawCommitTx4MainChain(cfg, notify, execName, feeRate)
		if err != nil {
			plog.Error("para get commit tx", "block height", notify.Status.Height)
			return nil, 0, err
		}
		rawTxs.Txs = append(rawTxs.Txs, tx)
	}

	txs, err := client.getTxsGroup(&rawTxs)
	if err != nil {
		return nil, 0, err
	}
	return txs, len(notifications), nil
}

func (client *commitMsgClient) singleCalcTx(notify *pt.ParacrossCommitAction, feeRate int64) (*types.Transaction, error) {
	cfg := client.paraClient.GetAPI().GetConfig()
	execName := client.getExecName(notify.Status.Height)
	tx, err := paracross.CreateRawCommitTx4MainChain(cfg, notify, execName, feeRate)
	if err != nil {
		plog.Error("para get commit tx", "block height", notify.Status.Height)
		return nil, err
	}
	tx.Sign(types.SECP256K1, client.privateKey)
	return tx, nil

}

func (client *commitMsgClient) setCurrentTx(tx *types.Transaction) {
	atomic.StorePointer(&client.currentTx, unsafe.Pointer(tx))
}

func (client *commitMsgClient) getCurrentTx() *types.Transaction {
	return (*types.Transaction)(atomic.LoadPointer(&client.currentTx))
}

func (client *commitMsgClient) isSendingCommitMsg() bool {
	return client.getCurrentTx() != nil
}

func (client *commitMsgClient) checkRollback(height int64) {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	if height < client.sendingHeight {
		client.resetSendEnv()
	}
}

func (client *commitMsgClient) sendCommitTxOut(tx *types.Transaction) error {
	if tx == nil {
		return nil
	}
	resp, err := client.paraClient.grpcClient.SendTransaction(context.Background(), tx)
	if err != nil {
		plog.Error("sendCommitTxOut send tx", "tx", common.ToHex(tx.Hash()), "err", err.Error())
		return err
	}

	if !resp.GetIsOk() {
		plog.Error("sendCommitTxOut send tx Nok", "tx", common.ToHex(tx.Hash()), "err", string(resp.GetMsg()))
		return errors.New(string(resp.GetMsg()))
	}

	return nil

}

func needResentErr(err error) bool {
	switch err {
	case nil, types.ErrBalanceLessThanTenTimesFee, types.ErrNoBalance, types.ErrDupTx, types.ErrTxExist, types.ErrTxExpire:
		return false
	default:
		return true
	}
}

func (client *commitMsgClient) sendCommitMsg() {
	var err error
	var tx *types.Transaction
	var resendTimer <-chan time.Time

out:
	for {
		select {
		case tx = <-client.sendMsgCh:
			err = client.sendCommitTxOut(tx)
			if err != nil && err == types.ErrTxFeeTooLow {
				err := client.GetProperFeeRate()
				if err == nil {
					client.resetNotify()
				}
				continue
			}
			if needResentErr(err) {
				resendTimer = time.After(time.Second * 2)
			}
		case <-resendTimer:
			if err != nil && tx != nil {
				client.sendCommitTxOut(tx)
			}
		case <-client.quit:
			break out
		}
	}

	client.paraClient.wg.Done()
}

func (client *commitMsgClient) getNodeStatus(start, end int64) ([]*pt.ParacrossNodeStatus, error) {
	var ret []*pt.ParacrossNodeStatus
	if start == 0 {
		geneStatus, err := client.getGenesisNodeStatus()
		if err != nil {
			return nil, err
		}
		ret = append(ret, geneStatus)
		start++
	}
	if end < start {
		return ret, nil
	}

	req := &types.ReqBlocks{Start: start, End: end}
	count := req.End - req.Start + 1
	nodeList := make(map[int64]*pt.ParacrossNodeStatus, count+1)
	keys := &types.LocalDBGet{}
	cfg := client.paraClient.GetAPI().GetConfig()
	for i := 0; i < int(count); i++ {
		key := paracross.CalcMinerHeightKey(cfg.GetTitle(), req.Start+int64(i))
		keys.Keys = append(keys.Keys, key)
	}

	r, err := client.paraClient.GetAPI().LocalGet(keys)
	if err != nil {
		return nil, err
	}
	if count != int64(len(r.Values)) {
		plog.Error("paracommitmsg get node status key", "expect count", count, "actual count", len(r.Values))
		return nil, err
	}
	for _, val := range r.Values {
		status := &pt.ParacrossNodeStatus{}
		err = types.Decode(val, status)
		if err != nil {
			return nil, err
		}
		if !(status.Height >= req.Start && status.Height <= req.End) {
			plog.Error("paracommitmsg decode node status", "height", status.Height, "expect start", req.Start,
				"end", req.End, "status", status)
			return nil, errors.New("paracommitmsg wrong key result")
		}
		nodeList[status.Height] = status

	}
	for i := 0; i < int(count); i++ {
		if nodeList[req.Start+int64(i)] == nil {
			plog.Error("paracommitmsg get node status key nil", "height", req.Start+int64(i))
			return nil, errors.New("paracommitmsg wrong key status result")
		}
	}

	v, err := client.paraClient.GetAPI().GetBlocks(req)
	if err != nil {
		return nil, err
	}
	if count != int64(len(v.Items)) {
		plog.Error("paracommitmsg get node status block", "expect count", count, "actual count", len(v.Items))
		return nil, err
	}
	for _, block := range v.Items {
		if !(block.Block.Height >= req.Start && block.Block.Height <= req.End) {
			plog.Error("paracommitmsg get node status block", "height", block.Block.Height, "expect start", req.Start, "end", req.End)
			return nil, errors.New("paracommitmsg wrong block result")
		}
		nodeList[block.Block.Height].BlockHash = block.Block.Hash(cfg)
		if !paracross.IsParaForkHeight(cfg, nodeList[block.Block.Height].MainBlockHeight, paracross.ForkLoopCheckCommitTxDone) {
			nodeList[block.Block.Height].StateHash = block.Block.StateHash
		}
	}

	var needSentTxs uint32
	for i := 0; i < int(count); i++ {
		ret = append(ret, nodeList[req.Start+int64(i)])
		needSentTxs += nodeList[req.Start+int64(i)].NonCommitTxCounts
	}

	if needSentTxs == 0 && len(ret) < int(types.MaxTxGroupSize) {
		plog.Debug("para commitmsg all self-consensus commit tx,send delay", "start", start, "end", end)
		return nil, nil
	}

	//clear flag
	for _, v := range ret {
		v.NonCommitTxCounts = 0
	}

	return ret, nil

}

func (client *commitMsgClient) getGenesisNodeStatus() (*pt.ParacrossNodeStatus, error) {
	var status pt.ParacrossNodeStatus
	req := &types.ReqBlocks{Start: 0, End: 0}
	v, err := client.paraClient.GetAPI().GetBlocks(req)
	if err != nil {
		return nil, err
	}
	block := v.Items[0].Block
	if block.Height != 0 {
		return nil, errors.New("block chain not return 0 height block")
	}
	cfg := client.paraClient.GetAPI().GetConfig()
	status.Title = cfg.GetTitle()
	status.Height = block.Height
	status.BlockHash = block.Hash(cfg)

	return &status, nil
}

//only sync once, as main usually sync, here just need the first sync status after start up
func (client *commitMsgClient) mainSync() error {
	req := &types.ReqNil{}
	reply, err := client.paraClient.grpcClient.IsSync(context.Background(), req)
	if err != nil {
		plog.Error("Paracross main is syncing", "err", err.Error())
		return err
	}
	if !reply.IsOk {
		plog.Error("Paracross main reply not ok")
		return err
	}

	plog.Info("Paracross main sync succ")
	return nil

}

func (client *commitMsgClient) getMainConsensusInfo() {
	ticker := time.NewTicker(time.Second * time.Duration(consensusInterval))
	isSync := false
	defer ticker.Stop()

out:
	for {
		select {
		case <-client.quit:
			break out
		case <-ticker.C:
			if !isSync {
				err := client.mainSync()
				if err != nil {
					continue
				}
				isSync = true
			}

			if client.authAccount != "" {
				client.GetProperFeeRate()
			}

			selfHeight := int64(-2)
			selfStatus, _ := client.getSelfConsensusStatus()
			if selfStatus != nil {
				selfHeight = selfStatus.Height
			}

			mainStatus, err := client.getMainConsensusStatus()
			if err != nil {
				continue
			}

			if mainStatus.Height < selfHeight {
				atomic.StoreInt32(&client.selfConsensError, 1)
			} else {
				atomic.StoreInt32(&client.selfConsensError, 0)
			}

			preHeight := atomic.LoadInt64(&client.consensHeight)
			atomic.StoreInt64(&client.consensHeight, mainStatus.Height)
			if mainStatus.Height < preHeight {
				client.resetNotify()
			}

			plog.Info("para consensusHeight", "mainHeight", mainStatus.Height, "selfHeight", selfHeight)
		}
	}

	client.paraClient.wg.Done()
}

func (client *commitMsgClient) GetProperFeeRate() error {
	feeRate, err := client.paraClient.grpcClient.GetProperFee(context.Background(), &types.ReqProperFee{})
	if err != nil {
		plog.Error("para commit.GetProperFee", "err", err.Error())
		return err
	}
	if feeRate == nil {
		plog.Error("para commit.GetProperFee return nil")
		return types.ErrInvalidParam
	}

	atomic.StoreInt64(&client.txFeeRate, feeRate.ProperFee)
	return nil
}

func (client *commitMsgClient) getSelfConsensus() (*pt.ParacrossStatus, error) {
	block, err := client.paraClient.getLastBlockInfo()
	if err != nil {
		return nil, err
	}
	ret, err := client.paraClient.GetAPI().QueryChain(&types.ChainExecutor{
		Driver:   "paracross",
		FuncName: "GetSelfConsOneStage",
		Param:    types.Encode(&types.Int64{Data: block.Height}),
	})
	if err != nil {
		plog.Debug("getSelfConsensusStatus.GetSelfConsOneStage ", "err", err.Error())
		return nil, err
	}
	stage, ok := ret.(*pt.SelfConsensStage)
	if !ok {
		plog.Error("getSelfConsensusStatus nok")
		return nil, types.ErrInvalidParam
	}
	if stage.Enable == pt.ParaConfigYes {
		resp, err := client.getSelfConsensusStatus()
		if err != nil {
			return nil, err
		}
		if resp.Height > stage.StartHeight {
			return resp, nil
		}
	}
	return nil, types.ErrNotFound
}

func (client *commitMsgClient) getSelfConsensusStatus() (*pt.ParacrossStatus, error) {
	cfg := client.paraClient.GetAPI().GetConfig()
	ret, err := client.paraClient.GetAPI().QueryChain(&types.ChainExecutor{
		Driver:   "paracross",
		FuncName: "GetTitle",
		Param:    types.Encode(&types.ReqString{Data: cfg.GetTitle()}),
	})
	if err != nil {
		plog.Error("getSelfConsensusStatus ", "err", err)
		return nil, err
	}
	resp, ok := ret.(*pt.ParacrossStatus)
	if !ok {
		plog.Error("getSelfConsensusStatus ParacrossStatus nok")
		return nil, types.ErrNotFound
	}
	return resp, nil

}

func (client *commitMsgClient) getMainConsensusStatus() (*pt.ParacrossStatus, error) {
	block, err := client.paraClient.getLastBlockInfo()
	if err != nil {
		return nil, err
	}
	cfg := client.paraClient.GetAPI().GetConfig()

	reply, err := client.paraClient.grpcClient.QueryChain(context.Background(), &types.ChainExecutor{
		Driver:   "paracross",
		FuncName: "GetTitleByHash",
		Param:    types.Encode(&pt.ReqParacrossTitleHash{Title: cfg.GetTitle(), BlockHash: block.MainHash}),
	})
	if err != nil {
		plog.Error("getMainConsensusStatus", "err", err.Error())
		return nil, err
	}
	if !reply.GetIsOk() {
		plog.Info("getMainConsensusStatus nok", "error", reply.GetMsg())
		return nil, types.ErrNotFound
	}
	var result pt.ParacrossStatus
	err = types.Decode(reply.Msg, &result)
	if err != nil {
		plog.Error("getMainConsensusStatus decode", "err", err.Error())
		return nil, err
	}
	return &result, nil

}

func (client *commitMsgClient) getNodeGroupAddrs() (string, error) {
	cfg := client.paraClient.GetAPI().GetConfig()
	ret, err := client.paraClient.GetAPI().QueryChain(&types.ChainExecutor{
		Driver:   "paracross",
		FuncName: "GetNodeGroupAddrs",
		Param:    types.Encode(&pt.ReqParacrossNodeInfo{Title: cfg.GetTitle()}),
	})
	if err != nil {
		plog.Error("commitmsg.getNodeGroupAddrs ", "err", err.Error())
		return "", err
	}
	resp, ok := ret.(*types.ReplyConfig)
	if !ok {
		plog.Error("commitmsg.getNodeGroupAddrs rsp nok")
		return "", err
	}

	return resp.Value, nil
}

func (client *commitMsgClient) onWalletStatus(status *types.WalletStatus) {
	if status == nil || client.authAccount == "" {
		plog.Info("para onWalletStatus", "status", status == nil, "auth", client.authAccount == "")
		return
	}
	if !status.IsWalletLock && client.privateKey == nil {
		plog.Info("para commit fetchPriKey try")
		client.fetchPriKey()
		plog.Info("para commit fetchPriKey ok")
	}

	if client.privateKey == nil {
		plog.Info("para commit wallet status prikey null", "status", status.IsWalletLock)
		return
	}

	if status.IsWalletLock {
		atomic.StoreInt32(&client.minerSwitch, 0)
	} else {
		atomic.StoreInt32(&client.minerSwitch, 1)
	}

}

func (client *commitMsgClient) onWalletAccount(acc *types.Account) {
	if acc == nil || client.authAccount == "" || client.authAccount != acc.Addr || client.privateKey != nil {
		return
	}
	plog.Error("para onWalletAccount try fetch prikey")
	err := client.fetchPriKey()
	if err != nil {
		plog.Error("para onWalletAccount", "err", err.Error())
		return
	}

	atomic.StoreInt32(&client.minerSwitch, 1)

}

func getSecpPriKey(key string) (crypto.PrivKey, error) {
	pk, err := common.FromHex(key)
	if err != nil && pk == nil {
		return nil, errors.Wrapf(err, "fromhex=%s", key)
	}

	secp, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		return nil, errors.Wrapf(err, "crypto=%s", key)
	}

	priKey, err := secp.PrivKeyFromBytes(pk)
	if err != nil {
		return nil, errors.Wrapf(err, "fromBytes=%s", key)
	}

	return priKey, nil
}

func (client *commitMsgClient) fetchPriKey() error {
	req := &types.ReqString{Data: client.authAccount}

	resp, err := client.paraClient.GetAPI().ExecWalletFunc("wallet", "DumpPrivkey", req)
	if err != nil {
		plog.Error("para fetchPriKey dump priKey", "err", err)
		return err
	}
	str := resp.(*types.ReplyString).Data
	priKey, err := getSecpPriKey(str)
	if err != nil {
		plog.Error("para fetchPriKey get priKey", "err", err)
		return err
	}

	client.privateKey = priKey
	client.paraClient.blsSignCli.setBlsPriKey(priKey.Bytes())

	return nil
}

func parseSelfConsEnableStr(selfEnables []string) ([]*paraSelfConsEnable, error) {
	var list []*paraSelfConsEnable
	for _, e := range selfEnables {
		ret, err := divideStr2Int64s(e, "-")
		if err != nil {
			return nil, err
		}
		list = append(list, &paraSelfConsEnable{ret[0], ret[1]})
	}
	return list, nil
}

//only for "0:50" or "0-50" with one sep
func divideStr2Int64s(s, sep string) ([]int64, error) {
	var r []int64
	a := strings.Split(s, sep)
	if len(a) != 2 {
		plog.Error("error format for config to separate", "s", s)
		return nil, types.ErrInvalidParam
	}

	for _, v := range a {
		val, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			plog.Error("error format for config to parse to int", "s", s)
			return nil, err
		}
		r = append(r, val)
	}
	return r, nil
}

func (client *commitMsgClient) setSelfConsEnable() error {
	cfg := client.paraClient.GetAPI().GetConfig()
	selfEnables := types.Conf(cfg, pt.ParaPrefixConsSubConf).GStrList(pt.ParaSelfConsConfPreContract)
	list, err := parseSelfConsEnableStr(selfEnables)
	if err != nil {
		return err
	}
	client.selfConsEnableList = append(client.selfConsEnableList, list...)
	return nil
}

func (client *commitMsgClient) isSelfConsEnable(height int64) bool {
	for _, v := range client.selfConsEnableList {
		if height >= v.startHeight && height <= v.endHeight {
			return true
		}
	}
	return false
}
