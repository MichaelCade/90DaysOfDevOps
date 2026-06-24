---
title: '#90DaysOfDevOps - Microsoft Azure Compute Models - Day 31'
published: false
description: 90DaysOfDevOps - Microsoft Azure Compute Models
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049040
---

## Microsoft Azure Compute Models

Following on from covering the basics around security models within Microsoft Azure yesterday today we are going to look into the various compute services available to us in Azure.

### Service Availability Options

This section is close to my heart given my role in Data Management. As with on-premises, it is critical to ensure the availability of your services.

- High Availability (Protection within a region)
- Disaster Recovery (Protection between regions)
- Backup (Recovery from a point in time)

Microsoft deploys multiple regions within a geopolitical boundary.

Two concepts with Azure for Service Availability. Both sets and zones.

Availability Sets - Provide resiliency within a datacenter

Availability Zones - Provide resiliency between data centres within a region.

### Virtual Machines

Most likely the starting point for anyone in the public cloud.

- Provides a VM from a variety of series and sizes with different capabilities (Sometimes an overwhelming) [Sizes for Virtual machines in Azure](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes)
- There are many different options and focuses for VMs from high performance, and low latency to high memory options VMs.
- We also have a burstable VM type which can be found under the B-Series. This is great for workloads where you can have a low CPU requirement for the most part but require that maybe once a month performance spike requirement.
- Virtual Machines are placed on a virtual network that can provide connectivity to any network.
- Windows and Linux guest OS support.
- There are also Azure-tuned kernels when it comes to specific Linux distributions. [Azure Tuned Kernals](https://docs.microsoft.com/en-us/azure/virtual-machines/linux/endorsed-distros#azure-tuned-kernels)

### Templating

I have mentioned before that everything behind or underneath Microsoft Azure is JSON.

There are several different management portals and consoles we can use to create our resources the preferred route is going to be via JSON templates.

Idempotent deployments in incremental or complete mode - i.e repeatable desired state.

There is a large selection of templates that can export deployed resource definitions. I like to think about this templating feature to something like AWS CloudFormation or could be Terraform for a multi-cloud option. We will cover Terraform more in the Infrastructure as code section.

### Scaling

Automatic scaling is a large feature of the Public Cloud, being able to spin down resources you are not using or spin up when you need them.

In Azure, we have something called Virtual Machine Scale Sets (VMSS) for IaaS. This enables the automatic creation and scale from a gold standard image based on schedules and metrics.

This is ideal for updating windows so that you can update your images and roll those out with the least impact.

Other services such as Azure App Services have auto-scaling built in.

### Containers

We have not covered containers as a use case and what and how they can and should be needed in our DevOps learning journey but we need to mention that Azure has some specific container-focused services to mention.

Azure Kubernetes Service (AKS) - Provides a managed Kubernetes solution, no need to worry about the control plane or management of the underpinning cluster management. More on Kubernetes also later on.

Azure Container Instances - Containers as a service with Per-Second Billing. Run an image and integrate it with your virtual network, no need for Container Orchestration.

Service Fabric - Has many capabilities but includes orchestration for container instances.

Azure also has the Container Registry which provides a private registry for Docker Images, Helm charts, OCI Artifacts and images. More on this again when we reach the containers section.

We should also mention that a lot of the container services may indeed also leverage containers under the hood but this is abstracted away from your requirement to manage.

These mentioned container-focused services we also find similar services in all other public clouds.

### Application Services

- Azure Application Services provides an application hosting solution that provides an easy method to establish services.
- Automatic Deployment and Scaling.
- Supports Windows & Linux-based solutions.
- Services run in an App Service Plan which has a type and size.
- Number of different services including web apps, API apps and mobile apps.
- Support for Deployment slots for reliable testing and promotion.

### Serverless Computing

Serverless for me is an exciting next step that I am extremely interested in learning more about.

The goal with serverless is that we only pay for the runtime of the function and do not have to have running virtual machines or PaaS applications running all the time. We simply run our function when we need it and then it goes away.

Azure Functions - Provides serverless code. If we remember back to our first look into the public cloud we will remember the abstraction layer of management, with serverless functions you are only going to be managing the code.

Event-Driven with massive scale, I have a plan to build something when I get some hands-on here hopefully later on.

Provides input and output binding to many Azure and 3rd Party Services.

Supports many different programming languages. (C#, NodeJS, Python, PHP, batch, bash, Golang and Rust. Or any Executable)

Azure Event Grid enables logic to be triggered from services and events.

Azure Logic App provides a graphical-based workflow and integration.

We can also look at Azure Batch which can run large-scale jobs on both Windows and Linux nodes with consistent management & scheduling.

## Resources

- [Hybrid Cloud and MultiCloud](https://www.youtube.com/watch?v=qkj5W98Xdvw)
- [Microsoft Azure Fundamentals](https://www.youtube.com/watch?v=NKEFWyqJ5XA&list=WL&index=130&t=12s)
- [Google Cloud Digital Leader Certification Course](https://www.youtube.com/watch?v=UGRDM86MBIQ&list=WL&index=131&t=10s)
- [AWS Basics for Beginners - Full Course](https://www.youtube.com/watch?v=ulprqHHWlng&t=5352s)

See you on [Day 32](day32.md)
