# Containers Vulnerability Scanning

[Yesterday](day25.md) we learned that vulnerability scanning is the process of scanning a network or system to identify any existing security vulnerabilities.
We also learned that Containers Vulnerability Scanning is a subset of Systems Vulnerability Scanning, e.g. we are only scanning the "containers" part of our system.

In [Day 14](day14.md) we learned what container image vulnerability scanning and how it makes us more secure.
Then in [Day 15](day15.md) we learned more about that and on Days [21](day21.md) and [22](day22.md) we learned how to integrate the scanning process into our CI/CD pipelines
so that it is automatic and enforced.

Today, we are going to look at other techniques of scanning and securing containers.
Vulnerability scanning is important, but is not a silver bullet and not a guarantee that you are secure.

There are a few reasons for that.

First, image scanning only shows you the list of _known_ vulnerabilities.
There might be many vulnerabilities which have not been discovered, but are still there and could be exploited.

Second, the security of our deployments depends not only on the image and number of vulnerabilities, but also on the way we deploy that image.
For example, if we deploy an insecure application on the open internet where everyone has access to it, or leave the default SSH port and password of our VM,
then it does not matter whether our container has vulnerabilities or not, because the attackers will use the other holes in our system to get in.

That is why today we are going to take a look at few other aspects of containers vulnerability scanning.

## Host Security

Containers run on hosts.

Docker containers run on hosts that have the Docker Daemon installed.
Same is true for containerd, podman, cri-o, and other container runtimes.

If your host is not secured, and someone manages to break it, they will probably have access to your containers and be able to start, stop, modify them, etc.

That is why it's important to secure the host and secure it well.

Securing VMs is a deep topic I will not go into today, but the most basic things you can do are:

- limit the visibility of the machine on the public network
- if possible use a Load Balancer to access your containers, and make the host machine not visible on the public internet
- close all unnecessary ports
- use strong password for SSH and RDP

In the bottom of the article I will link 2 articles from AWS and VMware about VM security.

## Network Security

Network security is another deep topic, which we will look into in better detail [tomorrow](day27.md).

At a minimum, you should not have network exposure you don't need.
E.g. if Container A does not need to make network calls to Container B, it should not be able to make this calls at a first place.

In Docker you can define [different network drivers](https://docs.docker.com/network/) that can help you with this.
In Kubernetes there are [network policies](https://kubernetes.io/docs/concepts/services-networking/network-policies/) that limit which container has access to what.

## Security misconfiguration

When working with containers, there are a few security misconfiguration which you can make that can put you in danger of being hacked.

### Capabilities

One such thing is giving your container excessive capabilities.

[Linux capabilities](https://man7.org/linux/man-pages/man7/capabilities.7.html) determine what syscalls you container can execute.

The best practice is to be aware of the capabilities your containers need and assign them only them.
That way you will be sure that a left-over capability that was never needed was not abused by an attacker.

In practice, it is hard to know what capabilities exactly your containers need, because that involves complex monitoring of your container over time.
Even the developers that wrote the code are probably not aware of what capabilities exactly are needed to perform the actions that they code is doing.
That is so, because capabilities are a low-level construct and developers usually write higher-level code.

However, it is good to know which capabilities you should avoid assigning to your containers, because they are too overpowered and give it too many permissions.

One such capability is `CAP_SYS_ADMIN` which is way overpowered and can do a lot of things.
Even the Linux docs of this capability warn you that you should not be using this capability if you can avoid it.

### Running as root

Running containers as root is a really bad practice and it should be avoided as much as possible.

Of course, there might be situations in which you _must_ run containers as root.
One such example are the core components of Kubernetes, which run as root containers, because they need to have a lot of priviledges on the host.

However, if you are running a simple web server, or something like this, you should not have the need to run the container as root.

Running a container as root means that basically you are throwing away all the isolation containers give you, as a root container have almost full control over the host.

A lot of container runtime vulnerabilities are only applicable if containers are running as root.

Tools like [falco](https://github.com/falcosecurity/falco) and [kube-bench](https://github.com/aquasecurity/kube-bench) will warn you if you are running containers as root, so that you can take actions and change that.

### Resource limits

Not defining resource limits for your containers can lead to a DDoS attack that brings down your whole infrastructure.

When you are being DDoS-ed the workload starts consuming more memory and CPU.
If that workload is a container with no limits, at some point it will drain all the available resources from the host and there will be none left for the other containers on that host.
At some point, the whole host might go down, which will lead to more pressure on your other hosts and can have a domino effect on your whole infra.

If you have sensible limits for your container, it will consume them, but the orchestrator would not give him more.
At some point, the container will die due to lack of resources, but nothing else will happen.
Your host and other containers will be safe.

## Summary

Containers Vulnerability Scanning is more than just scanning for CVEs.
It includes things like proper configuration, host security, network configuration, etc.

There is not one tool that can help with this, but there are open source solutions that you can combine to achieve the desired results.

Most of these lessons are useful no matter the orchestrator you are using.
You can be using Kubernetes, OpenShift, AWS ECS, Docker Compose, VMs with Docker, etc.
The basics are the same, and you should adapt them to the platform you are using.

Some orchestrators give you more features than others.
For example, Kubernetes has [dynamic admission controllers](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/) that lets you define custom checks for your resources.
As far as I am aware, Docker Compose does not have something like this, but if you know what you want to achieve it should not be difficult to write your own.

## Resources

[This article](https://sysdig.com/blog/container-security-best-practices/) by Sysdig contains many best practices for containers vulnerability scanning.

Some of them like container image scanning and Infrastructure-as-Code scanning we already mentioned in previous days.
It also includes other useful things like [Host scanning](https://sysdig.com/blog/vulnerability-assessment/#host), [real-time logging and monitoring](https://sysdig.com/blog/container-security-best-practices/#13) and [security misconfigurations](https://sysdig.com/blog/container-security-best-practices/#11).

More on VM security:

<https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-security.html>

<https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.security.doc/GUID-60025A18-8FCF-42D4-8E7A-BB6E14708787.html>
See you on [Day 27](day27.md).
