# Blocksage - Bitcoin Block-Explorer
[![Build Status](https://travis-ci.com/ccdle12/Blocksage.svg?branch=master)](https://travis-ci.com/ccdle12/Blocksage)

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
