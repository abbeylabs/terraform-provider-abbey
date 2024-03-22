.PHONY: test unit-test acceptance-test

test: unit-test acceptance-test

unit-test:
	@echo "Running unit tests..."
	@go test -tags=unit ./... -v $(TESTARGS) -timeout 120m

acceptance-test:
	@echo "Running acceptance tests..."
	@go test -tags=acceptance ./... -v $(TESTARGS) -timeout 120m
