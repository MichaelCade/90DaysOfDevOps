# Logspout extension

Logspout collects all Docker logs using the Docker logs API, and forwards them to Logstash without any additional
configuration.

## Usage

If you want to include the Logspout extension, run Docker Compose from the root of the repository with an additional
command line argument referencing the `logspout-compose.yml` file:

```bash
$ docker-compose -f docker-compose.yml -f extensions/logspout/logspout-compose.yml up
```

In your Logstash pipeline configuration, enable the `udp` input and set the input codec to `json`:

```logstash
input {
  udp {
    port  => 5000
    codec => json
  }
}
```

## Documentation

<https://github.com/looplab/logspout-logstash>
