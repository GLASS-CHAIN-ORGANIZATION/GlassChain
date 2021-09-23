// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
         ：   ，  ， 

  ：

1. gen(s)      : s = hash(hash(privkey)+nonce)
2.   ： hash(s), hash(s+  ), Lock    * 2 -> gameid (create)
3.   :  gameid，      ，lock    (match)
4.   :    s，     			 (close)
5.   ：      (   3 *   ) （close）
6.   ：      (cancel)
    ：       100 BTY，         BTY     

          ，           

status: Create 1 -> Match 2 -> Cancel 3 -> Close 4


//      
//1.       ，         （      ）
//2.             (        )
*/
//game      ：
// staus ==  1 (  ，      ）
// status == 2 (  ，  )
// status == 3 (  )
// status == 4 (Close   )
