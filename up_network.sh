# eval $(parse_yaml config/orgsconfig.yaml "config_")
eval $(minikube docker-env)

PROJECT=$1
PUBLIC_IP=$2
rm -rf projects/${PROJECT}

mkdir -p projects/${PROJECT}/network/deploy-kubernetes/config/connection-files
mkdir -p projects/${PROJECT}/network/deploy-kubernetes/kubernetes

python3 generate_network.py ${PROJECT} ${PUBLIC_IP}

cd projects/${PROJECT}/network/deploy-kubernetes/
chmod 755 scripts.sh
chmod 755 network.sh
chmod 755 channel.sh
./scripts.sh
