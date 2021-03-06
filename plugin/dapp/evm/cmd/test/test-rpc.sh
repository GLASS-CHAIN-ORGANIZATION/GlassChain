#!/usr/bin/env bash
# shellcheck disable=SC2128
# shellcheck source=/dev/null
source ../dapp-test-common.sh

MAIN_HTTP=""
evm_createContract_unsignedTx="0a0365766d129407228405608060405234801561001057600080fd5b50610264806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063b8e010de1461003b578063cc80f6f314610045575b600080fd5b6100436100c2565b005b61004d610109565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561008757818101518382015260200161006f565b50505050905090810190601f1680156100b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051808201909152600d8082527f5468697320697320746573742e000000000000000000000000000000000000006020909201918252610106916000916101a0565b50565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156101955780601f1061016a57610100808354040283529160200191610195565b820191906000526020600020905b81548152906001019060200180831161017857829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101e157805160ff191683800117855561020e565b8280016001018555821561020e579182015b8281111561020e5782518255916020019190600101906101f3565b5061021a92915061021e565b5090565b61019d91905b8082111561021a576000815560010161022456fea165627a7a72305820fec5dd5ca2cb47523ba08c04749bc5c14c435afee039f3047c2b7ea2faca737800293a8a025b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2273686f77222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d5d20c0c7ee04309aedc4bcfba5beca5f3a223139746a5335316b6a7772436f535153313355336f7765376759424c6653666f466d"
evm_createContract_para_unsignedTx="0a0f757365722e702e706172612e65766d129407228405608060405234801561001057600080fd5b50610264806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063b8e010de1461003b578063cc80f6f314610045575b600080fd5b6100436100c2565b005b61004d610109565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561008757818101518382015260200161006f565b50505050905090810190601f1680156100b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051808201909152600d8082527f5468697320697320746573742e000000000000000000000000000000000000006020909201918252610106916000916101a0565b50565b60008054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156101955780601f1061016a57610100808354040283529160200191610195565b820191906000526020600020905b81548152906001019060200180831161017857829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101e157805160ff191683800117855561020e565b8280016001018555821561020e579182015b8281111561020e5782518255916020019190600101906101f3565b5061021a92915061021e565b5090565b61019d91905b8082111561021a576000815560010161022456fea165627a7a7230582080ff1004de2195e6c08d0d0a65484b3d393c99c280e305cb383dbc89343cdd6a00293a8a025b7b22636f6e7374616e74223a66616c73652c22696e70757473223a5b5d2c226e616d65223a22736574222c226f757470757473223a5b5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a226e6f6e70617961626c65222c2274797065223a2266756e6374696f6e227d2c7b22636f6e7374616e74223a747275652c22696e70757473223a5b5d2c226e616d65223a2273686f77222c226f757470757473223a5b7b226e616d65223a22222c2274797065223a22737472696e67227d5d2c2270617961626c65223a66616c73652c2273746174654d75746162696c697479223a2276696577222c2274797065223a2266756e6374696f6e227d5d20c0c7ee0430e1c7facdc1f199956c3a2231483969326a67464a594e5167573350695468694337796b7a5663653570764b7478"
evm_creatorAddr="1PTXh2EZ8aRUzpuoDRASV19K86Kx3qQiPt"
evm_creatorAddr_key="0x4947ce3c4b845cfed59be2edf47320546116a3ff3af5715a7df094d116039b89"
evm_contractAddr=""
evm_addr=""
txHash=""

function evm_createContract() {
    validator=$1
    expectRes=$2
    if [ "$ispara" == "true" ]; then
        paraName="user.p.para."
        chain33_SignAndSendTx "${evm_createContract_para_unsignedTx}" "${evm_creatorAddr_key}" "$MAIN_HTTP"
    else
        chain33_SignAndSendTx "${evm_createContract_unsignedTx}" "${evm_creatorAddr_key}" "$MAIN_HTTP"
    fi
    txHash=$RAW_TX_HASH
    queryTransaction "${validator}" "${expectRes}"

    echo "CreateContract queryExecRes end"

    chain33_BlockWait 1 "$MAIN_HTTP"
}

function evm_addressCheck() {
    req='{"method":"Chain33.Query","params":[{"execer":"evm","funcName":"CheckAddrExists","payload":{"addr":"'${evm_contractAddr}'"}}]}'
    resok='(.result.contract == true ) and (.result.contractAddr == "'"${evm_contractAddr}"'")'
    chain33_Http "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}
function evm_callContract() {
    op=$1
    validator=$2
    expectRes=$3
    if [ "${op}" == "preExec" ]; then
        unsignedTx=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"evm.EvmCallTx","params":[{"fee":1,"caller":"'${evm_creatorAddr}'", "expire":"120s", "exec":"'${evm_addr}'", "abi": "set()"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r ".result")
    elif [ "${op}" == "Exec" ]; then
        unsignedTx=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"evm.EvmCallTx","params":[{"fee":1,"caller":"'${evm_creatorAddr}'", "expire":"120s", "exec":"'${evm_addr}'", "abi": "show()"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r ".result")
    else
        rst=1
        echo_rst "CallContract invalid param" "$rst"
        return
    fi

    chain33_SignAndSendTx "${unsignedTx}" "${evm_creatorAddr_key}" "$MAIN_HTTP"
    txHash=$RAW_TX_HASH
    queryTransaction "${validator}" "${expectRes}"
    echo "CallContract queryExecRes end"

    chain33_BlockWait 1 "$MAIN_HTTP"
}

function evm_abiGet() {
    req='{"method":"Chain33.Query","params":[{"execer":"evm","funcName":"QueryABI","payload":{"address":"'${evm_contractAddr}'"}}]}'
    chain33_Http "$req" ${MAIN_HTTP} "(.result.abi != null)" "$FUNCNAME"
}

function evm_transfer() {
    validator=$1
    expectRes=$2
    unsignedTx=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"evm.EvmTransferTx","params":[{"amount":1,"caller":"'${evm_creatorAddr}'","expire":"", "exec":"'${evm_addr}'", "paraName": "'${paraName}'"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r ".result")
    if [ "${unsignedTx}" == "" ]; then
        rst=1
        echo_rst "evm transfer create tx" "$rst"
        return
    fi

    chain33_SignAndSendTx "${unsignedTx}" "${evm_creatorAddr_key}" "$MAIN_HTTP"
    txHash=$RAW_TX_HASH
    queryTransaction "${validator}" "${expectRes}"
    echo "evm transfer queryExecRes end"

    chain33_BlockWait 2 "$MAIN_HTTP"
}

function evm_getBalance() {
    expectBalance=$1
    req='{"method":"Chain33.GetBalance","params":[{"addresses":["'${evm_creatorAddr}'"],"execer":"'${evm_addr}'", "paraName": "'${paraName}'"}]}'
    resok='(.result[0].balance == '$expectBalance') and (.result[0].addr == "'"$evm_creatorAddr"'")'
    chain33_Http "$req" ${MAIN_HTTP} "$resok" "$FUNCNAME"
}

function evm_withDraw() {
    validator=$1
    expectRes=$2
    unsignedTx=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"evm.EvmWithdrawTx","params":[{"amount":1,"caller":"'${evm_creatorAddr}'","expire":"", "exec":"'${evm_addr}'", "paraName":"'${paraName}'"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP} | jq -r ".result")
    if [ "${unsignedTx}" == "" ]; then
        rst=1
        echo_rst "evm withdraw create tx" "$rst"
        return
    fi

    chain33_SignAndSendTx "${unsignedTx}" "${evm_creatorAddr_key}" "$MAIN_HTTP"
    txHash=$RAW_TX_HASH
    queryTransaction "${validator}" "${expectRes}"
    echo "evm withdraw queryExecRes end"

    chain33_BlockWait 1 "$MAIN_HTTP"
}

#          
#        ???        ???  1:        2:       ???
function queryTransaction() {
    validators=$1
    expectRes=$2
    echo "txHash=${txHash}"

    res=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"Chain33.QueryTransaction","params":[{"hash":"'"${txHash}"'"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP})

    times=$(echo "${validators}" | awk -F '|' '{print NF}')
    for ((i = 1; i <= times; i++)); do
        validator=$(echo "${validators}" | awk -F '|' '{print $'$i'}')
        res=$(echo "${res}" | ${validator})
    done

    if [ "${res}" != "${expectRes}" ]; then
        return 1
    else
        res=$(curl -s --data-binary '{"jsonrpc":"2.0","id":2,"method":"Chain33.QueryTransaction","params":[{"hash":"'"${txHash}"'"}]}' -H 'content-type:text/plain;' ${MAIN_HTTP})
        if [ "${evm_addr}" == "" ]; then
            if [ "$ispara" == false ]; then
                evm_addr="user.evm.${txHash}"
            else
                evm_addr="user.p.para.user.evm.${txHash}"
            fi
        fi

        if [ "${evm_contractAddr}" == "" ]; then
            evm_contractAddr=$(curl -ksd '{"method":"Chain33.ConvertExectoAddr","params":[{"execname":"'"${evm_addr}"'"}]}' ${MAIN_HTTP} | jq -r ".result")

        fi
        return 0
    fi
}

function init() {
    ispara=$(echo '"'"${MAIN_HTTP}"'"' | jq '.|contains("8901")')
    echo "ipara=$ispara"

    local main_ip=${MAIN_HTTP//8901/8801}
    chain33_ImportPrivkey "0x4947ce3c4b845cfed59be2edf47320546116a3ff3af5715a7df094d116039b89" "1PTXh2EZ8aRUzpuoDRASV19K86Kx3qQiPt" "evm" "${main_ip}"

    local from="1PTXh2EZ8aRUzpuoDRASV19K86Kx3qQiPt"

    if [ "$ispara" == false ]; then
        chain33_applyCoins "$from" 12000000000 "${main_ip}"
        chain33_QueryBalance "${from}" "$main_ip"
    else
        chain33_applyCoins "$from" 1000000000 "${main_ip}"
        chain33_QueryBalance "${from}" "$main_ip"

        local para_ip="${MAIN_HTTP}"
        chain33_ImportPrivkey "0x4947ce3c4b845cfed59be2edf47320546116a3ff3af5715a7df094d116039b89" "1PTXh2EZ8aRUzpuoDRASV19K86Kx3qQiPt" "evm" "$para_ip"

        chain33_applyCoins "$from" 12000000000 "${para_ip}"
        chain33_QueryBalance "${from}" "$para_ip"
    fi

    chain33_BlockWait 2 "$MAIN_HTTP"

    local evm_addr=""
    if [ "$ispara" == "true" ]; then
        evm_addr=$(curl -ksd '{"method":"Chain33.ConvertExectoAddr","params":[{"execname":"user.p.para.evm"}]}' ${MAIN_HTTP} | jq -r ".result")
        chain33_SendToAddress "$from" "$evm_addr" 10000000000 "$MAIN_HTTP"
    else
        evm_addr=$(curl -ksd '{"method":"Chain33.ConvertExectoAddr","params":[{"execname":"evm"}]}' ${MAIN_HTTP} | jq -r ".result")
    fi
    echo "evm=$evm_addr"

    chain33_SendToAddress "$from" "$evm_addr" 10000000000 "$MAIN_HTTP"
    chain33_BlockWait 2 "$MAIN_HTTP"
}
function run_test() {
    local ip=$1
    evm_createContract "jq -r .result.receipt.tyName" "ExecOk"
    evm_addressCheck

    if [ "$ispara" == "true" ]; then
        evm_callContract preExec "jq -r .result.receipt.logs[0].tyName" "LogEVMStateChangeItem"
        evm_callContract Exec "jq -r .result.receipt.logs[0].log.jsonRet | jq -r .[0].value" "This is test."
    else
        evm_callContract preExec "jq -r .result.receipt.logs[1].tyName" "LogEVMStateChangeItem"
        evm_callContract Exec "jq -r .result.receipt.logs[1].log.jsonRet | jq -r .[0].value" "This is test."
    fi

    evm_abiGet
    evm_transfer "jq -r .result.receipt.tyName" "ExecOk"
    evm_getBalance 100000000
    evm_withDraw "jq -r .result.receipt.tyName" "ExecOk"
    evm_getBalance 0
}

function main() {
    chain33_RpcTestBegin evm
    local ip=$1
    MAIN_HTTP=$ip
    echo "main_ip=$MAIN_HTTP"

    init
    run_test "$MAIN_HTTP"
    chain33_RpcTestRst evm "$CASE_ERR"
}

chain33_debug_function main "$1"
