---
title: '#90DaysOfDevOps - The Big Picture: Data Management - Day 84'
published: false
description: 90DaysOfDevOps - The Big Picture Data Management
tags: 'devops, 90daysofdevops, learning'
cover_image: null
canonical_url: null
id: 1048747
---

## The Big Picture: Data Management

![](Images/Day84_Data1.png)

Data Management is by no means a new wall to climb, although we do know that data is more important than it maybe was a few years ago. Valuable and ever-changing it can also be a massive nightmare when we are talking about automation and continuously integrating, testing and deploying frequent software releases. Enter the persistent data and underlying data services are often the main culprit when things go wrong.

But before I get into Cloud-Native Data Management, we need to go up a level. We have touched on many different platforms throughout this challenge. Be it Physical, Virtual, Cloud or Cloud-Native obviously including Kubernetes there is none of these platforms that provide the lack of requirement for data management.

Whatever our business it is more than likely you will find a database lurking in the environment somewhere, be it for the most mission-critical system in the business or at least some cog in the chain is storing that persistent data on some level of the system.

### DevOps and Data

Much like the very start of this series where we spoke about the DevOps principles, for a better process when it comes to data you have to include the right people. This might be the DBAs but equally, that is going to include people that care about the backup of those data services as well.

Secondly, we also need to identify the different data types, domains, and boundaries that we have associated with our data. This way it is not just dealt with in a silo approach amongst Database administrators, storage engineers or Backup focused engineers. This way the whole team can determine the best route of action when it comes to developing and hosting applications for the wider business and focus on the data architecture vs it being an afterthought.

Now, this can span many different areas of the data lifecycle, we could be talking about data ingest, where and how will data be ingested into our service or application? How will the service, application or users access this data? But then it also requires us to understand how we will secure the data and then how will we protect that data.

### Data Management 101

Data management according to the [Data Management Body of Knowledge](https://www.dama.org/cpages/body-of-knowledge) is “the development, execution and supervision of plans, policies, programs and practices that control, protect, deliver and enhance the value of data and information assets.”

- Data is the most important aspect of your business - Data is only one part of your overall business. I have seen the term "Data is the lifeblood of our business" and most likely true. This then got me thinking about blood being pretty important to the body but alone it is nothing we still need the aspects of the body to make the blood something other than a liquid.

- Data quality is more important than ever - We are having to treat data as a business asset, meaning that we have to give it the considerations it needs and requires to work with our automation and DevOps principles.

- Accessing data in a timely fashion - Nobody has the patience to not have access to the right data at the right time to make effective decisions. Data must be available in a streamlined and timely manner regardless of presentation.

- Data Management has to be an enabler to DevOps - I mentioned streamlining previously, we have to include the data management requirements into our cycle and ensure not just availability of that data but also include other important policy-based protection of those data points along with fully tested recovery models with that as well.

### DataOps

Both DataOps and DevOps apply the best practices of technology development and operations to improve quality, increase speed, reduce security threats, delight customers and provide meaningful and challenging work for skilled professionals. DevOps and DataOps share goals to accelerate product delivery by automating as many process steps as possible. For DataOps, the objective is a resilient data pipeline and trusted insights from data analytics.

Some of the most common higher-level areas that focus on DataOps are going to be Machine Learning, Big Data and Data Analytics including Artificial Intelligence.

### Data Management is the management of information

My focus throughout this section is not going to be getting into Machine Learning or Artificial Intelligence but focus on the protecting the data from a data protection point of view, the title of this subsection is "Data management is the management of information" and we can relate that information = data.

Three key areas that we should consider along this journey with data are:

- Accuracy - Making sure that production data is accurate, equally we need to ensure that our data in the form of backups are also working and tested against recovery to be sure if a failure or a reason comes up we need to be able to get back up and running as fast as possible.
- Consistent - If our data services span multiple locations then for production we need to make sure we have consistency across all data locations so that we are getting accurate data, this also spans into data protection when it comes to protecting these data services, especially data services we need to ensure consistency at different levels to make sure we are taking a good clean copy of that data for our backups, replicas etc.

- Secure - Access Control but equally just keeping data, in general, is a topical theme at the moment across the globe. Making sure the right people have access to your data is paramount, again this leads to data protection where we must make sure that only the required personnel have access to backups and the ability to restore from those as well clone and provide other versions of the business data.

Better Data = Better Decisions

### Data Management Days

During the next 6 sessions we are going to be taking a closer look at Databases, Backup & Recovery, Disaster Recovery, and Application Mobility all with an element of demo and hands-on throughout.

## Resources

- [Kubernetes Backup and Restore made easy!](https://www.youtube.com/watch?v=01qcYSck1c4&t=217s)
- [Kubernetes Backups, Upgrades, Migrations - with Velero](https://www.youtube.com/watch?v=zybLTQER0yY)
- [7 Database Paradigms](https://www.youtube.com/watch?v=W2Z7fbCLSTw&t=520s)
- [Disaster Recovery vs. Backup: What's the difference?](https://www.youtube.com/watch?v=07EHsPuKXc0)
- [Veeam Portability & Cloud Mobility](https://www.youtube.com/watch?v=hDBlTdzE6Us&t=3s)

See you on [Day 85](day85.md)
