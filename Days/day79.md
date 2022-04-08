---
title: '#90DaysOfDevOps - The Big Picture: Log Management - Day 79'
published: false
description: 90DaysOfDevOps - The Big Picture Log Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1049057
---
## The Big Picture: Log Management

A continuation to the infrastructure monitoring challenges and solutions, log management is another puzzle peice to the overall observability jigsaw. 

### Log Management & Aggregation 

Let's talk about two core concepts the first of which is log aggregation and it's a way of collecting and tagging application logs from many different services and to a single dashboard that can easily be searched. 

One of the first systems that have to be built out in an application performance management system is log aggregation. Application performance management is the part of the devops lifecycle where things have been built and deployed and you need to make sure that they're continuously working so they have enough resources allocated to them and errors aren't being shown to users. In most production deployments there are many related events that emit logs across services at google a single search might hit ten different services before being returned to the user if you got unexpected search results that might mean a logic problem in any of the ten services and log aggregation helps companies like google diagnose problems in production, they've built a single dashboard where they can map every request to unique id so if you search something your search will get a unique id and then every time that search is passing through a different service that service will connect that id to what they're currently doing. 

This is the essence of a good log aggregation platform efficiently collect logs from everywhere that emits them and make them easily searchable in the case of a fault again. 

### Example App 

Our example application is a web app, we have a typical front end and backend storing our critical data to a MongoDB database. 

If a user told us the page turned all white and printed an error message we would be hard-pressed to diagnose the problem with our current stack the user would need to manually send us the error and we'd need to match it with relevant logs in the other three services. 

### ELK 

Let's take a look at ELK, a popular open source log aggregation stack named after its three components elasticsearch, logstash and kibana if we installed it in the same environment as our example app. 

The web application would connect to the frontend which then connects to the backend, the backend would send logs to logstash and then the way that these three components work 

### The components of elk 

Elasticsearch, logstash and Kibana is that all of  services send logs to logstash, logstash takes these logs which are text emitted by the application. For example the web application when you visit a web page, the web page might log this visitor access to this page at this time and that's an example of a log message those logs would be sent to logstash.

Logstash would then extract things from them so for that log message user did **thing**, at **time**. It would extract the time and extract the message and extract the user and include those all as tags so the message would be an object of tags and message so that you could search them easily you could find all of the requests made by a specific user but logstash doesn't store things itself it stores things in elasticsearch which is a efficient database for querying text and elasticsearch exposes the results as Kibana and Kibana is a web server that connects to elasticsearch and allows administrators as the devops person or other people on your team, the on-call engineer to view the logs in production whenever there's a major fault. You as the administrator would connect to Kibana, Kibana would query elasticsearch for logs matching whatever you wanted. 

You could say hey Kibana in the search bar I want to find errors and kibana would say elasticsearch find the messages which contain the string error and then elasticsearch would return results that had been populated by logstash. Logstash would have been sent those results from all of the other services.

### how would we use elk to diagnose a production problem

A user says i saw error code one two three four five six seven when i tried to do this with elk setup we'd have to go to kibana enter one two three four five six seven in the search bar press enter and then that would show us the logs that corresponded to that and one of the logs might say internal server error returning one two three four five six seven and we'd see that the service that emitted that log was the backend and we'd see what time that log was emitted at so we could go to the time in that log and we could look at the messages above and below it in the backend and then we could see a better picture of what happened for the user's request and we'd be able to repeat this process going to other services until we found what actually caused the problem for the user.

### Security and Access to Logs 

An important peice of the puzzle is ensuring that logs are only visible to administrators (or the users and groups that absolutely need to have access), logs can contain sensitive information like tokens it's important that only authenticated users can access them you wouldn't want to expose Kibana to the internet without some way of authenticating.

### Examples of Log Management Tools

Examples of log management platforms there's

- Elasticsearch 
- Logstash 
- Kibana 
- Fluentd - popular open source choice
- Datadog - hosted offering, commonly used at larger enterprises, 
- LogDNA - hosted offering 
- Splunk 

Cloud providers also provide logging such as AWS CloudWatch Logs, Microsoft Azure Monitor and Google Cloud Logging. 


Log Management is a key aspect of the overall observability of your applications and instracture environment for diagnosing problems in production it's relatively simple to install a turnkey solution like ELK or CloudWatch and it makes diagnosing and triaging problems in production significantly easier. 

## Resources 

- [The Importance of Monitoring in DevOps](https://www.devopsonline.co.uk/the-importance-of-monitoring-in-devops/)
- [Understanding Continuous Monitoring in DevOps?](https://medium.com/devopscurry/understanding-continuous-monitoring-in-devops-f6695b004e3b) 
- [DevOps Monitoring Tools](https://www.youtube.com/watch?v=Zu53QQuYqJ0) 
- [Top 5 - DevOps Monitoring Tools](https://www.youtube.com/watch?v=4t71iv_9t_4)
- [How Prometheus Monitoring works](https://www.youtube.com/watch?v=h4Sl21AKiDg) 
- [Introduction to Prometheus monitoring](https://www.youtube.com/watch?v=5o37CGlNLr8)
- [Promql cheat sheet with examples](https://www.containiq.com/post/promql-cheat-sheet-with-examples)
- [Log Management for DevOps | Manage application, server, and cloud logs with Site24x7](https://www.youtube.com/watch?v=J0csO_Shsj0)
- [Log Management what DevOps need to know](https://devops.com/log-management-what-devops-teams-need-to-know/)
- [What is ELK Stack?](https://www.youtube.com/watch?v=4X0WLg05ASw)
- [Fluentd simply explained](https://www.youtube.com/watch?v=5ofsNyHZwWE&t=14s) 

See you on [Day 80](day80.md)
