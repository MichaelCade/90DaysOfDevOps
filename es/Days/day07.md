---
title: '#90DaysOfDevOps - The Big Picture: Learning a Programming Language - Day 7'
published: false
description: 90DaysOfDevOps - The Big Picture DevOps & Learning a Programming Language
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048856
---

## The Big Picture: DevOps & Learning a Programming Language

I think it is fair to say to be successful in the long term as a DevOps engineer you've got to know at least one programming language at a foundational level. I want to take this first session of this section to explore why this is such a critical skill to have, and hopefully, by the end of this week or section, you are going to have a better understanding of the why, how and what to do to progress with your learning journey.

I think if I was to ask out on social do you need to have programming skills for DevOps related roles, the answer will be most likely a hard yes? Let me know if you think otherwise? Ok but then a bigger question and this is where you won't get such a clear answer which programming language? The most common answer I have seen here has been Python or increasingly more often, we're seeing Golang or Go should be the language that you learn.

To be successful in DevOps you have to have a good knowledge of programming skills is my takeaway from that at least. But we have to understand why we need it to choose the right path.

## Understand why you need to learn a programming language.

The reason that Python and Go are recommended so often for DevOps engineers is that a lot of the DevOps tooling is written in either Python or Go, which makes sense if you are going to be building DevOps tools. Now this is important as this will determine really what you should learn and that would likely be the most beneficial. If you are going to be building DevOps tools or you are joining a team that does then it would make sense to learn that same language, if you are going to be heavily involved in Kubernetes or Containers then it's more than likely that you would want to choose Go as your programming language. For me, the company I work for (Kasten by Veeam) is in the Cloud-Native ecosystem focused on data management for Kubernetes and everything is written in Go.

But then you might not have clear cut reasoning like that to choose you might be a student or transitioning careers with no real decision made for you. I think in this situation then you should choose the one that seems to resonate and fit with the applications you are looking to work with.

Remember I am not looking to become a software developer here I just want to understand a little more about the programming language so that I can read and understand what those tools are doing and then that leads to possibly how we can help improve things.

It is also important to know how you interact with those DevOps tools which could be Kasten K10 or it could be Terraform and HCL. These are what we will call config files and this is how you interact with those DevOps tools to make things happen, commonly these are going to be YAML. (We may use the last day of this section to dive a little into YAML)

## Did I just talk myself out of learning a programming language?

Most of the time or depending on the role, you will be helping engineering teams implement DevOps into their workflow, a lot of testing around the application and making sure that the workflow that is built aligns to those DevOps principles we mentioned over the first few days. But in reality, this is going to be a lot of the time troubleshooting an application performance issue or something along those lines. This comes back to my original point and reasoning, the programming language I need to know is the one that the code is written in? If their application is written in NodeJS it won’t help much if you have a Go or Python badge.

## Why Go

Why Golang is the next programming language for DevOps, Go has become a very popular programming language in recent years. According to the StackOverflow Survey for 2021 Go came in fourth for the most wanted Programming, scripting and markup languages with Python being top but hear me out. [StackOverflow 2021 Developer Survey – Most Wanted Link](https://insights.stackoverflow.com/survey/2021#section-most-loved-dreaded-and-wanted-programming-scripting-and-markup-languages)

As I have also mentioned some of the most known DevOps tools and platforms are written in Go such as Kubernetes, Docker, Grafana and Prometheus.

What are some of the characteristics of Go that make it great for DevOps?

## Build and Deployment of Go Programs

An advantage of using a language like Python that is interpreted in a DevOps role is that you don’t need to compile a python program before running it. Especially for smaller automation tasks, you don’t want to be slowed down by a build process that requires compilation even though, Go is a compiled programming language, **Go compiles directly into machine code**. Go is known also for fast compilation times.

## Go vs Python for DevOps

Go Programs are statically linked, this means that when you compile a go program everything is included in a single binary executable, and no external dependencies will be required that would need to be installed on the remote machine, this makes the deployment of go programs easy, compared to python program that uses external libraries you have to make sure that all those libraries are installed on the remote machine that you wish to run on.

Go is a platform-independent language, which means you can produce binary executables for \*all the operating systems, Linux, Windows, macOS etc and very easy to do so. With Python, it is not as easy to create these binary executables for particular operating systems.

Go is a very performant language, it has fast compilation and fast run time with lower resource usage like CPU and memory especially compared to python, numerous optimisations have been implemented in the Go language that makes it so performant. (Resources below)

Unlike Python which often requires the use of third party libraries to implement a particular python program, go includes a standard library that has the majority of functionality that you would need for DevOps built directly into it. This includes functionality file processing, HTTP web services, JSON processing, native support for concurrency and parallelism as well as built-in testing.

This is by no way throwing Python under the bus I am just giving my reasons for choosing Go but they are not the above Go vs Python it's generally because it makes sense as the company I work for develops software in Go so that is why.

I will say that once you have or at least I am told as I am not many pages into this chapter right now, is that once you learn your first programming language it becomes easier to take on other languages. You're probably never going to have a single job in any company anywhere where you don't have to deal with managing, architect, orchestrating, debug JavaScript and Node JS applications.

## Resources

- [StackOverflow 2021 Developer Survey](https://insights.stackoverflow.com/survey/2021)
- [Why we are choosing Golang to learn](https://www.youtube.com/watch?v=7pLqIIAqZD4&t=9s)
- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI&t=312s)
- [Techworld with Nana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
- [**NOT FREE** Nigel Poulton Pluralsight - Go Fundamentals - 3 hours 26 mins](https://www.pluralsight.com/courses/go-fundamentals)
- [FreeCodeCamp - Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=1025s)
- [Hitesh Choudhary - Complete playlist](https://www.youtube.com/playlist?list=PLRAV69dS1uWSR89FRQGZ6q9BR2b44Tr9N)

Now for the next 6 days of this topic, I intend to work through some of the resources listed above and document my notes for each day. You will notice that they are generally around 3 hours as a full course, I wanted to share my complete list so that if you have time you should move ahead and work through each one if time permits, I will be sticking to my learning hour each day.

See you on [Day 8](day08.md).
