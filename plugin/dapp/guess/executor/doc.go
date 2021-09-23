// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
     ：  

 、    ：
                ，             .
         ，                ，         ，         .
           ，                             。

 、       
1、              ，        ，  ，    、     ，       ，             ，  10BTY，      100 ，     2      (             )。
2、         ，                             ，       。

 、    
1、                  ，                    。
2、                        ，                       ，                   1/2，              1/2。
3、                     （                 ，      、     ），          ，             。
4、      ，                         ，         ，           。

 、       
1、    ，                        。
2、          
        （  ：         ）、
     （  ：A    B      ）、
       （  BTY）、
       （  5BTY）、
            （  100 BTY）、
         （          ，  ：     500000）
         （          ，  ：     1000000）
           （  20000 BTY）
3、        ，             ，          。            ，        。
4、         ，       。                   。
5、            ，                  。
6、（１）                 ，  ５‰      ，５‰     ，（２）              ，             ，                                       （  Ａ    ，   Ａ    10000 BTY，     Ａ  100BTY，      1/100）。
7、                     （                 ），          ，             。
8、      ，         ，                  ，              。
8、    ：
   start(   )->bet(  )->stopbet(   )->publish(   )
   start(   )->bet(  )->stopbet(   )->abort(   )
   start(   )->bet(  )->abort(   )
   start(   )->abort(   )
   start(   )->bet(  )->stopbet(   )->timeout->abort(   )
   start(   )->bet(  )->timeout->abort(   )
   start(   )->timeout->abort(   )
   start(   )->stopbet(   )->publish(   )
   start(   )->bet(  )->publish(   )
   start(   )->publish(   )

     ：            ，           ，             。
*/
