---
title: '#90DaysOfDevOps - The Big Picture: Configuration Management - Day 63'
published: false
description: 90DaysOfDevOps - The Big Picture Configuration Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048711
---

## The Big Picture: Configuration Management

Coming straight off the back of the section covering Infrastructure as Code, there is likely going to be some crossover as we talk about Configuration Management or Application Configuration Management.

Configuration Management is the process of maintaining applications, systems and servers in the desired state. The overlap with Infrastructure as code is that IaC is going to make sure your infrastructure is at the desired state but after that especially terraform is not going to look after the desired state of your OS settings or Application and that is where Configuration Management tools come in. Make sure that the system and applications perform the way it is expected as changes occur over Deane.

Configuration management keeps you from making small or large changes that go undocumented.

### Scenario: Why would you want to use Configuration Management

The scenario or why you'd want to use Configuration Management, meet Dean he's our system administrator and Dean is a happy camper pretty and
working on all of the systems in his environment.

What happens if their system fails, if there's a fire, a server goes down well? Dean knows exactly what to do he can fix that fire easily the problems become difficult for Dean however if multiple servers start failing particularly when you have large and expanding environments, this is why Dean needs to have a configuration management tool. Configuration Management tools can help make Dean look like a rockstar, all he has to do is configure the right codes that allow him to push out the instructions on how to set up each of the servers quickly effectively and at scale.

### Configuration Management tools

There are a variety of configuration management tools available, and each has specific features that make it better for some situations than others.

![](Images/Day63_config1.png)

At this stage, we will take a quickfire look at the options in the above picture before making our choice on which one we will use and why.

- **Chef**

  - Chef ensures configuration is applied consistently in every environment, at any scale with infrastructure automation.
  - Chef is an open-source tool developed by OpsCode written in Ruby and Erlang.
  - Chef is best suited for organisations that have a heterogeneous infrastructure and are looking for mature solutions.
  - Recipes and Cookbooks determine the configuration code for your systems.
  - Pro - A large collection of recipes is available
  - Pro - Integrates well with Git which provides a strong version control
  - Con - Steep learning curve, a considerable amount of time required.
  - Con - The main server doesn't have much control.
  - Architecture - Server / Clients
  - Ease of setup - Moderate
  - Language - Procedural - Specify how to do a task

- **Puppet**
  - Puppet is a configuration management tool that supports automatic deployment.
  - Puppet is built in Ruby and uses DSL for writing manifests.
  - Puppet also works well with heterogeneous infrastructure where the focus is on scalability.
  - Pro - Large community for support.
  - Pro - Well-developed reporting mechanism.
  - Con - Advance tasks require knowledge of the Ruby language.
  - Con - The main server doesn't have much control.
  - Architecture - Server / Clients
  - Ease of setup - Moderate
  - Language - Declarative - Specify only what to do
- **Ansible**

  - Ansible is an IT automation tool that automates configuration management, cloud provisioning, deployment and orchestration.
  - The core of Ansible playbooks is written in YAML. (Should do a section on YAML as we have seen this a few times)
  - Ansible works well when there are environments that focus on getting things up and running fast.
  - Works on playbooks which provide instructions to your servers.
  - Pro - No agents are needed on remote nodes.
  - Pro - YAML is easy to learn.
  - Con - Performance speed is often less than other tools (Faster than Dean doing it himself manually)
  - Con - YAML is not as powerful as Ruby but has less of a learning curve.
  - Architecture - Client Only
  - Ease of setup - Very Easy
  - Language - Procedural - Specify how to do a task

- **SaltStack**
  - SaltStack is a CLI-based tool that automates configuration management and remote execution.
  - SaltStack is Python based whilst the instructions are written in YAML or its DSL.
  - Perfect for environments with scalability and resilience as the priority.
  - Pro - Easy to use when up and running
  - Pro - Good reporting mechanism
  - Con - The setup phase is tough
  - Con - New web UI which is much less developed than the others.
  - Architecture - Server / Clients
  - Ease of setup - Moderate
  - Language - Declarative - Specify only what to do

### Ansible vs Terraform

The tool that we will be using for this section is going to be Ansible. (Easy to use and easier language basics required.)

I think it is important to touch on some of the differences between Ansible and Terraform before we look into the tooling a little further.

|                | Ansible                                                      | Terraform                                                        |
| -------------- | ------------------------------------------------------------ | ---------------------------------------------------------------- |
| Type           | Ansible is a configuration management tool                   | Terraform is an orchestration tool                             |
| Infrastructure | Ansible provides support for mutable infrastructure          | Terraform provides support for immutable infrastructure          |
| Language       | Ansible follows procedural language                          | Terraform follows a declarative language                          |
| Provisioning   | Ansible provides partial provisioning (VM, Network, Storage) | Terraform provides extensive provisioning (VM, Network, Storage) |
| Packaging      | Ansible provides complete support for packaging & templating | Terraform provides partial support for packaging & templating    |
| Lifecycle Mgmt | Ansible does not have lifecycle management                   | Terraform is heavily dependent on lifecycle and state mgmt       |

## Resources

- [What is Ansible](https://www.youtube.com/watch?v=1id6ERvfozo)
- [Ansible 101 - Episode 1 - Introduction to Ansible](https://www.youtube.com/watch?v=goclfp6a2IQ)
- [NetworkChuck - You need to learn Ansible right now!](https://www.youtube.com/watch?v=5hycyr-8EKs&t=955s)

See you on [Day 64](day64.md)
