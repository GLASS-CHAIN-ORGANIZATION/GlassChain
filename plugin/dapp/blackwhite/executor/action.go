// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"math"

	"bytes"
	"strconv"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/common"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	gt "github.com/33cn/plugin/plugin/dapp/blackwhite/types"
)

const (
	maxAmount      int64 = 100 * types.Coin
	minAmount      int64 = 1 * types.Coin
	minPlayerCount int32 = 3
	maxPlayerCount int32 = 100000
	lockAmount     int64 = types.Coin / 100 //       
	showTimeout    int64 = 60 * 5           //         
	maxPlayTimeout int64 = 60 * 60 * 24     //             
	minPlayTimeout int64 = 60 * 10          //             

	white = "0"
	black = "1"
)

type action struct {
	coinsAccount *account.DB
	db           dbm.KV
	txhash       []byte
	fromaddr     string
	blocktime    int64
	height       int64
	index        int32
	execaddr     string
	api          client.QueueProtocolAPI
}

type resultCalc struct {
	Addr    string
	amount  int64
	IsWin   bool
	IsBlack []bool
}

type addrResult struct {
	addr   string
	amount int64
}

func newAction(t *Blackwhite, tx *types.Transaction, index int32) *action {
	hash := tx.Hash()
	fromaddr := tx.From()
	return &action{t.GetCoinsAccount(), t.GetStateDB(), hash, fromaddr,
		t.GetBlockTime(), t.GetHeight(), index, dapp.ExecAddress(string(tx.Execer)), t.GetAPI()}
}

func (a *action) Create(create *gt.BlackwhiteCreate) (*types.Receipt, error) {
	if create.PlayAmount < minAmount || create.PlayAmount > maxAmount {
		return nil, types.ErrAmount
	}
	if create.PlayerCount < minPlayerCount || create.PlayerCount > maxPlayerCount {
		return nil, types.ErrInvalidParam
	}
	if create.Timeout < minPlayTimeout || create.Timeout > maxPlayTimeout {
		return nil, types.ErrInvalidParam
	}

	receipt, err := a.coinsAccount.ExecFrozen(a.fromaddr, a.execaddr, lockAmount)
	if err != nil {
		clog.Error("blackwhite create ", "addr", a.fromaddr, "execaddr", a.execaddr, "ExecFrozen amount", lockAmount)
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	logs = append(logs, receipt.Logs...)
	kv = append(kv, receipt.KV...)

	round := newRound(create, a.fromaddr)
	round.GameID = common.ToHex(a.txhash)
	round.CreateTime = a.blocktime
	round.Index = heightIndexToIndex(a.height, a.index)

	key := calcMavlRoundKey(round.GameID)
	value := types.Encode(round)
	kv = append(kv, &types.KeyValue{Key: key, Value: value})

	receiptLog := a.GetReceiptLog(round, round.GetCreateAddr())
	logs = append(logs, receiptLog)

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

func (a *action) Play(play *gt.BlackwhitePlay) (*types.Receipt, error) {
	//   GameID
	value, err := a.db.Get(calcMavlRoundKey(play.GameID))
	if err != nil {
		clog.Error("blackwhite play ", "addr", a.fromaddr, "execaddr", a.execaddr, "get round failed",
			play.GameID, "err", err)
		return nil, err
	}
	var round gt.BlackwhiteRound
	err = types.Decode(value, &round)
	if err != nil {
		clog.Error("blackwhite play ", "addr", a.fromaddr, "execaddr", a.execaddr, "decode round failed",
			play.GameID, "err", err)
		return nil, err
	}

	//       
	if gt.BlackwhiteStatusPlay != round.Status && gt.BlackwhiteStatusCreate != round.Status {
		err := gt.ErrIncorrectStatus
		clog.Error("blackwhite play ", "addr", a.fromaddr, "round status", round.Status, "status is not match, GameID ",
			play.GameID, "err", err)
		return nil, err
	}

	//        
	for _, addrResult := range round.AddrResult {
		if addrResult.Addr == a.fromaddr {
			err := gt.ErrRepeatPlayerAddr
			clog.Error("blackwhite play ", "addr", a.fromaddr, "execaddr", a.execaddr, "repeat address GameID",
				play.GameID, "err", err)
			return nil, err
		}
	}

	if play.Amount < round.PlayAmount {
		clog.Error("blackwhite play ", "addr", a.fromaddr, "execaddr", a.execaddr, "playAmount < roundAmount in this GameID ",
			play.GameID)
		return nil, types.ErrInvalidParam
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	receipt, err := a.coinsAccount.ExecFrozen(a.fromaddr, a.execaddr, play.Amount)
	if err != nil {
		clog.Error("blackwhite Play ", "addr", a.fromaddr, "execaddr", a.execaddr, "ExecFrozen amount", play.Amount)
		return nil, err
	}
	logs = append(logs, receipt.Logs...)
	kv = append(kv, receipt.KV...)

	round.Status = gt.BlackwhiteStatusPlay
	addrRes := &gt.AddressResult{
		Addr:       a.fromaddr,
		Amount:     play.Amount,
		HashValues: play.HashValues,
	}
	round.AddrResult = append(round.AddrResult, addrRes)
	round.CurPlayerCount++

	if round.CurPlayerCount >= round.PlayerCount {
		//          
		round.ShowTime = a.blocktime
		round.Status = gt.BlackwhiteStatusShow
		//           
		for _, addr := range round.AddrResult {
			if addr.Addr != round.CreateAddr { //            
				receiptLog := a.GetReceiptLog(&round, addr.Addr)
				logs = append(logs, receiptLog)
			}
		}
		receiptLog := a.GetReceiptLog(&round, round.CreateAddr)
		logs = append(logs, receiptLog)
	} else {
		receiptLog := a.GetReceiptLog(&round, a.fromaddr)
		logs = append(logs, receiptLog)
	}

	key1 := calcMavlRoundKey(round.GameID)
	value1 := types.Encode(&round)
	cfg := a.api.GetConfig()
	if cfg.IsDappFork(a.height, gt.BlackwhiteX, "ForkBlackWhiteV2") {
		//         ，              
		a.db.Set(key1, value1)
	}
	kv = append(kv, &types.KeyValue{Key: key1, Value: value1})

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

func (a *action) Show(show *gt.BlackwhiteShow) (*types.Receipt, error) {
	//   GameID
	value, err := a.db.Get(calcMavlRoundKey(show.GameID))
	if err != nil {
		clog.Error("blackwhite show ", "addr", a.fromaddr, "execaddr", a.execaddr, "get round failed",
			show.GameID, "err", err)
		return nil, err
	}
	var round gt.BlackwhiteRound
	err = types.Decode(value, &round)
	if err != nil {
		clog.Error("blackwhite show ", "addr", a.fromaddr, "execaddr", a.execaddr, "decode round failed",
			show.GameID, "err", err)
		return nil, err
	}
	//       
	if gt.BlackwhiteStatusShow != round.Status {
		err := gt.ErrIncorrectStatus
		clog.Error("blackwhite show ", "addr", a.fromaddr, "round status", round.Status, "status is not match, GameID ",
			show.GameID, "err", err)
		return nil, err
	}

	//            
	bIsExist := false
	index := 0
	for i, addrResult := range round.AddrResult {
		if addrResult.Addr == a.fromaddr {
			bIsExist = true
			index = i
			break
		}
	}
	if !bIsExist {
		err := gt.ErrNoExistAddr
		clog.Error("blackwhite show ", "addr", a.fromaddr, "execaddr", a.execaddr, "this addr is play in GameID",
			show.GameID, "err", err)
		return nil, err
	}
	//    
	if 0 == len(round.AddrResult[index].ShowSecret) {
		round.CurShowCount++ //  show   ，        
	}
	round.Status = gt.BlackwhiteStatusShow
	round.AddrResult[index].ShowSecret = show.Secret

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	if round.CurShowCount >= round.PlayerCount {
		//          
		round.Status = gt.BlackwhiteStatusDone
		receipt, err := a.StatTransfer(&round)
		if err != nil {
			clog.Error("blackwhite show fail", "StatTransfer err", err)
			return nil, err
		}
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
		//           
		for _, addr := range round.AddrResult {
			if addr.Addr != round.CreateAddr { //            
				receiptLog := a.GetReceiptLog(&round, addr.Addr)
				logs = append(logs, receiptLog)
			}
		}
		receiptLog := a.GetReceiptLog(&round, round.CreateAddr)
		logs = append(logs, receiptLog)
	} else {
		receiptLog := a.GetReceiptLog(&round, a.fromaddr)
		logs = append(logs, receiptLog)
	}

	key1 := calcMavlRoundKey(round.GameID)
	value1 := types.Encode(&round)
	cfg := a.api.GetConfig()
	if cfg.IsDappFork(a.height, gt.BlackwhiteX, "ForkBlackWhiteV2") {
		//         ，              
		a.db.Set(key1, value1)
	}
	kv = append(kv, &types.KeyValue{Key: key1, Value: value1})

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

func (a *action) TimeoutDone(done *gt.BlackwhiteTimeoutDone) (*types.Receipt, error) {
	value, err := a.db.Get(calcMavlRoundKey(done.GameID))
	if err != nil {
		clog.Error("blackwhite timeout done ", "addr", a.fromaddr, "execaddr", a.execaddr, "get round failed",
			done.GameID, "err", err)
		return nil, err
	}

	var round gt.BlackwhiteRound
	err = types.Decode(value, &round)
	if err != nil {
		clog.Error("blackwhite timeout done ", "addr", a.fromaddr, "execaddr", a.execaddr, "decode round failed",
			done.GameID, "err", err)
		return nil, err
	}

	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	//       
	if gt.BlackwhiteStatusPlay == round.Status {
		if a.blocktime >= round.Timeout+round.CreateTime {
			//    play    ，           ，      
			for i, addrRes := range round.AddrResult {
				receipt, err := a.coinsAccount.ExecActive(addrRes.Addr, a.execaddr, addrRes.Amount)
				if err != nil {
					//rollback
					for j, addrR := range round.AddrResult {
						if j < i {
							a.coinsAccount.ExecFrozen(addrR.Addr, a.execaddr, addrR.Amount)
						} else {
							break
						}
					}
					clog.Error("blackwhite timeout done", "addr", a.fromaddr, "execaddr", a.execaddr, "execActive all player GameID", done.GameID, "err", err)
					return nil, err
				}
				logs = append(logs, receipt.Logs...)
				kv = append(kv, receipt.KV...)
			}
			//         
			receipt, err := a.coinsAccount.ExecActive(round.CreateAddr, a.execaddr, lockAmount)
			if err != nil {
				for _, addrR := range round.AddrResult {
					a.coinsAccount.ExecFrozen(addrR.Addr, a.execaddr, addrR.Amount)
				}
				clog.Error("blackwhite timeout done", "addr", round.CreateAddr, "execaddr", a.execaddr, "execActive create lockAmount", lockAmount, "err", err)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)

			round.Status = gt.BlackwhiteStatusTimeout

		} else {
			err := gt.ErrNoTimeoutDone
			clog.Error("blackwhite timeout done ", "addr", a.fromaddr, "execaddr", a.execaddr, "is BlackwhiteStatusPlay GameID",
				done.GameID, "err", err)
			return nil, err
		}
	} else if gt.BlackwhiteStatusShow == round.Status {
		if a.blocktime >= showTimeout+round.ShowTime {
			//show    ,        
			round.Status = gt.BlackwhiteStatusDone
			receipt, err := a.StatTransfer(&round)
			if err != nil {
				clog.Error("blackwhite done fail", "StatTransfer err", err)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)

		} else {
			err := gt.ErrNoTimeoutDone
			clog.Error("blackwhite timeout done ", "addr", a.fromaddr, "execaddr", a.execaddr, "is blackwhiteStatusShow GameID",
				done.GameID, "err", err)
			return nil, err
		}
	} else {
		err := gt.ErrIncorrectStatus
		clog.Error("blackwhite timeout done ", "addr", a.fromaddr, "execaddr", a.execaddr, "status is not match GameID",
			done.GameID, "status", round.Status, "err", err)
		return nil, err
	}

	key1 := calcMavlRoundKey(round.GameID)
	value1 := types.Encode(&round)
	cfg := a.api.GetConfig()
	if cfg.IsDappFork(a.height, gt.BlackwhiteX, "ForkBlackWhiteV2") {
		//         ，              
		a.db.Set(key1, value1)
	}
	kv = append(kv, &types.KeyValue{Key: key1, Value: value1})

	//           
	for _, addr := range round.AddrResult {
		if addr.Addr != round.CreateAddr { //            
			receiptLog := a.GetReceiptLog(&round, addr.Addr)
			logs = append(logs, receiptLog)
		}
	}
	receiptLog := a.GetReceiptLog(&round, round.CreateAddr)
	logs = append(logs, receiptLog)

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil

}

func (a *action) StatTransfer(round *gt.BlackwhiteRound) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	winers, loopResults := a.getWinner(round)
	Losers := a.getLoser(round)
	var averAmount int64

	if len(winers) == 0 || len(Losers) == 0 {
		//           
		for i, addrRes := range round.AddrResult {
			receipt, err := a.coinsAccount.ExecActive(addrRes.Addr, a.execaddr, addrRes.Amount)
			if err != nil {
				//rollback
				for j, addrR := range round.AddrResult {
					if j < i {
						a.coinsAccount.ExecFrozen(addrR.Addr, a.execaddr, addrR.Amount)
					} else {
						break
					}
				}
				clog.Error("StatTransfer execActive no winers", "addr", a.fromaddr, "execaddr", a.execaddr, "err", err)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}
	} else {

		var sumAmount int64
		for i, Loser := range Losers {
			//               
			sumAmount += Loser.amount
			receipt, err := a.coinsAccount.ExecTransferFrozen(Loser.addr, blackwhiteAddr, a.execaddr, Loser.amount)
			if err != nil {
				//rollback
				for j, addrR := range Losers {
					if j < i {
						a.coinsAccount.ExecTransfer(blackwhiteAddr, addrR.addr, a.execaddr, addrR.amount)
						a.coinsAccount.ExecFrozen(addrR.addr, a.execaddr, addrR.amount)
					} else {
						break
					}
				}
				clog.Error("StatTransfer all losers to blackwhiteAddr", "addr", a.fromaddr, "execaddr", a.execaddr, "amount", Loser.amount)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}

		winNum := int64(len(winers))
		averAmount = sumAmount / winNum
		//              
		for i, winer := range winers {
			receipt, err := a.coinsAccount.ExecTransfer(blackwhiteAddr, winer.addr, a.execaddr, averAmount)
			if err != nil {
				//rollback
				for j, winer := range winers {
					if j < i {
						a.coinsAccount.ExecTransfer(winer.addr, blackwhiteAddr, a.execaddr, averAmount)
					} else {
						break
					}
				}
				for _, loser := range Losers {
					a.coinsAccount.ExecTransfer(blackwhiteAddr, loser.addr, a.execaddr, loser.amount)
					a.coinsAccount.ExecFrozen(loser.addr, a.execaddr, loser.amount)
				}
				clog.Error("StatTransfer one winer to any other winers fail", "addr", winer, "execaddr", a.execaddr, "err", err)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}

		//        
		for i, winer := range winers {
			receipt, err := a.coinsAccount.ExecActive(winer.addr, a.execaddr, winer.amount)
			if err != nil {
				//rollback
				for j, winer := range winers {
					if j < i {
						a.coinsAccount.ExecFrozen(winer.addr, a.execaddr, winer.amount)
					} else {
						break
					}
				}
				for _, winer := range winers {
					a.coinsAccount.ExecTransfer(winer.addr, blackwhiteAddr, a.execaddr, averAmount)
				}
				for _, loser := range Losers {
					a.coinsAccount.ExecTransfer(blackwhiteAddr, loser.addr, a.execaddr, loser.amount)
					a.coinsAccount.ExecFrozen(loser.addr, a.execaddr, loser.amount)
				}
				clog.Error("StatTransfer ExecActive have winers", "addr", a.fromaddr, "execaddr", a.execaddr, "err", err)
				return nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}
	}

	for _, winer := range winers {
		round.Winner = append(round.Winner, winer.addr)
	}

	//         
	receipt, err := a.coinsAccount.ExecActive(round.CreateAddr, a.execaddr, lockAmount)
	if err != nil {
		// rollback
		if len(winers) == 0 {
			for _, addrR := range round.AddrResult {
				a.coinsAccount.ExecFrozen(addrR.Addr, a.execaddr, addrR.Amount)
			}
		} else {
			for _, winer := range winers {
				a.coinsAccount.ExecFrozen(winer.addr, a.execaddr, winer.amount)
			}
			for _, winer := range winers {
				a.coinsAccount.ExecTransfer(winer.addr, blackwhiteAddr, a.execaddr, averAmount)
			}
			for _, loser := range Losers {
				a.coinsAccount.ExecTransfer(blackwhiteAddr, loser.addr, a.execaddr, loser.amount)
				a.coinsAccount.ExecFrozen(loser.addr, a.execaddr, loser.amount)
			}
		}
		clog.Error("StatTransfer ExecActive create ExecFrozen ", "addr", round.CreateAddr, "execaddr", a.execaddr, "amount", lockAmount)
		return nil, err
	}
	logs = append(logs, receipt.Logs...)
	kv = append(kv, receipt.KV...)

	//           
	logs = append(logs, &types.ReceiptLog{Ty: gt.TyLogBlackwhiteLoopInfo, Log: types.Encode(loopResults)})

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil

}

func (a *action) getWinner(round *gt.BlackwhiteRound) ([]*addrResult, *gt.ReplyLoopResults) {
	var loopRes gt.ReplyLoopResults
	var addresXs []*resultCalc

	loopRes.GameID = round.GetGameID()
	addrRes := round.AddrResult
	loop := int(round.Loop)

	cfg := a.api.GetConfig()
	for _, address := range addrRes {
		if len(address.ShowSecret) > 0 && len(address.HashValues) == loop {
			var isBlack []bool
			//         ：     ForkV25BlackWhite ForkV25BlackWhiteV2         ，  ForkV25BlackWhiteV2     ，
			//   ForkV25BlackWhite     
			if !cfg.IsDappFork(a.height, gt.BlackwhiteX, "ForkBlackWhiteV2") {
				for _, hash := range address.HashValues {
					if bytes.Equal(common.Sha256([]byte(address.ShowSecret+black)), hash) {
						isBlack = append(isBlack, true)
					} else if bytes.Equal(common.Sha256([]byte(address.ShowSecret+white)), hash) {
						isBlack = append(isBlack, false)
					} else {
						isBlack = append(isBlack, false)
					}
				}
			} else {
				for i, hash := range address.HashValues {
					if bytes.Equal(common.Sha256([]byte(strconv.Itoa(i)+address.ShowSecret+black)), hash) {
						isBlack = append(isBlack, true)
					} else if bytes.Equal(common.Sha256([]byte(strconv.Itoa(i)+address.ShowSecret+white)), hash) {
						isBlack = append(isBlack, false)
					} else {
						isBlack = append(isBlack, false)
					}
				}
			}
			addresX := &resultCalc{
				Addr:    address.Addr,
				amount:  address.Amount,
				IsWin:   true,
				IsBlack: isBlack,
			}
			addresXs = append(addresXs, addresX)
		}
	}

	for index := 0; index < loop; index++ {
		blackNum := 0
		whiteNum := 0
		for _, addr := range addresXs {
			if addr.IsWin {
				if addr.IsBlack[index] {
					blackNum++
				} else {
					whiteNum++
				}
			}
		}

		if blackNum < whiteNum {
			for _, addr := range addresXs {
				if addr.IsWin && 0 != blackNum && !addr.IsBlack[index] {
					addr.IsWin = false
				}
			}
		} else if blackNum > whiteNum {
			for _, addr := range addresXs {
				if addr.IsWin && 0 != whiteNum && addr.IsBlack[index] {
					addr.IsWin = false
				}
			}
		}

		winNum := 0
		var perRes gt.PerLoopResult //       
		for _, addr := range addresXs {
			if addr.IsWin {
				winNum++
				perRes.Winers = append(perRes.Winers, addr.Addr)
			} else {
				perRes.Losers = append(perRes.Losers, addr.Addr)
			}
		}

		loopRes.Results = append(loopRes.Results, &perRes)

		if 1 == winNum || 2 == winNum {
			break
		}
	}

	var results []*addrResult
	for _, addr := range addresXs {
		if addr.IsWin {
			result := &addrResult{
				addr:   addr.Addr,
				amount: addr.amount,
			}
			results = append(results, result)
		}
	}

	return results, &loopRes
}

func (a *action) getLoser(round *gt.BlackwhiteRound) []*addrResult {
	addrRes := round.AddrResult
	wins, _ := a.getWinner(round)

	addMap := make(map[string]bool)
	for _, win := range wins {
		addMap[win.addr] = true
	}

	var results []*addrResult
	for _, addr := range addrRes {
		if ok := addMap[addr.Addr]; !ok {
			result := &addrResult{
				addr:   addr.Addr,
				amount: addr.Amount,
			}
			results = append(results, result)
		}
	}

	return results
}

// GetReceiptLog         log
//     ：
// staus == BlackwhiteStatusCreate  (  ，    ）
// status == BlackwhiteStatusPlay (  )
// status == BlackwhiteStatusShow (    )
// status == BlackwhiteStatusTime (      )
// status == BlackwhiteStatusDone (    )
func (a *action) GetReceiptLog(round *gt.BlackwhiteRound, addr string) *types.ReceiptLog {
	log := &types.ReceiptLog{}
	r := &gt.ReceiptBlackwhiteStatus{}
	if round.Status == gt.BlackwhiteStatusCreate {
		log.Ty = gt.TyLogBlackwhiteCreate
		r.PrevStatus = -1
	} else if round.Status == gt.BlackwhiteStatusPlay {
		log.Ty = gt.TyLogBlackwhitePlay
		r.PrevStatus = gt.BlackwhiteStatusCreate
	} else if round.Status == gt.BlackwhiteStatusShow {
		log.Ty = gt.TyLogBlackwhiteShow
		r.PrevStatus = gt.BlackwhiteStatusPlay
	} else if round.Status == gt.BlackwhiteStatusTimeout {
		log.Ty = gt.TyLogBlackwhiteTimeout
		r.PrevStatus = gt.BlackwhiteStatusPlay
	} else if round.Status == gt.BlackwhiteStatusDone {
		log.Ty = gt.TyLogBlackwhiteDone
		r.PrevStatus = gt.BlackwhiteStatusShow
	}

	r.GameID = round.GameID
	r.Status = round.Status
	r.Addr = addr
	r.Index = round.Index

	log.Log = types.Encode(r)
	return log
}

func calcloopNumByPlayer(player int32) int32 {
	a := math.Log2(float64(player))
	a += 0.5
	num := int32(a)
	return num + 1
}
