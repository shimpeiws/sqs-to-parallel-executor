clean:
	rm -rf sqs-to-parallel-executor
	rm -rf examples/examples

build:
	GOOS=linux GOARCH=amd64 go build -o sqs-to-parallel-executor main.go
	GOOS=linux GOARCH=amd64 go build -o examples/examples ./examples

build-for-mac:
	GOOS=darwin GOARCH=amd64 go build -o sqs-to-parallel-executor main.go
	GOOS=darwin GOARCH=amd64 go build -o examples/examples ./examples
