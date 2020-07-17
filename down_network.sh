PROJECT=$1
eval $(minikube docker-env)

kubectl delete svc $(kubectl get svc | grep ".*${PROJECT}-" | awk '{print $1}')

sleep 3

kubectl delete deployment $(kubectl get deployment | grep ".*${PROJECT}-" | awk '{print $1}')

sleep 3

kubectl delete pod $(kubectl get pod | grep ".*${PROJECT}-" | awk '{print $1}')

sleep 3

kubectl delete pvc $(kubectl get pvc | grep ".*${PROJECT}-" | awk '{print $1}')

sleep 3

kubectl delete pv $(kubectl get pv | grep ".*${PROJECT}-" | awk '{print $1}')

exit 0