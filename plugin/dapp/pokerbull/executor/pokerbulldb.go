// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/33cn/chain33/client"

	"strings"

	"github.com/33cn/chain33/account"
	"github.com/33cn/chain33/common"
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/system/dapp"
	"github.com/33cn/chain33/types"
	pkt "github.com/33cn/plugin/plugin/dapp/pokerbull/types"
)

// Action actio 
type Action struct {
	api          client.QueueProtocolAPI
	coinsAccount *account.DB
	db           dbm.KV
	txhash       []byte
	fromaddr     string
	blocktime    int64
	height       int64
	execaddr     string
	localDB      dbm.Lister
	index        int
}

// NewAction action
func NewAction(pb *PokerBull, tx *types.Transaction, index int) *Action {
	hash := tx.Hash()
	fromaddr := tx.From()

	return &Action{pb.GetAPI(), pb.GetCoinsAccount(), pb.GetStateDB(), hash, fromaddr,
		pb.GetBlockTime(), pb.GetHeight(), dapp.ExecAddress(string(tx.Execer)), pb.GetLocalDB(), index}
}

// CheckExecAccountBalance 
func (action *Action) CheckExecAccountBalance(fromAddr string, ToFrozen, ToActive int64) bool {
	//  
	if ToFrozen == 0 {
		ToFrozen = pkt.MinPlayValue * action.api.GetConfig().GetCoinPrecision()
	}

	acc := action.coinsAccount.LoadExecAccount(fromAddr, action.execaddr)
	if acc.GetBalance() >= ToFrozen && acc.GetFrozen() >= ToActive {
		return true
	}
	return false
}

// Key ke 
func Key(id string) (key []byte) {
	key = append(key, []byte("mavl-"+pkt.PokerBullX+"-")...)
	key = append(key, []byte(id)...)
	return key
}

func readGame(db dbm.KV, id string) (*pkt.PokerBull, error) {
	data, err := db.Get(Key(id))
	if err != nil {
		logger.Error("query data have err:", "err", err)
		return nil, err
	}
	var game pkt.PokerBull
	//decode
	err = types.Decode(data, &game)
	if err != nil {
		logger.Error("decode game have err:", "err", err)
		return nil, err
	}
	return &game, nil
}

// Infos gameinfo
func Infos(db dbm.KV, infos *pkt.QueryPBGameInfos) (types.Message, error) {
	var games []*pkt.PokerBull
	for i := 0; i < len(infos.GameIds); i++ {
		id := infos.GameIds[i]
		game, err := readGame(db, id)
		if err != nil {
			return nil, err
		}
		games = append(games, game)
	}
	return &pkt.ReplyPBGameList{Games: games}, nil
}

func getGameListByAddr(db dbm.Lister, addr string, index int64) (types.Message, error) {
	var values [][]byte
	var err error
	if index == 0 {
		values, err = db.List(calcPBGameAddrPrefix(addr), nil, pkt.DefaultCount, pkt.ListDESC)
	} else {
		values, err = db.List(calcPBGameAddrPrefix(addr), calcPBGameAddrKey(addr, index), pkt.DefaultCount, pkt.ListDESC)
	}
	if err != nil {
		return nil, err
	}
	var gameIds []*pkt.PBGameRecord
	for _, value := range values {
		var record pkt.PBGameRecord
		err := types.Decode(value, &record)
		if err != nil {
			continue
		}
		gameIds = append(gameIds, &record)
	}
	if len(gameIds) == 0 {
		return nil, types.ErrNotFound
	}
	return &pkt.PBGameRecords{Records: gameIds}, nil
}

func getGameListByStatus(db dbm.Lister, status int32, index int64) (types.Message, error) {
	var values [][]byte
	var err error
	if index == 0 {
		values, err = db.List(calcPBGameStatusPrefix(status), nil, pkt.DefaultCount, pkt.ListDESC)
	} else {
		values, err = db.List(calcPBGameStatusPrefix(status), calcPBGameStatusKey(status, index), pkt.DefaultCount, pkt.ListDESC)
	}
	if err != nil {
		return nil, err
	}

	var gameIds []*pkt.PBGameIndexRecord
	for _, value := range values {
		var record pkt.PBGameIndexRecord
		err := types.Decode(value, &record)
		if err != nil {
			continue
		}
		gameIds = append(gameIds, &record)
	}

	return &pkt.PBGameIndexRecords{Records: gameIds}, nil
}

func queryGameListByStatusAndPlayer(db dbm.Lister, stat int32, player int32, value int64) ([]string, error) {
	values, err := db.List(calcPBGameStatusAndPlayerPrefix(stat, player, value), nil, pkt.DefaultCount, pkt.ListDESC)
	if err != nil {
		return nil, err
	}

	var gameIds []string
	for _, value := range values {
		var record pkt.PBGameIndexRecord
		err := types.Decode(value, &record)
		if err != nil {
			continue
		}
		gameIds = append(gameIds, record.GetGameId())
	}

	return gameIds, nil
}

func (action *Action) saveGame(game *pkt.PokerBull) (kvset []*types.KeyValue, err error) {
	value := types.Encode(game)
	err = action.db.Set(Key(game.GetGameId()), value)
	if err != nil {
		return nil, err
	}
	kvset = append(kvset, &types.KeyValue{Key: Key(game.GameId), Value: value})
	return kvset, nil
}

func (action *Action) getIndex(game *pkt.PokerBull) int64 {
	return action.height*types.MaxTxsPerBlock + int64(action.index)
}

// GetReceiptLog receip 
func (action *Action) GetReceiptLog(game *pkt.PokerBull) *types.ReceiptLog {
	log := &types.ReceiptLog{}
	r := &pkt.ReceiptPBGame{}
	r.Addr = action.fromaddr
	if game.Status == pkt.PBGameActionStart {
		log.Ty = pkt.TyLogPBGameStart
	} else if game.Status == pkt.PBGameActionContinue {
		log.Ty = pkt.TyLogPBGameContinue
	} else if game.Status == pkt.PBGameActionQuit {
		log.Ty = pkt.TyLogPBGameQuit
	} else if game.Status == pkt.PBGameActionPlay {
		log.Ty = pkt.TyLogPBGamePlay
	}

	r.GameId = game.GameId
	r.Status = game.Status
	r.Index = game.GetIndex()
	r.PrevIndex = game.GetPrevIndex()
	r.PlayerNum = game.PlayerNum
	r.Value = game.Value
	r.IsWaiting = game.IsWaiting
	if !r.IsWaiting {
		for _, v := range game.Players {
			r.Players = append(r.Players, v.Address)
		}
	}
	r.PreStatus = game.PreStatus
	r.Round = game.Round
	log.Log = types.Encode(r)
	return log
}

func (action *Action) readGame(id string) (*pkt.PokerBull, error) {
	data, err := action.db.Get(Key(id))
	if err != nil {
		return nil, err
	}
	var game pkt.PokerBull
	//decode
	err = types.Decode(data, &game)
	if err != nil {
		return nil, err
	}
	return &game, nil
}

func (action *Action) calculate(game *pkt.PokerBull) *pkt.PBResult {
	var handS HandSlice
	for _, player := range game.Players {
		hand := &pkt.PBHand{}
		hand.Cards = Deal(game.Poker, player.TxHash) / 
		hand.Result = Result(hand.Cards)             / 
		hand.Address = player.Address

		/ 
		player.Hands = append(player.Hands, hand)

		/ 
		handS = append(handS, hand)

		/ continu player
		player.Ready = false
	}

	// 
	if !sort.IsSorted(handS) {
		sort.Sort(handS)
	}
	winner := handS[len(handS)-1]

	// 
	result := &pkt.PBResult{}
	result.Winner = winner.Address
	//TODO Dealer 
	//result.Leverage = Leverage(winner)
	result.Hands = make([]*pkt.PBHand, len(handS))
	copy(result.Hands, handS)

	game.Results = append(game.Results, result)
	return result
}

func (action *Action) calculateDealer(game *pkt.PokerBull) *pkt.PBResult {
	var handS HandSlice
	var dealer *pkt.PBHand
	for _, player := range game.Players {
		hand := &pkt.PBHand{}
		hand.Cards = Deal(game.Poker, player.TxHash) / 
		hand.Result = Result(hand.Cards)             / 
		hand.Address = player.Address

		/ 
		player.Hands = append(player.Hands, hand)

		/ 
		handS = append(handS, hand)

		/ continu player
		player.Ready = false

		/ 
		if player.Address == game.DealerAddr {
			dealer = hand
		}
	}

	for _, hand := range handS {
		if hand.Address == game.DealerAddr {
			continue
		}

		if CompareResult(hand, dealer) {
			hand.IsWin = false
		} else {
			hand.IsWin = true
			hand.Leverage = Leverage(hand)
		}
	}

	// 
	result := &pkt.PBResult{}
	result.Dealer = game.DealerAddr
	result.DealerLeverage = Leverage(dealer)
	result.Hands = make([]*pkt.PBHand, len(handS))
	copy(result.Hands, handS)

	game.Results = append(game.Results, result)
	return result
}

func (action *Action) nextDealer(game *pkt.PokerBull) string {
	var flag = -1
	for i, player := range game.Players {
		if player.Address == game.DealerAddr {
			flag = i
		}
	}
	if flag == -1 {
		logger.Error("Get next dealer failed.")
		return game.DealerAddr
	}

	if flag == len(game.Players)-1 {
		return game.Players[0].Address
	}

	return game.Players[flag+1].Address
}

func (action *Action) settleDealerAccount(lastAddress string, game *pkt.PokerBull) ([]*types.ReceiptLog, []*types.KeyValue, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	result := action.calculateDealer(game)
	for _, hand := range result.Hands {
		// 
		if lastAddress != "" && hand.Address != lastAddress {
			receipt, err := action.coinsAccount.ExecActive(hand.Address, action.execaddr, game.GetValue()*PokerbullLeverageMax)
			if err != nil {
				logger.Error("GameSettleDealer.ExecActive", "GameID", game.GetGameId(), "addr", hand.Address,
					"execaddr", action.execaddr, "amount", game.GetValue(), "err", err)
				return nil, nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}

		/ 
		var receipt *types.Receipt
		var err error
		if hand.Address != result.Dealer {
			if hand.IsWin {
				receipt, err = action.coinsAccount.ExecTransfer(result.Dealer, hand.Address, action.execaddr, game.GetValue()*int64(hand.Leverage))
				if err != nil {
					action.coinsAccount.ExecFrozen(hand.Address, action.execaddr, game.GetValue()) // rollback
					logger.Error("GameSettleDealer.ExecTransfer", "GameID", game.GetGameId(), "addr", hand.Address,
						"execaddr", action.execaddr, "amount", game.GetValue()*int64(hand.Leverage), "err", err)
					return nil, nil, err
				}
			} else {
				receipt, err = action.coinsAccount.ExecTransfer(hand.Address, result.Dealer, action.execaddr, game.GetValue()*int64(result.DealerLeverage))
				if err != nil {
					action.coinsAccount.ExecFrozen(hand.Address, action.execaddr, game.GetValue()) // rollback
					logger.Error("GameSettleDealer.ExecTransfer", "GameID", game.GetGameId(), "addr", hand.Address,
						"execaddr", action.execaddr, "amount", game.GetValue()*int64(result.DealerLeverage), "err", err)
					return nil, nil, err
				}
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}
	}
	game.DealerAddr = action.nextDealer(game)

	return logs, kv, nil
}

func (action *Action) settleDefaultAccount(lastAddress string, game *pkt.PokerBull) ([]*types.ReceiptLog, []*types.KeyValue, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue
	result := action.calculate(game)

	for _, player := range game.Players {
		// 
		if lastAddress != "" && player.Address != lastAddress {
			receipt, err := action.coinsAccount.ExecActive(player.GetAddress(), action.execaddr, game.GetValue()*PokerbullLeverageMax)
			if err != nil {
				logger.Error("GameSettleDefault.ExecActive", "GameID", game.GetGameId(), "addr", player.GetAddress(),
					"execaddr", action.execaddr, "amount", game.GetValue()*PokerbullLeverageMax, "err", err)
				return nil, nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}

		/ 
		if player.Address != result.Winner {
			receipt, err := action.coinsAccount.ExecTransfer(player.Address, result.Winner, action.execaddr, game.GetValue() /**int64(result.Leverage)*/) //TODO Dealer 
			if err != nil {
				action.coinsAccount.ExecFrozen(result.Winner, action.execaddr, game.GetValue()) // rollback
				logger.Error("GameSettleDefault.ExecTransfer", "GameID", game.GetGameId(), "addr", result.Winner,
					"execaddr", action.execaddr, "amount", game.GetValue() /**int64(result.Leverage)*/, "err", err) //TODO Dealer 
				return nil, nil, err
			}
			logs = append(logs, receipt.Logs...)
			kv = append(kv, receipt.KV...)
		}
	}

	coinPrecision := action.api.GetConfig().GetCoinPrecision()
	// 
	receipt := action.defaultFeeTransfer(result.Winner, pkt.DeveloperAddress, int64(pkt.DeveloperFee*float64(coinPrecision)), game.GetValue())
	if receipt != nil {
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
	}

	// 
	receipt = action.defaultFeeTransfer(result.Winner, pkt.PlatformAddress, int64(pkt.PlatformFee*float64(coinPrecision)), game.GetValue())
	if receipt != nil {
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
	}

	return logs, kv, nil
}

// 
func (action *Action) defaultFeeTransfer(winner string, feeAddr string, fee int64, value int64) *types.Receipt {
	coinPrecision := action.api.GetConfig().GetCoinPrecision()
	receipt, err := action.coinsAccount.ExecTransfer(winner, feeAddr, action.execaddr, (value/coinPrecision)*fee /**int64(result.Leverage)*/) //TODO Dealer 
	if err != nil {
		action.coinsAccount.ExecFrozen(winner, action.execaddr, (value/coinPrecision)*fee) // rollback
		logger.Error("GameSettleDefault.ExecTransfer", "addr", winner, "execaddr", action.execaddr, "amount",
			(value/coinPrecision)*fee /**int64(result.Leverage)*/, "err", err) //TODO Dealer 
		return nil
	}

	return receipt
}

func (action *Action) settleAccount(lastAddress string, game *pkt.PokerBull) ([]*types.ReceiptLog, []*types.KeyValue, error) {
	if pkt.DefaultStyle == pkt.PlayStyleDealer {
		return action.settleDealerAccount(lastAddress, game)
	}
	return action.settleDefaultAccount(lastAddress, game)
}

func (action *Action) genTxRnd(txhash []byte) (int64, error) {
	randbyte := make([]byte, 6)
	for i := 0; i < 6; i++ {
		randbyte[i] = txhash[i]
	}

	randstr := common.ToHex(randbyte)
	randint, err := strconv.ParseInt(randstr, 0, 64)
	if err != nil {
		return 0, err
	}

	return randint, nil
}

func (action *Action) genTxRnds(txhash []byte, playnum int32) ([]int64, error) {
	rands := make([]int64, playnum)

	for i := 0; i < int(playnum); i++ {
		randbyte := make([]byte, 6)
		for j := 0; j < 6; j++ {
			randbyte[j] = txhash[i*6+j]
		}

		randstr := common.ToHex(randbyte)
		randint, err := strconv.ParseInt(randstr, 0, 64)
		if err != nil {
			return nil, err
		}

		rands[i] = randint
	}

	return rands, nil
}

func (action *Action) checkPlayerAddressExist(pbPlayers []*pkt.PBPlayer) bool {
	for _, player := range pbPlayers {
		if action.fromaddr == player.Address {
			return true
		}
	}

	return false
}

// 
func (action *Action) newGame(gameID string, start *pkt.PBGameStart) (*pkt.PokerBull, error) {
	var game *pkt.PokerBull

	//  
	if start.GetValue() == 0 {
		start.Value = pkt.MinPlayValue * action.api.GetConfig().GetCoinPrecision()
	}

	//TODO 
	if pkt.DefaultStyle == pkt.PlayStyleDealer {
		if !action.CheckExecAccountBalance(action.fromaddr, start.GetValue()*PokerbullLeverageMax*int64(start.PlayerNum-1), 0) {
			logger.Error("GameStart", "GameID", gameID, "addr", action.fromaddr, "execaddr", action.execaddr, "err", types.ErrNoBalance)
			return nil, types.ErrNoBalance
		}
	}

	game = &pkt.PokerBull{
		GameId:      gameID,
		Status:      pkt.PBGameActionStart,
		StartTime:   action.blocktime,
		StartTxHash: gameID,
		Value:       start.GetValue(),
		Poker:       NewPoker(),
		PlayerNum:   start.PlayerNum,
		Index:       action.getIndex(game),
		DealerAddr:  action.fromaddr,
		IsWaiting:   true,
		PreStatus:   0,
		Round:       1,
	}

	Shuffle(game.Poker, action.blocktime) / 
	logger.Info(fmt.Sprintf("Create a new game %s for player %s", game.GameId, action.fromaddr))

	return game, nil
}

// 
func (action *Action) selectGameFromIds(ids []string, value int64) *pkt.PokerBull {
	var gameRet *pkt.PokerBull
	for num := len(ids) - 1; num > -1; num-- {
		id := ids[num]
		game, err := action.readGame(id)
		if err != nil {
			logger.Error("Poker bull game start", "GameID", id, "addr", action.fromaddr, "execaddr", action.execaddr,
				"get game failed", "err", err)
			continue
		}

		//   locald ）
		if int32(len(game.Players)) == game.PlayerNum {
			continue
		}

		/ 
		if action.checkPlayerAddressExist(game.Players) {
			logger.Info(fmt.Sprintf("Player %s already exist in game %s", action.fromaddr, id))
			continue
		}

		/ 
		if value == 0 && game.GetValue() != (pkt.MinPlayValue*action.api.GetConfig().GetCoinPrecision()) {
			if !action.CheckExecAccountBalance(action.fromaddr, game.GetValue(), 0) {
				logger.Error("GameStart", "GameID", id, "addr", action.fromaddr, "execaddr", action.execaddr,
					"err", types.ErrNoBalance)
				continue
			}
		}

		gameRet = game
		logger.Info(fmt.Sprintf("Match a new game %s for player %s", id, action.fromaddr))
		break
	}
	return gameRet
}

//func (action *Action) checkPlayerExistInGame() bool {
//	values, err := action.localDB.List(calcPBGameAddrPrefix(action.fromaddr), nil, pkt.DefaultCount, pkt.ListDESC)
//	if err == types.ErrNotFound {
//		return false
//	}
//
//	var value pkt.PBGameRecord
//	length := len(values)
//	if length != 0 {
//		valueBytes := values[length-1]
//		err := types.Decode(valueBytes, &value)
//		if err == nil && value.Status == pkt.PBGameActionQuit {
//			return false
//		}
//	}
//	return true
//}

// GameStart 
func (action *Action) GameStart(start *pkt.PBGameStart) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	logger.Info(fmt.Sprintf("Pokerbull game match for %s", action.fromaddr))
	// 
	if start.PlayerNum <= 0 || start.Value < 0 {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr,
			"err", fmt.Sprintf("Invalid parameter"))
		return nil, types.ErrInvalidParam
	}

	if start.PlayerNum > pkt.MaxPlayerNum {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr,
			"err", fmt.Sprintf("The maximum player number is %d", pkt.MaxPlayerNum))
		return nil, types.ErrInvalidParam
	}

	gameID := common.ToHex(action.txhash)
	if !action.CheckExecAccountBalance(action.fromaddr, start.GetValue()*PokerbullLeverageMax, 0) {
		logger.Error("GameStart", "GameID", gameID, "addr", action.fromaddr, "execaddr", action.execaddr, "err", types.ErrNoBalance)
		return nil, types.ErrNoBalance
	}

	// 
	//if action.checkPlayerExistInGame() {
	//	logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr, "err", "Address is already in a game")
	//	return nil, fmt.Errorf("Address is already in a game")
	//}

	var game *pkt.PokerBull
	ids, err := queryGameListByStatusAndPlayer(action.localDB, pkt.PBGameActionStart, start.PlayerNum, start.Value)
	if err != nil || len(ids) == 0 {
		if err != types.ErrNotFound {
			return nil, err
		}

		game, err = action.newGame(gameID, start)
		if err != nil {
			return nil, err
		}
	} else {
		game = action.selectGameFromIds(ids, start.GetValue())
		if game == nil {
			//  
			game, err = action.newGame(gameID, start)
			if err != nil {
				return nil, err
			}
		}
	}

	/ txhash
	txrng, err := action.genTxRnd(action.txhash)
	if err != nil {
		return nil, err
	}

	/ 
	game.Players = append(game.Players, &pkt.PBPlayer{
		Address: action.fromaddr,
		TxHash:  txrng,
		Ready:   false,
	})

	//  
	if len(game.Players) == int(game.PlayerNum) {
		logger.Info(fmt.Sprintf("Game starting: %s round: %d", game.GameId, game.Round))
		logsH, kvH, err := action.settleAccount(action.fromaddr, game)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logsH...)
		kv = append(kv, kvH...)

		game.PrevIndex = game.Index
		game.Index = action.getIndex(game)
		game.Status = pkt.PBGameActionContinue // 
		game.PreStatus = pkt.PBGameActionStart
		game.IsWaiting = false
	} else {
		logger.Info(fmt.Sprintf("Game waiting: %s round: %d", game.GameId, game.Round))
		receipt, err := action.coinsAccount.ExecFrozen(action.fromaddr, action.execaddr, start.GetValue()*PokerbullLeverageMax) / , 
		if err != nil {
			logger.Error("GameCreate.ExecFrozen", "GameID", gameID, "addr", action.fromaddr, "execaddr", action.execaddr,
				"amount", start.GetValue(), "err", err.Error())
			return nil, err
		}
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
	}
	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	gamekv, err := action.saveGame(game)
	if err != nil {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr, "err", "save game to db failed")
		return nil, err
	}
	kv = append(kv, gamekv...)

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

func getReadyPlayerNum(players []*pkt.PBPlayer) int {
	var readyC = 0
	for _, player := range players {
		if player.Ready {
			readyC++
		}
	}
	return readyC
}

func getPlayerFromAddress(players []*pkt.PBPlayer, addr string) *pkt.PBPlayer {
	for _, player := range players {
		if player.Address == addr {
			return player
		}
	}
	return nil
}

// GameContinue 
func (action *Action) GameContinue(pbcontinue *pkt.PBGameContinue) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	game, err := action.readGame(pbcontinue.GetGameId())
	if err != nil {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "get game failed", "err", err)
		return nil, err
	}

	if game.Status != pkt.PBGameActionContinue {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "Status error")
		return nil, err
	}
	logger.Info(fmt.Sprintf("Continue pokerbull game %s from %s", game.GameId, action.fromaddr))

	//  
	checkValue := game.GetValue() * PokerbullLeverageMax
	if action.fromaddr == game.DealerAddr {
		checkValue = checkValue * int64(game.PlayerNum-1)
	}
	if !action.CheckExecAccountBalance(action.fromaddr, checkValue, 0) {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "err", types.ErrNoBalance)
		return nil, types.ErrNoBalance
	}

	// 
	pbplayer := getPlayerFromAddress(game.Players, action.fromaddr)
	if pbplayer == nil {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "get game player failed", "err", types.ErrNotFound)
		return nil, types.ErrNotFound
	}
	if pbplayer.Ready {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "player has been ready")
		return nil, fmt.Errorf("player %s has been ready", pbplayer.Address)
	}

	/ txhash
	txrng, err := action.genTxRnd(action.txhash)
	if err != nil {
		return nil, err
	}
	pbplayer.TxHash = txrng
	pbplayer.Ready = true

	if getReadyPlayerNum(game.Players) == int(game.PlayerNum) {
		logger.Info(fmt.Sprintf("Game starting: %s round: %d", game.GameId, game.Round))
		logsH, kvH, err := action.settleAccount(action.fromaddr, game)
		if err != nil {
			return nil, err
		}
		logs = append(logs, logsH...)
		kv = append(kv, kvH...)
		game.PrevIndex = game.Index
		game.Index = action.getIndex(game)
		game.IsWaiting = false
		game.PreStatus = pkt.PBGameActionContinue
	} else {
		logger.Info(fmt.Sprintf("Game waiting: %s round: %d", game.GameId, game.Round))
		// 
		if !game.IsWaiting {
			game.Round++
		}
		receipt, err := action.coinsAccount.ExecFrozen(action.fromaddr, action.execaddr, game.GetValue()*PokerbullLeverageMax) /  
		if err != nil {
			logger.Error("GameCreate.ExecFrozen", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr,
				"execaddr", action.execaddr, "amount", game.GetValue(), "err", err.Error())
			return nil, err
		}
		logs = append(logs, receipt.Logs...)
		kv = append(kv, receipt.KV...)
		game.IsWaiting = true
	}

	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	gamekv, err := action.saveGame(game)
	if err != nil {
		logger.Error("GameContinue", "GameID", pbcontinue.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "err", "save game to db failed")
		return nil, err
	}
	kv = append(kv, gamekv...)

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

// GameQuit 
func (action *Action) GameQuit(pbquit *pkt.PBGameQuit) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	logger.Info(fmt.Sprintf("Quit pokerbull game %s", pbquit.GameId))
	game, err := action.readGame(pbquit.GetGameId())
	if err != nil {
		logger.Error("GameQuit", "GameID", pbquit.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "get game failed", "err", err)
		return nil, err
	}

	if game.Status == pkt.PBGameActionQuit {
		logger.Error("Quit pokerbull game", "GameID", pbquit.GetGameId(), "value", game.Value, "err", "already game over")
		return nil, fmt.Errorf("already game over")
	}

	if !action.checkPlayerAddressExist(game.Players) {
		if action.fromaddr != pkt.PlatformSignAddress {
			logger.Error("GameQuit", "GameID", pbquit.GetGameId(), "addr", action.fromaddr, "execaddr",
				action.execaddr, "err", "permission denied")
			return nil, fmt.Errorf("permission denied")
		}
	}

	//  
	if game.IsWaiting {
		if game.Status == pkt.PBGameActionStart {
			for _, player := range game.Players {
				receipt, err := action.coinsAccount.ExecActive(player.Address, action.execaddr, game.GetValue()*PokerbullLeverageMax)
				if err != nil {
					logger.Error("GameSettleDealer.ExecActive", "GameID", pbquit.GetGameId(), "addr", player.Address,
						"execaddr", action.execaddr, "amount", game.GetValue(), "err", err)
					continue
				}
				logs = append(logs, receipt.Logs...)
				kv = append(kv, receipt.KV...)
			}
		} else if game.Status == pkt.PBGameActionContinue {
			for _, player := range game.Players {
				if !player.Ready {
					continue
				}

				receipt, err := action.coinsAccount.ExecActive(player.Address, action.execaddr, game.GetValue()*PokerbullLeverageMax)
				if err != nil {
					logger.Error("GameSettleDealer.ExecActive", "GameID", pbquit.GetGameId(), "addr", player.Address,
						"execaddr", action.execaddr, "amount", game.GetValue(), "err", err)
					continue
				}
				logs = append(logs, receipt.Logs...)
				kv = append(kv, receipt.KV...)
			}
		}
		game.IsWaiting = false
	}
	game.PreStatus = game.Status
	game.Status = pkt.PBGameActionQuit
	game.PrevIndex = game.Index
	game.Index = action.getIndex(game)
	game.QuitTime = action.blocktime
	game.QuitTxHash = common.ToHex(action.txhash)

	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	gamekv, err := action.saveGame(game)
	if err != nil {
		logger.Error("GameQuit", "GameID", pbquit.GetGameId(), "addr", action.fromaddr, "execaddr",
			action.execaddr, "err", "save game to db failed")
		return nil, err
	}
	kv = append(kv, gamekv...)
	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

// GamePlay 
func (action *Action) GamePlay(pbplay *pkt.PBGamePlay) (*types.Receipt, error) {
	var logs []*types.ReceiptLog
	var kv []*types.KeyValue

	logger.Info(fmt.Sprintf("Play pokerbull game %s, player:%s", pbplay.GameId, strings.Join(pbplay.Address, ",")))
	// 
	if action.fromaddr != pkt.PlatformSignAddress {
		logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
			pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "permission denied")
		return nil, fmt.Errorf("game signing address not support")
	}

	// 
	if pbplay.Round <= 0 || pbplay.Value <= 0 {
		logger.Error("GameStart", "addr", action.fromaddr, "execaddr", action.execaddr,
			"err", fmt.Sprintf("Invalid parameter"))
		return nil, types.ErrInvalidParam
	}

	// 
	if len(pbplay.Address) < pkt.MinPlayerNum || len(pbplay.Address) > pkt.MaxPlayerNum {
		logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
			pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "invalid player number")
		return nil, fmt.Errorf("Invalid player number")
	}

	// 
	for _, addr := range pbplay.Address {
		if !action.CheckExecAccountBalance(addr, pbplay.GetValue()*PokerbullLeverageMax, 0) {
			logger.Error("GamePlay", "addr", addr, "execaddr", action.execaddr, "id", pbplay.GetGameId(), "err", types.ErrNoBalance)
			return nil, types.ErrNoBalance
		}
	}

	//  
	game, _ := action.readGame(pbplay.GetGameId())
	if game != nil {
		if game.Status == pkt.PBGameActionQuit {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "already game over")
			return nil, fmt.Errorf("already game over")
		}

		if game.Round+1 != pbplay.Round {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "game round error")
			return nil, fmt.Errorf("game round error")
		}

		if game.Value != pbplay.Value {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "game value error")
			return nil, fmt.Errorf("game value error")
		}

		// 
		rands, err := action.genTxRnds(action.txhash, game.PlayerNum)
		if err != nil {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", err)
			return nil, err
		}

		// 
		for i, player := range game.Players {
			player.TxHash = rands[i]
		}

		game.Round++
		game.Status = pkt.PBGameActionContinue // 
		game.PreStatus = pkt.PBGameActionContinue
	} else {
		gameNew, err := action.newGame(pbplay.GameId, &pkt.PBGameStart{Value: pbplay.Value, PlayerNum: int32(len(pbplay.Address))})
		if err != nil {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", err)
			return nil, err
		}
		game = gameNew

		// 
		rands, err := action.genTxRnds(action.txhash, game.PlayerNum)
		if err != nil {
			logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
				pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", err)
			return nil, err
		}

		// 
		for i, addr := range pbplay.Address {
			player := &pkt.PBPlayer{
				Address: addr,
				TxHash:  rands[i],
			}
			game.Players = append(game.Players, player)
		}

		game.Status = pkt.PBGameActionQuit // 
		game.PreStatus = pkt.PBGameActionStart
		game.QuitTime = action.blocktime
		game.QuitTxHash = common.ToHex(action.txhash)
	}

	logger.Info(fmt.Sprintf("Game starting: %s round: %d", game.GameId, game.Round))
	logsH, kvH, err := action.settleAccount("", game)
	if err != nil {
		return nil, err
	}
	logs = append(logs, logsH...)
	kv = append(kv, kvH...)
	game.PrevIndex = game.Index
	game.Index = action.getIndex(game)
	game.IsWaiting = false

	receiptLog := action.GetReceiptLog(game)
	logs = append(logs, receiptLog)
	gamekv, err := action.saveGame(game)
	if err != nil {
		logger.Error("Pokerbull game play", "GameID", pbplay.GetGameId(), "round", pbplay.Round, "value",
			pbplay.Value, "players", strings.Join(pbplay.Address, ","), "err", "save game to db failed")
		return nil, err
	}
	kv = append(kv, gamekv...)

	return &types.Receipt{Ty: types.ExecOk, KV: kv, Logs: logs}, nil
}

// HandSlice 
type HandSlice []*pkt.PBHand

func (h HandSlice) Len() int {
	return len(h)
}

func (h HandSlice) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h HandSlice) Less(i, j int) bool {
	if i >= h.Len() || j >= h.Len() {
		logger.Error("length error. slice length:", h.Len(), " compare lenth: ", i, " ", j)
	}

	if h[i] == nil || h[j] == nil {
		logger.Error("nil pointer at ", i, " ", j)
	}

	return CompareResult(h[i], h[j])
}
