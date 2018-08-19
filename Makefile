GOCMD=go
GOTEST=$(GOCMD) test

DC=docker-compose
DE=docker exec -i -t
B=/bin/bash -c
A=api
BC=bitcoinclient
C=crawler
DE-API=$(DE) $(A) $(B)
DE-CRAWLER=$(DE) $(C) $(B)

CDEV=-f docker-compose.yml


# Run Project
default:
	$(DC) up

build:
	$(DC) $(CDEV) build

detach:
	$(DC) $(CDEV) up -d

adminer:
	open http://localhost:8080

# Tests
check:
	make test-api
	make test-crawler

# Run all Tests API
test-api: 
	make unit-test-api
	make integration-test-api

test-crawler: 
	make unit-test-crawler
	make integration-test-crawler

# Unit Tests
unit-test-api:
	$(DE-API) "${GOTEST} -v ./... -tags=unit"

unit-test-crawler:
	$(DE-CRAWLER) "${GOTEST} -v ./... -tags=unit"

# Integration Tests
integration-test-api:
	$(DE-API) "${GOTEST} -v ./... -tags=integration"

integration-test-crawler:
	$(DE-CRAWLER) "${GOTEST} -v ./... -tags=integration"

# Clean
clean:
	$(DC) stop
	$(DC) down
