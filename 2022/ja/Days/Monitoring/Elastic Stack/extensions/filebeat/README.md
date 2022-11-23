# Filebeat

Filebeat is a lightweight shipper for forwarding and centralizing log data. Installed as an agent on your servers,
Filebeat monitors the log files or locations that you specify, collects log events, and forwards them either to
Elasticsearch or Logstash for indexing.

## Usage

To include Filebeat in the stack, run Docker Compose from the root of the repository with an additional command line
argument referencing the `filebeat-compose.yml` file:

```console
$ docker-compose -f docker-compose.yml -f extensions/filebeat/filebeat-compose.yml up
```

## Configuring Filebeat

The Filebeat configuration is stored in [`config/filebeat.yml`](./config/filebeat.yml). You can modify this file with
the help of the [Configuration reference][filebeat-config].

Any change to the Filebeat configuration requires a restart of the Filebeat container:

```console
$ docker-compose -f docker-compose.yml -f extensions/filebeat/filebeat-compose.yml restart filebeat
```

Please refer to the following documentation page for more details about how to configure Filebeat inside a Docker
container: [Run Filebeat on Docker][filebeat-docker].

## See also

[Filebeat documentation][filebeat-doc]

[filebeat-config]: https://www.elastic.co/guide/en/beats/filebeat/current/filebeat-reference-yml.html
[filebeat-docker]: https://www.elastic.co/guide/en/beats/filebeat/current/running-on-docker.html
[filebeat-doc]: https://www.elastic.co/guide/en/beats/filebeat/current/index.html
