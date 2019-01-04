[![GoDoc](https://godoc.org/github.com/shazow/ssh-chat?status.svg)](https://godoc.org/github.com/nitishm/vegeta-server) [![Go Report Card](https://goreportcard.com/badge/gojp/goreportcard)](https://goreportcard.com/report/nitishm/vegeta-server) [![Build Status](https://travis-ci.org/shazow/ssh-chat.svg?branch=master)](https://travis-ci.org/nitishm/vegeta-server) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/nitishm/vegeta-server/blob/master/LICENSE) [![Coverage Status](https://coveralls.io/repos/github/nitishm/vegeta-server/badge.svg?branch=master)](https://coveralls.io/github/nitishm/vegeta-server?branch=master)
# Vegeta Server - A RESTful load-testing service

> NOTE: This is currently a work-in-progress. This means that features will be added/modified/removed at will. Do not use this in production.

A RESTful API server for [vegeta](https://github.com/tsenart/vegeta), a load testing tool written in [Go](https://github.com/golang/go).

Vegeta is a versatile HTTP load testing tool built out of a need to drill HTTP services with a constant request rate. The vegeta library is written in Go, which makes it ideal to implement server in Go.  

The REST API model enables users to, asynchronously, submit multiple `attack`s targeting the same (or varied) endpoints, lookup pending or historical `report`s and update/cancel/re-submit `attack`s, using a simple RESTful API.

## Getting Started

### Installing

```
make all
```
, which vendors all the dependencies, builds the `vegeta-server` binary and drops it in the `/bin` directory.

### Usage

Start the server using the `vegeta-server` binary generated after the previous step.

The `vegeta-server` supports `flags` to pass configuration options to the server such as , **scheme** [`http` / `https`], **host**, **port** and `TLS` configurations, among others.

```
Usage:
  vegeta-server [OPTIONS]

This is a RESTful API for the vegeta load-testing utility. Vegeta is a versatile HTTP load testing tool built out of a need to drill HTTP services with a constant request rate.


Application Options:
      --scheme=            the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec
      --cleanup-timeout=   grace period for which to wait before killing idle connections (default: 10s)
      --graceful-timeout=  grace period for which to wait before shutting down the server (default: 15s)
      --max-header-size=   controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request
                           body. (default: 1MiB)
      --socket-path=       the unix socket to listen on (default: /var/run/vegeta.sock)
      --host=              the IP to listen on (default: localhost) [$HOST]
      --port=              the port to listen on for insecure connections, defaults to a random value [$PORT]
      --listen-limit=      limit the number of outstanding requests
      --keep-alive=        sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download) (default: 3m)
      --read-timeout=      maximum duration before timing out read of the request (default: 30s)
      --write-timeout=     maximum duration before timing out write of the response (default: 60s)
      --tls-host=          the IP to listen on for tls, when not specified it's the same as --host [$TLS_HOST]
      --tls-port=          the port to listen on for secure connections, defaults to a random value [$TLS_PORT]
      --tls-certificate=   the certificate to use for secure connections [$TLS_CERTIFICATE]
      --tls-key=           the private key to use for secure conections [$TLS_PRIVATE_KEY]
      --tls-ca=            the certificate authority file to be used with mutual tls auth [$TLS_CA_CERTIFICATE]
      --tls-listen-limit=  limit the number of outstanding requests
      --tls-keep-alive=    sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)
      --tls-read-timeout=  maximum duration before timing out read of the request
      --tls-write-timeout= maximum duration before timing out write of the response

Help Options:
  -h, --help               Show this help message
```

#### Example 
*Serve `HTTP` traffic at `localhost:8000/api/v1`*
```
./bin/vegeta-server --scheme=http --host=localhost --port=8000
```

> **Bonus**
> 
> The `localhost:8000` example can be also started using the `run-local` target, which handles the build and starts the server.
> ```
> make run-local
> ```

### Running tests

```
make test
```

## Benchmark

*TODO*

## Guides & Tutorials

*TODO*

## Contributing

Link to [CONTRIBUTING.md](https://github.com/nitishm/vegeta-server/blob/master/CONTRIBUTING.md)

## Roadmap

*TODO* 

## License

Link to [LICENSE](https://github.com/nitishm/vegeta-server/blob/master/LICENSE)

## Support

Contact Author at nitish.malhotra@gmail.com

---

## Swagger - API Specification
The API's [swagger](https://swagger.io/) specification is formatted using the [OpenAPI 2.0](https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md) specification.

### Codegen
The server code is generated using [go-swagger](https://github.com/go-swagger/go-swagger), a tool, written in Go, that implements the Swagger 2.0 specification.