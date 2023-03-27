## Day 78 - Install and Test a Service Mesh
> **Tutorial**
> *Let's install our first service mesh.*

I spent enough time on some theory on Day 77. Let's dig right into getting a service mesh up and running.

### Tasks
- Install the Istio service mesh with Demo Profile
- Deploy Bookinfo Application into the service mesh
- Test it's functionality
- Externally expose the application

### Installation + Prequisites
I highly advise using something like minikube or a cloud-based K8s cluster that allows you to have load-balancer functionality.

- A Kubernetes cluster running 1.22, 1.23, 1.24, 1.25
    - KinD
    - Minikube
    - Civo K8s
    - EKS
- Access to a Loadbalancer service
    - Metallb
    - port-forwarding (not preferred)
    - Cloud Load-balancer
    - Inlets
- Linux or macOS to run istoctl

**Environment setup**
In my case, I spun up a Civo K3s cluster with 3-nodes, 2 CPU per node, and 4GB of RAM per node.
This is important because you will need enough resources to run the service mesh control plane, which, is Istiod in our case. If you need a cluster in a pinch register for free credit @ civo.com.

#### Install Istio
1. Verify your cluster is up and operational and make sure there aren't any errors. The commands below will output nodes and their IPs and OS info and the running pods in all namespaces, respectively.
    ```
    kubectl get nodes -o wide
    kubectl get pods -A
    ```
2. Download Istio, which will pick up the latest version (at the time of writing its 1.16.1)
    ```
    curl -L https://istio.io/downloadIstio | sh -
    ```
3. Change to the Istio directory
    ```
    cd istio-1.16.1
    ```
4. Add the istioctl binary to your path
    ```
    export PATH=$PWD/bin:$PATH
    ```
5. Check to verify that your environment is ready to go:
    ```
    istioctl x precheck
    ```
    And the output
    ```
    ✔ No issues found when checking the cluster. Istio is safe to install or upgrade!
    To get started, check out https://istio.io/latest/docs/setup/getting-started/    
    ```
6. Let's proceed to install Istio. We use the demo profile because it allows us to provide a minimal configuration to demonstrate functionality.
    ```
    istioctl install --set profile=demo -y
    ```
    You will see the following output:
    ```
    ✔ Istio core installed                                                                                                                                                          
    ✔ Istiod installed                                                                                                                                                              
    ✔ Egress gateways installed                                                                                                                                                     
    ✔ Ingress gateways installed                                                                                                                                                    
    ✔ Installation complete 
    ```
7. Next, we can verify all the components are deployed as expected by running the command below, which proceeds to output all the Istio components. 
    ```
    kubectl get all -n istio-system
    ```
    Your output should look similar in that all components are working. I changed my External-IP to *bring.your.LB.IP*, whcih means your IP will be different. Why do you need mine :P 
    ```
    NAME                                        READY   STATUS    RESTARTS   AGE
    pod/istiod-885df7bc9-f9k7c                  1/1     Running   0          31m
    pod/istio-ingressgateway-6688c7f65d-szxf9   1/1     Running   0          31m
    pod/istio-egressgateway-7475c75b68-mpxf7    1/1     Running   0          31m

    NAME                           TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)                                                                      AGE
    service/istiod                 ClusterIP      10.43.249.242   <none>         15010/TCP,15012/TCP,443/TCP,15014/TCP                                        31m
    service/istio-egressgateway    ClusterIP      10.43.75.47     <none>         80/TCP,443/TCP                                                               31m
    service/istio-ingressgateway   LoadBalancer   10.43.51.40     bring.your.LB.IP   15021:31000/TCP,80:32697/TCP,443:30834/TCP,31400:30953/TCP,15443:30733/TCP   31m

    NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
    deployment.apps/istiod                 1/1     1            1           31m
    deployment.apps/istio-ingressgateway   1/1     1            1           31m
    deployment.apps/istio-egressgateway    1/1     1            1           31m

    NAME                                              DESIRED   CURRENT   READY   AGE
    replicaset.apps/istiod-885df7bc9                  1         1         1       31m
    replicaset.apps/istio-ingressgateway-6688c7f65d   1         1         1       31m
    replicaset.apps/istio-egressgateway-7475c75b68    1         1         1       31m
    ```
#### Sidecar injection and Bookinfo deployment.
8. While everything looks good, we also want to deploy an application and simulataneously add it to the Service Mesh.
Let's label our default namespace with the *istio-injection=enabled* label. This tells Istiod to push the sidecar to any new microservice deployed to the namespace.
    ```
    kubectl label namespace default istio-injection=enabled
    ```
    Verify it
    ```
    kubectl get ns --show-labels
    ```
9. Let's deploy our app. Make sure you are the in the same directory as before. If not, change to *istio-1.16.1*
    ```
    kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml
    ```
    Let's give it one minute and check the application
    ```
    kubectl get all
    ```
    Let's see the output
    ```
    NAME                                 READY   STATUS    RESTARTS   AGE
    pod/details-v1-698b5d8c98-l8r5t      2/2     Running   0          7m30s
    pod/reviews-v3-6dc9897554-8pgtx      2/2     Running   0          7m30s
    pod/reviews-v1-9c6bb6658-lvzsr       2/2     Running   0          7m30s
    pod/reviews-v2-8454bb78d8-9tm2j      2/2     Running   0          7m30s
    pod/productpage-v1-bf4b489d8-q29jp   2/2     Running   0          7m29s
    pod/ratings-v1-5967f59c58-s9f88      2/2     Running   0          7m30s

    NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
    service/kubernetes    ClusterIP   10.43.0.1       <none>        443/TCP    80m
    service/details       ClusterIP   10.43.198.240   <none>        9080/TCP   7m31s
    service/ratings       ClusterIP   10.43.121.72    <none>        9080/TCP   7m30s
    service/reviews       ClusterIP   10.43.173.247   <none>        9080/TCP   7m30s
    service/productpage   ClusterIP   10.43.142.39    <none>        9080/TCP   7m30s

    NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
    deployment.apps/details-v1       1/1     1            1           7m30s
    deployment.apps/reviews-v3       1/1     1            1           7m30s
    deployment.apps/reviews-v1       1/1     1            1           7m30s
    deployment.apps/reviews-v2       1/1     1            1           7m30s
    deployment.apps/productpage-v1   1/1     1            1           7m30s
    deployment.apps/ratings-v1       1/1     1            1           7m30s

    NAME                                       DESIRED   CURRENT   READY   AGE
    replicaset.apps/details-v1-698b5d8c98      1         1         1       7m30s
    replicaset.apps/reviews-v3-6dc9897554      1         1         1       7m30s
    replicaset.apps/reviews-v1-9c6bb6658       1         1         1       7m30s
    replicaset.apps/reviews-v2-8454bb78d8      1         1         1       7m30s
    replicaset.apps/productpage-v1-bf4b489d8   1         1         1       7m30s
    replicaset.apps/ratings-v1-5967f59c58      1         1         1       7m30s
    ```
    The one thing to notice here is that all pods have *2/2* containers ready, meaning, the sidecar is now present.

#### Testing functionality
10. One test I'll run is to verify that I can connect to any one of these pods and get a response. Let's deploy a sleep pod. If you were in the same *istio-1.16.1* directory, then you can run this command.
    ```
    kubectl apply -f samples/sleep/sleep.yaml
    ```
    Let's check to see that the pod is deployed to the default namespace AND has the sidecar injected:
    ```
    kubectl get pods 
    ```
    ```
    sleep-75bbc86479-xw5lw           2/2     Running   0          36s
    ```

11. I'm going to use the sleep pod to curl over to the product-page. Please use yours which has a different name, to curl.
    ```
    kubectl exec sleep-75bbc86479-xw5lw -c sleep -- curl -sS productpage:9080
    ```
    And we should see something similar to the below which tells us the app is working!:
    ```
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
    100  1683  100  1683    0     0  69202      0 --:--:-- --:--:-- --:--:-- 70125
    <!DOCTYPE html>
    <html>
    <head>
        <title>Simple Bookstore App</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="static/bootstrap/css/bootstrap.min.css">

    <!-- Optional theme -->
    <link rel="stylesheet" href="static/bootstrap/css/bootstrap-theme.min.css">

    </head>
    <body>
        
        
    <p>
        <h3>Hello! This is a simple bookstore application consisting of three services as shown below</h3>
    </p>
    ```

#### Externally exposing the app

12. Let's deploy the Istio ingress gateway configuration and virtual-service configuration. The Ingress Gateway Configuration is an abstraction for the Istio-ingress-gateway and simply tells it which host and port is listening for requests. The virtual-service configuration tells the gateway how to route the request (or who to route it to). If you didn't navigate away from the original folder, run the following command:
    ```
    kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
    ```
    You will see two objects created:
    ```
    gateway.networking.istio.io/bookinfo-gateway created
    virtualservice.networking.istio.io/bookinfo created
    ```

13. Because we don't have external DNS setup, we'll gather the information we need and piece together our URL:
    Grab the External IP, and the HTTP2 port, which would be TCP port 80 from the Istio Ingress Gateway.
    ```
    kubectl get svc -n istio-system
    ```
    Store it in a variable
    ```
    export INGRESS_HOST=[IP_from_Istio_ingress_gateway]
    export INGRESS_PORT=80
    ```
    now curl this host:
    ```
    curl INGRESS_HOST:INGRESS_PORT/productpage
    ```
    And the RESULT:
    ```
    !DOCTYPE html>
    <html>
    <head>
        <title>Simple Bookstore App</title>
    ```

### Conclusion
I decided to jump into getting a service mesh up and online. It's easy enough if you have the right pieces in place, like a Kubernetes cluster and a load-balancer service. Using the demo profile, you can have Istiod, and the Ingress/Egress gateway deployed. Deploy a sample app with a service definition, and you can expose it via the Ingress-Gateway and route to it using a virtual service.

Want to get deeper into Service Mesh? Head over to [#70DaysofServiceMesh](https://github.com/distributethe6ix/70DaysOfServiceMesh)!

See you on [Day 79](day79.md) and beyond of #90DaysofServiceMesh
