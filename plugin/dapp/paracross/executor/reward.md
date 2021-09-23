#              
>                    ，      ，           

##       
1.                            ，            
1.                    2/3                

##         
1.                          ，                    
1.         4     A,B,C,D，  A,B,C，D         ，    2/3   ， C             ，       A,B,C， D          ，       

##         （     )
1.           ，         ，         coins paracross  ，            
1.                       
1.        coins            
1.                      ，               
1.         bind   coins              coins  ，    unbind              

##        
>          
```
    [mver.consensus.paracross]
    #        
    coinReward=18
    #      
    coinDevFund=12
    #              ，       coinBaseReward，    (coinReward-coinBaseReward)         
    coinBaseReward=3
    #           (   )
    unBindTime=24
```
1.                  coinDevFund+coinReward=30，     ，        12，   18 coin           
1.                      ，          coinBaseReward，    (coinReward-coinBaseReward)                      。

##       
1.      /          （   ）

    ```
    {
        "method" : "Chain33.CreateTransaction",
        "params" : [
            {
            "execer" : "{user.p.para}.paracross",
            "actionName" : "ParaBindMiner",
            "payload" : {
    　　　　　　　"bindAction":"1"
                "bindCoins" : 5,
                "targetNode" : "1KSBd17H7ZK8iT37aJztFB22XGwsPTdwE4",
            }
            }
        ],
    }
    ```

    **    ：**

    |  |  |  |
    |----|----|----|
    |method|string|Chain33.CreateTransaction|
    |execer|string|          user.p.para.paracross,title:user.p.para.    |
    |actionName|string|ParaBindMiner|
    |bindAction|string|  :1，   :2|
    |bindCoins|int|          ，          ，         |
    |targetNode|string|        ，           |


