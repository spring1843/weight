SHELL = /bin/sh
GOFLAGS ?= $(GOFLAGS:)

build:
	@go build $(GOFLAGS) .

install:
	@go install $(GOFLAGS) .

format:
	@gofmt -l -s -w .

lint:
	@golint ./...

optimize_imports:
	@goimports -l -w .

vet:
	@go vet ./...

race:
	@go test -race $(GOFLAGS) ./...

test: install
	@go test $(GOFLAGS) ./...

coverage: install
	@go test -coverprofile=profile.cov -covermode=count  $(GOFLAGS) ./...

race_loop:
	@for i in {1..100}; do make beautify audit; sleep 1;done

bench: install
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

docker_build:
	@docker build -t spring1843/weight .

fix:
	@go fix $(GOFLAGS) ./...

commit: beautify audit
	@git add -p .

audit: vet race lint

github_workflow : build beautify vet race lint fix coverage

all: install test

beautify: format optimize_imports
