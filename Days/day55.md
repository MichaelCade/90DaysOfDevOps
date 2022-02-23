## State and Ingress in Kubernetes
In this closing section of Kubernetes, we are going to take a look at State and ingress. 

Everything we have said so far is about stateless, stateless is really where our applications do not care which network it is using and does not need any permanent storage. Whereas stateful apps, databases for example for such an application to function correctly, you’ll need to ensure that pods can reach each other through a unique identity that does not change (hostnames, IPs...etc.). Examples of stateful applications include MySQL clusters, Redis, Kafka, MongoDB and others. Basically though any application that stores data. 

### Stateful Application 

StatefulSets represent a set of Pods with unique, persistent identities and stable hostnames that Kubernetes maintains regardless of where they are scheduled. The state information and other resilient data for any given StatefulSet Pod is maintained in persistent disk storage associated with the StatefulSet.





`kubectl port-forward svc/pacman 9090:80 -n pacman`


### Persistant Volumes | Claims 


### Ingress explained 


### LAMP Stack > Drupal 

