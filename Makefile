LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	go build -o bexs cmd/main.go

linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin

lint:
	golangci-lint run ./...

test:
	go test ./... -covermode=count -count 1 -v  
