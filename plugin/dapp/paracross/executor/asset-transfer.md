# paracross                 

##     

    ，             ，              。

    
 1.          :               
 1.          :                  
 1.      paracross    ：         
 1.      paracross    ：         

    
 1. A: conis/token ->   paracross    
 1. A:   paracross     ->    paracross    
 1. A:    paracross     ->   
 1. B(    A):    paracross     ->   paracross    

##       

        ，                
 1.       ，                title：  user.p.guodun.paracross，                   

         
 1.             ，                  
 1.             bitmap，            ，      ，      bit　      

##   

asset-transfer    ，     ，     


       transfer
 *   
   1.     paracross    ， balance -
   1.     paracross    ， balance +
 *    (        ，            )
   1.        paracross      balance +

     withdraw
 *    
   1.        paracross      balance -
 *   (      ，      )
   1. commit        
   1.     paracross    ， balance -
   1.     paracross    ， balance +

  <->        cross-transfer
>cross-transfer  transfer withdraw    transfer,            transfer  withdraw

*  　=　assetExec + assetSymbol         
  1.     ：coins+BTY,token+CCNY
  1.      :user.p.test.coins + FZM,
  1.             paracross    :   ：paracross　+ user.p.test.coins.FZM，　   : user.p.test.paracross + coins.BTY
  1.                  ，        ，       
  1.        title     transfer        
  :
```
				exec                    symbol                              tx.title=user.p.test1   tx.title=user.p.test2
1.       ：
				coins                   bty                                 transfer                 transfer
				paracross               user.p.test1.coins.fzm              withdraw                 transfer

2.        ：
				user.p.test1.coins      fzm                                 transfer                 NAN
                user.p.test1.paracross  coins.bty                           withdraw                 NAN
                user.p.test1.paracross  paracross.user.p.test2.coins.cny    withdraw                 NAN

  user.p.test1.paracross.paracross.user.p.test2.coins.cny    ：
user.p.test1.paracross.    paracross   ，　paracross.user.p.test2.coins.cny paracross      paracross     user.p.test2.coins.cny  
```

             
 1.   
 1.    
 1.    

### kv，  ，   

 1. kv     
    1.       
    1.         (   ，     )
    1.        
 1.   
    1.          
 1.   
    1.   ：       /       

```
                                                                      
account                           A            B                  A      B
1 A  5bty    5           0                5              0               0       0           0
 para  
2 B  5bty    10          0                5              5               0       0           0
 para  
3 A  4bty    10          4                1              5               0       0           0           
            10          4                1              5               4       4           0            
4 A 3bty      10          4                1              5               1       1           0
    
5  B  2bty   10          4                1              5               3       1           2
    ，   2bty    
6 B         10          4                1              5               3       1           2           
   1bty      10          4                1              5               2       1           1             
              10           3                1              6               2       1           1            
```

###   <->        cross-transfer　  
```
# Alice     ５coins-bty -> user.p.test.    :

                    coins       paracross:Addr(Alice)   paracross:Addr(user.p.test.paracross)    user.p.test.paracross-coins-bty:Addr(Alice) 
1 Alice                5
2 to                  0　　　　　　　　 5       
3 cross-transfer       0            5-5=0                   0+5=5                                          0+5=5

# Alice      ５paracross-coins.bty ->   
4 cross-transfer                    　5                   5-5=0                                       5-5=0
5 withdraw           　5              0

# Bob      5 user.p.test.coins.fzm ->   
                    paracross-user.p.test.coins.fzm:Addr(Bob)    user.p.test.coins.fzm      user.p.test.paracross:Addr(Bob)   user.p.test.paracross:Addr(paracross)
1 Bob                                                                       5
2 to paracross  　　            　　　　　　　　                               0                       5       
3 cross-transfer                  0+5=5                                                             5-5=0                             0+5=5     

# Bob     ５exec:paracross　symbol:user.p.test.coins.fzm ->    
4 cross-transfer                  5-5=0                                                             0+5=5                                5-5=0
5 withdraw                                                                  5                       5-5=0


```