## Steps taken to deploy Jenkins 

minikube start

kubectl create namespace jenkins or kubectl create -f jenkins-namespace.yml 

kubectl get namespaces

helm repo list

helm repo add jenkinsci https://charts.jenkins.io

helm repo update

kubectl apply -f jenkins-volume.yml 

kubectl apply -f jenkins-sa.yml  

chart=jenkinsci/jenkins
helm install jenkins -n jenkins -f jenkins-values.yml $chart

minikube ssh
sudo chown -R 1000:1000 /data/jenkins-volume

kubectl delete pod jenkins-0 -n jenkins

kubectl get pods -n jenkins -w

kubectl exec --namespace jenkins -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/chart-admin-password && echo

kubectl --namespace jenkins port-forward svc/jenkins 8080:8080

open browser and login to http://localhost:8080

perform plugin updates 


