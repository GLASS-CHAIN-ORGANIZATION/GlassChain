// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wallet

import (
	privacy "github.com/33cn/plugin/plugin/dapp/privacy/crypto"
	privacytypes "github.com/33cn/plugin/plugin/dapp/privacy/types"
)

type addrAndprivacy struct {
	PrivacyKeyPair *privacy.Privacy
	Addr           *string
}

// buildInputInfo 
type buildInputInfo struct {
	assetExec   string
	assetSymbol string
	sender      string
	amount      int64
	mixcount    int32
}

// txOutputInfo UTX 
type txOutputInfo struct {
	amount           int64
	utxoGlobalIndex  *privacytypes.UTXOGlobalIndex
	txPublicKeyR     []byte
	onetimePublicKey []byte
}

type walletUTXO struct {
	height  int64
	outinfo *txOutputInfo
}

type walletUTXOs struct {
	utxos []*walletUTXO
}
