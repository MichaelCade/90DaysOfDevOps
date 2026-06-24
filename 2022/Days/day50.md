---
title: '#90DaysOfDevOps - Choosing your Kubernetes platform - Day 50'
published: false
description: 90DaysOfDevOps - Choosing your Kubernetes platform
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049046
---

## Choosing your Kubernetes platform

I wanted to use this session to break down some of the platforms or maybe distributions is a better term to use here, one thing that has been a challenge in the Kubernetes world is removing complexity.

Kubernetes the hard way walks through how to build out from nothing to a full-blown functional Kubernetes cluster this is to the extreme but more and more at least the people I am speaking to are wanting to remove that complexity and run a managed Kubernetes cluster. The issue there is that it costs more money but the benefits could be if you use a managed service do you need to know the underpinning node architecture and what is happening from a Control Plane node point of view when generally you do not have access to this.

Then we have the local development distributions that enable us to use our systems and run a local version of Kubernetes so developers can have the full working environment to run their apps in the platform they are intended for.

The general basis of all of these concepts is that they are all a flavour of Kubernetes which means we should be able to freely migrate and move our workloads where we need them to suit our requirements.

A lot of our choice will also depend on what investments have been made. I mentioned the developer experience as well but some of those local Kubernetes environments that run our laptops are great for getting to grips with the technology without spending any money.

### Bare-Metal Clusters

An option for many could be running your Linux OS straight onto several physical servers to create our cluster, it could also be Windows but I have not heard much about the adoption rate around Windows, Containers and Kubernetes. If you are a business and you have made a CAPEX decision to buy your physical servers then this might be how you go when building out your Kubernetes cluster, the management and admin side here means you are going to have to build yourself and manage everything from the ground up.

### Virtualisation

Regardless of test and learning environments or enterprise-ready Kubernetes clusters virtualisation is a great way to go, typically the ability to spin up virtual machines to act as your nodes and then cluster those together. You have the underpinning architecture, efficiency and speed of virtualisation as well as leveraging that existing spend. VMware for example offers a great solution for both Virtual Machines and Kubernetes in various flavours.

My first ever Kubernetes cluster was built based on Virtualisation using Microsoft Hyper-V on an old server that I had which was capable of running a few VMs as my nodes.

### Local Desktop options

There are several options when it comes to running a local Kubernetes cluster on your desktop or laptop. This as previously said gives developers the ability to see what their app will look like without having to have multiple costly or complex clusters. Personally, this has been one that I have used a lot and in particular, I have been using minikube. It has some great functionality and adds-ons which changes the way you get something up and running.

### Kubernetes Managed Services

I have mentioned virtualisation, and this can be achieved with hypervisors locally but we know from previous sections we could also leverage VMs in the public cloud to act as our nodes. What I am talking about here with Kubernetes managed services are the offerings we see from the large hyperscalers but also from MSPs removing layers of management and control away from the end user, this could be removing the control plane from the end user this is what happens with Amazon EKS, Microsoft AKS and Google Kubernetes Engine. (GKE)

### Overwhelming choice

I mean the choice is great but there is a point where things become overwhelming and this is not a depth look into all options within each category listed above. On top of the above, we also have OpenShift which is from Red Hat and this option can be run across the options above in all the major cloud providers and probably today gives the best overall useability to the admins regardless of where clusters are deployed.

So where do you start from your learning perspective, as I said I started with the virtualisation route but that was because I had access to a physical server which I could use for the purpose, I appreciate and in fact, since then I no longer have this option.

My actual advice now would be to use Minikube as a first option or Kind (Kubernetes in Docker) but Minikube gives us some additional benefits which almost abstracts the complexity out as we can just use add-ons and get things built out quickly and we can then blow it away when we are finished, we can run multiple clusters, we can run it almost anywhere, cross-platform and hardware agnostic.

I have been through a bit of a journey with my learning around Kubernetes so I am going to leave the platform choice and specifics here to list out the options that I have tried to give me a better understanding of Kubernetes the platform and where it can run. What I might do with the below blog posts is take another look at these update them and bring them more into here vs them being linked to blog posts.

- [Kubernetes playground – How to choose your platform](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-1)
- [Kubernetes playground – Setting up your cluster](https://vzilla.co.uk/vzilla-blog/building-the-home-lab-kubernetes-playground-part-2)
- [Getting started with Amazon Elastic Kubernetes Service (Amazon EKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-amazon-elastic-kubernetes-service-amazon-eks)
- [Getting started with Microsoft Azure Kubernetes Service (AKS)](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-azure-kubernetes-service-aks)
- [Getting Started with Microsoft AKS – Azure PowerShell Edition](https://vzilla.co.uk/vzilla-blog/getting-started-with-microsoft-aks-azure-powershell-edition)
- [Getting started with Google Kubernetes Service (GKE)](https://vzilla.co.uk/vzilla-blog/getting-started-with-google-kubernetes-service-gke)
- [Kubernetes, How to – AWS Bottlerocket + Amazon EKS](https://vzilla.co.uk/vzilla-blog/kubernetes-how-to-aws-bottlerocket-amazon-eks)
- [Getting started with CIVO Cloud](https://vzilla.co.uk/vzilla-blog/getting-started-with-civo-cloud)
- [Minikube - Kubernetes Demo Environment For Everyone](https://vzilla.co.uk/vzilla-blog/project_pace-kasten-k10-demo-environment-for-everyone)

### What we will cover in the series on Kubernetes

- Kubernetes Architecture
- Kubectl Commands
- Kubernetes YAML
- Kubernetes Ingress
- Kubernetes Services
- Helm Package Manager
- Persistent Storage
- Stateful Apps

## Resources

- [Kubernetes Documentation](https://kubernetes.io/docs/home/)
- [TechWorld with Nana - Kubernetes Tutorial for Beginners [FULL COURSE in 4 Hours]](https://www.youtube.com/watch?v=X48VuDVv0do)
- [TechWorld with Nana - Kubernetes Crash Course for Absolute Beginners](https://www.youtube.com/watch?v=s_o8dwzRlu4)
- [Kunal Kushwaha - Kubernetes Tutorial for Beginners | What is Kubernetes? Architecture Simplified!](https://www.youtube.com/watch?v=KVBON1lA9N8)

See you on [Day 51](day51.md)
