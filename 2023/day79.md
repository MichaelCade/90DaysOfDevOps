## Day 79 - Comparing Different Service Meshes
> **Informational**
> *A comparison of Istio, Linkerd, Consul, AWS App Mesh and Cilium*

### Service Mesh Offerings
There are PLENTY of service mesh offerings out there. Some are highly proprietary while others are very open.
Here are offerings you should definitely look into:

Service Mesh | Open Source or Proprietary | Notes |
---|---|---|
Istio | Open Source | Widely adopted and abstracted
Linkerd | Open Source | Built by Buoyant
Consul | Open Source | Owned by Hashcorp, Cloud offering available
Kuma | Open Source | Maintained by Kong
Traefik Mesh | Open Source | Specialized Proxy
Open Service Mesh | Open Source | By Microsoft 
Gloo Mesh | Proprietary | Built by Solo.io ontop of Istio
AWS App Mesh | Proprietary | AWS specific services
OpenShift Service Mesh | Proprietary | Built by Redhad, based on Istio
Tanzu Service Mesh | Proprietary | SaaS based on Istio, built by VMware
Anthos Service Mesh | Proprietary | SaaS based on Istio, built by Google
Bouyant Cloud | Proprietary | SaaS based on Linkerd
Cilium Service Mesh | Open Source | Orginally a CNI


I'll quickly recap some of the key options I'll compare. This was taken from Day 1.

#### Istio 
Istio is an open-source service mesh built by Google, IBM, and Lyft, and currently actively developed on and maintained by companies such as Solo.io. It is based on the Envoy proxy which is adopted for the sidecar pattern. Istio offers a high degree of customization and extensibility with advanced traffic routing, observability, and security for microservices. A new mode of operation for sidecar-less service mesh, called Ambient Mesh, was launched in 2022.

#### AppMesh
AppMesh is a service mesh implementation that is proprietary to AWS but primarily focuses in on applications deployed to various AWS services such as ECS, EKS, EC2. Its tight-nit integration into the AWS ecosystem allows for quick onboarding of services into the mesh. 

#### Consul 
Consul is a serivce mesh offering from Hashicorp that also provides traffic routing, observability, and sercurity much like Istio does.

#### Linkerd
Linkerd is an open-source service mesh offering that is lightweight. Similar to Istio, it provides traffic management, observability, and security, using a similar architecture. Linkerd adopts a sidecar-pattern using a Rust-based proxy.

#### Cilium
Cilium is a Container Networking Interface that leverages eBPF to optimize packet processing using the Linux kernel. It offers some Service Mesh capabilities, and doesn't use the sidecar model. It proceeds to deploy a per-node instance of Envoy for any sort of Layer 7 processing of requests. 

### Comparsion Table

Feature | Istio | Linkerd | AppMesh | Consul | Cilium | 
---|---|---|---|---|---|
Current Version | 1.16.1 | 2.12 | N/A (it's AWS :D ) | 1.14.3 | 1.12
Project Creators | Google, Lyft, IBM, Solo | Buoyant | AWS | Hashicorp | Isovalent 
Service Proxy | Envoy, Rust-Proxy (experimental) | Linkerd2-proxy | Envoy | Interchangeable, Envoy default | Per-node Envoy
Ingress Capabilities | Yes via the Istio Ingress-Gateway | No; BYO | Yes via AWS | Envoy | Cilium-Based Ingress
Traffic Management (Load Balancing, Traffic Split) | Yes | Yes | Yes | Yes | Yes, but manual Envoy config required for traffic splits
Resiliency Capabilities (Circuit Breaking, Retries/Timeouts, Faults, Delays) | Yes | Yes, no Circuit Breaking or Delays | Yes, No Fault or Delays | Yes, No Fault or Delays | Circuit Breaking, Retries and Timeouts require manual Envoy configuration, no other resiliency capabilities
Monitoring | Access Logs, Kiali, Jaegar/Zikin, Grafana, Prometheus, LETS, OTEL | LETS, Prometheus, Grafana, OTEL | AWS X-RAY, and Cloud Watch provides these | Datadog, Jaegar, Zipkin, OpenTracing, OTEL, Honeycomb | Hubble, OTEL, Prometheus, Grafana
Security Capabilities (mTLS, External CA) | Yes | Yes | Yes | Yes | Yes, with Wireguard
Getting Started | Yes | Yes | Yes | Yes | Yes
Production Ready | Yes | Yes | Yes | Yes | Yes
Key Features | Sidecar and Sidecar-less, Wasm Extensibility, VM support, Multi-cloud Support, Data Plane extensions | Simplistic and non-invasive | Highly focused and tight integration into AWS Ecosystem | Tight integration into Nomad and Hashicorp Ecosystem | Usage of eBPF for enhanced packet processing, Cilium Control Plane used to manage Service Mesh, No sidecars
Limitations | Complex, learning curve | Strictly K8s, additional config for BYO Ingress | Limited to just AWS services | Storage tied to Consul and not K8s | Not a complete Service Mesh, requires manual configuration
Protocol Support (TCP, HTTP 1.1 & 2, gRPC) | Yes | Yes | Yes | Yes | Yes
Sidecar Modes | Sidecar and Sidecar-less | Sidecar | Sidecar | Sidecar | No sidecar
CNI Redirection | Istio CNI Plugin | linkerd-cni | ProxyConfiguration Required | Consul CNI | eBPF Kernel processing
Platform Support | K8s and VMs | K8s | EC2, EKS, ECS, Fargate, K8s on EC2 | K8s, Nomad, ECS, Lambda, VMs | K8s, VMs, Nomad
Multi-cluster Mesh | Yes | Yes | Yes, only AWS | Yes | Yes
Governance and Oversight | Istio Community | Linkered Community | AWS | Hashicorp | Cilium Community

### Conclusion 
Service Meshes have come a long way in terms of capabilities and the environments they support. Istio appears to be the most feature-complete service mesh, providing a balance of platform support, customizability, extensibility, and is most production ready. Linkered trails right behind with a lighter-weight approach, and is mostly complete as a service mesh. AppMesh is mostly feature-filled but specific to the AWS Ecosystem. Consul is a great contender to Istio and Linkered. The Cilium CNI is taking the approach of using eBPF and climbing up the networking stack to address Service Mesh capabilities, but it has a lot of catching up to do.

Want to get deeper into Service Mesh? Head over to [#70DaysofServiceMesh](https://github.com/distributethe6ix/70DaysOfServiceMesh)!

See you on [Day 80](day80.md) of #70DaysOfServiceMesh! 