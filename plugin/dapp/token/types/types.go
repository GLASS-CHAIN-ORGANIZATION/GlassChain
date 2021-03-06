// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"reflect"

	log "github.com/33cn/chain33/common/log/log15"

	"github.com/33cn/chain33/types"
)

var tokenlog = log.New("module", "execs.token.types")

func init() {
	types.AllowUserExec = append(types.AllowUserExec, []byte(TokenX))
	types.RegFork(TokenX, InitFork)
	types.RegExec(TokenX, InitExecutor)
}

//InitFork ...
func InitFork(cfg *types.Chain33Config) {
	cfg.RegisterDappFork(TokenX, "Enable", 100899)
	cfg.RegisterDappFork(TokenX, ForkTokenBlackListX, 190000)
	cfg.RegisterDappFork(TokenX, ForkBadTokenSymbolX, 184000)
	cfg.RegisterDappFork(TokenX, ForkTokenPriceX, 560000)
	cfg.RegisterDappFork(TokenX, ForkTokenSymbolWithNumberX, 1298600)
	cfg.RegisterDappFork(TokenX, ForkTokenCheckX, 1600000)
}

//InitExecutor ...
func InitExecutor(cfg *types.Chain33Config) {
	types.RegistorExecutor(TokenX, NewType(cfg))
}

// TokenType         
type TokenType struct {
	types.ExecTypeBase
}

// NewType        
func NewType(cfg *types.Chain33Config) *TokenType {
	c := &TokenType{}
	c.SetChild(c)
	c.SetConfig(cfg)
	return c
}

// GetName        
func (t *TokenType) GetName() string {
	return TokenX
}

// GetPayload   token action
func (t *TokenType) GetPayload() types.Message {
	return &TokenAction{}
}

// GetTypeMap   action name  type
func (t *TokenType) GetTypeMap() map[string]int32 {
	return map[string]int32{
		"Transfer":          ActionTransfer,
		"Genesis":           ActionGenesis,
		"Withdraw":          ActionWithdraw,
		"TokenPreCreate":    TokenActionPreCreate,
		"TokenFinishCreate": TokenActionFinishCreate,
		"TokenRevokeCreate": TokenActionRevokeCreate,
		"TransferToExec":    TokenActionTransferToExec,
		"TokenMint":         TokenActionMint,
		"TokenBurn":         TokenActionBurn,
	}
}

// GetLogMap   log       
func (t *TokenType) GetLogMap() map[int64]*types.LogInfo {
	return map[int64]*types.LogInfo{
		TyLogTokenTransfer:        {Ty: reflect.TypeOf(types.ReceiptAccountTransfer{}), Name: "LogTokenTransfer"},
		TyLogTokenDeposit:         {Ty: reflect.TypeOf(types.ReceiptAccountTransfer{}), Name: "LogTokenDeposit"},
		TyLogTokenExecTransfer:    {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecTransfer"},
		TyLogTokenExecWithdraw:    {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecWithdraw"},
		TyLogTokenExecDeposit:     {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecDeposit"},
		TyLogTokenExecFrozen:      {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecFrozen"},
		TyLogTokenExecActive:      {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecActive"},
		TyLogTokenGenesisTransfer: {Ty: reflect.TypeOf(types.ReceiptAccountTransfer{}), Name: "LogTokenGenesisTransfer"},
		TyLogTokenGenesisDeposit:  {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenGenesisDeposit"},
		TyLogPreCreateToken:       {Ty: reflect.TypeOf(ReceiptToken{}), Name: "LogPreCreateToken"},
		TyLogFinishCreateToken:    {Ty: reflect.TypeOf(ReceiptToken{}), Name: "LogFinishCreateToken"},
		TyLogRevokeCreateToken:    {Ty: reflect.TypeOf(ReceiptToken{}), Name: "LogRevokeCreateToken"},
		TyLogTokenMint:            {Ty: reflect.TypeOf(ReceiptTokenAmount{}), Name: "LogMintToken"},
		TyLogTokenBurn:            {Ty: reflect.TypeOf(ReceiptTokenAmount{}), Name: "LogBurnToken"},
	}
}

// RPC_Default_Process rpc     
func (t *TokenType) RPC_Default_Process(action string, msg interface{}) (*types.Transaction, error) {
	var create *types.CreateTx
	if _, ok := msg.(*types.CreateTx); !ok {
		return nil, types.ErrInvalidParam
	}
	create = msg.(*types.CreateTx)
	if !create.IsToken {
		return nil, types.ErrNotSupport
	}
	tx, err := t.AssertCreate(create)
	if err != nil {
		return nil, err
	}
	//to     ,       ???to         to
	cfg := t.GetConfig()
	if !cfg.IsPara() {
		tx.To = create.To
	}
	return tx, err
}

// CreateTx token     
func (t *TokenType) CreateTx(action string, msg json.RawMessage) (*types.Transaction, error) {
	tx, err := t.ExecTypeBase.CreateTx(action, msg)
	if err != nil {
		tokenlog.Error("token CreateTx failed", "err", err, "action", action, "msg", string(msg))
		return nil, err
	}
	cfg := t.GetConfig()
	if !cfg.IsPara() {
		var transfer TokenAction
		err = types.Decode(tx.Payload, &transfer)
		if err != nil {
			tokenlog.Error("token CreateTx failed", "decode payload err", err, "action", action, "msg", string(msg))
			return nil, err
		}
		if action == "Transfer" {
			tx.To = transfer.GetTransfer().To
		} else if action == "Withdraw" {
			tx.To = transfer.GetWithdraw().To
		} else if action == "TransferToExec" {
			tx.To = transfer.GetTransferToExec().To
		}
	}
	return tx, nil
}
