clean:
	rm -rf sqs-to-parallel-executor

build:
	GOOS=linux GOARCH=amd64 go build -o sqs-to-parallel-executor main.go
