# Application runtime policies

## Introduction 

Application runtime policies are a set of rules and controls that determine how an application behaves at runtime. These policies are implemented to ensure that an application behaves securely and within the constraints of its intended purpose.

Some common examples of runtime policies include access controls, and network restrictions. Access controls determine who has access to various parts of an application and what level of access they have. Application policies dictate what system resources an application can consume. Network restrictions can control which network resources an application can access and limit what kind of network traffic it can send or receive.

Runtime policies can be implemented in several ways, including through the use of security frameworks, application-specific configuration settings, or specialized security software. For example, some web application firewalls (WAFs) are designed to enforce runtime policies for web applications by analyzing incoming traffic and blocking or allowing requests based on predefined policies.

Effective implementation of application runtime policies can help mitigate security risks and ensure that an application behaves as intended. By monitoring and enforcing policies at runtime, organizations can reduce the likelihood of unauthorized access, data breaches, and other security incidents that can result in loss or damage to critical assets. Additionally, application runtime policies can help ensure that an application is performing optimally and can deliver its intended functionality without disruption or unexpected behavior.

Defining and maintaining application runtime policies can be a challenging task for a number of reasons.

1. Modern applications can be very complex, with multiple components and dependencies, making it difficult to define a clear set of policies that cover all possible scenarios. Adding to microservice architecture to this complexity is even increasing the challenge.

2. Rapidly changing applications and technology: Technology is constantly evolving, with new applications and platforms being developed and updated regularly. This means that policies that were once effective may quickly become outdated, requiring frequent updates.

3. Application policies must strike a balance between providing strong security measures and not hindering user productivity. This can be a delicate balance to achieve, as overly strict policies can make it difficult for users to perform their work, while overly permissive policies can leave the system vulnerable to security breaches.

4. Implementing and enforcing application policies can require significant resources, including time, money, and expertise. Organizations may face budgetary and staffing limitations that make it difficult to fully implement and maintain application policies.

In this session, we will see how application and network policies can be implemented and see an interesting approach: generating policies using monitoring application behavior.

## Kubernetes Pod security contexts

Let's start a new Minikube (in anticipation of the next part, we are already creating it with a CNI that implements network policies).
```bash
minikube start --cni cilium
```

Let's take the simple example of an Nginx web-server. We want to make it more secure than the default settings of the container. Kubernetes and container runtime gives an option to create `securityContext` configuration which limits the container runtime in different aspects:

### User and group ID of the container

Containers use user and group `0` (root) by default. Root on the host machine is not the same as root in the container. Defining the difference between these two is beyond our scope, but in short, the containerized root is confined by the container boundaries. Despite this, an attack that can penetrate a container that is running as root has much more attack surface for container escape than a container that runs as a non-root.

Therefore it is important to define a user ID different from than root in the `securityContext`. **Note:** you have to make sure that the container is built to run as non-root. Running `nginx:latest` as a non-root user will fail since the file permissions in the container image are built for root user.

Here is an example of creating a Nginx instance in Kubernetes which is not running as root.

```bash
kubectl apply -f - << EOF
apiVersion: v1
kind: Pod
metadata:
  name: nginx-non-root
spec:
  containers:
  - name: nginx
    image: nginxinc/nginx-unprivileged
    securityContext:
        runAsUser: 1000
        runAsGroup: 1000
        allowPrivilegeEscalation: false

EOF
```

Note the `runAsUser` and `runAsGroup` fields where we are limiting the user ID of who is running this container. There is an additional field that was set here called `allowPrivilegeEscalation`. It removes the capability of the processes running inside the container to escalate privileges using [sticky bit](https://en.wikipedia.org/wiki/Sticky_bit).


### Runtime system call policies

By default, every process is allowed to use any system calls of the kernel. The kernel might decide not to complete these system calls depending on its logic, but in general 500+ system calls are available for applications. 

If an attacker penetrated a container, she/he can try to use all these system calls to trick kernel to escape the container or do other damage. In practice, containerized applications are using only a limited number of system calls out of the 500+ and usually, they are less prone to vulnerabilities. For this reason, it is a good defense to limit the system calls a containerized application can do.

The configuration in `securityContext` enables users to implement restrictions on a container using [seccomp](https://kubernetes.io/docs/tutorials/security/seccomp/). From version 1.25, Kubernetes has a default application seccomp profile called `RuntimeDefault`. It is a permissive restriction, but it is a good start. See here how to create a Pod with this policy:

```bash
kubectl apply -f - << EOF
apiVersion: v1
kind: Pod
metadata:
  name: nginx-seccomp-confined
spec:
  securityContext:
    seccompProfile:
        type: RuntimeDefault
  containers:
  - name: nginx
    image: nginxinc/nginx-unprivileged
    securityContext:
        runAsUser: 1000
        runAsGroup: 1000
        allowPrivilegeEscalation: false

EOF
```

Note, this enables 90% of all system calls to the application and limits a few. This won't break most applications but will limit the attacker to some extent. 

See later how to tailor this to your application.

## Kubernetes native network policies

By default, Pods in a Kubernetes cluster have no limits on network communication. Any Pod can talk to any other Pod in the "Pod network". Creating micro-segmentation around Pods is an important way to limit the "blast radius" of an attack: only enable network connections which are required by the application.

Let's create another Nginx deployment and service in the cluster with:
```bash
kubectl create deployment nginx --image=nginx
kubectl expose deployment nginx --port=80
```

Now let's test the connection from another Pod:
```bash
kubectl run curl --rm -ti --image=curlimages/curl:latest -- sh
```
Now you can use curl to test the connection from the new Pod:
```bash
curl nginx
```

Now let's apply a network policy that only enables access to Nginx Pod from other Pods that are marked "nginx:client" label

```bash
kubectl apply -f - << EOF
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: nginx-client-access
spec:
  podSelector:
    matchLabels:
      app: nginx
  ingress:
  - from:
    - podSelector:
        matchLabels:
          nginx: "client"
EOF
```

Running the same test as above will fail:
```bash
kubectl run curl --rm -ti --image=curlimages/curl:latest -- sh
```
This should timeout:
```bash
curl --connect-timeout 1 nginx
```

If we want it to work, we need to add the label "nginx: client" to the curl Pod:
```bash
kubectl run curl --rm -ti --labels="nginx=client" --image=curlimages/curl:latest -- sh
```
Now the request will succeed:
```bash
curl --connect-timeout 1 nginx
```

🆒 😄

This is a simple example of how Kubernetes native network policies working.


## Generating policies from application behavior

As it was discussed in the intro, there is some considerable complexity in defining these policies.

On one hand, it takes time to define them properly and with changes, these policies tend to break applications. This causes practitioners to define lenient policies. 

On the other hand, if they are not defined strictly enough, they are less effective in protecting your systems.

There is hope though 😉

Newer technologies are striving to monitor application behavior (both network and runtime) and turn them automatically to policies.

A great example is [Inspektor Gadget](https://www.inspektor-gadget.io/).

We will see here how to install it and see how it generates network and seccomp profiles.

You can install the controller of Inspektor Gadget using [krew](https://krew.sigs.k8s.io/).
```bash
kubectl krew install gadget
```
and install Gadgets with
```bash
kubectl gadget deploy
```

Now you can start monitoring the above Nginx Pod and generating seccomp profile for it:
```bash
kubectl gadget advise seccomp-profile start -n default -p $(kubectl get pods | grep nginx | head -n 1 | awk '{print $1}')
```

This command started monitoring and returns a trace ID, if you think that you got enough activity stop the tracing with:
```bash
kubectl gadget advise seccomp-profile stop <traceid>
```

Example:
```bash
$ kubectl gadget advise seccomp-profile stop jd4VM2jWhnONfakF
{
  "defaultAction": "SCMP_ACT_ERRNO",
  "architectures": [
    "SCMP_ARCH_X86_64",
    "SCMP_ARCH_X86",
    "SCMP_ARCH_X32"
  ],
  "syscalls": [
    {
      "names": [
        "access",
        "arch_prctl",
        "brk",
        "chown",
        "clone",
        "close",
        "connect",
        "dup",
        "dup2",
        "execve",
        "exit_group",
        "faccessat",
        "fcntl",
        "fstat",
        "futex",
        "getcwd",
        "getdents64",
        "getegid",
        "geteuid",
        "getgid",
        "getpgrp",
        "getpid",
        "getppid",
        "getuid",
        "ioctl",
        "lseek",
        "mmap",
        "mprotect",
        "munmap",
        "openat",
        "pipe",
        "prlimit64",
        "pselect6",
        "read",
        "rt_sigaction",
        "rt_sigprocmask",
        "rt_sigreturn",
        "select",
        "set_robust_list",
        "set_tid_address",
        "setns",
        "setpgid",
        "socket",
        "stat",
        "statfs",
        "sysinfo",
        "uname",
        "wait4",
        "write"
      ],
      "action": "SCMP_ACT_ALLOW"
    }
  ]
}
```

The same tool helps you to generate network policy. Let's start network monitoring with this command:
```bash
kubectl gadget advise network-policy monitor -p $(kubectl get pods | grep nginx | head -n 1 | awk '{print $1}') --output /tmp/network.log
```

When you think the monitoring have seen enough activity, you can stop with `ctrl-c`. Then generate the policy yaml with this command:
```bash
kubectl gadget advise network-policy report --input /tmp/network.log 
```

This is the policy you get:
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  creationTimestamp: null
  name: nginx-network
  namespace: default
spec:
  ingress:
  - from:
    - podSelector: {}
    ports:
    - port: 80
      protocol: TCP
  podSelector:
    matchLabels:
      app: nginx
  policyTypes:
  - Ingress
  - Egress
status: {}
```
### Summary
These were examples of how to turn behavior to policy! Good stuff 😃

See you on [Day 34](day34.md).
