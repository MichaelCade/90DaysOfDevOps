# APM Server extension

The APM Server receives data from APM agents and transforms them into Elasticsearch documents that can be visualised in
Kibana.

## Usage

To include APM Server in the stack, run Docker Compose from the root of the repository with an additional command line
argument referencing the `apm-server-compose.yml` file:

```console
$ docker-compose -f docker-compose.yml -f extensions/apm-server/apm-server-compose.yml up
```

Meanwhile, you can navigate to the **APM** application in Kibana and follow the setup instructions to get started.

## Connecting an agent to APM Server

The most basic configuration to send traces to APM server is to specify the `SERVICE_NAME` and `SERVICE_URL`. Here is an
example Python Flask configuration:

```python
import elasticapm
from elasticapm.contrib.flask import ElasticAPM

from flask import Flask

app = Flask(__name__)
app.config['ELASTIC_APM'] = {
    # Set required service name. Allowed characters:
    # a-z, A-Z, 0-9, -, _, and space
    'SERVICE_NAME': 'PYTHON_FLASK_TEST_APP',

    # Set custom APM Server URL (default: http://localhost:8200)
    'SERVER_URL': 'http://apm-server:8200',

    'DEBUG': True,
}
```

Configuration settings for each supported language are available in the APM documentation: [APM Agents][apm-agents].

## Checking connectivity and importing default APM dashboards

1. On the Kibana home page, click `Add APM` under the _Observability_ panel.
1. Click `Check APM Server status` to confirm the server is up and running.
1. Click `Check agent status` to verify your agent has registered properly.
1. Click `Load Kibana objects` to create an index pattern for APM.
1. Click `Launch APM` to be taken to the APM dashboard.

## See also

[Running APM Server on Docker][apm-docker]

[apm-agents]: https://www.elastic.co/guide/en/apm/guide/current/components.html
[apm-docker]: https://www.elastic.co/guide/en/apm/guide/current/running-on-docker.html
