GOCMD=go
GOTEST=$(GOCMD) test

DC=docker-compose
DE=docker exec -i -t
B=/bin/bash -c
A=api
BC=bitcoinclient
DE-API=$(DE) $(A) $(B)

CDEV=-f docker-compose.yml


# Run Project
default:
	$(DC) $(CDEV) up

build:
	$(DC) $(CDEV) build

detach:
	$(DC) $(CDEV) up -d

adminer:
	open http://localhost:8080

# Tests
check:
	make test-api

# Run all Tests API
test-api: 
	make test-api-unit
	make test-api-integration

# Unit Tests
unit-test-api:
	$(DE-API) "${GOTEST} -v ./... -tags=unit"

# Integration Tests
integration-test-api:
	$(DE-API) "${GOTEST} -v ./... -tags=integration"

# Clean
clean:
	$(DC) $(CDEV) down
