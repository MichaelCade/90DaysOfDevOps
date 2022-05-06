---
title: '#90DaysOfDevOps - Introduction - Day 1'
published: false
description: 90DaysOfDevOps - Introduction
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048731
---
## 简介 - Day 1 

在90天中的第一天，我们开始学习DevOps的基本理解和工具。这些可以有助于建立DevOps的思维方式。
<!-- Day 1 of our 90 days and adventure to learn a good foundational understanding of DevOps and tools that help with a DevOps mindset.  -->

几年前，我开始学习相关内容，但我的关注于虚拟化平台和基于云的技术，主要研究基础设施即代码 (Infrastructure as Code, [IaC](https://www.ibm.com/cloud/learn/infrastructure-as-code)) 和Terraform和Chef的应用程序配置管理。
<!-- This learning journey started for me a few years back but my focus then was around virtualisation platforms and cloud based technologies, I was looking mostly into Infrastructure as Code and Application configuration management with Terraform and Chef.  -->

接着来到2021年3月，我得到了一个可以将精力集中在Veeam的Kasten的云原生部署上的机会。这个项目专注于Kubernetes和DepOps以及相关的社区。在开始学习后，我很快发现在除了Kubernetes和容器化的基础知识，那里还有一个非常广阔的世界。我开始在社区中交流学习更多关于DevOps文化、工具和流程，最终我想公开地分享这些想法。
<!-- Fast forward to March 2021, I was given an amazing opportunity to concentrate my efforts around the Cloud Native strategy at Kasten by Veeam. Which was going to be a massive focus on Kubernetes and DevOps and the community surrounding these technologies. I started my learning journey and quickly realised there was a very wide world aside from just learning the fundamentals of Kubernetes and Containerisation and it was then when I started speaking to the community and learning more and more about the DevOps culture, tooling and processes so I started documenting some of the areas I wanted to learn in public.  -->

[So you want to learn DevOps?](https://blog.kasten.io/devops-learning-curve)

## 开始我们的旅程吧

如果你阅读了以上的博客，你会发现这是我学习过程中的进阶内容。我认为我并不是以上以上任一领域的专家，但我希望分享一些免费和需付费的资源，我们可以按需选择。
<!-- If you read the above blog you will see this is a high level contents for my learning journey and I will say at this point I am no where near an expert in any of these sections but what I wanted to do was share some resources both FREE and some paid for but an option for both as we all have different circumstances.  -->

在接下来的90天里，我想记录这些资料并涵盖那些基础领域。我希望社区参与进来，分享你的相关经历和资源，以便我们一起学习共同进步。
<!-- Over the next 90 days I want to document these resources and cover those foundational areas, I would love for the community to also get involved share your journey and resources so we can learn in public and help each other.  -->

在该项目开头的README中，你会了解到我已经将内容拆分成多个小节，基本上是由12周加6天组成。前6天，我们会大致探讨DevOps的基础，后续再深入到一些特定领域。这份清单不是完美的，再次希望社区参与进来并一起帮助它成为有用的资源。
<!-- You will see from the opening readme in the project repository that I have split things into sections and it is basically 12 weeks plus 6 days. The first 6 days we will explore the fundamentals of DevOps in general before diving into some of the specific areas, by no way is this list exhaustive and again would love for the community to assist in making this a useful resource.  -->

在这里我会分享另一个资源，我认为每个人都应认真了解，或是根据自身需求制作自己的思维导图，它的地址如下：
<!-- Another resource I will share at this point that I think everyone should have a good look at and maybe create your own mind map for yourself and your interest and position is the following:  -->

[DevOps Roadmap](https://roadmap.sh/devops)

当我在创建这个初始清单和博客的时候，我发现这是个很好的资源。你也可以看到除了在我列出的12个专题以外的，其他领域更详细的信息。
<!-- I found this a great resource when I was creating my initial list and blog post on this topic. You can also see there are other areas that go into a lot more detail outside of the 12 topics I have listed here in this repository.  -->

## 第一步 - 什么是 DevOps? 

这里可以列出很多的博客和YouTube视频，但作为90天挑战的开始，以及我们每天花费约一小时来学习一些新的或关于DevOps的东西。我觉得从宏观的“什么是DevOps”开始是个不错的选择。
<!-- There are so many blog articles and YouTube videos to list here, but as we start the 90 day challenge and we focus on spending around an hour a day learning something new or about DevOps I thought it was good to get some of the high level of "what DevOps is" down to begin.  -->

首先，DevOps不是工具。你不能购买它，它不是可下载的软件sku或开源GitHub仓库。DevOps也不是编程语言或什么黑魔法。
<!-- Firstly, DevOps is not a tool. You cannot buy it, it is not a software sku or an open source GitHub repository you can download. It is also not a programming language, it is also not some dark art magic either.  -->

DevOps是一种在软件开发中更明智的做事方式。- 等一下... 但如果你不是一个软件开发人员，你现在应该关闭这个页面并离开吗？？不，继续读下去... 因为DevOps将软件开发和运维运营结合在了一起。我先前提及到，我更多关注的是虚拟机方面的工作，而这些通常属于运营。但在社区中，不同背景的人们都可以通过更好地了解DevOps来学习那些实践案例。DevOps将100%造福于个人、开发者、运维运营和QA工程师。
<!-- DevOps is a way to do smarter things in Software Development. - Hold up... But if you are not a software developer should you turn away right now and not dive into this project??? No Not at all, Stay... Because DevOps brings together a combination of software development and operations. I mentioned earlier that I was more on the VM side and that would generally fall under the Operations side of the house, but within the community there are people with all different backgrounds where DevOps is 100% going to benefit the individual, Developers, Operations and QA Engineers all can equally learn these best practices by having a better understanding of DevOps.  -->

DevOps是一系列有助于达成这一目标的实践：减少产品从构思到发布阶段，到最终用户或内部团队或客户的任何人所需要的时间。
<!-- DevOps is a set of practices that help to reach the goal of this movement: reducing the time between the ideation phase of a product and its release in production to the end-user or whomever it could be an internal team or customer.  -->

在这第一个星期，我们将展开讨论**敏捷方法论**(The Agile Methodology)。DevOps和Agile是被广泛使用的方法，为的是实现**应用程序app**的持续迭代更新。
<!-- Another area we will dive into in this first week is around **The Agile Methodology** DevOps and Agile are widely adopted together to achieve continuous delivery of your **Application**  -->

宏观层次的收获是，DevOps的思维方式是将漫长的软件发布过程从可能几年的时间拆分成更频繁的、较小的多次发布。另一个关键点是，DevOps打破了团队间的隔阂：开发人员、运维运营人员和QA工程师。
<!-- The high level take away is with a DevOps mindset or culture its about taking a way the long drawn out software release process from potentially years to being able to drop smaller releases more frequently. The other key fundamental to take away here is it's about breaking down silos between the teams I previously mentioned, Developers, Operations and QA.  -->

从DevOps的角度，**开发、测试、部署**都属于DevOps团队。
<!-- From a DevOps perspective, **Development, Testing and Deployment** all land with the DevOps team.  -->

最后一点，我们必须通过**自动化**使得整个过程尽可能有效和高效。
<!-- The final point I will make is to make this as effective and efficient as possible we must leverage **Automation**  -->

<!-- ## 信息来源 -->
## Resources 

I am always open to adding additional resources to these readme files as it is here as a learning tool.  

My advice is to watch all of the below and hopefully you also picked something up from the text and explanations above. 

- [DevOps in 5 Minutes](https://www.youtube.com/watch?v=Xrgk023l4lI)
- [What is DevOps? Easy Way](https://www.youtube.com/watch?v=_Gpe1Zn-1fE&t=43s)
- [DevOps roadmap 2022 | Success Roadmap 2022](https://www.youtube.com/watch?v=7l_n97Mt0ko)

If you made it this far then you will know if this is where you want to be or not. See you on [Day 2](day02.md).  
