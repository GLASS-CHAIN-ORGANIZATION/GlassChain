# accountmanager  

##   
        [             ](http://www.cfstc.org/bzgk/gk/view/yulan.jsp?i_id=1855)，           ，      ，       ，   chain33     
accountmanager  

##   
              ，                
      ：
  |  
----|----
    |    ，     ，            ， accountmanager   ，accountID     ，        
    |            ，              
       |          ，         ，           
       |       ，         ，          ，               
    |         ，      ，        ，            
    |      accountmanager          

    ,                    CreateTransaction Query  ，     
[CreateTransaction  ](https://github.com/33cn/chain33/blob/master/rpc/jrpchandler.go#L1101) [Query  ](https://github.com/33cn/chain33/blob/master/rpc/jrpchandler.go#L838)

      |  
-----|----
QueryAccountByID|    ID      ，       ID    
QueryAccountsByStatus|          
QueryExpiredAccounts|      
QueryAccountByAddr|            
QueryBalanceByID|    ID        


   account_test.go        ，          

##     

**     **

  |  |  |  |  
 ---|---|---|---|---
 account|index|accountID,addr,status|        |index      {expiretime*1e5+index(            )}  

**        **

   |  
----|----
Asset|    
op|       supervisor op 1   ，2   ，3     ,4    apply op 1         , 2       ，        
status|    ，0   ， 1    , 2     3,    
level|     0   ，           
index|        *1e5+           ，  15 %015d
