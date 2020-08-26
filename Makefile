.DEFAULT_GOAL := prepare

.PHONY: prepare
prepare:
	go get -v -t -d ./...

static-analysis:
	go vet ./...

.PHONY: test
test:
	go test -v ./...