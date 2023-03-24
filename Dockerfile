# syntax=docker/dockerfile:1
## Build
FROM golang:alpine as build
ARG DATE
WORKDIR /build
COPY . .
COPY ./storage-service/internal/db/migrations ./app/storage-service/migrations
RUN go mod download
RUN go build -ldflags="-X 'main.BuildDate=${DATE}'" -o ./app/storage-service/ ./storage-service/cmd/main.go
RUN go build -ldflags="-X 'main.BuildDate=${DATE}'" -o ./app/http-service/ ./http-service/cmd/main.go

## Deploy storage-service
FROM golang:alpine as storage-service
WORKDIR /usr/bin
COPY --from=build ./build/app/storage-service/ ./app/
EXPOSE 9000
ENV GRPC_SERVER_HOST=storage-service
ENV GRPC_SERVER_PORT=9000
ENV DB_FILE=./app/sqlite.db
ENV DB_MIGRATIONS_DIR=./app/migrations
RUN ls -R
ENTRYPOINT ["./app/main"]

## Deploy http-service
FROM golang:alpine as http-service
WORKDIR /usr/bin
COPY --from=build ./build/app/http-service/ ./app/
EXPOSE 9001
ENV HTTP_SERVER_HOST=http-service
ENV HTTP_SERVER_PORT=9001
ENV STORAGE_SERVICE_HOST=storage-service:9000
ENTRYPOINT ["./app/main"]


