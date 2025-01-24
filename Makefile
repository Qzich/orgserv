lint:
	golangci-lint run

test:
	go test ./...

vet:
	go vet ./...