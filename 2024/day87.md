## Hands-on Performance Testing with k6

Performance testing is the testing practice that involves verifying how the system behaves and performs, often regarding speed and reliability. 

But let's start with one of the common questions. What is the difference between performance and load testing, and what types of tests exist? Some authors have minor differences in their definitions and categorizations. **Performance testing** and **load testing** are often used interchangeably, referring to the same concept.  

Performance testing is the umbrella term. Load tests are performance tests that inject 'load' over a period, typically simulated through concurrent users or requests per second. 

Given the type of load and their purpose, we find the common load test types:
- **Smoke tests**: minimal load.
- **Average-load tests**: load with usual traffic.
- **Stress tests**: load exceeding the usual traffic.
- **Spike tests**: a short peak of load. 
- **Soak tests**: load for an extended period.
- **Breakpoint tests**: gradual increase of load to find the system's limits.
- **Scalability tests**: load to verify scalability and elasticity policies.

Performance tests that run on a single thread include **browser performance tests** and **synthetic tests**, often called synthetic monitoring. Additionally, among others, there are tests designed for measuring the execution time of specific code (**performance profiling**) or service time in distributed systems.


### Non-scriptable vs Scriptable tools

If you are looking to create a performance test with load — let's just call it a load test ;) — two types of tools are at your disposal. Non-scriptable tools generate the load test based on one or multiple requests, perfect for simple scenarios. Scriptable tools, on the other hand, provide a wider range of functionalities and allow you to implement a script to simulate more realistic scenarios. 

There are dozens of load testing tools available. Some of the most popular ones are:
- **Non-scriptable tools**: [ab](https://httpd.apache.org/docs/2.4/programs/ab.html), [hey](https://github.com/rakyll/hey), [vegeta](https://github.com/tsenart/vegeta), [siege](https://github.com/JoeDog/siege). 

- **Scriptable tools**: [Gatling (Scala)](https://gatling.io/), [k6 (Javascript/Go)](https://k6.io/), [Locust (Python)](https://locust.io/).

[Wrk](https://github.com/wg/wrk) and [Artillery](https://www.artillery.io/) fall somewhere in the middle and worth a mention; wrk accepts LuaJIT scripts, and Artillery instructions are written in YAML files, similar to [Drill](https://github.com/fcsonline/drill).  Another category is UI-based tools, which includes popular ones such as [Postman](https://www.postman.com/) and [JMeter](https://jmeter.apache.org/). We also find benchmarking tools designed for specific protocols, like [ghz(gRPC)](https://ghz.sh/), [HammerDB(SQL DBs)](https://www.hammerdb.com/), [AMQP Benchmark](https://github.com/chirino/amqp-benchmark), and many more. 

Typically, non-scriptable tools are designed to set the load by specifying a request rate. In our first example, we aim to test how the [https://httpbin.org/get](https://httpbin.org/get) endpoint performs under a load of 20 requests per second for 30 seconds. With [Vegeta](https://github.com/tsenart/vegeta), we would run the following command:

```bash
echo "GET https://httpbin.org/get" | vegeta attack -duration=30s -rate=20 | vegeta report
```

In the other category, scriptable tools are designed to also set the load by specifying a number of concurrent users, also known as virtual users.  For this tutorial, we’ll use [k6](https://k6.io/), an extensible open-source performance testing tool that is written in Go and scriptable in Javascript.

### Run a simple k6 test

For learning purposes, let’s create a simple k6 script to test the previous endpoint with 10 concurrent users. When testing with concurrent users, it is recommended to add pauses between user actions, just like real users interact with our apps. In the following test, each user will wait for one second after the endpoint responds.

```javascript
// script.js

import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 10, // virtual users
  duration: '30s', // total test duration
};

export default function () {
  http.get('https://httpbin.org/get');
 
 // pause for 1 second
  sleep(1);
}
```

Run the script with the following command:

```bash
k6 run script.js
```

By default, k6 outputs the test results to `stdout`. We’ll see test results later.


### Model the test load (workload)

Scriptable tools offer more flexibility to model the workload than non-scriptable tools.

Two concepts are key to understanding how to model the load in k6: Virtual Users (VUs) and iterations. A Virtual User is essentially a thread that executes your function, while an iteration refers to one complete execution of that function.

k6 acts as a scheduler of VUs and iterations based on the test load configuration, providing multiple options for configuring the test load:

**(1) [vus](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#vus) and [iterations](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#iterations)**: specifying the number of virtual users and total number of executions of the default function. 

In the following example, k6 schedules 10 virtual users that will execute the default function 30 times in total (between all VUs). 

```javascript
export const options = {
  iterations: 30, // the total number of executions of the default function
  vus: 10, // 10 virtual users 
};
```

Replace the options settings in the previous script and run the test again. 

**(2) [vus](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#vus) and [duration](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#duration)**:  specifying the number of virtual users and the total test duration.

```javascript
export const options = {
  vus: 10, // virtual users
  duration: '30s', // total test duration
};
```

In the following example, k6 schedules 10 virtual users that will execute the default function continuously during 30 seconds. 

Replace the options settings in the previous script and run the test again. 

**(3) [stages](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#stages)**: specifying a list of periods that must reach a specific VU target.

```javascript
export const options = {
  stages: [
    // ramp up from 0 to 20 VUs over the next 20 seconds
    { duration: '20s', target: 20 },
    // run 20 VUs over the next minute
    { duration: '1m', target: 20 },
    // ramp down from 20 to 0 VUs over the next 20 seconds
    { duration: '20s', target: 0 },
  ],
};
```

Replace the options settings in the previous script and run the test again. 

**(4) [Scenarios](https://grafana.com/docs/k6/latest/using-k6/scenarios/)**: provide more options to model the workload and the capability to execute multiple functions, each with distinct workloads.

This tutorial cannot cover all the different scenario options, so please refer to the [Scenario documentation for further details](https://grafana.com/docs/k6/latest/using-k6/scenarios/). However, we’ll use scenarios to demonstrate an example that sets the load in terms of requests per second—20 requests per second for 30 seconds like the previous Vegeta example.

```javascript
import http from 'k6/http';

export const options = {
  scenarios: {
    default: {
      executor: 'constant-arrival-rate',

      // How long the test lasts
      duration: '30s',

      // Iterations rate. By default, the `timeUnit` rate is 1s. 
      rate: 20,


      // Required. The value can be the same as the rate.
      // k6 warns during execution if more VUs are needed.
      preAllocatedVUs: 20,
    },
  },
};

export default function () {
  http.get('https://httpbin.org/get');
}
```

This test uses the [`constant-arrival-rate` scenario](https://grafana.com/docs/k6/latest/using-k6/scenarios/executors/constant-arrival-rate/) to schedule a rate of iterations, telling k6 to schedule the execution of 20 iterations (`rate`: 20) per second (the default `timeUnit`).

Given that each iteration executes only one request, the test will run 20 requests per second. Mathematically speaking:
- 20 requests per second = 20 iterations per second x 1 request per iteration 

Now, run this test using the `k6 run` command.  You can see the test request rate by looking at the `http_reqs` metric, which reports the number of http requests and its request rate. In our example, it is close to our goal of 20 requests per second.

```bash
http_reqs......................: 601  19.869526/s
```


### Performance test results

After running the test, it’s common to assess how well your system handled the traffic. Typically, you can examine the test result data in two places: either from your testing tool or from your observability/monitoring solution. 

Performance testing tools report the response times (latency) and response errors (failures) during the test. This data can indicate whether the system is struggling or failing to handle the traffic, but it won’t tell you what issues are happening in the system. For root-cause analysis, turn to your observability solution to understand what’s happening internally.

In this section, we cover the fundamentals aspects of k6 metrics and an overview of the various test result options.  Basically, during test execution, k6 collects time-series data for built-in and custom metrics and these can be exported in different ways for further analysis.

Let’s start enumerating a few important metrics that k6 collects by default, also known as [built-in k6 metrics](https://grafana.com/docs/k6/latest/using-k6/metrics/reference/):
- `http_reqs`, to measure the number of requests (request rate).
- `http_req_failed`, to measure the error rate (errors).
- `http_req_duration`, to measure response times (latency).
- `vus`, to measure the number of virtual users (traffic).

k6 provides other [built-in k6 metrics](https://grafana.com/docs/k6/latest/using-k6/metrics/reference/) dependent on the [k6 APIs](https://grafana.com/docs/k6/latest/javascript-api/) used in your script, including other HTTP, gRPC, Websocket, or Browser metrics. After completing the test run, k6 outputs the aggregated results of the metrics to `stdout`.

```bash
data_received..................: 490 kB 16 kB/s
data_sent......................: 49 kB  1.6 kB/s
http_req_blocked...............: avg=30ms     min=1µs      med=2µs      max=565.41ms p(90)=4µs      p(95)=425.64ms
http_req_connecting............: avg=9.34ms   min=0s       med=0s       max=173.06ms p(90)=0s       p(95)=130.03ms
http_req_duration..............: avg=189.53ms min=126.01ms med=158.87ms max=1.21s    p(90)=209.95ms p(95)=372.09ms
  { expected_response:true }...: avg=189.53ms min=126.01ms med=158.87ms max=1.21s    p(90)=209.95ms p(95)=372.09ms
http_req_failed................: 0.00%  ✓ 0         ✗ 600 
http_req_receiving.............: avg=518.35µs min=22µs     med=173µs    max=23.02ms  p(90)=300.1µs  p(95)=759.39µs
http_req_sending...............: avg=480.17µs min=96µs     med=397µs    max=7.45ms   p(90)=627.4µs  p(95)=1.02ms  
http_req_tls_handshaking.......: avg=20.35ms  min=0s       med=0s       max=337.45ms p(90)=0s       p(95)=292.16ms
http_req_waiting...............: avg=188.53ms min=123.86ms med=158.04ms max=1.21s    p(90)=209.46ms p(95)=371.52ms
http_reqs......................: 600    19.926569/s
iteration_duration.............: avg=219.99ms min=126.44ms med=161.65ms max=1.21s    p(90)=449.25ms p(95)=611.17ms
iterations.....................: 600    19.926569/s
vus............................: 3      min=3       max=12
vus_max........................: 40     min=40      max=40
```

However, aggregated results hide a lot of information. It is more useful to visualize time-series graphs to understand what happened at different stages of the test.

![Grafana dashboard showing k6 results in real-time](https://grafana.com/api/dashboards/19665/images/14906/image)


What options do we have for visualizing time-series results? With k6, you can send the time-series data of all the k6 metrics to any backend. k6 provides a few options, others are available through output extensions, and you also have the option to implement your custom output. 

To learn about all the available options, refer to the [k6 real time output documentation](https://grafana.com/docs/k6/latest/results-output/real-time/).

#### Store results in Prometheus and visualize with Grafana

If you want to send [k6 time-series data to a Prometheus instance](https://grafana.com/docs/k6/latest/results-output/real-time/prometheus-remote-write/) as part of this tutorial, try using the [QuickPizza Demo](https://github.com/grafana/quickpizza#local-setup) with Docker compose. It runs a simple web app and, optionally, a Prometheus instance.

```bash
git clone git@github.com:grafana/quickpizza.git
cd quickpizza
docker compose -f docker-compose-local.yaml up -d
```

Visit QuickPizza at [localhost:3333](http://localhost:3333/), the Grafana instance at [localhost:3000](http://localhost:3000/), and the Prometheus instance at [localhost:9090](https://localhost:9090).

Now, let’s run a test and stream k6 metrics as time-series data to our local Prometheus instance using the [`--out`](https://grafana.com/docs/k6/latest/using-k6/k6-options/reference/#results-output) option. Run either the previous test or one of the examples in the QuickPizza repository:

```bash
k6 run -o experimental-prometheus-rw script.js

# or

k6 run -o experimental-prometheus-rw k6/foundations/01.basic.js
```

To visualize the performance results, visit the Grafana instance([localhost:3000](http://localhost:3000/)) and select the `k6 Prometheus` dashboard. You can also query k6 metrics from the Prometheus web UI or by using [Grafana Explore](https://grafana.com/docs/grafana/latest/explore/).

For in-depth overview about the various options shown in this section, refer to the following resources: 

- [k6 results output documentation](https://grafana.com/docs/k6/latest/results-output/)
- [Ways to visualize k6 results](https://k6.io/blog/ways-to-visualize-k6-results/)


### Define Pass/Fail criteria in k6 tests

In testing, an assertion typically refers to verifying a particular condition in the test: Is this true or false? For most testing tools, if an assertion evaluates as false, the test fails. 

k6 works slightly different in this regard, having two different APIs for defining assertions:

1. [Thresholds](https://grafana.com/docs/k6/latest/using-k6/thresholds/): used to establish the Pass/Fail criteria of the test.
2. [Checks](https://grafana.com/docs/k6/latest/using-k6/checks/): used to create assertions and inform about their status without affecting the Pass/Fail outcome.

Most of the testing tools provide an API to assert boolean conditions. k6 provides [checks](https://grafana.com/docs/k6/latest/using-k6/checks/) to validate boolean conditions in our tests, similar to assertions in other testing frameworks.

```javascript
check(res, {
  "status is 200": (res) => res.status === 200,
  "body includes URL": (res) => res.body.includes("https://httpbin.org/get"),
});
```

Why doesn't a test fail due to a check failure? It’s a design choice. Production systems typically don’t aim for 100% reliability; instead, we define error budgets and 'nines of availability,' accepting a certain percentage of errors. 

Some failures are expected 'Under Load'. A regular load test can evaluate thousands or millions of assertions; thus, by default, k6 won’t fail a test due to check failures.

Now, practice adding some checks to one of the previous tests. When the test ends, k6 prints the check results to the terminal:

```diff
+ ✓ status is 200
+ ✓ body includes URL
```

The other API is [Thresholds](https://grafana.com/docs/k6/latest/using-k6/thresholds/), specifically designed to set Pass/Fail criteria in our tests. Let’s define the success criteria of the test based on two of the golden signals: latency and errors: 

- 95% of requests must have a latency below 600ms.
- Less than 0.5% of requests must respond with errors. 

Thresholds evaluate the Pass/Fail criteria by querying k6 metrics using stats functions. Earlier, we reviewed the k6 metrics for latency and errors:
- `http_req_duration`, to measure response times (latency).
- `http_req_failed`, to measure the error rate (errors).

The previous Pass/Fail criteria can be added to the `options` object as follows:

```javascript
export const options = {
  ....
  thresholds: {
    http_req_duration: ['p(95)<600'],
    http_req_failed: ['rate<0.005'],
  },
};
```


Add these thresholds to the previous test and run the test again.  

When the test ends, k6 reports the threshold status with a green checkmark ✅ or red cross ❌ near the metric name. 

```bash
...
  ✓ http_req_duration..............: avg=106.24ms min=103.11ms....
  ✓ http_req_failed................: 0.00%  ✓ 0    ✗ 9 
...
default ✓ [======================================] 1 VUs 10s
```

When the test fails, k6 returns a non-zero error code, which is necessary when integrating on CI/CD pipelines. Now, please practice by changing the criteria and make the test fail. k6 will then report something similar to:

```bash
...
  ✗ http_req_duration..............: avg=106.24ms min=103.11ms....
  ✓ http_req_failed................: 0.00%  ✓ 0    ✗ 9 
...
default ✓ [======================================] 1 VUs 10s
ERRO[0011] thresholds on metrics 'http_req_duration' have been crossed
```

Thresholds are very flexible, allowing various cases such as:



- Defining thresholds on [custom metrics](https://grafana.com/docs/k6/latest/using-k6/metrics/create-custom-metrics/).
- Setting multiple thresholds for the same metric.
- Querying [tags](https://grafana.com/docs/k6/latest/using-k6/tags-and-groups/#tags) in metrics. 
- Aborting the test when threshold is breached.

For further details and examples, please refer to the [Thresholds documentation](https://grafana.com/docs/k6/latest/using-k6/thresholds/).


### Wrapping up

This tutorial aimed to provide a practical foundation in performance testing, particularly in the use of k6.  Through hands-on examples, we delved into k6's capabilities, including modeling workload, analyzing results, and establishing Pass/Fail criteria to verify SLO compliance.

I encourage you to continue exploring the depths of this topic. But more importantly, **adopt a proactive approach in reliability testing and don’t wait to start until critical failures happen**. [Automating performance testing](https://grafana.com/docs/k6/latest/testing-guides/automated-performance-testing/) will help you in your reliability efforts. 

Thank you for joining today! I hope you have learned something new and have sparked the curiosity in performance and reliability testing. For further resources, please visit: 

- [k6 documentation](https://grafana.com/docs/k6/latest/)
- [Load test types](https://grafana.com/docs/k6/latest/testing-guides/test-types/)
- [QuickPizza Demo - k6 Testing environment](https://github.com/grafana/quickpizza)
- [k6-awesome](https://github.com/grafana/awesome-k6)
- [How they load](https://github.com/aliesbelik/how-they-load)