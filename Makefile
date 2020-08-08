GOFLAGS ?= $(GOFLAGS:)

all: install test

build:
	@go build $(GOFLAGS) .

format:
	@gofmt -l -s -w .

lint:
	@golint ./...

optimize_imports:
	@goimports -l -w .

beautify: format optimize_imports

vet:
	@go vet ./...

race:
	@go test -race $(GOFLAGS) ./...

audit: vet race lint

test: install
	@go test $(GOFLAGS) ./...

race_loop:
	@for i in {1..100}; do make beautify audit; sleep 1;done

bench: install
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...

fix:
	@go fix $(GOFLAGS) ./...

github_workflow : beautify vet lint

commit: beautify audit
	@git add -p .
