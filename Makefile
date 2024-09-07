BINARY_NAME=flowtool

build:
	go build -o ./bin/${BINARY_NAME} ./main.go

all:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin ./main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux ./main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows ./main.go

run:
	go run ./src/main/main.go

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

test:
	go test ./... -v

dep:
	go mod download
