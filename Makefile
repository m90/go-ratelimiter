default: test

test:
	@go test -race -cover ./...

.PHONY: test
