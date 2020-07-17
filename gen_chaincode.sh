eval $(minikube docker-env)

PROJECT=$1

mkdir -p projects/${PROJECT}/network/deploy-kubernetes/config/chaincode/${PROJECT}
python3 generate_chaincode.py ${PROJECT}

exit 0
