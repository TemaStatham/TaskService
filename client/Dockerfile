FROM golang:alpine AS build

WORKDIR /grpc/

COPY go.mod go.sum ./
RUN go mod download

COPY functions/ functions/
COPY client/ client/

RUN go build -v -o main ./client/cmd

FROM alpine:latest

WORKDIR /grpc/
EXPOSE 8080

COPY --from=build /grpc/main /grpc/main
COPY ./client/config.yaml ./config.yaml

RUN chmod +x /grpc/main

ENV PORT=50501
CMD ["/grpc/main", "--config=./config.yaml"]
