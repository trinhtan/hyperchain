COUNTER=1
MAX_RETRY=5
MIN_SLEEP=1
MAX_SLEEP=3

verifyResult() {
  if [ $1 -ne 0 ]; then
    echo "!!!!!!!!!!!!!!! "$2" !!!!!!!!!!!!!!!!"
    echo "=== ERROR !!! FAILED to execute End-2-End Scenario ==="
    echo
    exit 1
  fi
}

joinChannelWithRetry() {

    peer channel fetch newest -o ${ORDERER_URL} -c ${CHANNEL_NAME}
    set -x
    peer channel join -b ${CHANNEL_NAME}_newest.block >&log.txt
    res=$?
    set +x
    cat log.txt
    if [ $res -ne 0 -a $COUNTER -lt $MAX_RETRY ]; then
        rm -rf /${CHANNEL_NAME}_newest.block
        COUNTER=$(expr $COUNTER + 1)
        echo "${CORE_PEER_ADDRESS} failed to join the channel, Retry after 3 seconds"
        sleep $MAX_SLEEP
        joinChannelWithRetry
    else
        rm -rf /${CHANNEL_NAME}_newest.block
        COUNTER=1
    fi
    verifyResult $res "After $MAX_RETRY attempts, ${CORE_PEER_ADDRESS} has failed to join channel '${CHANNEL_NAME}' "
}


cd /fabric
export ORDERER_URL="edu-orderer:30006"
export CORE_PEER_ADDRESSAUTODETECT="false"
export CORE_PEER_NETWORKID="nid1"
export CORE_PEER_LOCALMSPID="StudentMSP"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/student.edu.com/users/Admin@student.edu.com/msp"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_ADDRESS="edu-student-peer0:30000"
set -x
peer channel create -o ${ORDERER_URL} -c ${CHANNEL_NAME} -f /fabric/${CHANNEL_NAME}.tx >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Channel creation failed"
echo "=== Channel '${CHANNEL_NAME}' created ==="
echo

sleep $MAX_SLEEP

cd /
export CORE_PEER_NETWORKID="nid1"
export ORDERER_URL="edu-orderer:30006"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_LOCALMSPID="StudentMSP"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/student.edu.com/users/Admin@student.edu.com/msp"

export CORE_PEER_ADDRESS="edu-student-peer0:30000"
# peer channel fetch newest -o ${ORDERER_URL} -c ${CHANNEL_NAME}
# peer channel join -b ${CHANNEL_NAME}_newest.block
# rm -rf /${CHANNEL_NAME}_newest.block
joinChannelWithRetry
echo "=== ${CORE_PEER_ADDRESS} joined channel '${CHANNEL_NAME}'===="
echo
sleep $MAX_SLEEP

export CORE_PEER_NETWORKID="nid1"
export ORDERER_URL="edu-orderer:30006"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_LOCALMSPID="AcademyMSP"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/academy.edu.com/users/Admin@academy.edu.com/msp"

export CORE_PEER_ADDRESS="edu-academy-peer0:30001"
# peer channel fetch newest -o ${ORDERER_URL} -c ${CHANNEL_NAME}
# peer channel join -b ${CHANNEL_NAME}_newest.block
# rm -rf /${CHANNEL_NAME}_newest.block
joinChannelWithRetry
echo "=== ${CORE_PEER_ADDRESS} joined channel '${CHANNEL_NAME}'===="
echo
sleep $MAX_SLEEP


export CORE_PEER_LOCALMSPID="StudentMSP"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/student.edu.com/users/Admin@student.edu.com/msp"
export CORE_PEER_ADDRESS="edu-student-peer0:30000"
export ORDERER_URL="edu-orderer:30006"
export FABRIC_LOGGING_LEVEL=debug
set -x
peer channel update -f /fabric/StudentMSPanchors.tx -c ${CHANNEL_NAME}  -o ${ORDERER_URL} >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Anchor peer ${CORE_PEER_ADDRESS} update failed"
echo "=== Anchor peers updated for org '${CORE_PEER_LOCALMSPID}' on channel '${CHANNEL_NAME}' ==="
echo
sleep $MAX_SLEEP
export CORE_PEER_LOCALMSPID="AcademyMSP"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/academy.edu.com/users/Admin@academy.edu.com/msp"
export CORE_PEER_ADDRESS="edu-academy-peer0:30001"
export ORDERER_URL="edu-orderer:30006"
export FABRIC_LOGGING_LEVEL=debug
set -x
peer channel update -f /fabric/AcademyMSPanchors.tx -c ${CHANNEL_NAME}  -o ${ORDERER_URL} >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Anchor peer ${CORE_PEER_ADDRESS} update failed"
echo "=== Anchor peers updated for org '${CORE_PEER_LOCALMSPID}' on channel '${CHANNEL_NAME}' ==="
echo
sleep $MAX_SLEEP
exit 0