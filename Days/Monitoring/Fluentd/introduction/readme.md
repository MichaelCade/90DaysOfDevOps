# Introduction to Fluentd

## Collecting logs from files

Reading logs from a file we need an application that writes logs to a file. <br/>
Lets start one:

```
cd monitoring\logging\fluentd\introduction\

docker-compose up -d file-myapp

```

To collect the logs, lets start fluentd

```
docker-compose up -d fluentd
```

## Collecting logs over HTTP (incoming)

```
cd monitoring\logging\fluentd\introduction\

docker-compose up -d http-myapp

```