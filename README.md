# Blocksage - Bitcoin Block-Explorer
[![Build Status](https://semaphoreci.com/api/v1/ccdle12/blocksage/branches/master/badge.svg)](https://semaphoreci.com/ccdle12/blocksage)

The project currently consists of:
* Front-End: VueJS
* API:       Golang
* Cloud:     Docker/Kubernetes/GKE
* Bitcoin Node: Mainnet

# Requirements
* [Docker](https://docs.docker.com/install/#supported-platforms)
* [Docker Compose](https://docs.docker.com/compose/install/)

# Setup Project
Run the Project
```
$ make
```

# Run Tests
## API Test
* Unit Tests
```
$ make test-api-unittest
```

* Integration Tests
```
$ make test-api-integration
```

* All tests
```
$ make check
```
