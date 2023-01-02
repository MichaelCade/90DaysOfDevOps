## The Big Picture: DevSecOps

Welcome to Day 2 of the 2023 edition here in this first module of the next 6 days we are going look at the foundational overview around DevSecOps. 

### What is DevSecOps? 

DevSecOps is a software development approach that aims to bring together development, security, and operations teams to build and maintain secure software applications. It is based on the principles of continuous integration, continuous delivery, and continuous deployment, which aim to deliver software updates and features more quickly and frequently. In DevSecOps, security is an integral part of the software development process, rather than an afterthought. This means that security testing, monitoring, and other security measures are built into the software development life cycle (SDLC) from the beginning, rather than being added later. DevSecOps aims to improve collaboration and communication between development, security, and operations teams, to create a more efficient and effective software development process.

### DevSecOps vs DevOps 

I use the "vs" lightly here again but if we think back to the 2022 edition and the goal of DevOps is to improve the speed, reliability, and quality of software releases.

DevSecOps is an extension of the DevOps philosophy that emphasizes the integration of security practices into the software development process. The goal of DevSecOps is to build security measures into the software development process so that security is an integral part of the software from the start, rather than an afterthought. This helps to reduce the risk of security vulnerabilities being introduced into the software and makes it easier to identify and fix any issues that do arise.

DevOps focuses on improving collaboration and communication between developers and operations staff to improve the speed, reliability, and quality of software releases, while DevSecOps focuses on integrating security practices into the software development process to reduce the risk of security vulnerabilities and improve the overall security of the software.

### Automated Security 

Automated security refers to the use of technology to perform security tasks without the need for human intervention. This can include things like security software that monitors a network for threats and takes action to block them, or systems that use artificial intelligence to analyse security footage and identify unusual activity. Automated security systems are designed to make security processes more efficient and effective, and to help reduce the workload on security personnel.

A key component of all things DevSecOps is the ability to automate a lot of the tasks at hand when creating and delivering software, when we add security from the start it means we also need to consider the automation aspect of security. 

### Security at Scale (Containers and Microservices)

We know that the scale and dynamic infrastructure that has been enabled by containerisation and microservices have changed the way that most organisations do business. 

This is also why we must bring that automated security into our DevOps principles to ensure that specific container security guidelines are met. 

What I mean by this is with cloud-native technologies we cannot only have static security policies and posture; our security model also must be dynamic with the workload in hand and how that is running. 

DevOps teams will need to include automated security to protect the overall environment and data, as well as continuous integration and continuous delivery processes. 

The below list is taken from a [RedHat blog post](https://www.redhat.com/en/topics/devops/what-is-devsecops)

- Standardise and automate the environment: Each service should have the least privilege possible to minimize unauthorized connections and access.

- Centralise user identity and access control capabilities: Tight access control and centralised authentication mechanisms are essential for securing microservices since authentication is initiated at multiple points.

- Isolate containers running microservices from each other and the network: This includes both in-transit and at-rest data since both can represent high-value targets for attackers.

- Encrypt data between apps and services: A container orchestration platform with integrated security features helps minimize the chance of unauthorized access.

- Introduce secure API gateways: Secure APIs increase authorization and routing visibility. By reducing exposed APIs, organizations can reduce surfaces of attacks.

### Security is HOT right now

One thing you will have seen regardless of your background is that security is hot all over the industry, this is partly to do with security breaches appearing in global news and big brands being affected by security vulnerabilities or following potential bad practices allowing bad actors into the networks of these companies. It is fair to say or at least from my perspective the creation of software is much more achievable and obtainable now than it ever has. But in creating software it is increasingly exposed with vulnerabilities and the like which allows the bad actors to cause havoc and sometimes hold data to ransom or shut down businesses causing mayhem. We have discussed so far what is DevSecOps but I think it is also worthwhile exploring the cybersecurity side of the attack vector and why we protect our software supply chain to help avoid these cyber-attacks. 

### Cybersecurity vs DevSecOps

As the heading goes it is not really a vs but more of a difference between the two topics. But I think it is important to raise this as really this will explain why Security must be part of that DevOps process, principles, and methodology. 

Cybersecurity is the practice of protecting computer systems and networks from digital attacks, theft, and damage. It involves identifying and addressing vulnerabilities, implementing security measures, and monitoring systems for threats.

DevSecOps, on the other hand, is a combination of development, security, and operations practices. It is a philosophy that aims to integrate security into the development process, rather than treating it as a separate step. This involves collaboration between development, security, and operations teams throughout the entire software development lifecycle (SDLC).

Some key differences between cybersecurity and DevSecOps include:

**Focus**: Cybersecurity is primarily focused on protecting systems from external threats, while DevSecOps focuses on integrating security into the development process.

**Scope**: Cybersecurity covers a wider range of topics, including network security, data security, application security, and more. DevSecOps, on the other hand, is specifically focused on improving the security of software development and deployment.

**Approach**: Cybersecurity typically involves implementing security measures after the development process is complete, while DevSecOps involves integrating security into the development process from the start.

**Collaboration**: Cybersecurity often involves collaboration between IT and security teams, while DevSecOps involves collaboration between development, security, and operations teams.

## Resources 

Over the course of the 90 Days, we will have a daily resources list that will bring relevant content that will help continue the topics and where you can go to find out more. 

- [TechWorld with Nana - What is DevSecOps? DevSecOps explained in 8 Mins](https://www.youtube.com/watch?v=nrhxNNH5lt0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=1&t=19s)

- [What is DevSecOps?](https://www.youtube.com/watch?v=J73MELGF6u0&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=2&t=1s)

- [freeCodeCamp.org - Web App Vulnerabilities - DevSecOps Course for Beginners](https://www.youtube.com/watch?v=F5KJVuii0Yw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=3&t=67s)

- [The Importance of DevSecOps and 5 Steps to Doing it Properly (DevSecOps EXPLAINED)](https://www.youtube.com/watch?v=KaoPQLyWq_g&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=4&t=13s)

- [Continuous Delivery - What is DevSecOps?](https://www.youtube.com/watch?v=NdvMUcWNlFw&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=5&t=6s)

- [Cloud Advocate - What is DevSecOps?](https://www.youtube.com/watch?v=a2y4Oj5wrZg&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=6)

- [Cloud Advocate - DevSecOps Pipeline CI Process - Real world example!](https://www.youtube.com/watch?v=ipe08lFQZU8&list=PLsKoqAvws1pvg7qL7u28_OWfXwqkI3dQ1&index=7&t=204s)

Hopefully this gave you a taster for what you can expect for this module and some of the resources above will help provide more depth on the topic, In the post on [Day 3](day03.md) we will be taking a look at what an attacker thinks like which is why we have to protect from the start.