# Smarter, Better, Faster, Stronger
#### Simulation Frameworks as the Future of Performance Testing
 

|      |                                                            |
| ----------- | -----------                                                     |
| By     | [Ada Lundhe](https://github.com/scorbettum/)                                                          |
| Where         | [Datavant](https://datavant.com/)                                               |
| Twitter    | [@sc_codeum](https://twitter.com/sc_codeum)                               |                       |
| Code Source    | [Hedra](https://github.com/scorbettum/hedra)                               |                       |
| Keywords    | devops, simulation-framework, distributed, graph, testing, performance


<br/>

## Overview

Performance testing has long been a critical part of a DevOps engineer's toolbox, allowing engineers to better understand the systems they build and maintain by simulating varying degrees of traffic at scale. However, existing performance frameworks are often limited in their ability to simulate realistic user scenarios at scale. Additionaly, performance testing frameworks are almost universally infamous for their poor developer/user experience, difficulty in integration with CI/CD and modern cloud environments, and lack of built-in quality reporting options.

<b>Simulation frameworks</b> represent the next step in performance testing, delivering on performance frameworks' core value proposition(s) while extending their functionality to embrace new testing techniques and modern developer needs. Simulation frameworks achieve this by:

- Allowing developers to write integration test like code using full programming languages that executes at the concurrency and with the speed of performance frameworks.

- Utilizing machine learning and statistical methods to provide features such as learned configuration values, A/B testing, etc.

- Providing chaos-testing facilities to allow request-level fault injection.

- Providing ample built-in reporting integrations to facilitate results submission to common DevOps analytics platforms like Datadog, Grafana, etc.

- Embracing multi-datacenter distributed execution as a core part of functionality while minimizing test code and configuration changes necessary to do so.

- Making developer/user experience a top priority via modern CLI interfaces, carefully constructed APIs, etc.

In doing so, simulation frameworks allow DevOps and infrastructure teams to:

- Reproduce the sophisticated usage patterns of actual users thus surfacing underlying infrastructure and/or application issues.

- Easily verify the impact of changes on different environments.

- Minimize the number of tests teams maintain by emphasizing parameterization while reducing time spent identifying "ideal" test parameters via learned configuration values.

- Ensure metrics from tests are available wherever it needs to be to maximize context and empower teams to make better decisions without having to maintain or compile external plugins or integrations.

- Minimize CI/CD or runtime environment complexity by reducing changes required to run between local, pipeline, and modern cloud environments.

<br/>

## Why Performance Testing?

Performance testing differs from this "functional testing" in that it focuses less on determining whether the target executes expected behaviors and more on how quickly and efficiently the target handles heavy amounts of traffic or usage. For distributed services, websites, and modern applications in general, this information can be critical in determining:

- Stability of services when a sudden spike in user traffic occurs

- CPU and memory usage of the target system under simulated increasing or defined traffic levels

- Identifying memory leaks or tricky transient issues

- Page load or API response times

as well as other target system behavior characteristics such as ready/write times to file or database, how quickly autoscaling adapts to simulated traffic or usage influx, etc. For these reasons, performance testing is now widely consider as a component of "chaos testing", or testing that determines how a system behaves under unexpected conditions.

However, we would argue that performance testing's domain and usefulness extend far beyond this area and are as critical a component of determining the quality, scalability, and durability of software as any type of functional testing. It's not enough that systems function, they must function well under a variety of usage and use cases.

> [A 2019 study by Portent](https://www.portent.com/blog/analytics/research-site-speed-hurting-everyones-revenue.htm) shows the impact of page load times on conversion rates. These load times are a product of both UI and API performance, and can increase drastically if a system is experiencing significant traffic.

Performance testing helps you provide essential metrics that illustrate what parts have the greatest impact on your application's the overall speed, stability, and efficiency. Hedra is one such tool that can provide these insights.

<br/>

## Limitations and Frustrations

Many load testing tools were developed back when running on bare-metal or small clusters of VMs was the predominant means of hosting/running software. Web applications were likewise in their infancy, so targeting and testing a handful of static URLs was perfectly acceptable.

Today's applications run in complicated cloud environments that are more sophisticated and distributed, such that targeting static URLs provides a woefully incomplete picture of application performance. Corequisitely, much of the tooling we use now needs to run in these resource-restrictive environments both to facilitate ease of use by developers and insure proper integration.

The majority of performance testing frameworks neglect or are outright incompatible these needs, placing the burden of customization and configuration on developers such to the point that performance test frameworks effectively require in-house engineering teams to deliver any sort of impactful insight. Frequent pain points include:

- Having to write a variety of custom execution engines to test with more modern protocols or libraries such as HTTP2, GRPC, UI testing via Playwright/Selenium, etc.

- Having to write custom integration reporting plugins or integrate clunky third-party options in order to publish test results to modern DevOps analytics platforms.

- Having to effectively micromanage deployment to avoid OOM or other compatability issues with CI/CD pipelines or distributed environments like Kubernetes.

- Having to "guess" appropriate testing configuration values not just for initial testing but after infrastructure or application changes, resulting in significant time wasted manually tweaking tests to get good signal on the impact of changes.

- Having to maintain a library of often similar tests due to performance frameworks not allowing for easy parameterization of tests and not offering A|B testing for simultaneously targeting multiple environments.

<br/>

## Simulaing Realism in Tests - Workflows as Graphs

When we want to execute a collection of tasks in some defined order, we often try to organize that work into contained steps and orchestrate them as a workflow. Computational frameworks such as Airflow ans Spark have particularly popularized and made evident the power of this approach for data science and data analysis.

What Airflow, Spark, and other "workflow" centeric tooling commonly share is the use of graphsto characterize the dependencies between tasks, determine and group execution order, and even provision required resources. Graphs are powerful data structures that make determining relationships between two disparate "things" computationally efficient.

The benefits of graphs in representing and managing workflows are numerous:

- Graphs make determining parallizable work easy

- Graphs make isolating and handling failure or errors in workflows efficient

- Graphs make authoring complex workflows natural

- Graphs make workflow progress visualization intuitive

Graphs also translate more naturally to distributed execution. Because Graphs allow us to better determine relationships between tasks we can then more easily isolate and delegate that work to disparte nodes in a cluster. We can also better handle failure not just of work but of nodes. Since graphs make keeping track of progress easy, we can simply decide whether we want a recovered node to resume that work, skip the work, or halt execution of the workflow as a whole.

Finally, graphs aow us to use a wide variety of interesting algorithms and possibilities - from shortest path algorithms helping us best determine how to optimize a workflow to probabilistic graphs allowing us to inject degrees of simulated human uncertainty.

<br/>

## Tests as Workflows

When writing tests, we are most familiar with writing test code as a series of steps to be executed in sequential order. When testing realistic usage of a system - while this representation might be suitable for simulating a user's surface level interactions - the underlying processes and events that triggered by those interactions are not sequential.

Consider a user submitting a form. The user enters some text into inputs, clicks some checkboxes, and then submits the form. On surface level this appears to be a perfectly sequential series of events, and thus can be simulated and tested as a sequential series of tasks in test code.

Behind the UI, modern applications and systems are executing a bevy of concurrent tasks for each sequential user "step" - from API calls to validate critical fields, to submitting user input to machine learning pipelines for analysis and recommendations as the user types, to capturing page interaction events, to batching calls to third party providers to return relevant advertisements, etc. Accurate tests capture more than the surface level interaction, they test and validate the complex interconnected work each interaction triggers.


```python

# An example simulating, from API level, an authorized user
# searching for a book by author.

import os
from hedra import (
    Workflow, 
    step,
)
from hedra.core.engines import (
    HTTPResult,
    HTTP2Result
)
from typing import Literal, Optional


class Test(Workflow):
    vus=1000
    duration='1m'
    username=os.getenv('USERNAME')
    password=os.getenv('PASSWORD')

    def get_book_title(self) -> Literal['Shakespeare Collected Works']:
        return 'Shakespeare Collected Works'
    
    @step()
    async def login_via_api(self) -> HTTP2Result:
        return await self.client.http2.post(
            'https://myapi.com/api/v1/login',
            auth=(
                self.username,
                self.password
            )
        )

    @step('login_via_api')
    async def get_book(
        self,
        auth_response: HTTP2Result
    ) -> HTTPResult:
        auth_token = response.headers.get('X-API-TOKEN')
        title = self.get_book_title()

        return await self.client.http.get(
            f'https://myapi.com/api/v1/books?title={title}',
            headers={
                'X-API-TOKEN': auth_token
            }
        )
    
    @step('login_via_api')
    async def get_author(
        self,
        auth_response: HTTP2Result
    ) -> HTTPResult:
        auth_token = response.headers.get('X-API-TOKEN')

        return await self.client.http.get(
            'https://myapi.com/api/v1/authors?author=william&shakespeare',
            headers={
                'X-API-TOKEN': auth_token
            }
        )

```

Given that the underlying pieces of work generated have clear and defined relationships with each other, it makes sense to both mentally and computationally model this work as a graph. The ideal test representitave of a system under usage is then also a graph composed of discrete tasks simulating and validating the functionality of the underlying components responsible for the work generated by user interaction events.

Ideally, integration and end-to-end tests would accomplish this. However these tests lack the granularity to validate the underlying work triggered by user interaction. Indvidual step code in integration or end-to-end tests becomes entagled and intertwined with other test code when attempting to accomplish this, and it becomes a difficult to maintain the distinct task boundares required by graph structures. Eventually, most integration and end-to-end tests devolve into sequential workflows for the sake of stability and scalability.

```
# As sequential steps
[Authorized User Search] -> 
[Get API Token] -> 
[Search for Author] -> 
[Search for Book]

# As a concurrent workflow

[Authorized User Search] -> 
[Get API Token] -> 
[Search cache for popular authors, send search analytics, hit third party search API via fan out] -> 
[Search cache for matching popular books, send search analytics, hit third party search API via fan out]
```

By contrast, unit tests are written to be fast, efficient, and discrete. This is because unit tests focus on validating the functionality of the smallest testable components of a system. When we break down the complex work generated by user interaction into its smallest components, we find that the work generated corresponds directly to orchestrated execution of these discrete components.

Instead of integration or end-to-end tests - to accurately test user impact on a system it makes more sense to compose unit tests into graph workflows.

However, the granularity of unit tests can make composing them into meaningful tests arduous compared to writing integration or end-to-end tests. The answer is a balance - tests that allow for assessing higher-level integrated functionality while retaining as much of the efficiency and independence of unit tests as possible. This sort of test allows our tests maximize the benefits of test workflow orchestration while ensuring test workflows do not become overly complex.

<br/>

### A|B Testing, Chaos, and Learned Configuration

One of the the most frequent (and awkward) questions engineers encounter when setting up performance testing is - "how should I setup this test"? While existing application analystics can provide insight into how an application performs now, we want to account for unexpected and future scenarios. 

Simply setting concurrency to maximum will likely cause the application to fail, but will provide no insight as to where issues <i>begin</i> to arise. Likewise, setting concurrency too low means the application will not be placed under proper stress, resulting in test results delivering no value. This challenge compounds when testing different environments (such as development or staging environments) which may have differing levels of resource allocation.

Conveniently, performance tests themselves contain a potential solution.Performance tests can involve millions of requests per-run and test results contain the contextual information we need (errors, status codes, etc.) to establish metrics that we can seek to maximize or minimize. For example, we could aim to maximize number of successful requests completed:

> <br/>
> R_success = R_total - R_error
>
> <br/>
<br/>

Or any other generally measurable outcome that could be expressed as some subset of the total set out output results. We could then express this via mathematical optimization as a loss function, where we seek to minimize the "distance" (error) from our goal:

> <br/>
> E_success = 1/(R_success) = 1/(R_total - R_error)
>
> <br/>
<br/>

As successful requests increase, the error decreases and trends towards a limit of zero.

More importantly, <i>most full programming languages contain libraries that allow for the automated optimization of this sort of function</i>. Then by repeatedly running short "subset" tests that output these error metrics, we can automate the identification of configuration values such as concurrency.

Frameworks can extend upon this functionality be embracing A|B testing to allow engineers to specify subsets of simulated traffic to divert to differing environments, API versions, etc. 

For example, we could define a test where:

- 20% of traffic is randomly diverted to the Development environment
- 40% randomly diverted to Staging
- 40% randomly diverted to Production

Using automated configuration, we can largely account for differences in environment resources since the framework can be set to automatically search for and find concurrency values that maximize successful requests. This then allows us to determine, from a single test, the actual impact changes in infrastructure resources first deployed to Development may have vs. the existing environments in Staging and Production.

We can also pair this with protocol level fault injection (i.e. sending a randomly selected subset of requests as intentionally malformed) to determine how changes adapt to and handle common attacks like request smuggling, etc. Although most libraries now protect from these sort of attacks, their handling can significantly slow processing of requests.

<br/>

## Integrations and Reporting

Simulating realism in tests takes more than orchestrating test execution a certain way - it requires being able to test an application at every level of the stack, from UI to DB.

The majority of existing performance frameowrks facilitate plain `HTTP/1.1` requests, with a subset supporting `HTTP/2`. However, modern applications often use UDP, Websockets, GraphQL, and more. Frameworks that facilitate additional protocols or libraries often only allow for the use of a single protocol during a test (i.e. only using `HTTP/1.1` <i>or</i> `HTTP/2.2` <i>or</i>  a third-party `Selenium` extension).

Simulation frameworks address these limitations by allowing for use of any supported protocols or libraries concurrently.

```python

import os
from hedra import (
    Workflow, 
    step,
)
from hedra.core.engines import (
    HTTPResult,
    HTTP2Result,
    PlaywrightResult
)
from typing import Literal, Optional


class Test(Workflow):
    vus=1000
    duration='1m'
    username=os.getenv('USERNAME')
    password=os.getenv('PASSWORD')

    def get_book_title(self) -> Literal['Shakespeare Collected Works']:
        return 'Shakespeare Collected Works'
    
    @step()
    async def login_via_api(self) -> HTTP2Result:
        return await self.client.http2.post(
            'https://myapi.com/api/v1/login',
            auth=(
                self.username,
                self.password
            )
        )
    
    @step()
    async def login_via_ui(self) -> PlaywrightResult:
        await self.client.playwright.goto('https://myapi.com/login')

        await self.client.playwright.input_text(
            '[data-test-id="username-input"]',
            self.username
        )

        await self.client.playwright.input_text(
            '[data-test-id="password-input"]',
            self.password
        )

        return await self.client.playwright.click('[data-test-id="login-button"]')

    @step('login_via_api')
    async def get_book(
        self,
        auth_response: HTTP2Result
    ) -> HTTPResult:
        auth_token = response.headers.get('X-API-TOKEN')
        title = self.get_book_title()

        return await self.client.http.get(
            f'https://myapi.com/api/v1/books?title={title}',
            headers={
                'X-API-TOKEN': auth_token
            }
        )
    
    @step('login_via_api')
    async def get_author(
        self,
        auth_response: HTTP2Result
    ) -> HTTPResult:
        auth_token = response.headers.get('X-API-TOKEN')

        return await self.client.http.get(
            'https://myapi.com/api/v1/authors?author=william&shakespeare',
            headers={
                'X-API-TOKEN': auth_token
            }
        )
    
    @step('login_via_ui')
    async def get_author_and_book_via_search(self) -> PlaywrightResult:
        await self.client.playwright.click('[data-test-id="author-search"]')
        await self.client.playwright.input_text(
            '[data-test-id="author-search-input"]',
            'William Shakespeare'
        )


        title = self.get_book_title()
        await self.client.playwright.click('[data-test-id="book-search"]')
        await self.client.playwright.input_text(
            '[data-test-id="book-search-input"]',
            title
        )

        return await self.client.playwright.click('[data-test-id="search-button"]')

```

While existing performance frameworks have made efforts to improve their reporting options, but often delegate the reponsibility of non-file/non-CLI output reporting to engineers, who must integrate or even write the extensions necessary. This increase in developer burden is often enough to prevent developers from integrating their tests into their workflows at all.

Simulation frameworks recognize the importance of integrations for reporting by offering plentiful options which are declared in-test and run concurrently. For example:

```python
import os
import statistics
from hedra import (
    Workflow, 
    step
    Metric
)
from hedra.core.engines import (
    HTTPResult,
    HTTP2Result,
    PlaywrightResult
)
from hedra.reporting (
    JSONResults,
    DatadogResults,
    KafkaResults
)
from typing import List


@depends(Test)
class SubmitResults(Workflow):
    reporters=[
        JSONResults(
            path='./events.json'
        ),
        DatadogResults(
            api_key=os.getenv('DD_API_KEY'),
            app_key=os.getenv('DD_APP_KEY')
        ),
        KafkaResults(
            host=os.getenv('KAFKA_HOST'),
            topic='myapi_testing_results'
        )
    ]

    @step()
    async def ui_vs_api_timing(
        self,
        results: List[
            HTTPResult |
            HTTP2Result |
            PlaywrightResult
        ]
    ) -> Metric['median_api_vs_ui_time_ratio']:

        avg_api_timing = statistics.median([
            result.total_time for result in results if isinstance(
                result,
                (HTTPResult, HTTP2Result)
            )
        ])

        ui_timings = statistics.median([
            result.total_time for result in results if isinstance(
                result,
                PlaywrightResult
            )
        ])

        return avg_api_timing/ui_timings

```

<br/>

## Developer Experience as a Priority

The primary source of adoption cost for existing performance testing frameworks is poor developer experience. Opaque CLI interfaces, sprawling APIs, and significant changes in test code or framework configuration to run tests via CI/CD or distributed vs local.

Drawing inspiration from modern web development tooling and frameworks K6, we can begin to improve by:

- Running tests via a single CLI whether locally or distributed.

- Embracing code generation to help test, devops, and application developers write tests more quickly via "starter" templates.

- Provide comprehensive, CLI-configurable metrics output of results to help provide additional visual feedback on test runs.

- Facilitating management of tests as "projects", collections of related work as opposed to offloading the entirety of test organization and management upon developers.

Code generation in particular is critical in helping developers rapidly prototype and develop tests. For example, a developer running the command

```bash
hedra test generate my-test --using http,http2,playwright --tags service=myapi.com,environment=staging
```

generates the following template code:


```python
# Generated by Hedra - my_test.py
import os
from hedra import (
    Workflow, 
    step,
)
from hedra.core.engines import (
    HTTPResult,
    HTTP2Result,
    PlaywrightResult
)


class MyTest(Workflow):
    vus=1000
    duration='1m'
    tags={
        'service': 'myapi.com',
        'environment': 'staging'
    }
    
    @step()
    async def get_http(self) -> HTTPResult:
        return await self.client.http.get('<ADD_URL_HERE>')

    @step()
    async def get_http2(self) -> HTTP2Result:
        return await self.client.http2.get('<ADD_URL_HERE>')

    @step()
    async def goto_url(self) -> PlaywrightResult:
        return await self.client.playwright.goto('<ADD_URL_HERE>')

```

When combined with linting:

```bash
hedra test lint my_test.py

Linting my_test.py...
OK!
```

project management features:
```bash
hedra test submit my_test.py --project github.com/myorg/tests

Submitting my_test.py to github.com/myorg/tests...
Repo updated!
```

And RPC remote execution:

```bash
hedra cloud test my_test.py --send staging

Sending to - staging - cluster at 155.020.313.33:6883
```

Allows developers to focus on value delivery as opposed to maintaining a plethora of of extensions, disorganized tests, and execution environments.

<br/>

## Summing it Up

While performance frameworks are valuable tools, their inherent limitations have made them difficult to adopt and their value proposition increasingly questionable. We can build upon their strengths by:

- Facilitating simulation of realistic user interactions via workflows
- Allowing for concurrent use of multiple protocols/libraries in a single test/workflow
- Embracing statistical frameworks like A|B testing, using optimization to automate configuration, and providing protocol-level fault injection for chaos testing
- Including modern developer experience features like starter template code generation, "unified experience" CLIs, test linting, and project management

to create a new class of tooling, simulation frameworks, that both deliver upon and exceed the value proposition of performance testing frameworks.