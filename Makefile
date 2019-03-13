.PHONY: all
all: build test

setup:
	mkdir -p $(GOPATH)/bin
	go get -u golang.org/x/lint/golint

build: compile fmt vet lint

compile:
	mkdir -p out/
	go build -race ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint -set_exit_status $$p; \
	done

test:
	ENVIRONMENT=test go test -race ./...
