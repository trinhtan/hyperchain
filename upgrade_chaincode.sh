# eval $(parse_yaml config/orgsconfig.yaml "config_")
eval $(minikube docker-env)

PROJECT=$1
CHAINCODE_VERSION=$2

# python3 generate_chaincode.py ${PROJECT}
python3 generate_upgrade_chaincode.py ${PROJECT} ${CHAINCODE_VERSION}

cd projects/${PROJECT}/network/deploy-kubernetes/
kubectl cp config/chaincode/ fabric-${PROJECT}-tools:/fabric/config/
kubectl cp ./upgrade_chaincode.sh fabric-certificate-tools:/

set -x
kubectl exec -it fabric-${PROJECT}-tools -- /bin/bash upgrade_chaincode.sh
res=$?
set +x

if [ $res -ne 0 ]; then
kubectl cp fabric-${PROJECT}-tools:log.txt ../../log.txt
exit 1
else
exit 0
fi