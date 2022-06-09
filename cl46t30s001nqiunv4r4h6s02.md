## 90DaysOfDevOps 
DAY 5.

> Today we are focusing  on the  steps of continuous software development life-cycle of an Applications in  DevOps.


## Plan:

It all starts with the planning process this is where the development team gets together and figures out what types of features and bug fixes they're going to roll out in their next sprint. This is an opportunity as a DevOps Engineer for you to get involved with that and learn what kinds of things are going to be coming your way that you need to be involved with and also influence their decisions or their path and kind of help them work with the infrastructure that you've built or steer them towards something that's going to work better for them in case they're not on that path and so one key thing to point out here is the developers or software engineering team is your customer as a DevOps engineer so this is your opportunity to work with your customer before they go down a bad path.

## Code:

Now once that planning phase is  done, they're going to go start writing the code you may or may not be involved with this process of code writing as a Devops engineer , it depends. but taking this opportunity may help you grasp the concept of  infrastructure as Code as well as the infrastructures in place. Knowing what services are available and how to interact with those services is an added plus.

## Build:

This is where we'll kick off the first of our automation processes because we're going to take the code and  build it depending on what language they used.  It might require complication first and you might be  creating a docker image from that code either way we're going to go through that process using our ci/cd pipeline.

## Testing:

Once we've built it we're going to run some tests on it now the development team usually writes the test you may have some input in what tests get written but we need to run those tests and the testing is a way for us to try and minimize bugs in production, it doesn't guarantee (100% bug free pipeline though)  but we want to get as close so as not to  introduce new bugs and two not breaking things. 

## Release:

Once those tests pass we're going to do the release process and depending again on what type of application you're working on .  This is more like the process of taking your compiled code or the docker image that you've built and putting it into a registry or a repository where it's accessible by your production servers for the deployment process.

## Deploy:

The part i love most, deployment is like the end game of this whole thing because deployments is  when we push  the code to production and it's not until we do that that our business realizes the value of all the time, effort and hard work that you and the development team have put into this product.

## Operate:

Once it's deployed we are going to operate it and operate it may involves feedback from end-users  that about the Application performance and then possibly auto-scale the application
 (which  increase the number of servers available during peak periods and decrease the number of servers during off-peak periods either way that's all operational type metrics).
> Another operational thing might include a feedback loop from production back to your ops team letting you know about key events that happened in production such as a deployment back one step on the deployment thing this may or may not get automated depending on your environment the goal is to always automate it when possible there are some environments where you possibly need to do a few steps before you're ready to do that but ideally you want to deploy automatically as part of your automation process but if you're doing that it might be a good idea to include in your operational steps some type of notification so that your ops team knows that a deployment has happened.

## Continuous Monitoring

In this stage, you continuously monitor the performance of your application. Here the vital information about the use of the application is recorded. For example, system errors like low memory, and servers not being reachable are resolved in this stage. In this stage, the problem is fixed as soon as they got detected.

The popular tools that are used for this stage are Splunk, ELK Stack, Nagios, NewRelic, and Sensu. This stage helps in improving productivity and increases the reliability of the systems, which led to reducing the IT support costs. That’s why all the major IT companies have shifted to DevOps for building their products.

## Nice, Now  Repeat:


![Day5_DevOps8.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1654766090486/qCUlEI9x-.png align="left")

 > Once that's in place you go right back to the beginning to the planning stage and go through the whole thing again
> Continuous:
> 
> Many tools help us achieve the above continuous process, all this code and the ultimate goal of being completely automated, cloud infrastructure or any environment is often described as Continuous Integration/ Continuous Delivery/Continous Deployment or “CI/CD” for short. We will spend a whole week on CI/CD later on in the 90 Days with some examples and walkthroughs to grasp the fundamentals.

## Continuous Delivery:

Continuous Delivery = Plan > Code > Build > Test

## Continuous Integration:

This is effectively the outcome of the Continuous Delivery phases above plus the outcome of the Release phase. This is the case for both failure and success but this is fed back into continuous delivery or moved to Continuous Deployment.

Continuous Integration = Plan > Code > Build > Test > Release


## Continuous Deployment:

If you have a successful release from your continuous integration then move to Continuous Deployment which brings in the following phases

CI Release is Success = Continuous Deployment = Deploy > Operate > Monitor

You can see these three Continuous processes above as the simple collection of phases of the DevOps Lifecycle.

These DevOps stages are carried out on loop continuously till you achieve the desired product quality and that is why the DevOps cycle you saw above is in the shape of Infinite...  Anonymous

## Resources
%[https://www.youtube.com/watch?v=a0-uE3rOyeU]
%[https://www.youtube.com/watch?v=9pZ2xmsSDdo&t=125s]
%[https://www.youtube.com/watch?v=5pxbp6FyTfk]

You are simply amazing, See you on Day 6.
