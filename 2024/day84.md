# Day 84 - Hacking Kubernetes For Beginners
[![Watch the video](thumbnails/day84.png)](https://www.youtube.com/watch?v=ZUHUEZKl0vc)

 In this scenario, we discussed various methods a potential attacker could use to compromise a Kubernetes cluster. Here's a summary of the attacks discussed and possible mitigation strategies:

1. Container Escaping: An attacker gains access to a container and tries to escape it to reach the worker node hosting that container. This can be mitigated by implementing security measures on the containers, such as limiting privileges and monitoring activity within them.

2. Image Poisoning: An attacker replaces a legitimate image in the registry with a malicious one containing a reverse shell or other malware. To prevent this, regularly scan images for any suspicious activities and compare their hash keys with known good ones. Replace any images that have been tampered with.

3. Kubernetes Administrator Access: An attacker gains access to the Kubernetes cluster as an administrator and uses it to jump into a host node with root privileges. To mitigate this, limit the number of users with administrative rights and closely monitor their actions for any suspicious activity.

In terms of resources for further learning, I recommend checking out "Hacking Kubernetes" and "Container Security" books, as well as resources on Kubernetes security, observability, and web application attack techniques. These resources will provide more in-depth information on various aspects of securing a Kubernetes cluster.
Let's summarize the key points from this presentation:

**Scenario 1: Attacking a Registry**

* An attacker can replace an image in a registry with a malicious version.
* This allows the attacker to gain access to any container that uses the poisoned image.

Mitigation strategy: The Blue Team can scan images for suspicious activity and compare hash keys to detect tampering.

**Scenario 2: Escaping a Container**

* A Kubernetes administrator can create a pod using a malicious image, which contains a reverse shell.
* The attacker can then use this reverse shell to gain access to any container that uses the poisoned image.

Mitigation strategy: The Blue Team can restrict administrative rights and closely monitor user activity to detect suspicious behavior.

**General Takeaways**

* Kubernetes clusters are vulnerable to various types of attacks, including registry poisoning, container escape, and host compromise.
* Defense is challenging because multiple attack vectors exist, and any successful attack can compromise the cluster.
* Key components that can be attacked include CRI-O, Kubefwd, Etcd, and the Host itself.

**Recommendations for Further Learning**

* "Hacking Kubernetes" book: Provides detailed examples of attacks and defenses.
* "Container Security": Offers in-depth information on container technology and security best practices.
* "Kubernetes Security and Observability": Provides insights into securing and monitoring Kubernetes clusters.
* "Exposing" series: Shows examples of web application attacks, which are relevant to hosting web applications in a Kubernetes cluster.
