---
title: '#90DaysOfDevOps - Responsibilities of a DevOps Engineer - Day 2'
published: false
description: 90DaysOfDevOps - Responsibilities of a DevOps Engineer
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048699
date: '2022-04-17T21:15:34Z'
---

## Responsibilities of a DevOps Engineer

Hopefully, you are coming into this off the back of going through the resources and posting on [Day1 of #90DaysOfDevOps](day01.md)

It was briefly touched on in the first post but now we must get deeper into this concept and understand that there are two main parts when creating an application. We have the **Development** part where software developers program the application and test it. Then we have the **Operations** part where the application is deployed and maintained on a server.

## DevOps is the link between the two

To get to grips with DevOps or the tasks which a DevOps engineer would be carrying out we need to understand the tools or the process and overview of those and how they come together.

Everything starts with the application! You will see so much throughout that it is all about the application when it comes to DevOps.

Developers will create an application, this can be done with many different technology stacks and let's leave that to the imagination for now as we get into this later. This can also involve many different programming languages, build tools, code repositories etc.

As a DevOps engineer you won't be programming the application but having a good understanding of the concepts of how a developer works and the systems, tools and processes they are using is key to success.

At a very high level, you are going to need to know how the application is configured to talk to all of its required services or data services and then also sprinkle a requirement of how this can or should be tested.

The application will need to be deployed somewhere, lets's keep it generally simple here and make this a server, doesn't matter where but a server. This is then expected to be accessed by the customer or end user depending on the application that has been created.

This server needs to run somewhere, on-premises, in a public cloud, serverless (Ok I have gone too far, we won't be covering serverless but its an option and more and more enterprises are heading this way) Someone needs to create and configure these servers and get them ready for the application to run. Now, this element might land to you as a DevOps engineer to deploy and configure these servers.

These servers run an operating system and generally speaking this is going to be Linux but we have a whole section or week where we cover some of the foundational knowledge you should gain here.

It is also likely that we need to communicate with other services in our network or environment, so we also need to have that level of knowledge around networking and configuring that, this might to some degree also land at the feet of the DevOps engineer. Again we will cover this in more detail in a dedicated section talking about all things DNS, DHCP, Load Balancing etc.

## Jack of all trades, Master of none

I will say at this point though, you don't need to be a Network or Infrastructure specialist you need a foundational knowledge of how to get things up and running and talking to each other, much the same as maybe having a foundational knowledge of a programming language but you don't need to be a developer. However, you might be coming into this as a specialist in an area and that is a great footing to adapt to other areas.

You will also most likely not take over the management of these servers or the application daily.

We have been talking about servers but the likelihood is that your application will be developed to run as containers, Which still runs on a server for the most part but you will also need an understanding of not only virtualisation, Cloud Infrastructure as a Service (IaaS) but also containerisation as well, The focus in these 90 days will be more catered towards containers.

## High-Level Overview

On one side we have our developers creating new features and functionality (as well as bug fixes) for the application.

On the other side, we have some sort of environment, infrastructure or servers which are configured and managed to run this application and communicate with all its required services.

The big question is how do we get those features and bug fixes into our products and make them available to those end users?

How do we release the new application version? This is one of the main tasks for a DevOps engineer, and the important thing here is not to just figure out how to do this once but we need to do this continuously and in an automated, efficient way which also needs to include testing!

This is where we are going to end this day of learning, hopefully, this was useful. Over the next few days, we are going to dive a little deeper into some more areas of DevOps and then we will get into the sections that dive deeper into the tooling and processes and the benefits of these.

## Resources

I am always open to adding additional resources to these readme files as it is here as a learning tool.

My advice is to watch all of the below and hopefully you also picked something up from the text and explanations above.

- [What is DevOps? - TechWorld with Nana](https://www.youtube.com/watch?v=0yWAtQ6wYNM)
- [What is DevOps? - GitHub YouTube](https://www.youtube.com/watch?v=kBV8gPVZNEE)
- [What is DevOps? - IBM YouTube](https://www.youtube.com/watch?v=UbtB4sMaaNM)
- [What is DevOps? - AWS](https://aws.amazon.com/devops/what-is-devops/)
- [What is DevOps? - Microsoft](https://docs.microsoft.com/en-us/devops/what-is-devops)

If you made it this far then you will know if this is where you want to be or not. See you on [Day 3](day03.md).
