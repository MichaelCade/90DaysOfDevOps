BINARY_NAME=90DaysOfDevOps

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}_0.2_darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}_0.2_linux main.go
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}_0.2_windows main.go
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}_0.2_linux_arm64 main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}_0.2_darwin_arm64 main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows