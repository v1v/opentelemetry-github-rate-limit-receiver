# opentelemetry-github-rate-limit-receiver

## Local Development

### Prerequisites

- [Go](https://golang.org/dl/)
- [Gh](https://cli.github.com/)

### Install tools

#### Ngrok

The [Ngrok](https://ngrok.com/download/) can be installed by running:

```shell
make install-ngrok
```

#### ocb

The [OpenTelemetry Collector Builder (OCB)](https://opentelemetry.io/docs/collector/custom-collector/) can be installed by running:

```shell
make install-ocb
```

### Build the collector

```shell
make build
```

### Run the collector

```shell
make run
```