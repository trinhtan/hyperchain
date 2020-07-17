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

cp -r /fabric/config/chaincode $GOPATH/src/
export CHAINCODE_NAME="certificate"
export CHAINCODE_VERSION="1.0"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/academy.certificate.com/users/Admin@academy.certificate.com/msp"
export CORE_PEER_LOCALMSPID="AcademyMSP"

export CORE_PEER_ADDRESS="certificate-academy-peer0:30000"
set -x
peer chaincode install -n ${CHAINCODE_NAME} -l golang -v ${CHAINCODE_VERSION} -p chaincode/certificate/ >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Chaincode installation on ${CORE_PEER_ADDRESS} has failed"
echo "=== Chaincode is installed on ${CORE_PEER_ADDRESS} ==="
echo
sleep $MAX_SLEEP
cp -r /fabric/config/chaincode $GOPATH/src/
export CHAINCODE_NAME="certificate"
export CHAINCODE_VERSION="1.0"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/student.certificate.com/users/Admin@student.certificate.com/msp"
export CORE_PEER_LOCALMSPID="StudentMSP"

export CORE_PEER_ADDRESS="certificate-student-peer0:30001"
set -x
peer chaincode install -n ${CHAINCODE_NAME} -l golang -v ${CHAINCODE_VERSION} -p chaincode/certificate/ >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Chaincode installation on ${CORE_PEER_ADDRESS} has failed"
echo "=== Chaincode is installed on ${CORE_PEER_ADDRESS} ==="
echo
sleep $MAX_SLEEP

export CHAINCODE_NAME="certificate"
export CHAINCODE_VERSION="1.0"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/academy.certificate.com/users/Admin@academy.certificate.com/msp"
export CORE_PEER_LOCALMSPID="AcademyMSP"
export CORE_PEER_ADDRESS="certificate-academy-peer0:30000"
export ORDERER_URL="certificate-orderer:30006"
export FABRIC_LOGGING_LEVEL=debug
set -x
peer chaincode instantiate -o ${ORDERER_URL} -C ${CHANNEL_NAME} -n ${CHAINCODE_NAME} -l golang -v ${CHAINCODE_VERSION} -c '{"Args":[]}' -P "OR ('AcademyMSP.peer', 'StudentMSP.peer') " >&log.txt
res=$?
set +x
cat log.txt
verifyResult $res "Chaincode instantiation on ${CORE_PEER_ADDRESS} on channel '${CHANNEL_NAME}' failed"
echo "=== Chaincode is instantiated on ${CORE_PEER_ADDRESS} on channel '${CHANNEL_NAME}' ==="
echo
sleep $MAX_SLEEP

export CHAINCODE_NAME="certificate"
export FABRIC_CFG_PATH="/etc/hyperledger/fabric"
export ORDERER_URL="certificate-orderer:30006"
export CORE_PEER_LOCALMSPID="AcademyMSP"
export CORE_PEER_MSPCONFIGPATH="/fabric/crypto-config/peerOrganizations/academy.certificate.com/users/Admin@academy.certificate.com/msp"
export CORE_PEER_ADDRESS="certificate-academy-peer0:30000"

#peer chaincode invoke --peerAddresses ${CORE_PEER_ADDRESS} -o ${ORDERER_URL} -C ${CHANNEL_NAME} -n ${CHAINCODE_NAME} -c '{"Args":["CreateSubject","Subject01","Trinh Van Tan", "abc"]}'
#sleep $MAX_SLEEP
#peer chaincode query -C ${CHANNEL_NAME} -n ${CHAINCODE_NAME} -c '{"Args": ["GetAllSubjects"]}'
exit 0