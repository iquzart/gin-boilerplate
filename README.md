# Gin Boilerplate
[![Build and Test](https://github.com/iquzart/gin-boilerplate/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/iquzart/gin-boilerplate/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/iquzart/gin-boilerplate)](https://goreportcard.com/report/github.com/iquzart/gin-boilerplate)
![GitHub](https://img.shields.io/github/license/iquzart/gin-boilerplate) 
![Metrics Support](https://img.shields.io/badge/Metrics%20Support-Prometheus-blue)


Golang Gin Boilerplate

Features
--------
1. Configurable Graceful Shutdown
2. Custom ports
2. Health check for Kubernetes
3. Prometheus Metrics

Environment Veriables
---------------------

| Variable | Description | Default |
| --- | --- | --- |
| PORT | Application port | 8080 |
| GIN_MODE | Gin Modes - debug, release, test | release |
| ENABLE_GRACEFUL_SHUTDOWN | To enable graceful shutdown of the application | true |
| API_VERSION | Set API version for the application  |  v1.0.0  |


License
-------

MIT


Author Information
------------------

Muhammed Iqbal <iquzart@hotmail.com>