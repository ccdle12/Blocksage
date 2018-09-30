.PHONY: env-up build adminer check clean

default: clean build env-up

# env-up will launch the project	
env-up:
	@echo "[Bringing up environment...]"
	docker-compose up -d

# build all the images in the docker-compose file
build:
	@echo "[Building containers...]"
	docker-compose -f docker-compose.yml build

# adminer will open adminer on the localhost
adminer:
	open http://localhost:8080

# check will run all tests for each service in the project
check: test-api test-crawler

# service specific tests
test-api: unit-test-api integration-test-crawler

test-crawler: unit-test-crawler integration-test-crawler

# api tests
utest-api:
	@docker exec -it api /bin/bash -c "go test -v ./... -tags=unit"

itest-api:
	@docker exec -it api /bin/bash -c "go test -v ./... -tags=integration"

# crawler tests
utest-crawler:
	@docker exec -it crawler /bin/bash -c "go test -v ./... -tags=unit"

itest-crawler:
	@echo "[Running crawler integration tests...]"
	@docker exec -it crawler /bin/bash -c "go test -v ./... -tags=integration"

# clean will tear down the docker network
clean:
	@echo "[Cleaning up project...]"
	@docker-compose down --volumes --remove-orphans
	@cd go-crawler && rm -rf tmp | true