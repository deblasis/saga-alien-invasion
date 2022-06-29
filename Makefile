GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test
GODOC=godoc
DOCKER=docker

.PHONY: run
run:
	$(GO) run ./... $(ARGS)

.PHONY: docker/build
docker/build:
	$(DOCKER) build -f ./docker/Dockerfile -t deblasis-saga-alien-invasion .

.PHONY: run/docker
run/docker: docker/build
	$(DOCKER) run --rm deblasis-saga-alien-invasion saga-alien-invasion $(ARGS)

.PHONY: build
build:
	$(GO) build -o ./saga-alien-invasion

.PHONY: test
test:
	$(GOTEST) -v ./... -cover

.PHONY: test/nocacke
test/nocache:
	$(GOTEST) -v -count=1 ./...

.PHONY: test/cover
test/cover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out -o coverage.html
	$(DOCKER) build -f ./docker/Dockerfile.nginx -t deblasis-saga-alien-invasion-cover .
	@echo "\n\nPlease navigate to http://localhost:8080/ to view the code coverage report\n\n"
	$(DOCKER) run --rm -p 8080:80 deblasis-saga-alien-invasion-cover

.PHONY: godoc
godoc:
	@echo "\n\nPlease navigate to http://localhost:6060/pkg/github.com/deblasis/saga-alien-invasion/ to access the documentation\n\n"
	$(DOCKER) run --rm -p 6060:6060 deblasis-saga-alien-invasion godoc -http=:6060
	$(GODOC)  -http=:6060