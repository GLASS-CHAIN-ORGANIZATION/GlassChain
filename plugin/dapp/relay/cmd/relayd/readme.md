# relayd     

##   

*          ，        ，    (supervisor)，   ，   。
*           ，      ，       。
* relayd   chain33  GRPC  ，  btcd    websockets，  HTTPS，      。
*       GRPC  HTTP   ，      。
*     btc      ，      chain33     ，          。
*     ，  chain33 btc  TCP  。

##       

*       ，   blockinfo.org    。
* SPV       。
*          。

## SPV
0.                block header   ；
1.       hash tx_hash；
2.       tx_hash     ，  block header            ；
3.         merkle tree   hash ；
4.     hash   merkle_root_hash；
5.       block header  merkle_root_hash  ，       。
6.    block header     ，              。

##     

*        hash，       
*   hash       。
*      address（  ），             。
*       。
*         timestamp(   )，    10m + hash.time <= hash.time <= hash.time + 2h    。
*          。 SPV   。

## relay executor      

*     ，       
* relay             ，        ：pending、locking、confirming、canceled、timeout、finished   。
*     :
    * pending:        。
    * locking：       ，     ，        。
    * confirming:        ，               。
    * finished:     ，        。
    * canceled:        ，    timeout、pending          。
    * timeout: pending          ，    ；lock            ，    ； conforming    ，          。
* unlock：        ，     locking   ，             ，     ，      pending  。
*      :
```


                      +--------------|---------------|
                    pending ---+ locking ---+ confirming ---+ finished
                    |                |               |
                    |                |               |
                    |----------------|---------------|------+ cancel
                    |                |               |
                    |                +               |
                     -----------+ timeout +----------


  ：
       ：pending
       ：locking、confirming
       ：finished、cancel、timeout.

    unlock       ，unlock            pending，      canceled  
     unlock  ,             pending

locking  unlock    12     12*6 BTC block    BTC         
confirming   ，      4*12    4*12*6 BTC block

finish     BTC n    ，n   6，    BTC     

  ：      ，            BTC Header  ，   testNet genesis addr,         

  ：
              。
```

##    （   ）

*     ，       ，    。
*            ，                        ，        ？，    relayd   ？

##   

*     ，     hash，      ？  ：    hash，relayd        ，    ，  chain33   unlock  。
*        ，   ？  ：  relayd         ，      。
*     ：  chain33        ?     ，chain33   ，  executor(relay)   ，   hash      ， ：hash  。          ，     。
*     ：             ，            。        hash    ？    ，    ？       ？
*     ：hash                ， ，      ，chain33         btc   ，     ，   ，btc       chain33    。
                 。
