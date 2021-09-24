package main

var gxl = `
Title = "user.p.gxlchain."
TestNet = false
CoinSymbol = "GXL"
EnableParaFork = true

[log]
  loglevel = "debug"
  logConsoleLevel = "info"
  logFile = "logs/GXL.para.log"
  maxFileSize = 300
  maxBackups = 100
  maxAge = 28
  localTime = true
  compress = true
  callerFile = false
  callerFunction = false

[blockchain]
  defCacheSize = 128
  maxFetchBlockNum = 128
  timeoutSeconds = 5
  batchBlockNum = 128
  driver = "leveldb"
  dbPath = "paradatadir"
  dbCache = 64
  isStrongConsistency = true
  batchsync = false
  isRecordBlockSequence = true
  isParaChain = true
  enableTxQuickIndex = true
  enableReExecLocal = true
  enableReduceLocaldb = false

[p2p]
  enable = false
  driver = "leveldb"
  dbPath = "paradatadir/addrbook"
  dbCache = 4
  grpcLogFile = "grpc33.log"

[rpc] 
  whitelist = ["*"]
  jrpcFuncWhitelist = ["*"]
  grpcFuncWhitelist = ["*"]

[mempool]
  name = "para"
  poolCacheSize = 10240
  minTxFeeRate = 100000
  maxTxNumPerAccount = 10000

[consensus]
  name = "para"
  genesisBlockTime = 1514533394
  genesis = "1LXzJV51Tps97C1qDExDJ6coDBWHenCFmy"
  minerExecs = ["paracross"]

  [consensus.sub]

    [consensus.sub.para]
      ParaRemoteGrpcClient = "116.63.168.132:8802,172.31.232.187:8802"
      startHeight = 5710000
      writeBlockSeconds = 2
      authAccount = ""
      genesisAmount = 2000000000
      mainBlockHashForkHeight = 209186
      mainForkParacrossCommitTx = 2270000
      mainLoopCheckCommitTxDoneForkHeight = 4320000
      emptyBlockInterval = ["0:50"]

[mver]

  [mver.consensus]
    fundKeyAddr = "1BQXS6TxaYYG5mADaWij4AxhZZUTpw95a5"
    powLimitBits = "0x1f00ffff"
    maxTxNumber = 1600

    [mver.consensus.paracross]
      coinReward = 18
      coinDevFund = 12

[store]
  name = "kvmvccmavl"
  driver = "leveldb"
  storedbVersion = "2.0.0"
  dbPath = "paradatadir/mavltree"
  dbCache = 128

  [store.sub]

    [store.sub.mavl]
      enableMavlPrefix = false
      enableMVCC = false
      enableMavlPrune = false
      pruneHeight = 10000
      enableMemTree = true
      enableMemVal = true
      tkCloseCacheLen = 100000

    [store.sub.kvmvccmavl]
      enableMVCCIter = true
      enableMavlPrefix = false
      enableMVCC = false
      enableMavlPrune = false
      pruneMavlHeight = 10000
      enableMVCCPrune = false
      pruneMVCCHeight = 10000
      enableMemTree = true
      enableMemVal = true
      tkCloseCacheLen = 100000
      enableEmptyBlockHandle = false

[wallet]
  minFee = 100000
  driver = "leveldb"
  dbPath = "parawallet"
  dbCache = 16
  signType = "secp256k1"
  minerdisable = true

[exec]
  enableStat = false
  enableMVCC = false

  [exec.sub]

    [exec.sub.relay]
      genesis = "1G8qVvuFbnvoAuDE75UR8Bz5TWMV2oVmw8"

    [exec.sub.manage]
      superManager = ["1G8qVvuFbnvoAuDE75UR8Bz5TWMV2oVmw8"]

    [exec.sub.token]
      saveTokenTxList = true
      tokenApprs = ["1G8qVvuFbnvoAuDE75UR8Bz5TWMV2oVmw8"]

    [exec.sub.paracross]
      paraConsensusStopBlocks = 30000

    [exec.sub.autonomy]
      total = "16htvcBNSEA7fZhAdLJphDwQRQJaHpyHTp"
      useBalance = false

[fork]

  [fork.system]
    ForkChainParamV1 = 0
    ForkCheckTxDup = 0
    ForkBlockHash = 1
    ForkMinerTime = 0
    ForkTransferExec = 0
    ForkExecKey = 0
    ForkTxGroup = 0
    ForkResetTx0 = 0
    ForkWithdraw = 0
    ForkExecRollback = 0
    ForkCheckBlockTime = 0
    ForkTxHeight = 0
    ForkTxGroupPara = 0
    ForkChainParamV2 = 0
    ForkMultiSignAddress = 0
    ForkStateDBSet = 0
    ForkLocalDBAccess = 0
    ForkBlockCheck = 0
    ForkBase58AddressCheck = 100000
    ForkEnableParaRegExec = 0
    ForkCacheDriver = 0
    ForkTicketFundAddrV1 = -1
    ForkRootHash = 7200000

  [fork.sub]

    [fork.sub.coins]
      Enable = 0

    [fork.sub.ticket]
      Enable = 0
      ForkTicketId = 0
      ForkTicketVrf = 0

    [fork.sub.retrieve]
      Enable = 0
      ForkRetrive = 0
      ForkRetriveAsset = 0

    [fork.sub.hashlock]
      Enable = 0
      ForkBadRepeatSecret = 0

    [fork.sub.manage]
      Enable = 0
      ForkManageExec = 0

    [fork.sub.token]
      Enable = 0
      ForkTokenBlackList = 0
      ForkBadTokenSymbol = 0
      ForkTokenPrice = 0
      ForkTokenSymbolWithNumber = 0
      ForkTokenCheck = 0

    [fork.sub.trade]
      Enable = 0
      ForkTradeBuyLimit = 0
      ForkTradeAsset = 0
      ForkTradeID = 0
      ForkTradeFixAssetDB = 0
      ForkTradePrice = 0

    [fork.sub.paracross]
      Enable = 0
      ForkParacrossWithdrawFromParachain = 0
      ForkParacrossCommitTx = 0
      ForkLoopCheckCommitTxDone = 0
      ForkParaSelfConsStages = 0
      ForkParaAssetTransferRbk = 0

    [fork.sub.evm]
      Enable = 0
      ForkEVMState = 0
      ForkEVMABI = 0
      ForkEVMFrozen = 0
      ForkEVMKVHash = 0

    [fork.sub.blackwhite]
      Enable = 0
      ForkBlackWhiteV2 = 0

    [fork.sub.cert]
      Enable = 0

    [fork.sub.guess]
      Enable = 0

    [fork.sub.lottery]
      Enable = 0

    [fork.sub.oracle]
      Enable = 0

    [fork.sub.relay]
      Enable = 0

    [fork.sub.norm]
      Enable = 0

    [fork.sub.pokerbull]
      Enable = 0

    [fork.sub.privacy]
      Enable = 0

    [fork.sub.game]
      Enable = 0

    [fork.sub.multisig]
      Enable = 0

    [fork.sub.unfreeze]
      Enable = 0
      ForkTerminatePart = 0
      ForkUnfreezeIDX = 0

    [fork.sub.autonomy]
      Enable = 0

    [fork.sub.jsvm]
      Enable = 0

    [fork.sub.issuance]
      Enable = 0
      ForkIssuanceTableUpdate = 0

    [fork.sub.collateralize]
      Enable = 0
      ForkCollateralizeTableUpdate = 0

    [fork.sub.store-kvmvccmavl]
      ForkKvmvccmavl = 0

    [fork.sub.storage] 
      Enable=0 
      ForkStorageLocalDB=0 

[pprof]
  listenAddr = "localhost:6061"

[metrics]
  enableMetrics = false
  dataEmitMode = "influxdb"

  [metrics.sub]

    [metrics.sub.influxdb]
      duration = 1000000000
      url = "http://influxdb:8086"
      database = "chain33metrics"
      username = ""
      password = ""
      namespace = ""
`
