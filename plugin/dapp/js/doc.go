// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package js

/*
Java Script VM contract

cli    
Available Commands:
  call        call java script contract
  create      create java script contract
  query       query java script contract

cli jsvm create
Flags:
  -c, --code string   path of js file,it must always be in utf-8.
  -h, --help          help for create
  -n, --name string   contract name

cli jsvm call
Flags:
  -a, --args string       json str of args
  -f, --funcname string   java script contract funcname
  -h, --help              help for call
  -n, --name string       java script contract name

cli jsvm query
Flags:
  -a, --args string       json str of args
  -f, --funcname string   java script contract funcname
  -h, --help              help for query
  -n, --name string       java script contract name

    ：
   ：    
cli seed save -p heyubin1234 -s "voice leisure mechanic tape cluster grunt receive joke nurse between monkey lunch save useful cruise"

cli wallet unlock -p heyubin1234

cli account import_key  -l miner -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944

   ：    test js  ，      test.js    （   utf-8  ）
cli send jsvm create -c "../plugin/dapp/js/executor/test.js"  -n test -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt

   ：  test   hello  
cli send jsvm  call -a "{\"hello\": \"world\"}" -f hello -n test -k 14KEKbYtKKQm4wMthSK9J4La4nAiidGozt

   ：query test  hello  
cli jsvm  query -a "{\"hello\": \"world\"}" -f hello -n test
*/
