MIN_SLEEP=1
MAX_SLEEP=3
replace() {
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORG_}/$1/" \
        -e "s/\${PRIVATE_KEY}/$2/" \
        -e "s/\${PORT}/$3/" \
        kubernetes/edu-ca_deploy.yaml | sed -e $'s/\\\\n/\\\n        /g'
}

kubectl apply -f kubernetes/fabric-edu_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/fabric-edu_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-edu-student-peer0_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-edu-student-peer0_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-edu-academy-peer0_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-edu-academy-peer0_pvc.yaml
sleep $MIN_SLEEP


kubectl apply -f kubernetes/backup-edu-orderer_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-edu-orderer_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/fabric-edu-tools.yaml
sleep $MAX_SLEEP

kubectl exec -i fabric-edu-tools -- mkdir /fabric/config
sleep $MIN_SLEEP

kubectl cp ./network.sh fabric-edu-tools:/
sleep $MIN_SLEEP

kubectl cp ./channel.sh fabric-edu-tools:/
sleep $MIN_SLEEP

kubectl cp config/configtx.yaml fabric-edu-tools:/fabric/config/
sleep $MIN_SLEEP

kubectl cp config/crypto-config.yaml fabric-edu-tools:/fabric/config/
sleep $MIN_SLEEP

kubectl exec -i fabric-edu-tools -- mkdir -p /opt/gopath/src/github.com/hyperledger
sleep $MIN_SLEEP

kubectl cp ~/go/src/github.com/hyperledger/fabric  fabric-edu-tools:/opt/gopath/src/github.com/hyperledger/
sleep $MIN_SLEEP

kubectl cp ~/go/src/github.com/golang  fabric-edu-tools:/opt/gopath/src/github.com/
sleep $MIN_SLEEP

kubectl exec -i fabric-edu-tools -- /bin/bash network.sh
sleep $MIN_SLEEP

kubectl cp fabric-edu-tools:/fabric/crypto-config/ordererOrganizations/edu.com/tlsca/tlsca.edu.com-cert.pem ./config/connection-files/tlsca.edu.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-edu-tools:/fabric/crypto-config/peerOrganizations/student.edu.com/tlsca/tlsca.student.edu.com-cert.pem ./config/connection-files/student/tlsca.student.edu.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-edu-tools:/fabric/crypto-config/peerOrganizations/student.edu.com/ca/ca.student.edu.com-cert.pem ./config/connection-files/student/ca.student.edu.com-cert.pem
sleep $MIN_SLEEP
kubectl cp fabric-edu-tools:/fabric/crypto-config/peerOrganizations/academy.edu.com/tlsca/tlsca.academy.edu.com-cert.pem ./config/connection-files/academy/tlsca.academy.edu.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-edu-tools:/fabric/crypto-config/peerOrganizations/academy.edu.com/ca/ca.academy.edu.com-cert.pem ./config/connection-files/academy/ca.academy.edu.com-cert.pem
sleep $MIN_SLEEP

cd config/connection-files
chmod 777 *
./ccp-generate.sh
cd ../..


export PRIVATE_KEY=$(kubectl exec -i fabric-edu-tools -- ls /fabric/crypto-config/peerOrganizations/student.edu.com/ca/ | grep ".*_sk")
echo "$(replace student $PRIVATE_KEY 7054)" > kubernetes/edu-student-ca_deploy.yaml

kubectl apply -f kubernetes/edu-student-ca_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-student-ca_svc.yaml
sleep $MIN_SLEEP


export PRIVATE_KEY=$(kubectl exec -i fabric-edu-tools -- ls /fabric/crypto-config/peerOrganizations/academy.edu.com/ca/ | grep ".*_sk")
echo "$(replace academy $PRIVATE_KEY 7054)" > kubernetes/edu-academy-ca_deploy.yaml

kubectl apply -f kubernetes/edu-academy-ca_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-academy-ca_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-orderer_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-orderer_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-student-peer0_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-student-peer0_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-academy-peer0_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/edu-academy-peer0_svc.yaml
sleep $MIN_SLEEP

sleep $MAX_SLEEP

set -x
kubectl exec -i fabric-edu-tools -- /bin/bash channel.sh
res=$?
set +x
if [ $res -ne 0 ]; then
kubectl cp fabric-edu-tools:/fabric/log.txt ../../log.txt
exit 1
else
exit 0
fi