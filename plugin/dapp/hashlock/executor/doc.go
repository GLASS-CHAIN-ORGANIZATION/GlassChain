// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
  hashlocking        ：lock ，unlock ，send

  lock  ：         
   coins         priv   
            header，    Version ，ParentHash，TxHash ，StateHash，Height，BlockTime，TxCount，Hash，Signature
           label      string  
 genaddress()        addrto    privkey，
privkey  type PrivKey interface {
	Bytes() []byte
	Sign(msg []byte) Signature
	PubKey() PubKey
	Equals(PrivKey) bool
}    
  addrto    PubKey()        ，
addrto  type Address struct {
	Version  byte
	Hash160  [20]byte // For a stealth address: it's HASH160
	Checksum []byte   // Unused for a stealth address
	Pubkey   []byte   // Unused for a stealth address
	Enc58str string
}     

   privkey    label         wallet ，    Account  ，
type Account struct {
	Currency int32  `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	Balance  int64  `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	Frozen   int64  `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	Addr     string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}
              ，     addrto_b

   sendtoaddress（） priv                       1e10  

sendtoaddress（）。。。。。            amount   fee         tx ，    ，   nil
    addrto      ，accs WalletAccounts    ，     Wallets WalletAccount        ，Wallets  
Acc Label      ，Acc Account       
type Account struct {
	Currency int32  `protobuf:"varint,1,opt,name=currency" json:"currency,omitempty"`
	Balance  int64  `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
	Frozen   int64  `protobuf:"varint,3,opt,name=frozen" json:"frozen,omitempty"`
	Addr     string `protobuf:"bytes,4,opt,name=addr" json:"addr,omitempty"`
}

  sendtolock（）      ，   hashlock        ，                      coins,    
showAccount（）  ，      ，      lock.

  ：   lock       ：
                      a  ，          
   a hashlock       ，       ，      
           lock


*/
