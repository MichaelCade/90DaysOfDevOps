---
title: '#90DaysOfDevOps - 概述：DevOps 与学习一门编程语言 - 第七天'
published: false
description: 90DaysOfDevOps - 概述：DevOps 与学习一门编程语言
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---
## 概述：DevOps与学习一门编程语言

我认为，要想在长期成为一名成功的DevOps工程师，你需要了解至少一种编程语言的基本用法。我想通过本节的第一部分来探究，为什么这是一项重要的技能。同时也希望在本周或本节结束时，你会更好地理解在这个学习过程中，为什么、如何做和做什么。

我想如果我在社交平台上提问，从事与DevOps的相关工作是否需要具备编程能力，得到的答案很可能是肯定的？如果你不这么认为，欢迎告诉我。但一个更大的问题是，你无法明确知道所需的是哪种编程语言？一个最常见的回答是，我看到Python正变得越来越热门，而我们应该选择学习Golang或Go语言。

为了在DevOps中获得成功，至少在我看来，你应该具备良好的编程知识。但我们也该知道为什么我们需要它来选择正确的方向。

## 了解为什么需要学习编程语言

许多DevOps的工具是用Python或Go编写的，如果你要构建DevOps工具，这将为你提供便利。这也成为了DevOps推荐Python和Go的原因。如今这会影响你决定学习哪一种编程语言，并可能是对你最有益的。如果你想构建DevOps工具或是加入一个从事相关工作的团队，选择学习与之相同的语言将是有意义的。如果你需要大量使用Kubernetes或Containers，那么你很可能会将Go作为你的编程语言。对我来说，我工作的公司(Kasten by Veeam) 位于云原生态系统领域(Cloud-Native ecosystem)，专注于Kubernetes的数据管理并且所有工作都用Go来编写。

但或许你是一名学生或过渡职业，可能没有像这样明确的方向来帮助你做出选择。我觉得在这个情况下，你应该选择一个与你感兴趣的应用程序有相近特点的。

请记住，我在这里并不是为了成为一名程序开发者。我只是想去更多地了解编程语言，从而让我能够阅读和理解那些工具在做些什么，进而有可能启发我们如何改进相关的工作。

另一个重要的点是了解如何与DevOps工具(Kasten K10, Terraform 又或是HCL)进行交互。这些就是我们所说的配置文件(config files)，它就是帮助你与那些DevOps工具成功交互的东西，通常它们会以YAML的格式出现。(我们可能会在本节的最后一天稍微讲解YAML)

## 我只是自说自话而不是学习编程语言吗？

大多数时候或根据担任的角色，你会帮助工程团队将DevOps在他们的工作流程中实现。大量围绕应用程序的测试，并确保被构建的工作流程符合我们前几天提到的那些DevOps原则。但实际上，这个过程将有大量时间花费在寻找程序性能或类似的问题上。这回到了我最初的观点，我需要了解那个被用于编写程序源码的编程语言吗？如果那个应用程序是用NodeJS编写的，而你熟悉的是Go或Python，那么这些知识并不能帮助你很多。

## 为什么选Go

为什么Golang是DevOps的下一个编程语言，Go已经成为近年非常流行的编程语言。根据2021年StackOverflow的调查，Go在最受欢迎的编程、脚本和标记语言中排名第四，其中Python位于榜首，但也请继续看完。[StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

正如我提到的那样，一些最出名的DevOps工具和平台是用Go来编写的，例如Kubernetes、Docker、Grafana和Prometheus。

那么Go具备哪些适合DevOps的特性呢？

## Go的构建和部署

一个优势是就像Python那样具备解释性，并且在DevOps工作中，你无需在运行程序之前进行编译。特别是对于小规模的自动化任务，你不希望在构建的过程中被编译流程拖后腿。Go是一个编译性的编程语言，**Go直接完成编译变成机器码**。Go也是出了名的编译速度快。

## DevOps的Go vs Python

Go程序是静态链接的，这意味着当你编译一个go程序时，所有的东西都会被放在一个二进制执行文件里，并且不需要在远程机器上安装外部依赖。对比在运行使用了外部库的Python程序时，它需要确保所用到的库都已安装在这台远程计算机上，这一特点让go程序的部署变得简单。

Go是一种独立于平台的语言，这意味着你可以很轻松地为\*所有操作系统 Linux、Windows、macOS等等生成二进制可执行文件。而对于Python来说，为特定操作系统制作二进制可执行文件就没那么简单了。

Go是一个具备非常高性能的语言，它可以快速完成编译，并且比Python占用更少的CPU、内存等资源。许多优化已经被应用于Go语言中，使其能达到高性能。(详见文末Resources)

与常需要使用第三方库来实现特定程序的Python不同，Go包含了一个标准库，其中有DevOps所需的大部分功能。包括文件处理功能、HTTP Web服务、JSON处理、对并发和并行的本机支持和内置测试。

这篇文章不是让你放弃Python，我只是给出我自己选择Go的理由。但这些理由不一定是上面提到的内容，通常是因为我工作中用到了Go开发软件，所以这是我的原因。

我会说，一旦你学习了你的第一门编程语言，学习其他的语言将变得更简单。你可能不会永远只在一家公司里做一个岗位的工作，你很有可能会接触到管理、架构、编排、调试JavaScript和NodeJS的应用程序。

## 相关资料

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s) 
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I) 
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals) 
- [FreeCodeCamp -  Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s) 
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N) 

在未来的6天里，我打算通过上述资料来帮助我组织每天的笔记。你会注意到，它们作为一整个课程通常需要3个小时来了解。我想分享我的完成列表，如果时间允许，你应该去每个链接看看，我也会每天坚持学习。

让我们[第八天](day08.md)再见。
