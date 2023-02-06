.PHONY: docs
docs:
	go generate

.PHONY: lint
lint:
	golangci-lint run ./... -v
