# Designing for Resilience, Redundancy and Reliability

We now have an application which uses asynchronous queue based messaging to communicate. This gives us some real
flexibility on how we design our system to withstand failures. We are going to look at the concept of Failure Zones,
Replication, Contract testing, Logging and Tracing to build more robust systems.

## Failure  Zones

Imagine building an application and deploying everything onto a single VM or server. What happens when, inevitably, this
server fails. Your application goes offline and your customers won’t be happy! There’s a fix to this. Spread your
workloads over multiple points of failure. This doesn’t just go for your application instances but you can build
multiple redundant points for every aspect of your system.

Ever wonder what some of the things large cloud providers do to keep their services running despite external and
unpredictable factors? For starters they will have generators on-site for when they inevitably suffer a power cut. They
have at least two internet connections into the datacentres and they often provide multiple availability zones in each
region. Take Amazon’s eu-west-2 (London) region. This has 3 availability zones, eu-west-2a, eu-west-2b and eu-west-2c.
There are 3 separate datacenters (physically separated) that all inter-connect to provide a single region. This means by
deploying services across these availability zones (AZs) we increase our redundancy and resilience over factors such as
a fire in one of these facilities. If you run a homelab you could distribute work over failure zones by running more
than one computer, placing these computers in separate parts of the house or buying 2 internet connections to stay
online if one goes down (This does mean checking that these connections don’t just run on the same infrastructure as
soon as they leave your door. At my house I can get fibre to my house and also a connection over the old phone lines)

## Replication

This links nicely into having multiple replicas of our “stuff”, this doesn’t just mean running 2 or 3 copies of our
application over our previously mentioned availability zones we also need to consider our data - If we ran a database in
AZ1 and our application code over AZ1, AZ2 and AZ3 what do you think would happen if AZ1 was to go offline, potentially
permanently? There have been instances of cloud datacenters being completely destroyed by fire and many customers had
not been backing up their data or running across multiple AZs. Your data in this case is… gone. Do you think the
business you workin in could survive if their data, or their customers data, just disappeared overnight?

Many of our data storage tools come with the ability to be configured for replication across multiple failure zones.
NATS.io that we used previously can be configured to replicate messages across multiple instances to survive failure of
one or more zones. Postgresql databases can be configured to stream their WAL (Write ahead log), which stores an
in-order history of all the transactions, to a standby instance running somewhere else. If our primary fails then the
most we would lose would be the amount of data in our WAL that was not transferred to the replica. Much less data loss
than if we didn’t have any replication.

## Contract Testing

We are going to change direction now and look at making our applications more reliable. This means reducing the change
of failures. You may appreciate that the time most likely to cause failures in your system is during deployments. New
code hits our system and if it has not been tested thoroughly in production-like environments then we may end up with
undefined behaviours.

There’s a concept called Contract testing where we can test the interfaces between our applications at development and
build time. This allows us to rapidly get feedback (a core principle of DevOps) on our progress.

## Async programming and queues

We have already discussed how breaking the dependencies without our system can lead to increased reliability. Our
changes become smaller, less likely to impact other areas of our application and easy to roll-back.

If our users are not expecting an immediate tactile response to an event, such as requesting a PDF be generated then we
can always place a message onto a queue and just wait for the consumer of that message to eventually get round to it. We
could see a situation where thousands of people request their PDF at once and an application that does this
synchronously would just fall over, run out of memory and collapse. This would leave all our customers without their
PDFs and needing to take an action again to wait for them to be generated. Without developer intervention we may not get
back to a state where the service can recover.

Using a queue we can slowly work away in the background to generate these PDFs and could even scale the service in the
background automatically when we noticed the queue getting longer. Each new application would just take the next message
from the queue and work through the backlog. Once we were seeing less demand we could automatically scale these services
down when our queue depth reached 0.

## Logging and Tracing

It goes without saying that our applications and systems will fail. What we are doing in this section of 90daysOfDevOps
is thinking about what we can do to make our lives less bad when things do. Logging and tracing are some really
important tools in our toolbox to keep ourselves happy.

If we log frequently with both success and failure messages we can follow the progress of our requests and customer
journeys through our system then when things go wrong we can quickly rule out specific services or focus down on
applications that are logging warning or error messages. My general rule is that you can’t log too much - its not
possible! It is however important to use a log framework that lets you tune the log level that gets printed to your
logs. For example if i have 5 log levels (as is common), TRACE, DEBUG, INFO, WARN and ERROR we should have a mechanism
for ever application to set the level of logs we want to print. Most of the time we only want WARN and ERROR to be
visible in logs, to quickly show us theres and issue. However if we are in development or debugging a specific
application its important to be able to turn up the verbosity of our logs and see those INFO/DEBUG or TRACE levels.

Tracing is a concept of being able to attack a unique identifier to a request in our system that then gets populated and
logged throughout that requests journey, we could see a HTTP request hit a LoadBalancer, get a correlationID and then we
want to see that correlationID in ever log line as our request’s actions percolate through our system.

## Conclusion

It’s hard to build a fully fault-tolerant system. It involves learning from our, and other people’s, mistakes. I have
personally made many changes to company infrastructure after we discovered a previously unknown failure zone in our
application.

We could change our application running in Kubernetes to tolerate our underlying infrastructure’s failure zones by
leveraging [topology spread constraints](https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/)
to spread replicas around. We won’t in this example as we have much more to cover!
