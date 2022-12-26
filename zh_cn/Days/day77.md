---
title: '#90DaysOfDevOps - The Big Picture: Monitoring - Day 77'
published: false
description: 90DaysOfDevOps - The Big Picture Monitoring
tags: "devops, 90daysofdevops, learning"
cover_image: null
canonical_url: null
id: 1048715
---
## The Big Picture: Monitoring

In this section we are going to talk about monitoring, what is it why do we need it? 

### What is Monitoring? 

Monitoring is the process of keeping a close eye on the entire infrastructure  

### and why do we need it? 

Let's assume we're managing a thousand servers these include a variety of specialised servers like application servers, database servers and web servers. We could also complicate this further with additional services and different platforms including public cloud offerings and Kubernetes. 

![](Images/Day77_Monitoring1.png)

We are responsible for ensuring that all the services, applications and resources on the servers are running as they should be. 

![](Images/Day77_Monitoring2.png)

How do we do it? there are three ways: 

- Login manually to all of our servers and check all the data pertaining to services processes and resources. 
- Write a script that logs in to the servers for us and checks on the data.  

Both of these options would require considerable amount of work on our part, 

The third option is easier, we could use a monitoring solution that is available in the market.  

Nagios and Zabbix are possible solutions that are readily available which allow us to upscale our monitoring infrastructure to include as many servers as we want. 

### Nagios

Nagios is an infrastructure monitoring tool that is made by a company that goes by the same name. The open-source version of this tool is called Nagios core while the commercial version is called Nagios XI. [Nagios Website](https://www.nagios.org/)

The tool allows us to monitor our servers and see if they are being sufficiently utilised or if there are any tasks of failure that need addressing. 

![](Images/Day77_Monitoring3.png)

Essentially monitoring allows us to achieve these two goals, check the status of our servers and services and determine the health of our infrastructure it also gives us a 40,000ft view of the complete infrastructure to see if our servers are up and running, if the applications are working properly and the web servers are reachable or not. 

It will tell us that our disk has been increasing by 10 percent for the last 10 weeks in a particular server, that it will exhaust entirely within the next four or five days and we'll fail to respond soon it will alert us when your disk or server is in a critical state so that we can take appropriate actions to avoid possible outages. 

In this case we can free up some disk space and ensure that our servers don't fail and that our users are not affected. 

The difficult question for most monitoring engineers is what do we monitor? and alternately what do we not? 

Every system has a number of resources, which of these should we keep a close eye on and which ones can we turn a blind eye to for instance is it necessary to monitor CPU usage the answer is yes obviously nevertheless it is still a decision that has to be made is it necessary to monitor the number of open ports in the system we may or may not have to depending on the situation if it is a general-purpose server we probably won't have to but then again if it is a webserver we probably would have to.  

### Continous Monitoring

Monitoring is not a new item and even continous monitoring has been an ideal that many enterprises have adopted for many years. 

There are three key areas of focus when it comes to monitoring. 

- Infrastructure Monitoring
- Application Monitoring 
- Network Monitoring 

The important thing to note is that there are many tools available we have mentioned two generic systems and tools in this session but there are lots. The real benefit of a monitoring solution comes when you have really spent the time making sure you are answering that question of what should we be monitoring and what shouldn't we? 

We could turn on a monitoring solution in any of our platforms and it will start grabbing information but if that information is simply too much then you are going to struggle to benefit from that solution, you have to spend the time to configure. 

In the next session we will get hands on with a monitoring tool and see what we can start monitoring. 

## Resources 

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b) 
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0) 
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg) 
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)

See you on [Day 78](day78.md)
