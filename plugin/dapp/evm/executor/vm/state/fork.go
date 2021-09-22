// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package state

import (
	"fmt"

	"github.com/33cn/chain33/types"
	"github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	evmtypes "github.com/33cn/plugin/plugin/dapp/evm/types"
)

type BlockData struct {
	blockHeight int64
	testnet     bool
	txs map[string]*TxData
}

type TxData struct {
	KV map[string][]byte
	Logs map[string][]byte
}

var forkData map[int64]*BlockData

func newBlockData(blockHeight int64, testnet bool) *BlockData {
	data := BlockData{blockHeight: blockHeight, testnet: testnet}
	data.txs = make(map[string]*TxData)
	return &data
}

func (block *BlockData) newTxData(txHash string) *TxData {
	data := TxData{}
	data.KV = make(map[string][]byte)
	data.Logs = make(map[string][]byte)
	block.txs[txHash] = &data
	return &data
}

func makeLogReceiptKey(logType int32, logIndex int) string {
	return fmt.Sprintf("%v_%v", logType, logIndex)
}

func InitForkData() {
	forkData = make(map[int64]*BlockData)

	data := newBlockData(556151, true)
	txData := data.newTxData("0xc79c9113a71c0a4244e20f0780e7c13552f40ee30b05998a38edb08fe617aaa5")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x1a2039c5e33bc3f4dca79c611d979a87c24aebc7e7f056cdc8845dae0a2c418fd35a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 1)] = common.FromHex("0x1a2039c5e33bc3f4dca79c611d979a87c24aebc7e7f056cdc8845dae0a2c418fd35a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb")
	forkData[556151] = data

	data = newBlockData(556255, true)
	txData = data.newTxData("0xa2e3a06322ce7561ec3c7e442dbc0a6b12618f661c80d16a1af0ffda3e8c2dd8")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x1a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 3)] = common.FromHex("0x1a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c80022660a4230783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303112200000000000000000000000000000000000000000000000000000000000002710")
	forkData[556255] = data

	data = newBlockData(556294, true)
	txData = data.newTxData("0x206c26b2a00751b13df4f44924f46bd353d4d2b48595687a29c9c9c1c34f6d3f")
	txData.KV["0x6d61766c2d65766d2d73746174653a20314761634d39335374725a76654d72506a58446f7a355478616a4b61394c4d354847"] = common.FromHex("0x10011a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	txData.Logs[makeLogReceiptKey(evmtypes.TyLogContractState, 1)] = common.FromHex("0x10011a208515259f99f2e464df971c367832cee5a554733e72ff988c034534f9a762b05a22660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030311220000000000000000000000000000000000000000000000000000000000000271022660a423078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030301220000000000000000000000000e383514ca69827156779875ae63c5a3b87d7b7eb22660a42307864623638613236356461623130373965663837633133353662343737353831616537353830616563643962353132323162663465653661613331383165306632122000000000000000000000000000000000000000000000000000000004a817c800")
	forkData[556294] = data
}

func ProcessFork(cfg *types.Chain33Config, blockHeight int64, txHash []byte, receipt *types.Receipt) {
	if cfg.IsLocal() {
		return
	}

	if !cfg.IsTestNet() {
		return
	}

	if block, ok := forkData[blockHeight]; ok {
		strHash := common.Bytes2Hex(txHash)

		if tx, ok := block.txs[strHash]; ok {

			for i, v := range receipt.KV {
				if val, ok := tx.KV[common.Bytes2Hex(v.Key)]; ok {
					receipt.KV[i].Value = val
				}
			}

			for i, v := range receipt.Logs {
				if val, ok := tx.Logs[makeLogReceiptKey(v.Ty, i)]; ok {
					receipt.Logs[i].Log = val
				}
			}
		}
	}
}
