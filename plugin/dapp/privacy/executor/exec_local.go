// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"encoding/hex"

	"github.com/33cn/chain33/common"
	"github.com/33cn/chain33/types"
	ty "github.com/33cn/plugin/plugin/dapp/privacy/types"
)

func (p *privacy) execLocal(receiptData *types.ReceiptData, tx *types.Transaction, index int) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	txhashstr := hex.EncodeToString(tx.Hash())
	localDB := p.GetLocalDB()
	for _, item := range receiptData.Logs {
		if item.Ty != ty.TyLogPrivacyOutput {
			continue
		}
		var receiptPrivacyOutput ty.ReceiptPrivacyOutput
		err := types.Decode(item.Log, &receiptPrivacyOutput)
		if err != nil {
			privacylog.Error("PrivacyTrading ExecLocal", "txhash", txhashstr, "Decode item.Log error ", err)
			panic(err) /  
		}

		assetExec := receiptPrivacyOutput.GetAssetExec()
		assetSymbol := receiptPrivacyOutput.GetAssetSymbol()
		txhashInByte := tx.Hash()
		txhash := common.ToHex(txhashInByte)
		for outputIndex, keyOutput := range receiptPrivacyOutput.Keyoutput {
			//kv1 UTXO toke   txhas UTXO
			key := CalcPrivacyUTXOkeyHeight(assetExec, assetSymbol, keyOutput.Amount, p.GetHeight(), txhash, index, outputIndex)
			localUTXOItem := &ty.LocalUTXOItem{
				Height:        p.GetHeight(),
				Txindex:       int32(index),
				Outindex:      int32(outputIndex),
				Txhash:        txhashInByte,
				Onetimepubkey: keyOutput.Onetimepubkey,
			}
			value := types.Encode(localUTXOItem)
			kv := &types.KeyValue{Key: key, Value: value}
			dbSet.KV = append(dbSet.KV, kv)

			//kv2 k  UTXO
			var amountTypes ty.AmountsOfUTXO
			key2 := CalcprivacyKeyTokenAmountType(assetExec, assetSymbol)
			value2, err := localDB.Get(key2)
			/ toke 
			if err == nil && value2 != nil {
				err := types.Decode(value2, &amountTypes)
				if err == nil {
					/  
					amount, ok := amountTypes.AmountMap[keyOutput.Amount]
					if !ok {
						amountTypes.AmountMap[keyOutput.Amount] = 1
					} else {
						//todo 
						amountTypes.AmountMap[keyOutput.Amount] = amount + 1
					}
					kv := &types.KeyValue{Key: key2, Value: types.Encode(&amountTypes)}
					dbSet.KV = append(dbSet.KV, kv)
					/ quer  amou kv 
					localDB.Set(key2, types.Encode(&amountTypes))
				} else {
					privacylog.Error("PrivacyTrading ExecLocal", "txhash", txhashstr, "value2 Decode error ", err)
					panic(err)
				}
			} else {
				/ toke 
				amountTypes.AmountMap = make(map[int64]int64)
				amountTypes.AmountMap[keyOutput.Amount] = 1
				kv := &types.KeyValue{Key: key2, Value: types.Encode(&amountTypes)}
				dbSet.KV = append(dbSet.KV, kv)
				localDB.Set(key2, types.Encode(&amountTypes))
			}

			//kv3 toke 
			assetKey := calcExecLocalAssetKey(assetExec, assetSymbol)
			var tokenNames ty.TokenNamesOfUTXO
			key3 := CalcprivacyKeyTokenTypes()
			value3, err := localDB.Get(key3)
			if err == nil && len(value3) != 0 {
				err := types.Decode(value3, &tokenNames)
				if err == nil {
					if _, ok := tokenNames.TokensMap[assetKey]; !ok {
						tokenNames.TokensMap[assetKey] = txhash
						kv := &types.KeyValue{Key: key3, Value: types.Encode(&tokenNames)}
						dbSet.KV = append(dbSet.KV, kv)
						localDB.Set(key3, types.Encode(&tokenNames))
					}
				}
			} else {
				tokenNames.TokensMap = make(map[string]string)
				tokenNames.TokensMap[assetKey] = txhash
				kv := &types.KeyValue{Key: key3, Value: types.Encode(&tokenNames)}
				dbSet.KV = append(dbSet.KV, kv)
				localDB.Set(key3, types.Encode(&tokenNames))
			}
		}
	}
	return dbSet, nil
}

// ExecLocal_Public2Privacy local execute public to privacy transaction
func (p *privacy) ExecLocal_Public2Privacy(payload *ty.Public2Privacy, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return p.execLocal(receiptData, tx, index)
}

// ExecLocal_Privacy2Privacy local execute privacy to privacy transaction
func (p *privacy) ExecLocal_Privacy2Privacy(payload *ty.Privacy2Privacy, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return p.execLocal(receiptData, tx, index)
}

// ExecLocal_Privacy2Public local execute privacy to public trasaction
func (p *privacy) ExecLocal_Privacy2Public(payload *ty.Privacy2Public, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	return p.execLocal(receiptData, tx, index)
}