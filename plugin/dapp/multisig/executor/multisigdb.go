// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	dbm "github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/types"
	mty "github.com/33cn/plugin/plugin/dapp/multisig/types"
)

// statedb    
func getMultiSigAccFromDb(db dbm.KV, multiSigAddr string) (*mty.MultiSig, error) {

	//   statedb   MultiSigAccAddr     
	value, err := db.Get(calcMultiSigAccountKey(multiSigAddr))
	if err != nil {
		multisiglog.Error("getMultiSigAccFromDb", "MultiSigAccAddr", multiSigAddr, "err", err)
		return nil, err
	}

	//         ErrNotFound
	if len(value) == 0 || err == types.ErrNotFound {
		return nil, types.ErrNotFound
	}

	var multiSigAccount mty.MultiSig
	err = types.Decode(value, &multiSigAccount)
	if err != nil {
		multisiglog.Error("getMultiSigAccFromDb", "MultiSigAccAddr", multiSigAddr, "types.Decode err", err)
		return nil, err
	}
	return &multiSigAccount, nil
}

func setMultiSigAccToDb(db dbm.KV, multiSigAcc *mty.MultiSig) ([]byte, []byte) {
	key := calcMultiSigAccountKey(multiSigAcc.MultiSigAddr)
	value := types.Encode(multiSigAcc)

	//     db ，               
	err := db.Set(key, value)
	if err != nil {
		multisiglog.Error("setMultiSigAccToDb", "multiSigAcc", multiSigAcc, "err", err)
	}
	return key, value
}

//  db           txid       
func getMultiSigAccTxFromDb(db dbm.KV, multiSigAddr string, txid uint64) (*mty.MultiSigTx, error) {

	//   statedb   MultiSigAccTx     
	value, err := db.Get(calcMultiSigAccTxKey(multiSigAddr, txid))
	if err != nil {
		multisiglog.Error("getMultiSigAccTxFromDb", "MultiSigAccAddr", multiSigAddr, "err", err)
		return nil, err
	}
	//         ErrNotFound
	if len(value) == 0 || err == types.ErrNotFound {
		return nil, types.ErrNotFound
	}

	var multiSigAccTx mty.MultiSigTx
	err = types.Decode(value, &multiSigAccTx)
	if err != nil {
		multisiglog.Error("getMultiSigAccTxFromDb", "MultiSigAccAddr", multiSigAddr, "types.Decode err", err)
		return nil, err
	}
	return &multiSigAccTx, nil
}
func setMultiSigAccTxToDb(db dbm.KV, multiSigTx *mty.MultiSigTx) ([]byte, []byte) {
	key := calcMultiSigAccTxKey(multiSigTx.MultiSigAddr, multiSigTx.Txid)
	value := types.Encode(multiSigTx)
	err := db.Set(key, value)
	if err != nil {
		multisiglog.Error("setMultiSigAccTxToDb", "multiSigTx", multiSigTx, "err", err)
	}
	return key, value
}

// localdb    
func getMultiSigAccCountKV(count int64) *types.KeyValue {
	tempcount := &types.Int64{Data: count}
	countbytes := types.Encode(tempcount)
	kv := &types.KeyValue{Key: calcMultiSigAccCountKey(), Value: countbytes}
	return kv
}

//            
func getMultiSigAccCount(db dbm.KVDB) (int64, error) {
	count := types.Int64{}
	value, err := db.Get(calcMultiSigAccCountKey())
	if err != nil && err != types.ErrNotFound {
		return 0, err
	}

	if len(value) == 0 || err == types.ErrNotFound {
		return 0, nil
	}

	err = types.Decode(value, &count)
	if err != nil {
		return 0, err
	}

	return count.Data, nil
}

//            
func setMultiSigAccCount(db dbm.KVDB, count int64) error {
	value := &types.Int64{Data: count}
	valuebytes := types.Encode(value)
	return db.Set(calcMultiSigAccCountKey(), valuebytes)
}

//          
func updateMultiSigAccCount(cachedb dbm.KVDB, isadd bool) (*types.KeyValue, error) {
	count, err := getMultiSigAccCount(cachedb)
	if err != nil {
		return nil, err
	}
	if isadd {
		count++
	} else {
		if count == 0 {
			return nil, mty.ErrAccCountNoMatch
		}
		count--
	}
	err = setMultiSigAccCount(cachedb, count)
	if err != nil {
		multisiglog.Error("updateMultiSigAccCount:setMultiSigAccCount ", "count", count, "err", err)
	}
	//keyvalue
	return getMultiSigAccCountKV(count), nil
}

//          key 
func getMultiSigAccount(db dbm.KVDB, addr string) (*mty.MultiSig, error) {
	multiSigAcc := &mty.MultiSig{}
	value, err := db.Get(calcMultiSigAcc(addr))
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	//       ErrNotFound   ，
	if len(value) == 0 || err == types.ErrNotFound {
		return nil, nil
	}

	err = types.Decode(value, multiSigAcc)
	if err != nil {
		return nil, err
	}

	return multiSigAcc, nil
}

//           db   key ,
func setMultiSigAccount(db dbm.KVDB, multiSig *mty.MultiSig, isadd bool) error {
	valuebytes := types.Encode(multiSig)
	if isadd {
		return db.Set(calcMultiSigAcc(multiSig.MultiSigAddr), valuebytes)
	}
	return db.Set(calcMultiSigAcc(multiSig.MultiSigAddr), nil)
}

//         kv 
func getMultiSigAccountKV(multiSig *mty.MultiSig, isadd bool) *types.KeyValue {
	accountbytes := types.Encode(multiSig)
	var kv *types.KeyValue
	if isadd {
		kv = &types.KeyValue{Key: calcMultiSigAcc(multiSig.MultiSigAddr), Value: accountbytes}
	} else {
		kv = &types.KeyValue{Key: calcMultiSigAcc(multiSig.MultiSigAddr), Value: nil}
	}
	return kv
}

//          
func updateMultiSigAccList(db dbm.KVDB, addr string, index int64, isadd bool) (*types.KeyValue, error) {
	oldaddr, err := getMultiSigAccList(db, index)
	if err != nil {
		return nil, err
	}
	if isadd && oldaddr != "" { //  
		multisiglog.Error("UpdateMultiSigAccList:getMultiSigAccList", "addr", addr, "oldaddr", oldaddr, "index", index, "err", err)
		return nil, mty.ErrAccCountNoMatch
	} else if !isadd && oldaddr == "" { //   
		multisiglog.Error("UpdateMultiSigAccList:getMultiSigAccList", "addr", addr, "index", index, "err", err)
		return nil, mty.ErrAccCountNoMatch
	}

	if isadd { //  
		err = db.Set(calcMultiSigAllAcc(index), []byte(addr))
		if err != nil {
			multisiglog.Error("UpdateMultiSigAccList add", "addr", addr, "index", index, "err", err)
		}
		kv := &types.KeyValue{Key: calcMultiSigAllAcc(index), Value: []byte(addr)}
		return kv, nil
	}
	//   
	err = db.Set(calcMultiSigAllAcc(index), nil)
	if err != nil {
		multisiglog.Error("UpdateMultiSigAccList del", "addr", addr, "index", index, "err", err)
	}
	kv := &types.KeyValue{Key: calcMultiSigAllAcc(index), Value: nil}
	return kv, nil
}

func getMultiSigAccList(db dbm.KVDB, index int64) (string, error) {
	value, err := db.Get(calcMultiSigAllAcc(index))
	if err != nil && err != types.ErrNotFound {
		return "", err
	}

	if len(value) == 0 || err == types.ErrNotFound {
		return "", nil
	}
	return string(value), nil
}

//MultiSigTx:
//             
func getMultiSigTx(db dbm.KVDB, addr string, txid uint64) (*mty.MultiSigTx, error) {
	multiSigTx := &mty.MultiSigTx{}
	value, err := db.Get(calcMultiSigAccTx(addr, txid))
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	//     nil
	if len(value) == 0 || err == types.ErrNotFound {
		return nil, nil
	}

	err = types.Decode(value, multiSigTx)
	if err != nil {
		return nil, err
	}

	return multiSigTx, nil
}

//             db   key ,          
func setMultiSigTx(db dbm.KVDB, multiSigTx *mty.MultiSigTx, isadd bool) error {
	valuebytes := types.Encode(multiSigTx)
	if isadd {
		return db.Set(calcMultiSigAccTx(multiSigTx.MultiSigAddr, multiSigTx.Txid), valuebytes)
	}
	return db.Set(calcMultiSigAccTx(multiSigTx.MultiSigAddr, multiSigTx.Txid), nil)
}

//           kv 
func getMultiSigTxKV(multiSigTx *mty.MultiSigTx, isadd bool) *types.KeyValue {
	accountbytes := types.Encode(multiSigTx)
	var kv *types.KeyValue
	if isadd {
		kv = &types.KeyValue{Key: calcMultiSigAccTx(multiSigTx.MultiSigAddr, multiSigTx.Txid), Value: accountbytes}
	} else {
		kv = &types.KeyValue{Key: calcMultiSigAccTx(multiSigTx.MultiSigAddr, multiSigTx.Txid), Value: nil}
	}
	return kv
}

//        ,         
func updateAddrReciver(cachedb dbm.KVDB, addr, execname, symbol string, amount int64, isadd bool) (*types.KeyValue, error) {
	recv, err := getAddrReciver(cachedb, addr, execname, symbol)
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	if isadd {
		recv += amount
	} else {
		recv -= amount
	}
	err = setAddrReciver(cachedb, addr, execname, symbol, recv)
	if err != nil {
		multisiglog.Error("updateAddrReciver setAddrReciver", "addr", addr, "execname", execname, "symbol", symbol, "err", err)
	}
	//keyvalue
	return getAddrReciverKV(addr, execname, symbol, recv), nil
}

func getAddrReciverKV(addr, execname, symbol string, reciverAmount int64) *types.KeyValue {
	assets := &mty.Assets{
		Execer: execname,
		Symbol: symbol,
	}
	reciver := &mty.AccountAssets{
		MultiSigAddr: addr,
		Assets:       assets,
		Amount:       reciverAmount,
	}
	amountbytes := types.Encode(reciver)
	kv := &types.KeyValue{Key: calcAddrRecvAmountKey(addr, execname, symbol), Value: amountbytes}
	return kv
}

func getAddrReciver(db dbm.KVDB, addr, execname, symbol string) (int64, error) {
	reciver := mty.AccountAssets{}
	addrReciver, err := db.Get(calcAddrRecvAmountKey(addr, execname, symbol))
	if err != nil && err != types.ErrNotFound {
		return 0, err
	}
	if len(addrReciver) == 0 {
		return 0, nil
	}
	err = types.Decode(addrReciver, &reciver)
	if err != nil {
		return 0, err
	}
	return reciver.Amount, nil
}

func setAddrReciver(db dbm.KVDB, addr, execname, symbol string, reciverAmount int64) error {
	kv := getAddrReciverKV(addr, execname, symbol, reciverAmount)
	return db.Set(kv.Key, kv.Value)
}

//MultiSigAccAddress:
//           MultiSigAddress
func getMultiSigAddress(db dbm.KVDB, createAddr string) (*mty.AccAddress, error) {
	address := &mty.AccAddress{}
	value, err := db.Get(calcMultiSigAccCreateAddr(createAddr))
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	//     nil
	if len(value) == 0 || err == types.ErrNotFound {
		return address, nil
	}

	err = types.Decode(value, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

//             
func setMultiSigAddress(db dbm.KVDB, createAddr, multiSigAddr string, isadd bool) *types.KeyValue {
	accAddress, err := getMultiSigAddress(db, createAddr)
	if err != nil {
		return nil
	}

	var found = false
	var foundindex int
	for index, addr := range accAddress.Address {
		if multiSigAddr == addr {
			found = true
			foundindex = index
			break
		}
	}
	if isadd && !found {
		accAddress.Address = append(accAddress.Address, multiSigAddr)
	} else if !isadd && found {
		accAddress.Address = append(accAddress.Address[0:foundindex], accAddress.Address[foundindex+1:]...)
	}

	key := calcMultiSigAccCreateAddr(createAddr)
	value := types.Encode(accAddress)

	err = db.Set(key, value)
	if err != nil {
		multisiglog.Error("setMultiSigAddress", "key", string(key), "err", err)
	}
	return &types.KeyValue{Key: key, Value: value}
}

//           MultiSigAddress
func getMultiSigAccAllAddress(db dbm.KVDB, createAddr string) (*mty.AccAddress, error) {
	address := &mty.AccAddress{}
	value, err := db.Get(calcMultiSigAccCreateAddr(createAddr))
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	//     nil
	if len(value) == 0 || err == types.ErrNotFound {
		return address, nil
	}

	err = types.Decode(value, address)
	if err != nil {
		return nil, err
	}

	return address, nil
}

//                
func getMultiSigAccAllAssets(db dbm.KVDB, addr string) ([][]byte, error) {
	values, err := db.List(calcAddrRecvAmountPrefix(addr), nil, 0, 0)
	if err != nil && err != types.ErrNotFound {
		return nil, err
	}
	return values, nil
}
