# exchange  

##   
      chain33            ，       ，                     ，         。

##   
                    ，                  
    ,                    CreateTransaction Query  ，     
[CreateTransaction  ](https://github.com/33cn/chain33/blob/master/rpc/jrpchandler.go#L1101) [Query  ](https://github.com/33cn/chain33/blob/master/rpc/jrpchandler.go#L838)

      |  
-----|----
QueryMarketDepth|             
QueryHistoryOrderList|                  
QueryOrder|  orderID            
QueryOrderList|           （ordered,completed,revoked)，              

   exchange_test.go        ，  limitOrder  revokeOrder        

##     
        ：

  |  
---|----
1|       
2|       ，         
3|       ，           
4|                
5|        ，       100 ，       1e8,    bty

**     **

  |  |  |  |  
 ---|---|---|---|---
 depth|price|nil|        |  price      {leftAsset}:{rightAsset}:{op}:{price}  
 order|orderID|market_order,addr_status|              |market_order      {leftAsset}:{rightAsset}:{op}:{price}:{orderID},addr_status      {addr}:{status}，          ，           order      
 history|index|name,addr_status|                     (revoked         )|name      {leftAsset}:{rightAsset}  , addr_status      {addr}:{status}

**        **

   |  
----|----
leftAsset|         
rightAsset|         
op|     1  ，2  
status|    ，0 ordered, 1 completed,2 revoked
price|    ，  16 %016d,           ，       ，        1e8。                 0.25，     25000000，price     1<=price<=1e16   
orderID|  ，       ，  ，  22 %022d
index|       index，  22 %022d

