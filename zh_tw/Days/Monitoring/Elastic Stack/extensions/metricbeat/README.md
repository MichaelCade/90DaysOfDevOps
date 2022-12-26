# Metricbeat

Metricbeat is a lightweight shipper that you can install on your servers to periodically collect metrics from the
operating system and from services running on the server. Metricbeat takes the metrics and statistics that it collects
and ships them to the output that you specify, such as Elasticsearch or Logstash.

## Usage

To include Metricbeat in the stack, run Docker Compose from the root of the repository with an additional command line
argument referencing the `metricbeat-compose.yml` file:

```console
$ docker-compose -f docker-compose.yml -f extensions/metricbeat/metricbeat-compose.yml up
```

## Configuring Metricbeat

The Metricbeat configuration is stored in [`config/metricbeat.yml`](./config/metricbeat.yml). You can modify this file
with the help of the [Configuration reference][metricbeat-config].

Any change to the Metricbeat configuration requires a restart of the Metricbeat container:

```console
$ docker-compose -f docker-compose.yml -f extensions/metricbeat/metricbeat-compose.yml restart metricbeat
```

Please refer to the following documentation page for more details about how to configure Metricbeat inside a
Docker container: [Run Metricbeat on Docker][metricbeat-docker].

## See also

[Metricbeat documentation][metricbeat-doc]

[metricbeat-config]: https://www.elastic.co/guide/en/beats/metricbeat/current/metricbeat-reference-yml.html
[metricbeat-docker]: https://www.elastic.co/guide/en/beats/metricbeat/current/running-on-docker.html
[metricbeat-doc]: https://www.elastic.co/guide/en/beats/metricbeat/current/index.html
