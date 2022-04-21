---
title: '#90DaysOfDevOps - Application Focused - Day 3'
published: false
description: 90DaysOfDevOps - Application Focused
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048825
---
## DevOps Lifecycle - Application Focused

As we continue through these next few weeks we are 100% going to come across these titles (Continuous Development, Testing, Deployment, Monitor) over and over again, If you are heading towards the DevOps Engineer role then repeatability will be something you will get used to but constantly enhancing each time is another thing that keeps things interesting. 

In this hour we are going to take a look at the high level view of the application from start to finish and then back round again like a constant loop. 

### Development 
Let's take a brand new example of an Application, to start with we have nothing created, maybe as a developer you have to discuss with your client or end user on the requirements and come up with some sort of plan or requirements for your Application. We then need to create from the requirements our brand new application. 

In regards to tooling at this stage there is no real requirement here other than choosing your IDE and the programming language you wish to use to write your application. 

As a DevOps engineer, remember you are probably not the one creating this plan or coding the application for the end user, this will be a skilled developer. 

But it also would not hurt for you to be able to read some of the code so that you can make the best infrastructure decisions moving forward for your application.

We previously mentioned that this application can be written in any language. Importantly this should be maintained using a version control system, this is something we will cover also in detail later on and in particular we will dive into **Git**. 

It is also likely that it will not be one developer working on this project although this could be the case but even so best practices would require a code repository to store and collaborate on the code, this could be private or public and could be hosted or privately deployed generally speaking you would hear the likes of **GitHub or GitLab** being used as a code repository. Again we will cover these as part of our section on **Git** later on. 


### Testing 
At this stage we have our requirements and we have our application being developed. But we need to make sure we are testing our code in all the various different environments that we have available to us or specifically maybe to the programming language chosen. 

This phase enables QA to test for bugs, more frequently we see containers being used for simulating the test environment which overall can improve on cost overheads of physical or cloud infrastructure. 

This phase is also likely going to be automated as part of the next area which is Continuous Integration.

The ability to automate this testing vs 10s,100s or even 1000s of QA engineers having to do this manually speaks for itself, these engineers can focus on something else within the stack to ensure you are moving faster and developing more functionality vs testing bugs and software which tends to be the hold up on most traditional software releases that use a waterfall methodology. 

### Integration 

Quite importantly Integration is at the middle of the DevOps lifecycle. It is the practice of in which developers require to commit changes to the source code more frequently. This could be on a daily or weekly basis. 

With every commit your application can go through the automated testing phases and this allows for early detection of issues or bugs before the next phase. 

Now you might at this stage be saying "but we don't create applications, we buy it off the shelf from a software vendor" Don't worry many companies do this and will continue to do this and it will be the software vendor that is concentrating on the above 3 phases but you might want to still adopt the final phase as this will enable for faster and more efficient deployments of your off the shelf deployments. 

I would also suggest just having this above knowledge is very important as you might buy off the shelf software today, but what about tomorrow or down the line... next job maybe? 

### Deployment 
Ok so we have our application built and tested against the requirements of our end user and we now need to go ahead and deploy this application into production for our end users to consume. 

This is the stage where the code is deployed to the production servers, now this is where things get extremely interesting and it is where the rest of our 86 days dives deeper into these areas. Because different applications require different possibly hardware or configurations. This is where **Application Configuration Management** and **Infrastructure as Code** could play a key part in your DevOps lifecycle. It might be that your application is **Containerised** but also available to run on a virtual machine. Which then also leads us onto platforms like **Kubernetes** which would be orchestrating those containers and making sure you have the desired state available to your end users. 

All of these bold topics we will go into more detail over the next few weeks to get a better foundational knowledge of what they are and when to use them. 

### Monitoring 

Things are moving fast here and we have our Application that we are continuously updating with new features and functionality and we have our testing making sure no gremlins are being found. We have the application running in our environment that can be continually keeping the required configuration and performance. 

But now we need to be sure that our end users are getting the experience they require. Here we need to make sure that our Application Performance is continuously being monitored, this phase is going to allow your developers to make better decisions about enhancements to the application in future releases to better serve the end users. 

This section is also where we are going to capture that feedback wheel about the features that have been implemented and how the end users would like to make these better for them. 

Reliability is a key factor here as well, at the end of the day we want our Application to be available all the time it is required. This then lends to other **observability, security and data management** areas that should be continuously monitored and feedback can always be used to better enhance, update and release the application continuously. 

Some input from the community here specifcally [@_ediri](https://twitter.com/_ediri) mentioned also part of this continous process we should also have the FinOps teams involved. Apps & Data are running and stored somewhere you should be monitoring this continously to make sure if things change from a resources point of view your costs are not causing some major financial pain on your Cloud Bills.

I think it is also a good time to bring up the "DevOps Engineer" mentions above, albeit there are many DevOps Engineer positions in the wild that people hold, this is not really the ideal way of positioning the process of DevOps. What I mean is from speaking to others in the community the title of DevOps Engineer should not be the goal for anyone because really any position should be adopting DevOps processes and the culture explained here. DevOps should be used in many different positions such as Cloud-Native engineer/architect, virtualisation admin, cloud architect/engineer, infrastructure admin. This is to name a few but the reason for using DevOps Engineer above was really to highlight the scope or the process used by any of the above positions and more. 

## Resources 

I am always open to adding additional resources to these readme files as it is here as a learning tool.  

My advice is to watch all of the below and hopefully you also picked something up from the text and explanations above. 

- [Continuous Development](https://www.youtube.com/watch?v=UnjwVYAN7Ns) I will also add that this is focused on manufacturing but the lean culture can be closely followed with DevOps. 
- [Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM)
- [Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI)
- [Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)
- [The Remote Flow](https://www.notion.so/The-Remote-Flow-d90982e77a144f4f990c135f115f41c6)
- [FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)
- [**NOT FREE** The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.co.uk/Phoenix-Project-DevOps-Helping-Business-ebook/dp/B00AZRBLHO)

If you made it this far then you will know if this is where you want to be or not. See you on [Day 4](day04.md).  
