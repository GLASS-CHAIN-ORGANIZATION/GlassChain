// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"bytes"
	"encoding/hex"
	"strings"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/common/crypto"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	"github.com/33cn/chain33/util"
	pt "github.com/33cn/plugin/plugin/dapp/paracross/types"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

type action struct {
	coinsAccount *account.DB
	db           dbm.KV
	localdb      dbm.KVDB
	txhash       []byte
	fromaddr     string
	blocktime    int64
	height       int64
	execaddr     string
	api          client.QueueProtocolAPI
	tx           *types.Transaction
	exec         *Paracross
}

func newAction(t *Paracross, tx *types.Transaction) *action {
	hash := tx.Hash()
	fromaddr := tx.From()
	return &action{t.GetCoinsAccount(), t.GetStateDB(), t.GetLocalDB(), hash, fromaddr,
		t.GetBlockTime(), t.GetHeight(), dapp.ExecAddress(string(tx.Execer)), t.GetAPI(), tx, t}
}

func getNodes(db dbm.KV, key []byte) (map[string]struct{}, []string, error) {
	item, err := db.Get(key)
	if err != nil {
		clog.Info("getNodes", "get db key", string(key), "failed", err)
		if isNotFound(err) {
			err = pt.ErrTitleNotExist
		}
		return nil, nil, errors.Wrapf(err, "db get key:%s", string(key))
	}
	var config types.ConfigItem
	err = types.Decode(item, &config)
	if err != nil {
		return nil, nil, errors.Wrap(err, "decode config")
	}

	value := config.GetArr()
	if value == nil {
		//       ，      ，          
		return map[string]struct{}{}, nil, nil
	}
	var nodesArray []string
	nodesMap := make(map[string]struct{})
	for _, v := range value.Value {
		if _, exist := nodesMap[v]; !exist {
			nodesMap[v] = struct{}{}
			nodesArray = append(nodesArray, v)
		}
	}

	return nodesMap, nodesArray, nil
}

func getConfigManageNodes(db dbm.KV, title string) (map[string]struct{}, []string, error) {
	key := calcManageConfigNodesKey(title)
	return getNodes(db, key)
}

func getParacrossNodes(db dbm.KV, title string) (map[string]struct{}, []string, error) {
	key := calcParaNodeGroupAddrsKey(title)
	return getNodes(db, key)
}

func validTitle(cfg *types.Chain33Config, title string) bool {
	if cfg.IsPara() {
		return cfg.GetTitle() == title
	}
	return len(title) > 0
}

func validNode(addr string, nodes map[string]struct{}) bool {
	if _, exist := nodes[addr]; exist {
		return exist
	}
	return false
}

func checkCommitInfo(cfg *types.Chain33Config, commit *pt.ParacrossNodeStatus) error {
	if commit == nil {
		return types.ErrInvalidParam
	}
	clog.Debug("paracross.Commit check input", "height", commit.Height, "mainHeight", commit.MainBlockHeight,
		"mainHash", common.ToHex(commit.MainBlockHash), "blockHash", common.ToHex(commit.BlockHash))

	if commit.Height == 0 {
		if len(commit.Title) == 0 || len(commit.BlockHash) == 0 {
			return types.ErrInvalidParam
		}
		return nil
	}

	if !pt.IsParaForkHeight(cfg, commit.MainBlockHeight, pt.ForkLoopCheckCommitTxDone) {
		if len(commit.MainBlockHash) == 0 || len(commit.Title) == 0 || commit.Height < 0 ||
			len(commit.PreBlockHash) == 0 || len(commit.BlockHash) == 0 ||
			len(commit.StateHash) == 0 || len(commit.PreStateHash) == 0 {
			return types.ErrInvalidParam
		}
		return nil
	}

	if len(commit.MainBlockHash) == 0 || len(commit.BlockHash) == 0 ||
		commit.MainBlockHeight < 0 || commit.Height < 0 {
		return types.ErrInvalidParam
	}

	if !validTitle(cfg, commit.Title) {
		return pt.ErrInvalidTitle
	}
	return nil
}

//        float  
func isCommitDone(nodes, mostSame int) bool {
	return 3*mostSame > 2*nodes
}

func makeCommitReceipt(addr string, commit *pt.ParacrossNodeStatus, prev, current *pt.ParacrossHeightStatus) *types.Receipt {
	key := calcTitleHeightKey(commit.Title, commit.Height)
	log := &pt.ReceiptParacrossCommit{
		Addr:    addr,
		Status:  commit,
		Prev:    prev,
		Current: current,
	}
	return &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{
			{Key: key, Value: types.Encode(current)},
		},
		Logs: []*types.ReceiptLog{
			{
				Ty:  pt.TyLogParacrossCommit,
				Log: types.Encode(log),
			},
		},
	}
}

func makeCommitStatReceipt(current *pt.ParacrossHeightStatus) *types.Receipt {
	key := calcTitleHeightKey(current.Title, current.Height)
	return &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{
			{Key: key, Value: types.Encode(current)},
		},
		Logs: nil,
	}
}

func makeRecordReceipt(addr string, commit *pt.ParacrossNodeStatus) *types.Receipt {
	log := &pt.ReceiptParacrossRecord{
		Addr:   addr,
		Status: commit,
	}
	return &types.Receipt{
		Ty: types.ExecOk,
		KV: nil,
		Logs: []*types.ReceiptLog{
			{
				Ty:  pt.TyLogParacrossCommitRecord,
				Log: types.Encode(log),
			},
		},
	}
}

func makeDoneReceipt(cfg *types.Chain33Config, execMainHeight, execHeight int64, commit *pt.ParacrossNodeStatus,
	most, commitCount, totalCount int32) *types.Receipt {

	log := &pt.ReceiptParacrossDone{
		TotalNodes:      totalCount,
		TotalCommit:     commitCount,
		MostSameCommit:  most,
		Title:           commit.Title,
		Height:          commit.Height,
		BlockHash:       commit.BlockHash,
		TxResult:        commit.TxResult,
		MainBlockHeight: commit.MainBlockHeight,
		MainBlockHash:   commit.MainBlockHash,
		ChainExecHeight: execHeight,
	}
	key := calcTitleKey(commit.Title)
	status := &pt.ParacrossStatus{
		Title:     commit.Title,
		Height:    commit.Height,
		BlockHash: commit.BlockHash,
	}
	if execMainHeight >= pt.GetDappForkHeight(cfg, pt.ForkLoopCheckCommitTxDone) {
		status.MainHeight = commit.MainBlockHeight
		status.MainHash = commit.MainBlockHash
	}
	return &types.Receipt{
		Ty: types.ExecOk,
		KV: []*types.KeyValue{
			{Key: key, Value: types.Encode(status)},
		},
		Logs: []*types.ReceiptLog{
			{
				Ty:  pt.TyLogParacrossCommitDone,
				Log: types.Encode(log),
			},
		},
	}
}

//GetMostCommit ...
func GetMostCommit(commits [][]byte) (int, string) {
	stats := make(map[string]int)
	n := len(commits)
	for i := 0; i < n; i++ {
		if _, ok := stats[string(commits[i])]; ok {
			stats[string(commits[i])]++
		} else {
			stats[string(commits[i])] = 1
		}
	}
	most := -1
	var hash string
	for k, v := range stats {
		if v > most {
			most = v
			hash = k
		}
	}
	return most, hash
}

//   ForkLoopCheckCommitTxDone   
func getMostResults(mostHash []byte, stat *pt.ParacrossHeightStatus) []byte {
	for i, hash := range stat.BlockDetails.BlockHashs {
		if bytes.Equal(mostHash, hash) {
			return stat.BlockDetails.TxResults[i]
		}
	}
	return nil
}

func hasCommited(addrs []string, addr string) (bool, int) {
	for i, a := range addrs {
		if a == addr {
			return true, i
		}
	}
	return false, 0
}

func getConfigNodes(db dbm.KV, title string) (map[string]struct{}, []string, []byte, error) {
	key := calcParaNodeGroupAddrsKey(title)
	nodes, nodesArray, err := getNodes(db, key)
	if err != nil {
		if errors.Cause(err) != pt.ErrTitleNotExist {
			return nil, nil, nil, errors.Wrapf(err, "getNodes para for title:%s", title)
		}
		key = calcManageConfigNodesKey(title)
		nodes, nodesArray, err = getNodes(db, key)
		if err != nil {
			return nil, nil, nil, errors.Wrapf(err, "getNodes manager for title:%s", title)
		}
	}

	return nodes, nodesArray, key, nil
}

func (a *action) getNodesGroup(title string) (map[string]struct{}, []string, error) {
	cfg := a.api.GetConfig()
	if a.exec.GetMainHeight() < pt.GetDappForkHeight(cfg, pt.ForkCommitTx) {
		nodes, nodesArray, err := getConfigManageNodes(a.db, title)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "getNodes for title:%s", title)
		}
		return nodes, nodesArray, nil
	}

	nodes, nodesArray, _, err := getConfigNodes(a.db, title)
	return nodes, nodesArray, err

}

func (a *action) isValidSuperNode(addr string) error {
	cfg := a.api.GetConfig()
	nodes, _, err := getParacrossNodes(a.db, cfg.GetTitle())
	if err != nil {
		return errors.Wrapf(err, "getNodes for title:%s", cfg.GetTitle())
	}
	if !validNode(addr, nodes) {
		return errors.Wrapf(pt.ErrParaNodeAddrNotExisted, "invalid node=%s", addr)
	}
	return nil
}

//   BlockHash，       
func updateCommitBlockHashs(stat *pt.ParacrossHeightStatus, commit *pt.ParacrossNodeStatus) {
	if stat.BlockDetails == nil {
		stat.BlockDetails = &pt.ParacrossStatusBlockDetails{}
	}
	for _, blockHash := range stat.BlockDetails.BlockHashs {
		if bytes.Equal(blockHash, commit.BlockHash) {
			return
		}
	}
	stat.BlockDetails.BlockHashs = append(stat.BlockDetails.BlockHashs, commit.BlockHash)
	stat.BlockDetails.TxResults = append(stat.BlockDetails.TxResults, commit.TxResult)

}

//  nodes         addrs
func updateCommitAddrs(stat *pt.ParacrossHeightStatus, nodes map[string]struct{}) {
	details := &pt.ParacrossStatusDetails{}
	for i, addr := range stat.Details.Addrs {
		if _, ok := nodes[addr]; ok {
			details.Addrs = append(details.Addrs, addr)
			details.BlockHash = append(details.BlockHash, stat.Details.BlockHash[i])
		}
	}
	stat.Details = details

}

//        ，               ，                  ，           a.height
//           100，          ，  100      80~99 20     ， 20    100      ，       ，        80   100
//      commit.Status.Height  ，               ，      100  
func paraCheckSelfConsOn(cfg *types.Chain33Config, db dbm.KV, commit *pt.ParacrossNodeStatus) (bool, *types.Receipt, error) {
	if !cfg.IsDappFork(commit.Height, pt.ParaX, pt.ForkParaSelfConsStages) {
		return true, nil, nil
	}

	//    ，key   ，               
	isSelfConsOn, err := isSelfConsOn(db, commit.Height)
	if err != nil && errors.Cause(err) != pt.ErrKeyNotExist {
		return false, nil, err
	}
	if !isSelfConsOn {
		clog.Debug("paracross.Commit self consens off", "height", commit.Height)
		return false, &types.Receipt{Ty: types.ExecOk}, nil
	}
	return true, nil, nil
}

func (a *action) preCheckCommitInfo(commit *pt.ParacrossNodeStatus, commitAddrs []string) error {
	cfg := a.api.GetConfig()
	err := checkCommitInfo(cfg, commit)
	if err != nil {
		return err
	}

	titleStatus, err := getTitle(a.db, calcTitleKey(commit.Title))
	if err != nil {
		return errors.Wrapf(err, "getTitle=%s", commit.Title)
	}
	if titleStatus.Height+1 == commit.Height && commit.Height > 0 && !pt.IsParaForkHeight(cfg, commit.MainBlockHeight, pt.ForkLoopCheckCommitTxDone) {
		if !bytes.Equal(titleStatus.BlockHash, commit.PreBlockHash) {
			clog.Error("paracross.Commit", "check PreBlockHash", common.ToHex(titleStatus.BlockHash),
				"commit tx", common.ToHex(commit.PreBlockHash), "commitheit", commit.Height, "from", commitAddrs)
			return pt.ErrParaBlockHashNoMatch
		}
	}

	//     ，            commit        ，        ，     
	//      （1）Bn1        （3） rollback-Bn1   （4） commit-done in Bn2
	//             （2）commit                                 （5）          
	//           
	var dbMainHash []byte
	if !cfg.IsPara() {
		blockHash, err := getBlockHash(a.api, commit.MainBlockHeight)
		if err != nil {
			clog.Error("paracross.Commit getBlockHash", "err", err, "commit tx height", commit.MainBlockHeight, "from", commitAddrs)
			return err
		}
		dbMainHash = blockHash.Hash

	} else {
		block, err := getBlockInfo(a.api, commit.Height)
		if err != nil {
			clog.Error("paracross.Commit getBlockInfo", "err", err, "height", commit.Height, "from", commitAddrs)
			return err
		}
		dbMainHash = block.MainHash
	}

	//    ，           blockhash   commit   
	//     ，     commit      height block   mainHash            mainHash  ，    hash           blockhash    
	if !bytes.Equal(dbMainHash, commit.MainBlockHash) && commit.Height > 0 {
		clog.Error("paracross.Commit blockHash not match", "isMain", !cfg.IsPara(), "db", common.ToHex(dbMainHash),
			"commit", common.ToHex(commit.MainBlockHash), "commitHeight", commit.Height,
			"commitMainHeight", commit.MainBlockHeight, "from", commitAddrs)
		return types.ErrBlockHashNoMatch
	}

	return nil
}

func getValidAddrs(nodes map[string]struct{}, addrs []string) []string {
	var ret []string
	for _, addr := range addrs {
		if !validNode(addr, nodes) {
			clog.Error("paracross.Commit getValidAddrs not valid", "addr", addr)
			continue
		}
		ret = append(ret, addr)
	}
	return ret
}

//bls               3ms (2~4ms)
func (a *action) procBlsSign(nodesArry []string, commit *pt.ParacrossCommitAction) ([]string, error) {
	signAddrs := util.GetAddrsByBitMap(nodesArry, commit.Bls.AddrsMap)
	var pubs []string
	for _, addr := range signAddrs {
		pub, err := getAddrBlsPubKey(a.db, commit.Status.Title, addr)
		if err != nil {
			return nil, errors.Wrapf(err, "pubkey not exist to addr=%s", addr)
		}
		pubs = append(pubs, pub)
	}
	err := verifyBlsSign(a.exec.cryptoCli, pubs, commit)
	if err != nil {
		clog.Error("paracross.Commit bls sign verify", "addr", signAddrs, "nodes", nodesArry, "from", a.fromaddr)
		return nil, err
	}
	return signAddrs, nil
}

func verifyBlsSign(cryptoCli crypto.Crypto, pubs []string, commit *pt.ParacrossCommitAction) error {
	t1 := types.Now()
	//1.   addr   bls   
	pubKeys := make([]crypto.PubKey, 0)
	for _, p := range pubs {
		k, err := common.FromHex(p)
		if err != nil {
			return errors.Wrapf(err, "pub FromHex=%s", p)
		}
		pub, err := cryptoCli.PubKeyFromBytes(k)
		if err != nil {
			return errors.Wrapf(err, "DeserializePublicKey=%s", p)
		}
		pubKeys = append(pubKeys, pub)

	}

	//2.　       , deserial 300us
	sign, err := cryptoCli.SignatureFromBytes(commit.Bls.Sign)
	if err != nil {
		return errors.Wrapf(err, "DeserializeSignature,key=%s", common.ToHex(commit.Bls.Sign))
	}
	//3.        msg
	msg := types.Encode(commit.Status)

	//4. verify 1ms, total 2ms
	agg, err := crypto.ToAggregate(cryptoCli)
	if err != nil {
		return errors.Wrap(err, "ToAggregate")
	}
	err = agg.VerifyAggregatedOne(pubKeys, msg, sign)
	if err != nil {
		clog.Error("paracross.Commit bls sign verify", "title", commit.Status.Title, "height", commit.Status.Height,
			"addrsMap", common.ToHex(commit.Bls.AddrsMap), "sign", common.ToHex(commit.Bls.Sign), "data", common.ToHex(msg))
		return pt.ErrBlsSignVerify
	}
	clog.Debug("paracross procBlsSign success", "title", commit.Status.Title, "height", commit.Status.Height, "time", types.Since(t1))
	return nil
}

//  commit　msg　  
func (a *action) Commit(commit *pt.ParacrossCommitAction) (*types.Receipt, error) {
	cfg := a.api.GetConfig()
	//    ，          
	if cfg.IsPara() {
		isSelfConsOn, receipt, err := paraCheckSelfConsOn(cfg, a.db, commit.Status)
		if !isSelfConsOn {
			return receipt, errors.Wrap(err, "checkSelfConsOn")
		}
	}

	nodesMap, nodesArry, err := a.getNodesGroup(commit.Status.Title)
	if err != nil {
		return nil, errors.Wrap(err, "getNodesGroup")
	}

	//  commitAddrs, bls sign            
	commitAddrs := []string{a.fromaddr}
	if commit.Bls != nil {
		addrs, err := a.procBlsSign(nodesArry, commit)
		if err != nil {
			return nil, errors.Wrap(err, "procBlsSign")
		}
		commitAddrs = addrs
	}

	validAddrs := getValidAddrs(nodesMap, commitAddrs)
	if len(validAddrs) <= 0 {
		return nil, errors.Wrapf(err, "getValidAddrs nil commitAddrs=%s", strings.Join(commitAddrs, ","))
	}

	return a.proCommitMsg(commit.Status, nodesMap, validAddrs)
}

func (a *action) proCommitMsg(commit *pt.ParacrossNodeStatus, nodes map[string]struct{}, commitAddrs []string) (*types.Receipt, error) {
	cfg := a.api.GetConfig()

	err := a.preCheckCommitInfo(commit, commitAddrs)
	if err != nil {
		return nil, err
	}

	titleStatus, err := getTitle(a.db, calcTitleKey(commit.Title))
	if err != nil {
		return nil, errors.Wrapf(err, "getTitle:%s", commit.Title)
	}

	//          ，    record log，              
	if commit.Height <= titleStatus.Height {
		clog.Debug("paracross.Commit record", "node", commitAddrs, "titile", commit.Title, "height", commit.Height)
		return makeRecordReceipt(strings.Join(commitAddrs, ","), commit), nil
	}

	//      ，             
	stat, err := getTitleHeight(a.db, calcTitleHeightKey(commit.Title, commit.Height))
	if err != nil && !isNotFound(err) {
		clog.Error("paracross.Commit getTitleHeight failed", "err", err)
		return nil, err
	}

	var receipt *types.Receipt
	var copyStat *pt.ParacrossHeightStatus
	if isNotFound(err) {
		stat = &pt.ParacrossHeightStatus{
			Status:  pt.ParacrossStatusCommiting,
			Title:   commit.Title,
			Height:  commit.Height,
			Details: &pt.ParacrossStatusDetails{},
		}
		if pt.IsParaForkHeight(cfg, a.exec.GetMainHeight(), pt.ForkCommitTx) {
			stat.MainHeight = commit.MainBlockHeight
			stat.MainHash = commit.MainBlockHash
		}
	} else {
		copyStat = proto.Clone(stat).(*pt.ParacrossHeightStatus)
	}

	for _, addr := range commitAddrs {
		//     ，            commit  
		found, index := hasCommited(stat.Details.Addrs, addr)
		if found {
			stat.Details.BlockHash[index] = commit.BlockHash
		} else {
			stat.Details.Addrs = append(stat.Details.Addrs, addr)
			stat.Details.BlockHash = append(stat.Details.BlockHash, commit.BlockHash)
		}
	}

	// commit.MainBlockHeight      ，   a.exec.MainHeight   ，      MainHeight       tx，
	//   loopCommitTxDone                   
	if pt.IsParaForkHeight(cfg, commit.MainBlockHeight, pt.ForkLoopCheckCommitTxDone) {
		updateCommitBlockHashs(stat, commit)
	}
	receipt = makeCommitReceipt(strings.Join(commitAddrs, ","), commit, copyStat, stat)

	//   fork pt.ForkCommitTx=0,   ForkCommitTx   nodegroup，     dappFork   true
	if cfg.IsDappFork(commit.MainBlockHeight, pt.ParaX, pt.ForkCommitTx) {
		updateCommitAddrs(stat, nodes)
	}
	saveTitleHeight(a.db, calcTitleHeightKey(stat.Title, stat.Height), stat)
	//fork     stat     nodes     
	if pt.IsParaForkHeight(cfg, stat.MainHeight, pt.ForkLoopCheckCommitTxDone) {
		r := makeCommitStatReceipt(stat)
		receipt = mergeReceipt(receipt, r)
	}

	if commit.Height > titleStatus.Height+1 {
		saveTitleHeight(a.db, calcTitleHeightKey(commit.Title, commit.Height), stat)
		//            ，           ，    0  
		allow, err := a.isAllowConsensJump(commit, titleStatus)
		if err != nil {
			clog.Error("paracross.Commit allowJump", "err", err)
			return nil, err
		}
		if !allow {
			return receipt, nil
		}
	}
	r, err := a.commitTxDone(commit, stat, titleStatus, nodes)
	if err != nil {
		return nil, err
	}
	receipt = mergeReceipt(receipt, r)
	return receipt, nil
}

//    stat      blockhash   ，  crossTxHash   ，    stat     mostCommitStatus
func (a *action) commitTxDone(nodeStatus *pt.ParacrossNodeStatus, stat *pt.ParacrossHeightStatus, titleStatus *pt.ParacrossStatus,
	nodes map[string]struct{}) (*types.Receipt, error) {
	receipt := &types.Receipt{}

	clog.Debug("paracross.Commit commit", "stat.title", stat.Title, "stat.height", stat.Height, "notes", len(nodes))
	for i, v := range stat.Details.Addrs {
		clog.Debug("paracross.Commit commit detail", "addr", v, "hash", common.ToHex(stat.Details.BlockHash[i]))
	}

	mostCount, mostHash := GetMostCommit(stat.Details.BlockHash)
	if !isCommitDone(len(nodes), mostCount) {
		return receipt, nil
	}
	clog.Debug("paracross.Commit commit ----pass", "most", mostCount, "mostHash", common.ToHex([]byte(mostHash)))
	stat.Status = pt.ParacrossStatusCommitDone
	saveTitleHeight(a.db, calcTitleHeightKey(stat.Title, stat.Height), stat)

	//     stat      
	cfg := a.api.GetConfig()
	if pt.IsParaForkHeight(cfg, stat.MainHeight, pt.ForkLoopCheckCommitTxDone) {
		r := makeCommitStatReceipt(stat)
		receipt = mergeReceipt(receipt, r)
	}

	//add commit done receipt
	receiptDone := makeDoneReceipt(cfg, a.exec.GetMainHeight(), a.height, nodeStatus, int32(mostCount), int32(len(stat.Details.Addrs)), int32(len(nodes)))
	receipt = mergeReceipt(receipt, receiptDone)

	r, err := a.commitTxDoneStep2(nodeStatus, stat, titleStatus)
	if err != nil {
		return nil, err
	}
	receipt = mergeReceipt(receipt, r)
	return receipt, nil
}

func (a *action) commitTxDoneStep2(nodeStatus *pt.ParacrossNodeStatus, stat *pt.ParacrossHeightStatus, titleStatus *pt.ParacrossStatus) (*types.Receipt, error) {
	receipt := &types.Receipt{}

	titleStatus.Title = nodeStatus.Title
	titleStatus.Height = nodeStatus.Height
	titleStatus.BlockHash = nodeStatus.BlockHash
	cfg := a.api.GetConfig()
	if pt.IsParaForkHeight(cfg, a.exec.GetMainHeight(), pt.ForkLoopCheckCommitTxDone) {
		titleStatus.MainHeight = nodeStatus.MainBlockHeight
		titleStatus.MainHash = nodeStatus.MainBlockHash
	}
	saveTitle(a.db, calcTitleKey(titleStatus.Title), titleStatus)

	clog.Debug("paracross.Commit commit done", "height", nodeStatus.Height, "statusBlockHash", common.ToHex(nodeStatus.BlockHash))

	//parallel chain not need to process cross commit tx here
	if cfg.IsPara() {
		//        
		selfBlockHash, err := getBlockHash(a.api, nodeStatus.Height)
		if err != nil {
			clog.Error("paracross.CommitDone getBlockHash", "err", err, "commit tx height", nodeStatus.Height, "tx", common.ToHex(a.txhash))
			return nil, err
		}
		//     blockhash   hash   ，         
		if !bytes.Equal(selfBlockHash.Hash, nodeStatus.BlockHash) {
			clog.Error("paracross.CommitDone mosthash not match", "height", nodeStatus.Height,
				"blockHash", common.ToHex(selfBlockHash.Hash), "mosthash", common.ToHex(nodeStatus.BlockHash))
			return nil, types.ErrConsensusHashErr
		}

		//         
		rewardReceipt, err := a.reward(nodeStatus, stat)
		//                  
		if err != nil {
			clog.Error("paracross mining reward err", "height", nodeStatus.Height,
				"blockhash", common.ToHex(nodeStatus.BlockHash), "err", err)
			return nil, err
		}
		receipt = mergeReceipt(receipt, rewardReceipt)
		return receipt, nil
	}

	//  ，      
	r, err := a.procCrossTxs(nodeStatus)
	if err != nil {
		return nil, err
	}
	receipt = mergeReceipt(receipt, r)
	return receipt, nil
}

func isHaveCrossTxs(cfg *types.Chain33Config, status *pt.ParacrossNodeStatus) bool {
	//ForkLoopCheckCommitTxDone        txResult   ，                 tx
	if pt.IsParaForkHeight(cfg, status.MainBlockHeight, pt.ForkLoopCheckCommitTxDone) {
		return true
	}

	haveCrossTxs := len(status.CrossTxHashs) > 0
	//ForkCommitTx ，CrossTxHashs[][]             hash，     [0] nil
	if status.Height > 0 && pt.IsParaForkHeight(cfg, status.MainBlockHeight, pt.ForkCommitTx) && len(status.CrossTxHashs[0]) == 0 {
		haveCrossTxs = false
	}
	return haveCrossTxs

}

func (a *action) procCrossTxs(status *pt.ParacrossNodeStatus) (*types.Receipt, error) {
	cfg := a.api.GetConfig()
	if enableParacrossTransfer && status.Height > 0 && isHaveCrossTxs(cfg, status) {
		clog.Debug("paracross.Commit commitDone do cross", "height", status.Height)
		crossTxReceipt, err := a.execCrossTxs(status)
		if err != nil {
			return nil, err
		}
		return crossTxReceipt, nil
	}
	return nil, nil
}

//                 ，         statedb，  tx              
func (a *action) loopCommitTxDone(title string) (*types.Receipt, error) {
	receipt := &types.Receipt{}

	nodes, _, err := getParacrossNodes(a.db, title)
	if err != nil {
		return nil, errors.Wrapf(err, "getNodes for title:%s", title)
	}
	//           
	titleStatus, err := getTitle(a.db, calcTitleKey(title))
	if err != nil {
		return nil, errors.Wrapf(err, "getTitle:%s", title)
	}
	//             ，    
	cfg := a.api.GetConfig()
	if !pt.IsParaForkHeight(cfg, titleStatus.GetMainHeight(), pt.ForkLoopCheckCommitTxDone) {
		return nil, errors.Wrapf(pt.ErrForkHeightNotReach,
			"titleHeight:%d,forkHeight:%d", titleStatus.MainHeight, pt.GetDappForkHeight(cfg, pt.ForkLoopCheckCommitTxDone))
	}

	loopHeight := titleStatus.Height

	for {
		loopHeight++

		stat, err := getTitleHeight(a.db, calcTitleHeightKey(title, loopHeight))
		if err != nil {
			clog.Error("paracross.loopCommitTxDone getTitleHeight failed", "title", title, "height", loopHeight, "err", err.Error())
			return receipt, err
		}
		//      
		if stat.MainHeight > a.exec.GetMainHeight() {
			return receipt, nil
		}

		r, err := a.checkCommitTxDone(title, stat, nodes)
		if err != nil {
			clog.Error("paracross.loopCommitTxDone checkExecCommitTxDone", "para title", title, "height", stat.Height, "error", err)
			return receipt, nil
		}
		if r == nil {
			return receipt, nil
		}
		receipt = mergeReceipt(receipt, r)

	}
}

func (a *action) checkCommitTxDone(title string, stat *pt.ParacrossHeightStatus, nodes map[string]struct{}) (*types.Receipt, error) {
	status, err := getTitle(a.db, calcTitleKey(title))
	if err != nil {
		return nil, errors.Wrapf(err, "getTitle:%s", title)
	}

	//    stat       status  +1，       ，  ，               ，            
	if stat.Height > status.Height+1 {
		return nil, nil
	}

	return a.commitTxDoneByStat(stat, status, nodes)

}

//   stat    commitDone      commitMostStatus     
func (a *action) commitTxDoneByStat(stat *pt.ParacrossHeightStatus, titleStatus *pt.ParacrossStatus, nodes map[string]struct{}) (*types.Receipt, error) {
	receipt := &types.Receipt{}
	clog.Debug("paracross.commitTxDoneByStat", "stat.title", stat.Title, "stat.height", stat.Height, "notes", len(nodes))
	for i, v := range stat.Details.Addrs {
		clog.Debug("paracross.commitTxDoneByStat detail", "addr", v, "hash", common.ToHex(stat.Details.BlockHash[i]))
	}

	updateCommitAddrs(stat, nodes)
	commitCount := len(stat.Details.Addrs)
	most, mostHash := GetMostCommit(stat.Details.BlockHash)
	if !isCommitDone(len(nodes), most) {
		return nil, nil
	}
	clog.Debug("paracross.commitTxDoneByStat ----pass", "most", most, "mostHash", common.ToHex([]byte(mostHash)))
	stat.Status = pt.ParacrossStatusCommitDone
	saveTitleHeight(a.db, calcTitleHeightKey(stat.Title, stat.Height), stat)
	r := makeCommitStatReceipt(stat)
	receipt = mergeReceipt(receipt, r)

	txRst := getMostResults([]byte(mostHash), stat)
	mostStatus := &pt.ParacrossNodeStatus{
		MainBlockHash:   stat.MainHash,
		MainBlockHeight: stat.MainHeight,
		Title:           stat.Title,
		Height:          stat.Height,
		BlockHash:       []byte(mostHash),
		TxResult:        txRst,
	}

	//add commit done receipt
	cfg := a.api.GetConfig()
	receiptDone := makeDoneReceipt(cfg, a.exec.GetMainHeight(), a.height, mostStatus, int32(most), int32(commitCount), int32(len(nodes)))
	receipt = mergeReceipt(receipt, receiptDone)

	r, err := a.commitTxDoneStep2(mostStatus, stat, titleStatus)
	if err != nil {
		return nil, err
	}
	receipt = mergeReceipt(receipt, r)
	return receipt, nil
}

//        ：             -1，        ，         
func (a *action) isAllowMainConsensJump(commit *pt.ParacrossNodeStatus, titleStatus *pt.ParacrossStatus) bool {
	cfg := a.api.GetConfig()
	if cfg.IsDappFork(a.exec.GetMainHeight(), pt.ParaX, pt.ForkLoopCheckCommitTxDone) {
		if titleStatus.Height == -1 {
			return true
		}
	}

	return false
}

//            ：1，        ，2：commit                                ，       ，
// 1.     ，            １  ，           ， 0    ，     １     
// 2.     ，  stage.blockHeight== commit.height，   stage          
func (a *action) isAllowParaConsensJump(commit *pt.ParacrossNodeStatus, titleStatus *pt.ParacrossStatus) (bool, error) {
	cfg := a.api.GetConfig()
	if cfg.IsDappFork(a.height, pt.ParaX, pt.ForkParaSelfConsStages) {
		stage, err := getSelfConsOneStage(a.db, commit.Height)
		if err != nil && errors.Cause(err) != pt.ErrKeyNotExist {
			return false, err
		}
		if stage == nil {
			return false, nil
		}
		return stage.StartHeight == commit.Height, nil
	}

	//       １    
	return titleStatus.Height == -1, nil
}

func (a *action) isAllowConsensJump(commit *pt.ParacrossNodeStatus, titleStatus *pt.ParacrossStatus) (bool, error) {
	cfg := a.api.GetConfig()
	if cfg.IsPara() {
		return a.isAllowParaConsensJump(commit, titleStatus)
	}
	return a.isAllowMainConsensJump(commit, titleStatus), nil

}

func execCrossTx(a *action, cross *types.TransactionDetail, crossTxHash []byte) (*types.Receipt, error) {
	if !bytes.HasSuffix(cross.Tx.Execer, []byte(pt.ParaX)) {
		return nil, nil
	}
	var payload pt.ParacrossAction
	err := types.Decode(cross.Tx.Payload, &payload)
	if err != nil {
		clog.Crit("paracross.Commit Decode Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
		return nil, err
	}

	if payload.Ty == pt.ParacrossActionCrossAssetTransfer {
		act, err := getCrossAction(payload.GetCrossAssetTransfer(), string(cross.Tx.Execer))
		if err != nil {
			clog.Crit("paracross.Commit getCrossAction Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
			return nil, err
		}
		if act == pt.ParacrossMainAssetWithdraw || act == pt.ParacrossParaAssetTransfer {
			receipt, err := a.crossAssetTransfer(payload.GetCrossAssetTransfer(), act, cross.Tx)
			if err != nil {
				clog.Crit("paracross.Commit crossAssetTransfer Tx failed", "error", err, "act", act, "txHash", common.ToHex(crossTxHash))
				return nil, err
			}
			clog.Debug("paracross.Commit crossAssetTransfer done", "act", act, "txHash", common.ToHex(crossTxHash))
			return receipt, nil
		}

	}

	//     ，      withdraw,    CrossAssetTransfer     action
	if payload.Ty == pt.ParacrossActionAssetWithdraw {
		receiptWithdraw, err := a.assetWithdraw(payload.GetAssetWithdraw(), cross.Tx)
		if err != nil {
			clog.Crit("paracross.Commit withdraw Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
			return nil, errors.Cause(err)
		}

		clog.Debug("paracross.Commit WithdrawCoins", "txHash", common.ToHex(crossTxHash))
		return receiptWithdraw, nil
	}
	return nil, nil

}

func rollbackCrossTx(a *action, cross *types.TransactionDetail, crossTxHash []byte) (*types.Receipt, error) {
	if !bytes.HasSuffix(cross.Tx.Execer, []byte(pt.ParaX)) {
		return nil, nil
	}
	var payload pt.ParacrossAction
	err := types.Decode(cross.Tx.Payload, &payload)
	if err != nil {
		clog.Crit("paracross.Commit.rollbackCrossTx Decode Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
		return nil, err
	}

	if payload.Ty == pt.ParacrossActionCrossAssetTransfer {
		act, err := getCrossAction(payload.GetCrossAssetTransfer(), string(cross.Tx.Execer))
		if err != nil {
			clog.Crit("paracross.Commit.rollbackCrossTx getCrossAction failed", "error", err, "txHash", common.ToHex(crossTxHash))
			return nil, err
		}
		//     ，            transfer  
		if act == pt.ParacrossMainAssetTransfer {
			receipt, err := a.assetTransferRollback(payload.GetCrossAssetTransfer(), cross.Tx)
			if err != nil {
				clog.Crit("paracross.Commit crossAssetTransfer rbk failed", "error", err, "txHash", common.ToHex(crossTxHash))
				return nil, errors.Cause(err)
			}

			clog.Debug("paracross.Commit crossAssetTransfer rollbackCrossTx", "txHash", common.ToHex(crossTxHash), "mainHeight", a.height)
			return receipt, nil
		}
		//     ，             withdraw  
		if act == pt.ParacrossParaAssetWithdraw {
			receipt, err := a.paraAssetWithdrawRollback(payload.GetCrossAssetTransfer(), cross.Tx)
			if err != nil {
				clog.Crit("paracross.Commit rbk paraAssetWithdraw Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
				return nil, errors.Cause(err)
			}

			clog.Debug("paracross.Commit paraAssetWithdraw rollbackCrossTx", "txHash", common.ToHex(crossTxHash), "mainHeight", a.height)
			return receipt, nil
		}
	}

	//     ，            transfer  
	if payload.Ty == pt.ParacrossActionAssetTransfer {
		cfg := payload.GetAssetTransfer()
		transfer := &pt.CrossAssetTransfer{
			AssetSymbol: cfg.Cointoken,
			Amount:      cfg.Amount,
			Note:        string(cfg.Note),
			ToAddr:      cfg.To,
		}

		receipt, err := a.assetTransferRollback(transfer, cross.Tx)
		if err != nil {
			clog.Crit("paracross.Commit rbk asset transfer Tx failed", "error", err, "txHash", common.ToHex(crossTxHash))
			return nil, errors.Cause(err)
		}

		clog.Debug("paracross.Commit assetTransfer rollbackCrossTx", "txHash", common.ToHex(crossTxHash), "mainHeight", a.height)
		return receipt, nil
	}
	return nil, nil

}

func getCrossTxHashsByRst(api client.QueueProtocolAPI, status *pt.ParacrossNodeStatus) ([][]byte, []byte, error) {
	//     tx
	cfg := api.GetConfig()
	rst, err := hex.DecodeString(string(status.TxResult))
	if err != nil {
		clog.Error("getCrossTxHashs decode rst", "CrossTxResult", string(status.TxResult), "paraHeight", status.Height)
		return nil, nil, types.ErrInvalidParam
	}
	clog.Debug("getCrossTxHashsByRst", "height", status.Height, "txResult", string(status.TxResult))

	if !cfg.IsDappFork(status.MainBlockHeight, pt.ParaX, pt.ForkParaAssetTransferRbk) {
		if len(rst) == 0 {
			return nil, nil, nil
		}
	}

	blockDetail, err := GetBlock(api, status.MainBlockHash)
	if err != nil {
		return nil, nil, err
	}

	//            
	paraAllTxs := FilterTxsForPara(cfg, blockDetail.FilterParaTxsByTitle(cfg, status.Title))
	var baseHashs [][]byte
	for _, tx := range paraAllTxs {
		baseHashs = append(baseHashs, tx.Hash())
	}
	paraCrossHashs := FilterParaCrossTxHashes(paraAllTxs)
	crossRst := util.CalcBitMapByBitMap(paraCrossHashs, baseHashs, rst)
	clog.Debug("getCrossTxHashsByRst.crossRst", "height", status.Height, "txResult", common.ToHex(crossRst), "len", len(paraCrossHashs))

	return paraCrossHashs, crossRst, nil

}

func getCrossTxHashs(api client.QueueProtocolAPI, status *pt.ParacrossNodeStatus) ([][]byte, []byte, error) {
	cfg := api.GetConfig()
	if pt.IsParaForkHeight(cfg, status.MainBlockHeight, pt.ForkLoopCheckCommitTxDone) {
		return getCrossTxHashsByRst(api, status)
	}
	if !pt.IsParaForkHeight(cfg, status.MainBlockHeight, pt.ForkCommitTx) {
		return status.CrossTxHashs, status.CrossTxResult, nil
	}

	if len(status.CrossTxHashs) == 0 {
		clog.Error("getCrossTxHashs len=0", "paraHeight", status.Height,
			"mainHeight", status.MainBlockHeight, "mainHash", common.ToHex(status.MainBlockHash))
		return nil, nil, types.ErrCheckTxHash
	}

	blockDetail, err := GetBlock(api, status.MainBlockHash)
	if err != nil {
		return nil, nil, err
	}
	//  
	paraBaseTxs := FilterTxsForPara(cfg, blockDetail.FilterParaTxsByTitle(cfg, status.Title))
	paraCrossHashs := FilterParaCrossTxHashes(paraBaseTxs)
	var baseHashs [][]byte
	for _, tx := range paraBaseTxs {
		baseHashs = append(baseHashs, tx.Hash())
	}
	baseCheckTxHash := CalcTxHashsHash(baseHashs)
	crossCheckHash := CalcTxHashsHash(paraCrossHashs)
	if !bytes.Equal(status.CrossTxHashs[0], crossCheckHash) {
		clog.Error("getCrossTxHashs para hash not equal", "paraHeight", status.Height,
			"mainHeight", status.MainBlockHeight, "mainHash", common.ToHex(status.MainBlockHash),
			"main.crossHash", common.ToHex(crossCheckHash), "commit.crossHash", common.ToHex(status.CrossTxHashs[0]),
			"main.baseHash", common.ToHex(baseCheckTxHash), "commit.baseHash", common.ToHex(status.TxHashs[0]))
		for _, hash := range baseHashs {
			clog.Error("getCrossTxHashs base tx hash", "txhash", common.ToHex(hash))
		}
		for _, hash := range paraCrossHashs {
			clog.Error("getCrossTxHashs paracross tx hash", "txhash", common.ToHex(hash))
		}
		return nil, nil, types.ErrCheckTxHash
	}

	//     tx
	rst, err := hex.DecodeString(string(status.CrossTxResult))
	if err != nil {
		clog.Error("getCrossTxHashs decode rst", "CrossTxResult", string(status.CrossTxResult), "paraHeight", status.Height)
		return nil, nil, types.ErrInvalidParam
	}

	return paraCrossHashs, rst, nil

}

func crossTxProc(a *action, txHash []byte, fn func(*action, *types.TransactionDetail, []byte) (*types.Receipt, error)) (*types.Receipt, error) {
	tx, err := GetTx(a.api, txHash)
	if err != nil {
		clog.Crit("paracross.Commit Load Tx failed", "error", err, "txHash", common.ToHex(txHash))
		return nil, err
	}
	if tx == nil {
		clog.Error("paracross.Commit Load Tx nil", "error", err, "txHash", common.ToHex(txHash))
		return nil, types.ErrHashNotExist
	}
	receiptCross, err := fn(a, tx, txHash)
	if err != nil {
		clog.Error("paracross.Commit execCrossTx", "error", err)
		return nil, errors.Cause(err)
	}
	return receiptCross, nil
}

func (a *action) execCrossTxs(status *pt.ParacrossNodeStatus) (*types.Receipt, error) {
	var receipt types.Receipt

	crossTxHashs, crossTxResult, err := getCrossTxHashs(a.api, status)
	if err != nil {
		clog.Error("paracross.Commit getCrossTxHashs", "err", err.Error())
		return nil, err
	}

	for i := 0; i < len(crossTxHashs); i++ {
		clog.Debug("paracross.Commit commitDone", "do cross number", i, "hash", common.ToHex(crossTxHashs[i]),
			"res", util.BitMapBit(crossTxResult, uint32(i)))
		if util.BitMapBit(crossTxResult, uint32(i)) {
			receiptCross, err := crossTxProc(a, crossTxHashs[i], execCrossTx)
			if err != nil {
				clog.Error("paracross.Commit execCrossTx", "para title", status.Title, "para height", status.Height,
					"para tx index", i, "error", err)
				return nil, errors.Cause(err)
			}
			if receiptCross == nil {
				continue
			}
			receipt.KV = append(receipt.KV, receiptCross.KV...)
			receipt.Logs = append(receipt.Logs, receiptCross.Logs...)
		} else {
			clog.Error("paracross.Commit commitDone", "do cross number", i, "hash",
				common.ToHex(crossTxHashs[i]), "para res", util.BitMapBit(crossTxResult, uint32(i)))
			cfg := a.api.GetConfig()
			if cfg.IsDappFork(a.height, pt.ParaX, pt.ForkParaAssetTransferRbk) {
				receiptCross, err := crossTxProc(a, crossTxHashs[i], rollbackCrossTx)
				if err != nil {
					clog.Error("paracross.Commit rollbackCrossTx", "para title", status.Title, "para height", status.Height,
						"para tx index", i, "error", err)
					return nil, errors.Cause(err)
				}
				if receiptCross == nil {
					continue
				}
				receipt.KV = append(receipt.KV, receiptCross.KV...)
				receipt.Logs = append(receipt.Logs, receiptCross.Logs...)
			}
		}
	}

	return &receipt, nil
}

func (a *action) assetTransferMainCheck(cfg *types.Chain33Config, transfer *types.AssetsTransfer) error {
	//      nodegroup  ，      ,      ，        
	if cfg.IsDappFork(a.height, pt.ParaX, pt.ForkParaAssetTransferRbk) {
		if len(transfer.To) == 0 {
			return errors.Wrap(types.ErrInvalidParam, "toAddr should not be null")
		}
		return a.isAllowTransfer()
	}
	return nil
}

func (a *action) AssetTransfer(transfer *types.AssetsTransfer) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec", "AssetTransfer", transfer.Cointoken, "transfer", "")
	cfg := a.api.GetConfig()
	isPara := cfg.IsPara()

	//      nodegroup  ，      ,      ，        
	if !isPara {
		err := a.assetTransferMainCheck(cfg, transfer)
		if err != nil {
			return nil, errors.Wrap(err, "AssetTransfer check")
		}
	}

	receipt, err := a.assetTransfer(transfer)
	if err != nil {
		return nil, errors.Wrap(err, "AssetTransfer")
	}
	return receipt, nil
}

func (a *action) assetWithdrawMainCheck(cfg *types.Chain33Config, withdraw *types.AssetsWithdraw) error {
	if !cfg.IsDappFork(a.height, pt.ParaX, "ForkParacrossWithdrawFromParachain") {
		if withdraw.Cointoken != "" {
			return errors.Wrapf(types.ErrNotSupport, "not support,token=%s", withdraw.Cointoken)
		}
	}

	//rbk fork 　    nodegroup　conf，      
	if cfg.IsDappFork(a.height, pt.ParaX, pt.ForkParaAssetTransferRbk) {
		if len(withdraw.To) == 0 {
			return errors.Wrap(types.ErrInvalidParam, "toAddr should not be null")
		}
		err := a.isAllowTransfer()
		if err != nil {
			return errors.Wrap(err, "AssetWithdraw not allow")
		}
	}
	return nil
}

func (a *action) AssetWithdraw(withdraw *types.AssetsWithdraw) (*types.Receipt, error) {
	//      ，          
	cfg := a.api.GetConfig()
	isPara := cfg.IsPara()
	if !isPara {
		err := a.assetWithdrawMainCheck(cfg, withdraw)
		if err != nil {
			return nil, err
		}
	}

	if !isPara {
		//         ，      ，    
		return nil, nil
	}
	clog.Debug("paracross.AssetWithdraw isPara", "execer", string(a.tx.Execer),
		"txHash", common.ToHex(a.tx.Hash()), "token name", withdraw.Cointoken)
	receipt, err := a.assetWithdraw(withdraw, a.tx)
	if err != nil {
		return nil, errors.Wrap(err, "AssetWithdraw failed")
	}
	return receipt, nil
}

func (a *action) crossAssetTransferMainCheck(transfer *pt.CrossAssetTransfer) error {
	err := a.isAllowTransfer()
	if err != nil {
		return errors.Wrap(err, "not Allow")
	}

	if len(transfer.AssetExec) == 0 || len(transfer.AssetSymbol) == 0 || transfer.Amount == 0 || len(transfer.ToAddr) == 0 {
		return errors.Wrapf(types.ErrInvalidParam, "exec=%s, symbol=%s, amount=%d,toAddr=%s should not be null",
			transfer.AssetExec, transfer.AssetSymbol, transfer.Amount, transfer.ToAddr)
	}
	return nil
}

func (a *action) CrossAssetTransfer(transfer *pt.CrossAssetTransfer) (*types.Receipt, error) {
	cfg := a.api.GetConfig()
	isPara := cfg.IsPara()

	//      
	if !isPara {
		err := a.crossAssetTransferMainCheck(transfer)
		if err != nil {
			return nil, err
		}
	}

	//    ForkRootHash     crossAssetTransfer
	if isPara && !cfg.IsFork(a.exec.GetMainHeight(), "ForkRootHash") {
		return nil, errors.Wrap(types.ErrNotSupport, "not Allow before ForkRootHash")
	}

	act, err := getCrossAction(transfer, string(a.tx.Execer))
	if act == pt.ParacrossNoneTransfer {
		return nil, errors.Wrap(err, "non action")
	}
	//         ，      ，    
	if !isPara && (act == pt.ParacrossMainAssetWithdraw || act == pt.ParacrossParaAssetTransfer) {
		return nil, nil
	}
	receipt, err := a.crossAssetTransfer(transfer, act, a.tx)
	if err != nil {
		return nil, errors.Wrap(err, "CrossAssetTransfer failed")
	}
	return receipt, nil
}

func getTitleFrom(exec []byte) ([]byte, error) {
	last := bytes.LastIndex(exec, []byte("."))
	if last == -1 {
		return nil, types.ErrNotFound
	}
	//         . ，    title     `.`    
	return exec[:last+1], nil
}

func (a *action) isAllowTransfer() error {
	//1.     nodegroup　   
	tempTitle, err := getTitleFrom(a.tx.Execer)
	if err != nil {
		return errors.Wrapf(types.ErrInvalidParam, "not para chain exec=%s", string(a.tx.Execer))
	}
	nodes, _, err := a.getNodesGroup(string(tempTitle))
	if err != nil {
		return errors.Wrapf(err, "nodegroup not config,title=%s", tempTitle)
	}
	if len(nodes) == 0 {
		return errors.Wrapf(types.ErrNotSupport, "nodegroup not create,title=%s", tempTitle)
	}
	//2.          
	if !types.IsParaExecName(string(a.tx.Execer)) {
		return errors.Wrapf(types.ErrInvalidParam, "exec=%s,should prefix with user.p.", string(a.tx.Execer))
	}

	return nil
}

/*
func (a *Paracross) CrossLimits(tx *types.Transaction, index int) bool {
	if tx.GroupCount < 2 {
		return true
	}

	txs, err := a.GetTxGroup(index)
	if err != nil {
		clog.Error("crossLimits", "get tx group failed", err, "hash", common.ToHex(tx.Hash()))
		return false
	}

	titles := make(map[string] struct{})
	for _, txTmp := range txs {
		title, err := getTitleFrom(txTmp.Execer)
		if err == nil {
			titles[string(title)] = struct{}{}
		}
	}
	return len(titles) <= 1
}
*/

func (a *action) Transfer(transfer *types.AssetsTransfer, tx *types.Transaction, index int) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec Transfer", "symbol", transfer.Cointoken, "amount",
		transfer.Amount, "to", tx.To)
	from := tx.From()

	cfg := a.api.GetConfig()
	acc, err := account.NewAccountDB(cfg, pt.ParaX, transfer.Cointoken, a.db)
	if err != nil {
		clog.Error("Transfer failed", "err", err)
		return nil, err
	}
	//to   execs     
	if dapp.IsDriverAddress(tx.GetRealToAddr(), a.height) {
		return acc.TransferToExec(from, tx.GetRealToAddr(), transfer.Amount)
	}
	return acc.Transfer(from, tx.GetRealToAddr(), transfer.Amount)
}

func (a *action) Withdraw(withdraw *types.AssetsWithdraw, tx *types.Transaction, index int) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec Withdraw", "symbol", withdraw.Cointoken, "amount",
		withdraw.Amount, "to", tx.To)
	cfg := a.api.GetConfig()
	acc, err := account.NewAccountDB(cfg, pt.ParaX, withdraw.Cointoken, a.db)
	if err != nil {
		clog.Error("Withdraw failed", "err", err)
		return nil, err
	}
	if dapp.IsDriverAddress(tx.GetRealToAddr(), a.height) || dapp.ExecAddress(withdraw.ExecName) == tx.GetRealToAddr() {
		return acc.TransferWithdraw(tx.From(), tx.GetRealToAddr(), withdraw.Amount)
	}
	return nil, types.ErrToAddrNotSameToExecAddr
}

func (a *action) TransferToExec(transfer *types.AssetsTransferToExec, tx *types.Transaction, index int) (*types.Receipt, error) {
	clog.Debug("Paracross.Exec TransferToExec", "symbol", transfer.Cointoken, "amount",
		transfer.Amount, "to", tx.To)
	from := tx.From()

	cfg := a.api.GetConfig()
	acc, err := account.NewAccountDB(cfg, pt.ParaX, transfer.Cointoken, a.db)
	if err != nil {
		clog.Error("TransferToExec failed", "err", err)
		return nil, err
	}
	//to   execs     
	if dapp.IsDriverAddress(tx.GetRealToAddr(), a.height) || dapp.ExecAddress(transfer.ExecName) == tx.GetRealToAddr() {
		return acc.TransferToExec(from, tx.GetRealToAddr(), transfer.Amount)
	}
	return nil, types.ErrToAddrNotSameToExecAddr
}
