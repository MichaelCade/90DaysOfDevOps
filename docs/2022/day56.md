---
title: '#90DaysOfDevOps - The Big Picture: IaC - Day 56'
published: false
description: 90DaysOfDevOps - The Big Picture IaC
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048709
---

## The Big Picture: IaC

Humans make mistakes! Automation is the way to go!

How do you build your systems today?

What would be your plan if you were to lose everything today, Physical machines, Virtual Machines, Cloud VMs, Cloud PaaS etc etc.?

How long would it take you to replace everything?

Infrastructure as code provides a solution to be able to do this whilst also being able to test this, we should not confuse this with backup and recovery but in terms of your infrastructure and environments, your platforms we should be able to spin them up and treat them as cattle vs pets.

The TLDR; is that we can use code to rebuild our entire environment.

If we also remember from the start we said about DevOps in general is a way in which to break down barriers to deliver systems into production safely and rapidly.

Infrastructure as code helps us deliver the systems, we have spoken a lot of processes and tools. IaC brings us more tools to be familiar with to enable this part of the process.

We are going to concentrate on Infrastructure as code in this section. You might also hear this mentioned as Infrastructure from code or configuration as code. I think the most well-known term is likely Infrastructure as code.

### Pets vs Cattle

If we take a look at pre-DevOps, if we had the requirement to build a new Application, we would need to prepare our servers manually for the most part.

- Deploy VMs | Physical Servers and install the operating system
- Configure networking
- Create routing tables
- Install software and updates
- Configure software
- Install database

This would be a manual process performed by Systems Administrators. The bigger the application the more resource and servers required the more manual effort it would take to bring up those systems. This would take a huge amount of human effort and time but also as a business you would have to pay for that resource to build out this environment. As I opened the section with "Humans make mistakes! Automation is the way to go!"

Ongoing from the above initial setup phase you then have maintenance of these servers.

- Update versions
- Deploy new releases
- Data Management
- Recovery of Applications
- Add, Remove and Scale Servers
- Network Configuration

Add the complexity of multiple test and dev environments.

This is where Infrastructure as Code comes in, the above was very much a time where we would look after those servers as if they were pets, people even called them servers pet names or at least named them something because they were going to be around for a while, they were going to hopefully be part of the "family" for a while.

With Infrastructure as Code, we can automate all these tasks end to end. Infrastructure as code is a concept and some tools carry out this automated provisioning of infrastructure, at this point if something bad happens to a server you throw it away and you spin up a new one. This process is automated and the server is exactly as defined in the code. At this point we don't care what they are called they are there in the field serving their purpose until they are no longer in the field and we have another to replace it either because of a failure or because we updated part or all of our application.

This can be used in almost all platforms, virtualisation, cloud-based workloads and also cloud-native infrastructures such as Kubernetes and containers.

### Infrastructure Provisioning

Not all IaC cover all of the below, You will find that the tool we are going to be using during this section only really covers the first 2 areas below; Terraform is that tool we will be covering and this allows us to start from nothing and define in code what our infrastructure should look like and then deploy that, it will also enable us to manage that infrastructure and also initially deploy an application but at that point it is going to lose track of the application which is where the next section comes in and something like Ansible as a configuration management tool might work better on that front.

Without jumping ahead tools like chef, puppet and ansible are best suited to deal with the initial application setup and then manage those applications and their configuration.

Initial installation & configuration of software

- Spinning up new servers
- Network configuration
- Creating load balancers
- Configuration on the infrastructure level

### Configuration of provisioned infrastructure

- Installing applications on servers
- Prepare the servers to deploy your application.

### Deployment of Application

- Deploy and Manage Application
- Maintain phase
- Software updates
- Reconfiguration

### Difference between IaC tools

Declarative vs procedural

Procedural

- Step-by-step instruction
- Create a server > Add a server > Make this change

Declarative

- declare the result
- 2 Servers

Mutable (pets) vs Immutable (cattle)

Mutable

- Change instead of replacing
- Generally long lived

Immutable

- Replace instead of change
- Possibly short-lived

This is really why we have lots of different options for Infrastructure as Code because there is no one tool to rule them all.

We are going to be mostly using terraform and getting hands-on as this is the best way to start seeing the benefits of Infrastructure as Code when it is in action. Getting hands-on is also the best way to pick up the skills you are going to be writing code.

Next up we will start looking into Terraform with a 101 before we get some hands-on getting used.

## Resources

I have listed a lot of resources down below and I think this topic has been covered so many times out there, If you have additional resources be sure to raise a PR with your resources and I will be happy to review and add them to the list.

- [What is Infrastructure as Code? Difference of Infrastructure as Code Tools](https://www.youtube.com/watch?v=POPP2WTJ8es)
- [Terraform Tutorial | Terraform Course Overview 2021](https://www.youtube.com/watch?v=m3cKkYXl-8o)
- [Terraform explained in 15 mins | Terraform Tutorial for Beginners](https://www.youtube.com/watch?v=l5k1ai_GBDE)
- [Terraform Course - From BEGINNER to PRO!](https://www.youtube.com/watch?v=7xngnjfIlK4&list=WL&index=141&t=16s)
- [HashiCorp Terraform Associate Certification Course](https://www.youtube.com/watch?v=V4waklkBC38&list=WL&index=55&t=111s)
- [Terraform Full Course for Beginners](https://www.youtube.com/watch?v=EJ3N-hhiWv0&list=WL&index=39&t=27s)
- [KodeKloud - Terraform for DevOps Beginners + Labs: Complete Step by Step Guide!](https://www.youtube.com/watch?v=YcJ9IeukJL8&list=WL&index=16&t=11s)
- [Terraform Simple Projects](https://terraform.joshuajebaraj.com/)
- [Terraform Tutorial - The Best Project Ideas](https://www.youtube.com/watch?v=oA-pPa0vfks)
- [Awesome Terraform](https://github.com/shuaibiyy/awesome-terraform)

See you on [Day 57](day57.md)
