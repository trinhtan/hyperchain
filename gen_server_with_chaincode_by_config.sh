eval $(minikube docker-env)

PROJECT=$1

mkdir -p projects/${PROJECT}/server
mkdir -p projects/${PROJECT}/server/cli
mkdir -p projects/${PROJECT}/server/fabric

python3 generate_server.py ${PROJECT} 'yes'
cp -r projects/${PROJECT}/network/deploy-kubernetes/config/connection-files projects/${PROJECT}/server

cd projects/${PROJECT}
zip -r server.zip server

exit 0
