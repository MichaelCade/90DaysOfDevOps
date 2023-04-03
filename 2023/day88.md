# Monitoring, Alerting and On-Call

As we have discussed there is no such thing as a perfect system that will never fail. Therefore we need to think about what to do when things go wrong. First and foremost the most important thing is not to panic. What tools and processes can we put in place to make incidents less stressful, easier to resolve and less likely to occur again?

## On-Call Practices

If you have never been “on-call” then you are an extremely lucky human. Being on-call for a system means being alerted when something is going wrong and being responsible for logging on, diagnosing the problem and helping to fix it. 

When we build software that people and businesses rely on we become liable for keeping those systems going. Our customers expect a certain level of uptime and availability for the services they are paying for. Let’s look at two examples, a Business-To-Business internal app for tracking month-end accounting and a real-time payments system with thousands of transactions a second with cumulative values of millions of pounds a minute. 


## Alerts: the good, bad and ugly

How do you even know if your application is working? 

Alerts.  These are notifications that inform you about the status or performance of your application. They can be configured to trigger when specific events or conditions occur, such as errors, timeouts, or spikes in traffic. I have seen `ERROR` log lines being used to trigger alerts as well as issues such as 4XX and 5XX http responses. You can practically alert on anything! If you wanted to know if your applications need to be scaled you could look at your message queue depth (number of messages in the queue) - if its growing or beyond a threshold you can either alert, or even better, you can automatically scale your application.
Alerts can be very useful in detecting and resolving issues quickly, which can help prevent downtime or other negative consequences. They allow you to proactively monitor your application and respond to issues before they become critical.
However, there can also be downsides to using alerts. If alerts are not properly configured or managed, they can generate a high volume of false alarms, which can be overwhelming and distract from real issues. This can lead to alert fatigue and decreased effectiveness in responding to issues.
It is important to carefully consider the types of alerts that are necessary for your application, and to regularly review and adjust your alert settings as needed. You should also have processes in place for responding to alerts, such as escalation procedures and clear documentation on how to resolve common issues.

## The cost of on-call

I have worked on systems similar to both of these hypothetical examples and I can tell you that the approach to on-call and incident management was wildly different. The month-end accounting software is only used by a handful of well-known customers sporadically - they may not log-in for many weeks before there being a flurry of activity for 2 days a month. The payments platform has constant and relentless use with many thousands of impacted people per minute. 

There’s an important message here around knowing your users, their tolerance for downtime and their usage patterns.
For the payments platform the on-call rota was 24-7 365 days a year. Every second counts and they expected an engineer to be online within a few minutes of being notified of a problem. This puts strain on your engineers and incident staff as even simple daily tasks become challenging - I can’t walk my dog today because I’m on-call and if called I need to be online within a few minutes. 

There’s very little fun about being on-call. You are tied to a location or always carrying your laptop, you get woken in the middle of your night and it's wildly unpredictable. If you can cope with some delay in fixing your system I highly recommend cutting your teams some slack and having good grace periods for getting online or muting alerts overnight. Obviously this isn't possible for all systems so where you need fast on-call cover make sure to keep an eye on your colleagues and look out for signs of stress and burnout.

