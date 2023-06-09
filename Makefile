format:
	@go fmt ./...

test:
	@go test ./... -coverprofile cover.out

cover:
	@go tool cover -html cover.out -o cover.html

build:
	@go build -v ./...

install:
	@go install -v ./...

vulncheck:
	@govulncheck ./...

pre-commit:
	@pip3 install pre-commit
	@pre-commit install
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.1
