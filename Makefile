 GOCMD=go
 GOTEST=$(GOCMD) test

test-api-all: 
	docker exec -i -t api /bin/bash -c "${GOTEST} -v ./... "

# Unit Tests
test-api-unittest:
	make test-api-injector

test-api-injector: 
	docker exec -i -t api /bin/bash -c "cd injector && ${GOTEST} -v"

# Integration Tests
test-api-integration:
	docker exec -i -t api /bin/bash -c "cd bitcoinclient && ${GOTEST} -v"
