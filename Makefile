GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test
GODOC=godoc
DOCKER=docker


.PHONY: run
run:
	$(GO) run ./... $(ARGS)

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

.PHONY: godoc
godoc:
	$(DOCKER) build -f ./docker/Dockerfile.godoc -t deblasis-godoc .
	@echo "\n\nPlease navigate to http://localhost:6060/pkg/github.com/deblasis/saga-alien-invasion/ to access the documentation\n\n"
	$(DOCKER) run --rm -p 6060:6060 deblasis-godoc godoc -http=:6060
	$(GODOC)  -http=:6060