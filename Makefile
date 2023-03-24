go-init: 
	go mod tidy

go-install:
	go get -u google.golang.org/protobuf/proto 2>/dev/null && go install google.golang.org/protobuf/proto
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@latest && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go-gen: go-install
	go generate ./...
	go get -u -t -v ./... || :
	go mod tidy

go-run-storage:
	go run ./storage-service/cmd/main.go

go-run-http:
	go run ./http-service/cmd/main.go

compose-run:
	docker-compose build --build-arg DATE=24.03.23
	docker-compose up

compose-delete:
	docker-compose rm -f
