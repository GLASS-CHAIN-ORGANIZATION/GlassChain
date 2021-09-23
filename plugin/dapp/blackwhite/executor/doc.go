// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
         ：   

  ：

1. gen(s)      : s = hash(hash(privkey)+nonce)
2.   ：    ，    ，    ，        ，  gameid
3.   :  gameid，      hash(s+ )，lock    (match)，                       ，     24*60*60s,         
4.     :    s，     5  ，                ，     ；    5          ，       

    ：       20 BTY


status: create -> play -> show -> done(timeout done)


//      
//1.       ，         （      ）
//2.             (        )
*/
