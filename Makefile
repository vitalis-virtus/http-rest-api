APP_NAME=app

.PHONY: build
build: 	## build main.go
	go build -v main.go

.PHONY: run
run: ## run main.go
	go run main.go

.PHONY: test
test:	## run test for app
	go test -v -race -timeout 30s ./...

.PHONY: lint
lint:	## run golangci lint
	golangci-lint run ./... --config=./.golangci.yml

.PHONY: lint-fast
lint-fast:	## run golangci lint fast
	golangci-lint run ./... --fast --config=./.golangci.yml

.PHONY: docker-build
docker-build:  ## run container in development mode
	docker-compose build --no-cache ${APP_NAME} && docker-compose run $(APP_NAME)

.PHONY: docker-compose
docker-compose: ## spin up the project
	docker-compose up

.PHONY: docker-stop
docker-stop: ## stop running containers
	docker stop

.PHONY: docker-rm
docker-rm:  ## stop and remove running containers
	docker rm $(APP_NAME)

.DEFAULT_GOAL := build