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

run-detach:
	$(DC) $(CDEV) up -d

# Tests
check:
	make test-api

# Run all Tests API
test-api: 
	make test-api-unit
	make test-api-integration

# Unit Tests
test-api-unit:
	$(DE-API) "${GOTEST} ./... -tags=unit"

# Integration Tests
test-api-integration:
	$(DE-API) "${GOTEST} ./... -tags=integration"

# Clean
clean:
	$(DC) $(CDEV) down