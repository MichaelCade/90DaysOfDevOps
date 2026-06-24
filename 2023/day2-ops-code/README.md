# Getting started

This repo expects you to have a working kubernetes cluster already setup and
available with kubectl

We expect you already have a kubernetes cluster setup and available with kubectl and helm.

I like using (Civo)[https://www.civo.com/] for this as it is easy to setup and run clusters

The code is available in this folder to build/push your own images if you wish - there are no instructions for this.

## Start the Database
```shell
kubectl apply -f database/mysql.yaml
```


## deploy the day1 - sync
```shell
kubectl apply -f synchronous/k8s.yaml
```

Check your logs 
```shell 
kubectl logs deploy/generator

kubectl logs deploy/requestor
```

## deploy nats
helm repo add nats https://nats-io.github.io/k8s/helm/charts/
helm install my-nats nats/nats

## deploy day 2 - async
```shell
kubectl apply -f async/k8s.yaml
```

Check your logs
```shell 
kubectl logs deploy/generator

kubectl logs deploy/requestor
```
