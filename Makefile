SERVER_IMAGE := docker.io/ac5tin/miniq:0.2.1-beta

test: 
	@echo "- Spinning Up Server for Testing -"
	docker run --rm -d -p 8080:8080 ${SERVER_IMAGE}
	@echo "- Running test cases -"
	go test ./... -v -cover
	@echo "- Running cleanup -"
	docker stop $$(docker ps -a -q --filter ancestor=${SERVER_IMAGE} --format="{{.ID}}")

build:
	go build -o bin/app

run:
	go run main.go

clean:
	rm -rf ./bin

protoc:
	protoc --go_out=. --go-grpc_out=. ./proto/queue.proto
