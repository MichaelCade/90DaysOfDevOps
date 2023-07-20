# Getting Hands-On with HashiCorp Vault 

A quick note to thank Bryan for the opening aspects of this Secrets Management section, @MichaelCade1 here to wrap things up and get some hands-on scenarios so that we can get some practical touch points and scenarios with HashiCorp Vault. 

A lot of what I will cover here can be found in existing [HashiCorp Tutorials](https://developer.hashicorp.com/vault) These resources from HashiCorp will go into a lot more different scenarios as well. 

## HashiCorp Vault CLI 

The first thing I want to walk through is getting the HashiCorp Vault CLI installed on our local machine. We have a great walkthrough from [HashiCorp on getting the CLI installed on your local machine.](https://developer.hashicorp.com/vault/tutorials/getting-started/getting-started-install) 

I am on my MacOS device but other instructions are also shown in the above link. 

`brew tap hashicorp/tap`

`brew install hashicorp/tap/vault`

Now when you run the `vault` command you should have access to the vault binary and sub commands. 

![](images\day39-1.png)

## Getting Vault up and running on Kubernetes

If you have been following along to any other sections in this project you will likely have seen that I am a big fan of Minikube when it comes to getting a local Kubernetes cluster up and running for learning or local development purposes. 

Vault is not exclusive in anyway to Kubernetes but I wanted to cover this scenario, HashiCorp also have a cloud SaaS option for Vault but it can be deployed many places. 

I use the below command to get my minikube up and running, you could just use `minikube start` but this is my go to for other demonstrations. 

`minikube start --addons volumesnapshots,csi-hostpath-driver --apiserver-port=6443 --container-runtime=containerd -p demo --kubernetes-version=1.26.0`

Now we have our cluster hopefully up and running, we will now use helm to deploy vault to our cluster. A quick google for helm installation steps will provide you with the how to make that happen. If you are on macos though then you can use homebrew to install 

`brew install helm`

Add the Hashicorp helm repo. 

`helm repo add hashicorp https://helm.releases.hashicorp.com`

There are several options and charts available to us within the repository, by running `helm search repo hashicorp` you will find the available hashicorp products and charts. 

![](images\day39-2.png)

For the rest of this walkthrough we will be taking this [tutorial](https://developer.hashicorp.com/vault/tutorials/kubernetes/kubernetes-minikube-raft) However we are going to be deploying vault to a dedicated namespace vs the default namespace. 

We will first create this file using the following command. 

```
cat > helm-vault-raft-values.yml <<EOF
server:
  affinity: ""
  ha:
    enabled: true
    raft: 
      enabled: true
EOF
```

I have added this file to a folder called day39 in the repository. 

We will use the following command now to deploy vault using helm. 

`helm install vault hashicorp/vault --namespace vault --values helm-vault-raft-values.yml --create-namespace`

By running `kubectl get pods -n vault` you should see the below output, it is expected to see the 3 pods in a 0/1 running state we are going to work through this next. 

![](images/day39-3.png)


Next we will initialise vault-0 with the following command, 

***"This command generates a root key that it disassembles into key shares -key-shares=1 and then sets the number of key shares required to unseal Vault -key-threshold=1. These key shares are written to the output as unseal keys in JSON format"***

```
kubectl exec vault-0 -n vault -- vault operator init \
    -key-shares=1 \
    -key-threshold=1 \
    -format=json > cluster-keys.json
```

We can then display the unseal key with jq again on macos if you want to get that installed you can do this with `brew install jq`

On your local machine you will now see a file called cluster-keys.json. The command to display the key is: 

`jq -r ".unseal_keys_b64[]" cluster-keys.json`

The output should look like this, 

![](images/day39-4.png)

We will then create a variable based on that key 

`VAULT_UNSEAL_KEY=$(jq -r ".unseal_keys_b64[]" cluster-keys.json)`

We can then go ahead and unseal our vault-0 pod with the following command: 

`kubectl exec vault-0 -n vault -- vault operator unseal $VAULT_UNSEAL_KEY`

This should display 

![](images/day39-5.png)

At this stage if you run `kubectl get pods -n vault` you should now see the vault-0 pod in a 1/1 running state. 

We will now join our other two pods vault-1 and vault-2 to our raft cluster using the following command. 

```
kubectl exec -ti vault-1 -n vault -- vault operator raft join http://vault-0.vault-internal:8200

kubectl exec -ti vault-2 -n vault -- vault operator raft join http://vault-0.vault-internal:8200
```

We can now unseal the two vault pods mentioned above with the following command. 

```
kubectl exec vault-1 -n vault -- vault operator unseal $VAULT_UNSEAL_KEY

kubectl exec vault-2 -n vault -- vault operator unseal $VAULT_UNSEAL_KEY
```

Take another look at the `kubectl get pods -n vault` command. Below is an output of all the recent commands. 

![](images/day39-6.png)

## Enable Key-Value secret engine 

In order for us to enable the secret engine we need to use the root token. This was also exported to the cluster-keys.json file we saw and used earlier for our unseal keys. If not following along you can see this JSON file in the day39 folder. 

`jq -r ".root_token" cluster-keys.json`

We must now exec into our vault-0 pod to enable the secret engine. 

`kubectl exec --stdin=true --tty=true vault-0 -n vault -- /bin/sh`

We will now have to authenticate and login using `vault login` and provide the token we discovered with root_token in a previous step. 

![](images/day39-7.png)

We will now run the following commands the first will enable the secret engine and the second will create secret at the path.  

```
vault secrets enable -path=secret kv-v2

vault kv put secret/webapp/config username="static-user" password="static-password"
```

You can then verify with the following command 

`vault kv get secret/webapp/config` 

So far we have used our root token this root user can peform any operation at any path and as you can expect best practices states that we dont or should not use this account other than initial setup and configuration. 

You should still be in your vault-0 pod. We are going to enable the Kubernetes authentication method with the following command: 

`vault auth enable kubernetes`

***Vault accepts this service token from any client within the Kubernetes cluster. During authentication, Vault verifies that the service account token is valid by querying a configured Kubernetes endpoint.***

Next we need to configure the Kubernetes authentication method to use the location of the Kubernetes API.

```
vault write auth/kubernetes/config \
    kubernetes_host="https://$KUBERNETES_PORT_443_TCP_ADDR:443"
```

## Creating a Vault Policy 

For a client or application to access the secret data defined, at secret/webapp/config, requires that the read capability be granted for the path secret/data/webapp/config.

```
vault policy write webapp - <<EOF
path "secret/data/webapp/config" {
  capabilities = ["read"]
}
EOF
```

Our Application shortly will be defined as webapp 

With the following command we will create a kubernetes authentication role 

```
vault write auth/kubernetes/role/webapp \
        bound_service_account_names=vault \
        bound_service_account_namespaces=webapp \
        policies=webapp \
        ttl=24h
```

now `exit` from the vault-0 pod and back to the local machine. 

## Deploying our Application 

I am again going to be using the web app that is used in the HashiCorp tutorial mentioned earlier. 

We will create a deployment yaml that looks like the following. 

```
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp
  labels:
    app: webapp
spec:
  replicas: 1
  selector:
    matchLabels:
      app: webapp
  template:
    metadata:
      labels:
        app: webapp
    spec:
      serviceAccountName: vault
      containers:
        - name: app
          image: hashieducation/simple-vault-client:latest
          imagePullPolicy: Always
          env:
            - name: VAULT_ADDR
              value: 'http://vault.vault.svc.cluster.local:8200/'
            - name: JWT_PATH
              value: '/var/run/secrets/kubernetes.io/serviceaccount/token'
            - name: SERVICE_PORT
              value: '8080'
```

Create the webapp namespace 

`kubectl create ns webapp`

Our YAML consists of our simple web app and the service account. 

`kubectl create -f deployment-01-webapp.yml -n webapp`

I also want to note that the helm chart for vault will deploy 

You can check that the authentication has worked by checking pods in the webapp namespace, if they are not in a running state or not there at all then something is not right as this is communicating with vault to make sure that this service is running. 

Once the pod is running, we need to port forward our webapp 
Find the pod name and then port forward that. 
```
kubectl get pods -n webapp 
kubectl port-forward <PODNAME> -n webapp 8080:8080
```

