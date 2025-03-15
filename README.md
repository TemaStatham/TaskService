cd functions
docker build -t profileserviceproto-go:v2.0 .
docker tag profileserviceproto-go:v2.0 temastatham/profileserviceproto-go:v2.0
docker login
docker push temastatham/profileserviceproto-go:v2.0

docker run -it --rm -v $PWD:/opt/functions -w /opt/functions temastatham/profileserviceproto-go:v2.0
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative functions.proto
ls -tal
exit

cd ../

```
docker run -it --rm -v $PWD:/opt/functions -w /opt/functions davarski/protoc-go
root@4eb705f9eea5:/opt/functions# protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative functions.proto
root@4eb705f9eea5:/opt/functions# ls -tal
total 36
-rw-r--r-- 1 root root 6393 Aug  1 08:22 functions.pb.go
-rw-r--r-- 1 root root 8564 Aug  1 08:22 functions_grpc.pb.go
drwxr-xr-x 1 root root 4096 Aug  1 08:22 ..
drwxr-xr-x 2 1000 1000 4096 Aug  1 08:17 .
-rw-r--r-- 1 1000 1000  421 Aug  1 08:07 Dockerfile
-rw-r--r-- 1 1000 1000  413 Aug  1 08:07 functions.proto
root@4eb705f9eea5:/opt/functions# exit
```






docker build -f server/Dockerfile -t profileservice-go:v2.0 .
docker build -f client/Dockerfile -t taskservice-go:v2.0 .

docker network create grpc

docker tag profileservice-go:v2.0 temastatham/profileservice-go:v2.0
docker tag taskservice-go:v2.0 temastatham/taskservice-go:v2.0

docker run -d --net grpc --name grpc-server -p 50501:50501 profileservice-go:v2.0
docker run --net grpc taskservice-go:v2.0

docker login
docker push temastatham/profileservice-go:v1.0
docker push temastatham/taskservice-go:v1.0






kind create cluster --name devops --config cluster-config.yaml
kind get kubeconfig --name="devops" > admin.conf
export KUBECONFIG=./admin.conf
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
kubectl -n kube-system set env daemonset/calico-node FELIX_IGNORELOOSERPF=true

kubectl apply -f kubernetes/setup/*
kubectl apply -f kubernetes/server.yaml
kubectl apply -f kubernetes/client.yaml

kubectl get all -n grpc-go

kubectl logs pod/grpc-server-66886c6698-b7mmk  -n grpc-go

kind delete cluster --name=devops
