// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multisig

/*
          ：
           ，   owner    
      ：
	owner add/del/modify/replace
	         
	       
             ，              。

          ：                  ，     owner     。owner              ，            。         

            ：   ，to           ，from            ；
					    ，from           ，to            ；           

cli         ：account      ，owner      tx     
cli multisig
Available Commands:
  account     multisig account
  owner       multisig owner
  tx          multisig tx

cli multisig  account
Available Commands:
  address     get multisig account address
  assets      get assets of multisig account
  count       get multisig account count
  create      Create a multisig account transaction
  creator     get all multisig accounts created by the address
  dailylimit  Create a modify assets dailylimit transaction
  info        get multisig account info
  owner       get multisig accounts by the owner
  unspent     get assets unspent today amount
  weight      Create a modify required weight transaction

cli multisig  owner
Available Commands:
  add         Create a add owner  transaction
  del         Create a del owner transaction
  modify      Create a modify owner weight transaction
  replace     Create a replace owner transaction

cli multisig  tx
Available Commands:
  confirm          Create a confirm transaction
  confirmed_weight get the weight of the transaction confirmed.
  count            get multisig tx count
  info             get multisig account tx info
  transfer_in      Create a transfer to multisig account transaction
  transfer_out     Create a transfer from multisig account transaction
  txids            get multisig txids


      ：
cli seed save -p heyubin1234 -s "voice leisure mechanic tape cluster grunt receive joke nurse between monkey lunch save useful cruise"

cli wallet unlock -p heyubin1234

cli account import_key  -l miner -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944

cli account create -l heyubin
cli account create -l heyubin1
cli account create -l heyubin2
cli account create -l heyubin3
cli account create -l heyubin4
cli account create -l heyubin5
cli account create -l heyubin6
cli account create -l heyubin7
cli account create -l heyubin8


cli send coins transfer -a 100 -n test  -t 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t 1Kkgztjcni3xKw95y2VZHwPpsSHDEH5sXF -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t 1N8LP5gBufZXCEdf3hyViDhWFqeB7WPGdv -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t 1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt

cli send coins transfer -a 100 -n test  -t "1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj" -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t "17a5NQTf9M2Dz9qBS8KiQ8VUg8qhoYeQbA" -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t "1DeGvSFX8HAFsuHxhaVkLX56Ke3FzFbdct" -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t "166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf" -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt
cli send coins transfer -a 100 -n test  -t "1KHwX7ZadNeQDjBGpnweb4k2dqj2CWtAYo" -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt



   ：1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd           ，owner：1Kkgztjcni3xKw95y2VZHwPpsSHDEH5sXF  1N8LP5gBufZXCEdf3hyViDhWFqeB7WPGdv
//    
cli send multisig account create -d 10 -e coins -s BTY -a "1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK-1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj" -w "20-10" -r 15 -k 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd

//         
cli multisig account count

//    index      
cli multisig account address -e 0 -s 0

//    addr      
cli multisig account info -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"


   ， multisig      1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd

cli exec addr -e "multisig"
14uBEP6LSHKdFvy97pTYRPVPAqij6bteee

//   
cli send coins transfer -a 50 -n test  -t 14uBEP6LSHKdFvy97pTYRPVPAqij6bteee -k 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd

cli account balance -a 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd


   ：               1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd   --》 "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"

cli send multisig tx transfer_in -a 40 -e coins -s BTY  -t "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -n test -k 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd


//  1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd   40    
cli account balance -a 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd


//        40   
cli account balance -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"

cli multisig  account assets  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"


   ：           "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"  --》1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj  owner:1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK  
cli send multisig  tx transfer_out  -a 11 -e coins -s BTY -f "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -t 1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj -n test -k "1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK"


      
cli multisig account info -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"

cli multisig  account assets  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"

           
cli multisig  account assets  -a 1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj


//       
cli multisig  tx count  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47"

//     txid
cli multisig   tx txids  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -s 0 -e 0

//       
cli multisig  tx info  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -i 0


//owner "1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj,"   5   1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj    owner:1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj  
cli send multisig  tx transfer_out  -a 5 -e coins -s BTY -f "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -t 1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj -n test -k "1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj"


   ：  add/del owner        owner     "1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK",  add 1KHwX7ZadNeQDjBGpnweb4k2dqj2CWtAYo

cli send multisig owner add  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -o 1KHwX7ZadNeQDjBGpnweb4k2dqj2CWtAYo -w 5 -k  "1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK"


  owner   
cli multisig  account info -a 13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47

cli multisig  tx info  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -i 0

//del owner
cli send multisig  owner del  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -o "1KHwX7ZadNeQDjBGpnweb4k2dqj2CWtAYo"  -k 1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK
// modify  dailylimit
cli send multisig  account dailylimit -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -e coins -s BTY -d 12 -k 1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK
// modify weight
cli send multisig  account weight -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -w 16 -k 1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK
//replace owner
cli send multisig  owner replace  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -n 166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf -o 1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj -k  1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK
// modify owner
cli send multisig  owner modify  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -o "166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf" -w 11 -k 1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK

//                 
cli multisig account creator -a 1DkrXbz2bK6XMpY4v9z2YUnhwWTXT6V5jd

//                 
cli multisig  account unspent  -a 13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47 -e coins -s BTY

   ：          
//      ，owner：166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf
cli send multisig  tx transfer_out  -a 10 -e coins -s BTY -f "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -t 1LDGrokrZjo1HtSmSnw8ef3oy5Vm1nctbj -n test -k "166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf"


//          
cli send   multisig tx confirm  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -i 8 -c f  -k 166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf


//      
cli send multisig tx confirm  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -i 8 -k "166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf"

cli send multisig tx confirm  -a "13q53Ga1kquDCqx7EWF8FU94tLUK18Zd47" -i 8 -k "1C5xK2ytuoFqxmVGMcyz9XFKFWcDA8T3rK"

//   owner           ，                        
cli  multisig account owner -a 166po3ghRbRu53hu8jBBQzddp7kUJ9Ynyf
*/
