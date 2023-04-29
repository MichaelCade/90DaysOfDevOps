# Zero Downtime Deployments

Another important part of your application lifecycle is deployment time. There are lots of strategies for deploying
software. Like with anything there are pros and cons to the various strategies so I will run through a few options from
least complex to most complex, and as you may imagine the most complex deployment types tend come with the highest
guarantees of uptime and least disruption to your customer.

You may be asking why it's important to consider how we deploy our applications as the vast majority of our application
lifecycle time will be in the “running” state and therefore we could focus our time on strategies that support our
running application’s resilience. My answer is: Have you ever been on-call? Almost all incidents are due to code
releases or changes. The first thing I do when im on-call and called to an incident is see what was recently deployed -
I focus my main attention on that component and more often than not it was to blame.

We do also need to consider that some of these deployment strategies will require us to make specific code changes or
application architecture decisions to allow us to support the specific deployment in question.

### Rolling Deployments

One of the simplest deployment strategies is a rolling deployment. This is where we slowly, one by one (or many be many,
depending on how many instances of a service you have) we replace old deployments with their new tasks. We can check
that the new deployments are healthy before moving onto the next, only have a few tasks not healthy at a time.

This is the default deployment strategy in Kubernetes. It actually borrows some characteristics from Surge, which is
coming next. It starts slightly more new tasks and waits for them to be healthy before removing old ones.

### Surge Deployments

Surge deployments are exactly what they sound like. We start a large number of new tasks before cutting over traffic to
those tasks and then draining traffic from our old tasks. This is a good strategy when you have high usage applications
that may not cope well with reducing their availability at all. Usually surge deployments can be configured to run a
certain percentage more than the existing tasks and then wait for them to be healthy before doing a cutover.

The problem with surge deployments is that we need a large capacity of spare compute resources to spin up a lot of new
tasks before rolling over and removing the old ones. This can work well where you have very elastic compute such as AWS
Fargate where you don’t need to provision more compute yourself.

### Blue/Green

The idea behind a Blue/Green deployment is that your entire stack (or application) is spun up, tested and then finally
once you are happy you change config to send traffic to the entire new deployment. Sometimes companies will always have
both a Blue and a Green stack running. This is a good strategy where you need very fast rollback and recovery to a known
good state. You can leave your “old” stack running for any amount of time once you are running on your new stack.

### Canary

Possibly one of the most complicated deployment strategies. This involves deploying a small number of your new
application and then sending a small portion of load to the new service, checking that nothing has broken by monitoring
application performance and metrics such as 4XX or 5XX error rates and then deciding if we continue with the deployment.
In advanced setups the canary controllers can do automatic rollbacks if error thresholds are exceeded.

This approach does involve a lot more configuration, code and effort.

Interestingly the name comes from from coal mining and the phrase "canary in the coal mine." Canary birds have a lower
tolerance to toxic gases than humans, so they were used to alert miners when these gases reached dangerous levels inside
the mine.

We use our metrics and monitoring to decide if our “canary” application is healthy and if it is, we then proceed with a
larger deployment.

## Application design considerations

You may have worked out by now that the more advanced deployment strategies require you to have both old and new
versions of your application running at once. This means that we need to ensure backwards compatibility with all the
other software running at the time. For instance, you couldn't use a database migration to rename a table or column
because the old deployment would no longer work.

Additionally, our canary deployment strategy requires our application to have health checks, metrics, good logging and
monitoring so that we can detect a problem in our specific canary application deployment. Without these metrics we would
be unable to programmatically decide if our new application works.

Both these considerations, along with others, mean that we need to spend extra time both on our application code,
deployment code and our monitoring and alerting stacks to take advantage of the most robust deployments.
