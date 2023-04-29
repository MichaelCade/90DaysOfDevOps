# Network security in runtime

Network security is crucial in the context of runtime protection as it aims to safeguard the communication channels between different components of an application or system. In a modern microservices-based architecture, where an application is comprised of several smaller services communicating with each other over the network, securing the network becomes even more critical.

One of the main reasons network security is essential in this context is to prevent attackers from exploiting network vulnerabilities, spreading in a system and gaining unauthorized access to sensitive data or systems. For instance, if an attacker is successful in intercepting the communication between two microservices, they can easily eavesdrop on the sensitive information exchanged between them or even manipulate the data, potentially leading to a significant security breach.

Moreover, network security also helps in protecting against distributed denial-of-service (DDoS) attacks, where the attacker overwhelms the network with a large number of requests or traffic, making the application or system unavailable. By implementing network security controls such as firewalls, intrusion detection/prevention systems (IDS/IPS), and load balancers, organizations can mitigate the risk of DDoS attacks and ensure the availability of their services.

Additionally, network security is essential in compliance with industry and regulatory standards, such as the Payment Card Industry Data Security Standard (PCI DSS), General Data Protection Regulation (GDPR), and Health Insurance Portability and Accountability Act (HIPAA). These standards mandate the protection of sensitive data during its transmission over the network, and organizations must implement appropriate security controls to comply with these standards.

## Network security basics
Network security is a broad field, but it's important to cover the basics so everyone is on the same page. This could include topics such as network topologies, network protocols (such as TCP/IP), firewalls, and access control mechanisms like ACLs and RBAC. You should also cover the OSI model, which is a conceptual model that describes how network protocols work. Understanding the OSI model is essential for troubleshooting network issues and implementing network security measures.

## Kubernetes network architecture
Kubernetes uses a complex network architecture to manage network traffic between Pods and Services. It's important to understand this architecture in order to secure it. At a high level, the Kubernetes network consists of a cluster-wide Pod network and a Service network. The Pod network allows Pods to communicate with each other, while the Service network provides a stable IP address and DNS name for a set of Pods. Kubernetes also supports Ingress, which is used to route external traffic into the cluster. Ingress is typically secured using TLS encryption and client authentication.

## Network policies
Kubernetes Network Policies allow you to control network traffic between Pods and Services. You can use Network Policies to limit the attack surface of your cluster, for example by blocking traffic from Pods that are not part of your application. Network Policies use selectors to specify which Pods or Services they apply to, and can be configured to allow or deny traffic based on a variety of criteria, such as IP addresses, port numbers, or protocols.

## Container networking
Containers are isolated from each other by default, but they can still communicate over the network. Container networking can be complex, especially when using container orchestration platforms like Kubernetes. It's important to understand how container networking works in order to secure it. Kubernetes uses a variety of networking plugins to enable communication between Pods, such as Calico, Weave Net, and Flannel. These plugins can be configured to use different networking modes, such as host networking, bridge networking, or overlay networking.

## Service mesh security
Service meshes like Istio and Linkerd are becoming increasingly popular in Kubernetes environments. Service meshes provide advanced traffic management features, such as load balancing, traffic shaping, and circuit breaking. They also provide security features such as mutual TLS authentication, RBAC, and traffic management. It's important to understand the principles of service mesh security, including how to configure and manage these features.

## Network monitoring and logging
Network monitoring and logging are critical for detecting and responding to security incidents. You should cover best practices for network monitoring and logging in Kubernetes, including the use of tools like Falco, Prometheus, and Grafana. Falco is a runtime security tool that can be used to detect suspicious behavior in Kubernetes, while Prometheus and Grafana can be used for monitoring and alerting.

## Zero trust networking
Zero trust networking is a security model that assumes all network traffic is untrusted, and requires authentication and authorization for every connection. Zero trust networking can be implemented in Kubernetes using tools like Istio, which provides mutual TLS authentication and RBAC features. You should cover the principles of zero trust networking, including how to implement it in Kubernetes, and how it can help to improve security in your environment.

# Kubernetes native network policies

Kubernetes native network policies are a mechanism for enforcing network segmentation and controlling the communication between pods in a Kubernetes cluster. Network policies allow administrators to define rules that specify which pods can communicate with each other based on their labels.

Network policies use labels to select pods and define rules to allow or deny traffic between them. The labels can be applied to the pods themselves or to their associated services. By default, Kubernetes denies all incoming and outgoing network traffic between pods, and network policies allow administrators to selectively open up communication between specific pods or groups of pods.

In the context of network security, ingress and egress policies are used to control traffic that enters and exits a network or a specific network segment. In Kubernetes, these policies are implemented using Network Policies.

### Ingress
An ingress policy specifies the rules that control incoming traffic to a network or network segment. These policies define what types of traffic are allowed to enter the network, from which sources, and through which ports. In Kubernetes, an ingress policy is applied to a specific set of pods or namespaces, and it can be used to allow or deny incoming traffic from other pods or external sources.


### Egress 
An egress policy, on the other hand, controls the outgoing traffic from a network or network segment. Egress policies are used to restrict what types of traffic are allowed to leave the network, and to which destinations. In Kubernetes, an egress policy is also applied to a specific set of pods or namespaces, and it can be used to allow or deny outgoing traffic to other pods or external destinations.

### Creating policies
To create a network policy, you first need to create a YAML file that defines the policy. The YAML file should specify the pods or pod selectors to which the policy applies, and the rules to allow or deny traffic between them. Here's an example YAML file for a network policy that allows traffic between pods with the label `app=frontend` and denies traffic to all other pods:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: frontend-policy
spec:
  podSelector:
    matchLabels:
      app: frontend
  policyTypes:
  - Ingress
  ingress:
  - from:
    - podSelector: {}
```

In this example, the `podSelector` section specifies that the policy applies to all pods with the label `app=frontend`. The `policyTypes` section specifies that the policy applies to incoming traffic (Ingress), and the ingress section specifies that all pods can receive traffic from pods that match the `podSelector`.

Once you've created the YAML file, you can apply the network policy to your Kubernetes cluster using the kubectl apply command:

```bash
kubectl apply -f <filename>.yaml
```

After applying the network policy, traffic will be restricted between pods based on the rules defined in the policy. Kubernetes network policies are enforced by network plugins, which may vary depending on the platform and environment in which Kubernetes is running. Some popular network plugins that support Kubernetes network policies include Calico, Cilium, and Weave Net. Note, that if you want to test this on Minikube, you will need to enable Calico CNI (`minikube start --cni calico`) since the default CNI implementation does not support it.

In summary, Kubernetes native network policies are a mechanism for enforcing network segmentation and controlling communication between pods in a Kubernetes cluster. Network policies use labels to select pods and define rules to allow or deny traffic between them. You can create a network policy by defining a YAML file that specifies the policy's rules, and then apply the policy to your cluster using the kubectl apply command.

# Network monitoring in Kubernetes

There are several tools available to monitor network traffic in Kubernetes, each with their own set of features and capabilities. Here are a few popular tools that you can use to monitor network traffic in Kubernetes:

* Cilium: Cilium is a powerful networking and security solution for Kubernetes that provides observability, security, and networking services for your Kubernetes cluster. Cilium can monitor network traffic at the container and pod level, and provide detailed insights into the traffic flows between pods and services. Cilium also provides network policies and can enforce them at the kernel level using eBPF technology.

* Istio: Istio is a service mesh for Kubernetes that provides traffic management, observability, and security features for your Kubernetes applications. Istio includes a powerful sidecar proxy called Envoy that can capture and monitor network traffic between your Kubernetes services, and provides rich observability features such as distributed tracing and metrics.

* Calico: Calico is a popular networking and security solution for Kubernetes that provides network policies, observability, and security features for your Kubernetes cluster. Calico includes a powerful network policy engine that allows you to define and enforce policies at the pod and namespace level, and provides detailed insights into network traffic flows.

* Weave Scope: Weave Scope is a Kubernetes monitoring and visualization tool that allows you to monitor your Kubernetes network and applications in real time. Weave Scope includes a network traffic view that allows you to visualize the traffic flows between your Kubernetes services and pods, and provides rich insights into your Kubernetes infrastructure.

These are just a few examples of the many tools available for monitoring network traffic in Kubernetes. When selecting a tool, consider the specific needs of your application and infrastructure, and choose a tool that provides the features and capabilities that best fit your requirements.

See you on [Day 32](day32.md).