update:
	@go mod tidy

install:
	@go build -o ./.log/go-password ./cmd/main.go
	@sudo mv ./.log/go-password /usr/local/bin/go-password
	@go-password -h

