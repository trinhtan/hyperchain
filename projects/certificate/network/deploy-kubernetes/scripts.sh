MIN_SLEEP=1
MAX_SLEEP=3
replace() {
    sed -e "s/\${ORG}/$1/" \
        -e "s/\${ORG_}/$1/" \
        -e "s/\${PRIVATE_KEY}/$2/" \
        -e "s/\${PORT}/$3/" \
        kubernetes/certificate-ca_deploy.yaml | sed -e $'s/\\\\n/\\\n        /g'
}

kubectl apply -f kubernetes/fabric-certificate_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/fabric-certificate_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-certificate-academy-peer0_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-certificate-academy-peer0_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-certificate-student-peer0_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-certificate-student-peer0_pvc.yaml
sleep $MIN_SLEEP


kubectl apply -f kubernetes/backup-certificate-orderer_pv.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/backup-certificate-orderer_pvc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/fabric-certificate-tools.yaml
sleep $MAX_SLEEP

kubectl exec -i fabric-certificate-tools -- mkdir /fabric/config
sleep $MIN_SLEEP

kubectl cp ./network.sh fabric-certificate-tools:/
sleep $MIN_SLEEP

kubectl cp ./channel.sh fabric-certificate-tools:/
sleep $MIN_SLEEP

kubectl cp config/configtx.yaml fabric-certificate-tools:/fabric/config/
sleep $MIN_SLEEP

kubectl cp config/crypto-config.yaml fabric-certificate-tools:/fabric/config/
sleep $MIN_SLEEP

kubectl exec -i fabric-certificate-tools -- mkdir -p /opt/gopath/src/github.com/hyperledger
sleep $MIN_SLEEP

kubectl cp ~/go/src/github.com/hyperledger/fabric  fabric-certificate-tools:/opt/gopath/src/github.com/hyperledger/
sleep $MIN_SLEEP

kubectl cp ~/go/src/github.com/golang  fabric-certificate-tools:/opt/gopath/src/github.com/
sleep $MIN_SLEEP

kubectl exec -i fabric-certificate-tools -- /bin/bash network.sh
sleep $MIN_SLEEP

kubectl cp fabric-certificate-tools:/fabric/crypto-config/ordererOrganizations/certificate.com/tlsca/tlsca.certificate.com-cert.pem ./config/connection-files/tlsca.certificate.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-certificate-tools:/fabric/crypto-config/peerOrganizations/academy.certificate.com/tlsca/tlsca.academy.certificate.com-cert.pem ./config/connection-files/academy/tlsca.academy.certificate.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-certificate-tools:/fabric/crypto-config/peerOrganizations/academy.certificate.com/ca/ca.academy.certificate.com-cert.pem ./config/connection-files/academy/ca.academy.certificate.com-cert.pem
sleep $MIN_SLEEP
kubectl cp fabric-certificate-tools:/fabric/crypto-config/peerOrganizations/student.certificate.com/tlsca/tlsca.student.certificate.com-cert.pem ./config/connection-files/student/tlsca.student.certificate.com-cert.pem
sleep $MIN_SLEEP

kubectl cp fabric-certificate-tools:/fabric/crypto-config/peerOrganizations/student.certificate.com/ca/ca.student.certificate.com-cert.pem ./config/connection-files/student/ca.student.certificate.com-cert.pem
sleep $MIN_SLEEP

cd config/connection-files
chmod 777 *
./ccp-generate.sh
cd ../..


export PRIVATE_KEY=$(kubectl exec -i fabric-certificate-tools -- ls /fabric/crypto-config/peerOrganizations/academy.certificate.com/ca/ | grep ".*_sk")
echo "$(replace academy $PRIVATE_KEY 7054)" > kubernetes/certificate-academy-ca_deploy.yaml

kubectl apply -f kubernetes/certificate-academy-ca_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-academy-ca_svc.yaml
sleep $MIN_SLEEP


export PRIVATE_KEY=$(kubectl exec -i fabric-certificate-tools -- ls /fabric/crypto-config/peerOrganizations/student.certificate.com/ca/ | grep ".*_sk")
echo "$(replace student $PRIVATE_KEY 7054)" > kubernetes/certificate-student-ca_deploy.yaml

kubectl apply -f kubernetes/certificate-student-ca_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-student-ca_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-orderer_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-orderer_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-academy-peer0_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-academy-peer0_svc.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-student-peer0_deploy.yaml
sleep $MIN_SLEEP

kubectl apply -f kubernetes/certificate-student-peer0_svc.yaml
sleep $MIN_SLEEP

sleep $MAX_SLEEP

set -x
kubectl exec -i fabric-certificate-tools -- /bin/bash channel.sh
res=$?
set +x
if [ $res -ne 0 ]; then
kubectl cp fabric-certificate-tools:/fabric/log.txt ../../log.txt
exit 1
else
exit 0
fi