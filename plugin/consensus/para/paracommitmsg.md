# paracross        ，         

##      
 1.            title      
 1.         ，                   ，         。（      ，            ）
 1.          ，           ，   20 ，10  。          ，        
 1.         ExecOk,        ok ，      
 1.      ExecPack，     ，                ，                 ，    LogErr，    ，     
 1.                 ，          tx，                           ，       。

##     
 1.   tick（16s）     grpc      height
    *            ，   -1，   sync  ，      
    *                ，           ，      ，       ，        ，     ，        ，   
               ，     ,               ,  sync  ，    

##         （    ）
   1.     ，          ，               ，                  
   1.    ，             ，                ，          。

##   ，    
 1. delete           ，       ，              
 1.             finish  ，      finish      ，              finish  
 1.           ，               ，       ，                

##     
 1.      block，          block       ，       pack，     ，     。

##   
 1.         wallet    ，            。          ，                ，         2s    ，
          ，    ，     。

##     
 1. grpc    ， 1s    ，       ，    tx        ，   tx       tx，    mempool      ，tx nonce  
 1.      ，    
 1.        commit msg，    ，     commit msg    ，       ，        sending  
 1.        ，                 ，      ，         
 1.      ，               ，  debug
 1.     ，           ，         ，            ，           ，     ，       2/3  ，  
         ，           ，       ，  debug
 1.               ，          ，      ，       ，    .           ，           ，
          ，            ，               
 1.       ，    2/3           ，              ，              ，                  
          ，         ，           ，    done，             commit    done，        ，  
                  ，             
 1.        ，               ，        ，         ，                 ，           ，
              20tx    ，    10 height   ，       -1，       ，        10       ，       
     0          ，         -1          

##        
 1.             ，             ，      ，      ，    ，             ，  2       
    sending     ，           ，    sending tx        ，                 。           
                 ，              。
 1.                       ，             ，      ，          ，          ，      
        ，              ，     ，     ，                 
 
##     
 1.             docker    ，          120s  ，       8     
 1. 6   ，4      ，        4，    3，      
 1. 6   ，4      ,        4，  3，      
 1. 6   ，4      ，2    ，    ，          ，     
 1. 6   ，4      ，2    ，    ，          ，     ，       ，    ，             ，    
 1. 6   ，4      ，     ，    10     ，            ，            
 1. 6   ，4      ，           ，              ，                       
 1. 6   ，4      ，    ，  a         ，b     ，    ，  b ，   a  b ，   a     ，b        ，
           ，       ， a    ，b       ，      ，b         

##      
 1.          tx，           ，      -1，              ，        ，     
 1.        ，     N，commit msg    >N       ，             ，       
 1.        ，   N，           ，    ，                     ，         ，          。
 
##          
 1. docker-compose.sh   CLI4       ，    
    1. miner CLI4
    1. transfer CLI4
 1. nginx.conf    server chain30:8802   ，    pause chain30  
 1.           ，   paracross testcase    MainParaSelfConsensusForkHeight fork.go local    
    pracross ForkParacrossCommitTx    ，  MainParaSelfConsensusForkHeight            