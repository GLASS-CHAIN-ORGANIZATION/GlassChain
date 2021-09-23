// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/types"
	"github.com/33cn/plugin/plugin/dapp/evm/executor/vm/common"
	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/sha3"
)

// ValidateSignatureValues           
func ValidateSignatureValues(r, s *big.Int) bool {
	if r.Cmp(common.Big1) < 0 || s.Cmp(common.Big1) < 0 {
		return false
	}
	return true
}

// Ecrecover          ，          
func Ecrecover(hash, sig []byte) ([]byte, error) {
	pub, err := SigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	bytes := (*btcec.PublicKey)(pub).SerializeUncompressed()
	return bytes, err
}

// SigToPub           
func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	btcsig := make([]byte, 65)
	btcsig[0] = sig[64] + 27
	copy(btcsig[1:], sig)

	pub, _, err := btcec.RecoverCompact(btcec.S256(), btcsig, hash)
	return (*ecdsa.PublicKey)(pub), err
}

// RandomContractAddress           ，           
func RandomContractAddress() *common.Address {
	c, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		return nil
	}
	key, err := c.GenKey()
	if err != nil {
		return nil
	}

	acc := address.PubKeyToAddress(key.PubKey().Bytes())
	ret := common.StringToAddress(address.ExecAddress(acc.String()))
	return ret
}

// Keccak256       Keccak256   
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

// Keccak256Hash calculates and returns the Keccak256 hash of the input data,
// converting it to an internal Hash data structure.
func Keccak256Hash(data ...[]byte) (h common.Hash) {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	d.Sum(h[:0])
	return h
}
