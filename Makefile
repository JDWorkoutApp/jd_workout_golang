test:
	go test -cover ./... --coverprofile=coverage.out -race -covermode=atomic -cover=true
lint:
	golangci-lint run
swagger:
	swag init -g cmd/main.go