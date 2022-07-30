---
title: '#90DaysOfDevOps - Network Automation - Day 24'
published: false
description: 90DaysOfDevOps - Network Automation
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048805
---

## Network Automation

### Basics of network automation

Primary drivers for Network Automation

- Achieve Agility
- Reduce Cost
- Eliminate Errors
- Ensure Compliance
- Centralised Management

The automation adoption process is specific to each business. There's no one size fits all when it comes to deploying automation, the ability to identify and embrace the approach that works best for your organisation is critical in advancing towards maintaining or creating a more agile environment, the focus should always be on business value and end-user experience. (We said something similar right at the start in regards to the whole of DevOps and the culture change and the automated process that this brings)

To break this down you would need to identify how the task or process that you're trying to automate is going to achieve and improve the end-user experience or business value whilst following a step-by-step systematic approach.

"If you don't know where you are going, any road will take you there."

Have a framework or design structure that you're trying to achieve know what your end goal is and then work step by step towards achieving that goal measuring the automation success at various stages based on the business outcomes.

Build concepts modelled around existing applications there's no need to design the concepts around automation in a bubble because they need to be applied to your application, your service, and your infrastructure, so begin to build the concepts and model them around your existing infrastructure, you're existing applications.

### Approach to Networking Automation

We should identify the tasks and perform a discovery on network change requests so that you have the most common issues and problems to automate a solution to.

- Make a list of all the change requests and workflows that are currently being addressed manually.
- Determine the most common, time-consuming and error-prone activities.
- Prioritise the requests by taking a business-driven approach.
- This is the framework for building an automation process, what must be automated and what must not.

We should then divide tasks and analyse how different network functions work and interact with each other.

- The infrastructure/Network team receives change tickets at multiple layers to deploy applications.
- Based on Network services, divide them into different areas and understand how they interact with each other.
  - Application Optimisation
  - ADC (Application Delivery Controller)
  - Firewall
  - DDI (DNS, DHCP, IPAM etc)
  - Routing
  - Others
- Identify various dependencies to address business and cultural differences and bring in cross-team collaboration.

Reusable policies, define and simplify reusable service tasks, processes and input/outputs.

- Define offerings for various services, processes and input/outputs.
- Simplifying the deployment process will reduce the time to market for both new and existing workloads.
- Once you have a standard process, it can be sequenced and aligned to individual requests for a multi-threaded approach and delivery.

Combine the policies with business-specific activities. How does implementing this policy help the business? Saves time? Saves Money? Provides a better business outcome?

- Ensure that service tasks are interoperable.
- Associate the incremental service tasks so that they align to create business services.
- Allow for the flexibility to associate and re-associate service tasks on demand.
- Deploy Self-Service capabilities and pave the way for improved operational efficiency.
- Allow for the multiple technology skillsets to continue to contribute with oversight and compliance.

**Iterate** on the policies and process, adding and improving while maintaining availability and service.

- Start small by automating existing tasks.
- Get familiar with the automation process, so that you can identify other areas that can benefit from automation.
- iterate your automation initiatives, adding agility incrementally while maintaining the required availability.
- Taking an incremental approach paves the way for success!

Orchestrate the network service!

- Automation of the deployment process is required to deliver applications rapidly.
- Creating an agile service environment requires different elements to be managed across technology skillsets.
- Prepare for an end to end orchestration that provides for control over automation and the order of deployments.

## Network Automation Tools

The good news here is that for the most part, the tools we use here for Network automation are generally the same that we will use for other areas of automation or what we have already covered so far or what we will cover in future sessions.

Operating System - As I have throughout this challenge, I am focusing on doing most of my learning with a Linux OS, those reasons were given in the Linux section but almost all of the tooling that we will touch albeit cross-OS platforms maybe today they all started as Linux based applications or tools, to begin with.

Integrated Development Environment (IDE) - Again not much to say here other than throughout I would suggest Visual Studio Code as your IDE, based on the extensive plugins that are available for so many different languages.

Configuration Management - We have not got to the Configuration management section yet, but it is very clear that Ansible is a favourite in this area for managing and automating configurations. Ansible is written in Python but you do not need to know Python.

- Agentless
- Only requires SSH
- Large Support Community
- Lots of Network Modules
- Push only model
- Configured with YAML
- Open Source!

[Link to Ansible Network Modules](https://docs.ansible.com/ansible/2.9/modules/list_of_network_modules.html)

We will also touch on **Ansible Tower** in the configuration management section, see this as the GUI front end for Ansible.

CI/CD - Again we will cover more about the concepts and tooling around this but it's important to at least mention here as this spans not only networking but all provisioning of service and platform.

In particular, Jenkins provides or seems to be a popular tool for Network Automation.

- Monitors git repository for changes and then initiates them.

Version Control - Again something we will dive deeper into later on.

- Git provides version control of your code on your local device - Cross-Platform
- GitHub, GitLab, BitBucket etc are online websites where you define your repositories and upload your code.

Language | Scripting - Something we have not covered here is Python as a language, I decided to dive into Go instead as the programming language based on my circumstances, I would say that it was a close call between Golang and Python and Python it seems to be the winner for Network Automation.

- Nornir is something to mention here, an automation framework written in Python. This seems to take the role of Ansible but specifically around Network Automation. [Nornir documentation](https://nornir.readthedocs.io/en/latest/)

Analyse APIs - Postman is a great tool for analysing RESTful APIs. Helps to build, test and modify APIs.

- POST >>> To create resources objects.
- GET >>> To retrieve a resources.
- PUT >>> To create or replace the resources.
- PATCH >>> To create or update the resources object.
- Delete >>> To delete a resources

[Postman tool Download](https://www.postman.com/downloads/)

### Other tools to mention

[Cisco NSO (Network Services Orchestrator)](https://www.cisco.com/c/en/us/products/cloud-systems-management/network-services-orchestrator/index.html)

[NetYCE - Simplify Network Automation](https://netyce.com/)

[Network Test Automation](https://pubhub.devnetcloud.com/media/genie-feature-browser/docs/#/)

Over the next 3 days, I am planning to get more hands-on with some of the things we have covered and put some work in around Python and Network automation.

We have nowhere near covered all of the networking topics so far but wanted to make this broad enough to follow along and still keep learning from the resources I am adding below.

## Resources

- [3 Necessary Skills for Network Automation](https://www.youtube.com/watch?v=KhiJ7Fu9kKA&list=WL&index=122&t=89s)
- [Computer Networking full course](https://www.youtube.com/watch?v=IPvYjXCsTg8)
- [Practical Networking](http://www.practicalnetworking.net/)
- [Python Network Automation](https://www.youtube.com/watch?v=xKPzLplPECU&list=WL&index=126)

See you on [Day 25](day25.md)
