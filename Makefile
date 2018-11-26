default: test

vet:
			@go vet ./...

test: vet
			go test ./...

build: 
			go build ./...

.PHONY: test setup build
