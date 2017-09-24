default: vet test

test:
	@go test -race -cover ./...

vet:
	@go vet ./...

.PHONY: test vet
