# Day 87 - Hands-on Performance Testing with k6
[![Watch the video](thumbnails/day87.png)](https://www.youtube.com/watch?v=Jg4GRzRHX9M)

The session covers (k6) which collects both built-in and custom metrics. The built-in metrics include:

1. Current/Active Virtual Users: The current number of active virtual users in the test.
2. Iteration Information: Information about the iteration such as the type.
3. HTTP Request Metrics (Rate Metric): Information about the HTTP requests, including the request rate and response time (time spend on the request excluding the connection time).
4. Other HTTP Request Metrics (Red Metric): Information about percentiles of the HTTP requests collected from all data points and giving a percent statistic. Other counter metrics like HTTP Rex give real-time data on the terminal.

Custom metrics are those that you define in your test using the k6 module with the type of metric set as either Trend or Counter. For example, in quick pix sample number four, we defined a custom metric 'pixza have more than six ingredient' using the check API to verify a condition and as many checks as desired can be added for any request or object.

Assertions in K6 can be defined using either the check API or the StresScore API. Checks don't make the test fail even if they do, but you can use StresScore to define the pass/fail criteria of your test based on Casic metrics. In the example provided, a stress test was defined to specify that only 1% of HTTP requests should fail and that 95% of the request should be below half a second and 99% of the request should be below 1 second for response times, while also setting a threshold based on the custom metric (average should be below two). If this condition fails, the test will report a failure and an informative message.

You can learn more about K6's features, testing guides, and extensions from the Casic documentation. There are also many community-built extensions available to help with specific use cases.
**IDENTITY and PURPOSE**

**Custom Metrics and Purposes**

* Custom metrics can be defined using the `counter` and `trend` types.
* An example of a custom metric is tracking the number of pixels (Quick Pixa) and ingredients returned in a test.
* The purpose of these metrics is to gain insights into system behavior, identify trends, and optimize performance.

**End-of-Test Results**

* K6 provides three options for displaying end-of-test results: terminal output, CSV files, and custom summaries.
* Custom summaries allow users to customize the format and content of the test results.

**Real-time Data Collection and Visualization**

* K6 provides real-time data collection capabilities through its Output module.
* This allows users to send data points to various destinations, such as Prometheus or Grafana.
* These tools enable visualization and monitoring of system performance in real-time.

**Assertions and Thresholds**

* Assertions are used to define pass/fail criteria for test results.
* Two types of assertions are available: `check` and `stress`.
* Stress assertions allow users to set custom failure thresholds based on specific metrics or conditions.
