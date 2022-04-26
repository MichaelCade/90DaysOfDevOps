---
title: '#90DaysOfDevOps - The Big Picture: Learning a Programming Language - Day 7'
published: false
description: 90DaysOfDevOps - The Big Picture DevOps & Learning a Programming Language
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---
## 介绍：DevOps & 学习编程语言
## The Big Picture: DevOps & Learning a Programming Language

我认为，要想在长期成为一名成功的DevOps工程师，你需要了解至少一种编程语言的基本用法。我想通过本节的第一部分来探究，为什么这是一项重要的技能。同时也希望在本周或本节结束时，你会更好地理解在这个学习过程中，为什么、如何做和做什么。
I think it is fair to say to be successful in the long term as a DevOps engineer you've got to know at least one programming language at a foundational level. I want to take this first session of this section to explore why this is such a critical skill to have, and hopefully, by the end of this week or section, you are going to have a better understanding of the why, how and what to do to progress with your learning journey. 

我想如果我在社交平台上提问，从事与DevOps的相关工作是否需要具备编程能力，得到的答案很可能是肯定的？如果你不这么认为，欢迎告诉我。但一个更大的问题是，你无法明确知道所需的是哪种编程语言？一个最常见的回答是，我看到Python正变得越来越热门，而我们应该选择学习Golang或Go语言。
I think if I was to ask out on social do you need to have programming skills for DevOps related roles, the answer will be most likely a hard yes? Let me know if you think otherwise? Ok but then a bigger question and this is where you won't get such a clear answer is which programming language?  The most common answer I have seen here has been Python or increasingly more often, we're seeing Golang or Go should be the language that you learn. 

为了在DevOps中获得成功，至少在我看来，你应该具备良好的编程知识。但我们也该知道为什么我们需要它来选择正确的方向。
To be successful in DevOps you have to have a good knowledge of programming skills is my takeaway from that at least. But we have to understand why we need it to choose the right path. 

## 了解为什么需要学习编程语言
## Understand why you need to learn a programming language. 

许多DevOps的工具是用Python或Go编写的，如果你要构建DevOps工具，这将为你提供便利。这也成为了DevOps推荐Python和Go的原因。如今这会影响你决定学习哪一种编程语言，并可能是对你最有益的。如果你想构建DevOps工具或是加入一个从事相关工作的团队，选择学习与之相同的语言将是有意义的。如果你需要大量使用Kubernetes或Containers，那么你很可能会将Go作为你的编程语言。对我来说，我工作的公司(Kasten by Veeam) 位于云原生态系统领域(Cloud-Native ecosystem)，专注于Kubernetes的数据管理并且所有工作都用Go来编写。
The reason that Python and Go are recommended so often for DevOps engineers is that a lot of the DevOps tooling is written in either Python or Go, which makes sense if you are going be build DevOps tools. Now this is important as this will determine really what you should learn and that would likely be the most beneficial. If you are going to be building DevOps tools or you are joining a team that does then it would make sense to learn that same language, if you are going to be heavily involved in Kubernetes or Containers then it's more than likely that you would want to choose Go as your programming language. For me, the company I work for (Kasten by Veeam) is in the Cloud-Native ecosystem focused on data management for Kubernetes and everything is written in Go. 

但或许你是一名学生或过渡职业，可能没有像这样明确的方向来帮助你做出选择。我觉得在这个情况下，你应该选择一个与你感兴趣的应用程序有相近特点的。
But then you might not have clear cut reasoning like that to choose you might be a student or transitioning careers with no real decision made for you. I think in this situation then you should choose the one that seems to resonate and fit with the applications you are looking to work with. 

请记住，我在这里并不是为了成为一名程序开发者。我只是想去更多地了解编程语言，从而让我能够阅读和理解那些工具在做些什么，进而有可能启发我们如何改进相关的工作。
Remember I am not looking to become a software developer here I just want to understand a little more about the programming language so that I can read and understand what those tools are doing and then that leads to possibly how we can help improve things. 

另一个重要的点是了解如何与DevOps工具(Kasten K10, Terraform 又或是HCL)进行交互。这些就是我们所说的配置文件(config files)，它就是帮助你与那些DevOps工具成功交互的东西，通常它们会以YAML的格式出现。(我们可能会在本节的最后一天稍微讲解YAML)
I would also it is also important to know how you interact with those DevOps tools which could be Kasten K10 or it could be Terraform and HCL. These are what we will call config files and this is how you interact with those DevOps tools to make things happen, commonly these are going to be YAML. (We may use the last day of this section to dive a little into YAML) 

## 我只是自说自话而不是学习编程语言吗？
## Did I just talk myself out of learning a programming language?

大多数时候或根据担任的角色，你会帮助工程团队将DevOps在他们的工作流程中实现。大量围绕应用程序的测试，并确保被构建的工作流程符合我们前几天提到的那些DevOps原则。但实际上，这个过程将有大量时间花费在寻找程序性能或类似的问题上。这回到了我最初的观点，我需要了解那个被用于编写程序源码的编程语言吗？如果那个应用程序是用NodeJS编写的，而你熟悉的是Go或Python，那么这些知识并不能帮助你很多。
Most of the time or depending on the role, you will be helping engineering teams implement DevOps into their workflow, a lot of testing around the application and making sure that the workflow that is built aligns to those DevOps principles we mentioned over the first few days. But in reality, this is going to be a lot of the time troubleshooting an application performance issue or something along those lines. This comes back to my original point and reasoning, the programming language I need to know is the one that the code is written in? If their application is written in NodeJS it won’t help much if you have a Go or Python badge. 

## 为什么选Go
## Why Go 

为什么Golang是DevOps的下一个编程语言，Go已经成为近年非常流行的编程语言。根据2021年StackOverflow的调查，Go在最受欢迎的编程、脚本和标记语言中排名第四，其中Python位于榜首，但也请继续看完。[StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)
Why Golang is the next programming language for DevOps, Go has become a very popular programming language in recent years. According to the StackOverflow Survey for 2021 Go came in fourth for the most wanted Programming, scripting and markup languages with Python being top but hear me out. [StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

正如我提到的那样，一些最出名的DevOps工具和平台是用Go来编写的，例如Kubernetes、Docker、Grafana和Prometheus。
As I have also mentioned some of the most known DevOps tools and platforms are written in Go such as Kubernetes, Docker, Grafana and Prometheus. 

那么Go具备哪些适合DevOps的特性呢？
What are some of the characteristics of Go that make it great for DevOps?

## Go的构建和部署
## Build and Deployment of Go Programs 

一个优势是就像Python那样具备解释性，并且在DevOps工作中，你无需在运行程序之前进行编译。特别是对于小规模的自动化任务，你不希望在构建的过程中被编译流程拖后腿。Go是一个编译性的编程语言，**Go直接完成编译变成机器码**。Go也是出了名的编译速度快。
An advantage of using a language like Python that is interpreted in a DevOps role is that you don’t need to compile a python program before running it. Especially for smaller automation tasks, you don’t want to be slowed down by a build process that requires compilation even though, Go is a compiled programming language, **Go compiles directly into machine code**.  Go is known also for fast compilation times. 

## DevOps的Go vs Python
## Go vs Python for DevOps 

Go程序是静态链接的，这意味着当你编译一个go程序时，所有的东西都会被放在一个二进制执行文件里，并且不需要在远程机器上安装外部依赖。对比在运行使用了外部库的Python程序时，它需要确保所用到的库都已安装在这台远程计算机上，这一特点让go程序的部署变得简单。
Go Programs are statically linked, this means that when you compile a go program everything is included in a single binary executable, no external dependencies will be required that would need to be installed on the remote machine, this makes the deployment of go programs easy, compared to python program that uses external libraries you have to make sure that all those libraries are installed on the remote machine that you wish to run on. 

Go是一种独立于平台的语言，这意味着你可以很轻松地为*所有操作系统 Linux、Windows、macOS等等生成二进制可执行文件。而对于Python来说，为特定操作系统制作二进制可执行文件就没那么简单了。
Go is a platform-independent language, which means you can produce binary executables for *all the operating systems, Linux, Windows, macOS etc and very easy to do so. With Python, it is not as easy to create these binary executables for particular operating systems. 

Go是一个具备非常高性能的语言，它可以快速完成编译，并且比Python占用更少的CPU、内存等资源。许多优化已经被应用于Go语言中，使其能达到高性能。(详见文末Resources)
Go is a very performant language, it has fast compilation and fast run time with lower resource usage like CPU and memory especially compared to python, numerous optimisations have been implemented in the Go language that makes it so performant. (Resources below) 

与常需要使用第三方库来实现特定程序的Python不同，Go包含了一个标准库，其中有DevOps所需的大部分功能。包括文件处理功能、HTTP Web服务、JSON处理、对并发和并行的本机支持和内置测试。
Unlike Python which often requires the use of third party libraries to implement a particular python program, go includes a standard library that has the majority of functionality that you would need for DevOps built directly into it. This includes functionality file processing, HTTP web services, JSON processing, native support for concurrency and parallelism as well as built-in testing. 

这篇文章不是让你放弃Python，我只是给出我自己选择Go的理由。但这些理由不一定是上面提到的内容，通常是因为我工作中用到了Go开发软件，所以这是我的原因。
This is by no way throwing Python under the bus I am just giving my reasons for choosing Go but they are not the above Go vs Python it's generally because it makes sense as the company I work for develops software in Go so that is why. 

我会说，一旦你学习了你的第一门编程语言，学习其他的语言将变得更简单。你可能不会永远只在一家公司里做一个岗位的工作，你很有可能会接触到管理、架构、编排、调试JavaScript和NodeJS的应用程序。
I will say that once you have or at least I am told as I am not many pages into this chapter right now, is that once you learn your first programming language it becomes easier to take on other languages. You're probably never going to have a single job in any company anywhere where you don't have to deal with manage, architect, orchestrating, debug JavaScript and Node JS applications. 

## Resources

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s) 
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I) 
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals) 
- [FreeCodeCamp -  Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s) 
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N) 

Now for the next 6 days of this topic my intention is to work through some of the resources listed above and document my notes for each day. You will notice that they are generally around 3 hours as a full course, I wanted to share my complete list so that if you have time you should move ahead and work through each one if time permits, I will be sticking to my learning hour each day. 

See you on [Day 8](day08.md). 
