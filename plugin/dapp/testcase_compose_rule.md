# dapp ci test cases compose rule/ ci test cases     

##     
  ：  dapp           ，  relay,paracross
  dapp testcase  CI         Dockerfile          cmd/build   
  paracross plugin/dapp/paracross/cmd/build  

cmd        Makefile build.sh    make   build     copy    build/ci    “dapp”      
  paracross copy chain33/build/ci/paracross   

  make docker-compose DAPP="dapp"       (  dapp  dapp      paracross)
     build/ci/dapp         dapp-ci  ， dapp    bin        copy       
        dockers      ，     docker exec build_chain33_1 /root/chain33-cli [cmd...]      
  make docker-compose-down DAPP="dapp"     dapp     ，  docker

     dapp     all run   dapp， all       pass dapp   

 1. make docker-compose [proj=xx] [dapp=xx]
    1.   proj  dapp      make docker-compose,   run    test case，  run  dapp
    1.   proj   ，       build     docker-compose service   ，         ，
         proj    docker compose   
    1.   dapp   ，  run  dapp，    ，  run    dapp，run        make docker-compose-down dapp=xx  
    1.   dapp=all   ALL，  run     testcase dapp
 1. make docker-compose down [proj=xx] [dapp=xx] 
      clean make docker-compose  make fork-test     docker  ， proj  dapp    
 1. make fork-test [proj=xx] [dapp=xx]       
    1.    make docker-compose     


## dapp/cmd/build      
build       CI      
 1. Dockerfile，    dapp    Dockerfile    ，     ，       ，          ，        ，      ，
          docker  Dockerfile，         Dockerfile-xxx，        docker-compose yml    
          Dockerfile    ，    
 1. docker-compose yml     docker service   ，chain33      2 docker    ，   Dockerfile      ，   
        ，docker-compose           ，    dapp           ，       ， dapp     。
    docker-compose yml     docker service     
       docker-compose    ，      
         docker-compose.yml      ，dapp compose      docker-compose-$dapp.yml       
 1. testcase.sh            ，     testcase.sh     ，       。
       testcase    step：
    1. init  docker              
    1. config： docker    dapp      ，    ，     
    1. test： ci      
    testcase          dapp     function      ，  step     
    ```
     function paracross() {
         if [ "${2}" == "init" ]; then
             para_init
         elif [ "${2}" == "config" ]; then
             para_transfer
             para_set_wallet
         elif [ "${2}" == "test" ]; then
             para_test 
         fi
     
     }
     ```    
 1. fork-test.sh       ，          ，       
    fork-test dapp       testcase  ，  source  testcase.sh import  
    fork test     5 step     
    1. forkInit: docker           
    1. forkConfig: docker service         ，   
    1. forkAGroupRun:           group A B              ，   Agroup            
    1. forkBGroupRun: B group    dapp   
    1. forkChekRst:                  
    ```
     function privacy() {
         if [ "${2}" == "forkInit" ]; then
             privacy_init
         elif [ "${2}" == "forkConfig" ]; then
             initPriAccount
         elif [ "${2}" == "forkAGroupRun" ]; then
             genFirstChainPritx
             genFirstChainPritxType4
         elif [ "${2}" == "forkBGroupRun" ]; then
             genSecondChainPritx
             genSecondChainPritxType4
         elif [ "${2}" == "forkCheckRst" ]; then
             checkPriResult
         fi
    
     }
     ```
 


