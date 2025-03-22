FROM golang:alpine AS build

WORKDIR /grpc/

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY proto/notification/ proto/notification/
COPY notificationservice/ notificationservice/

RUN go build -v -o main ./notificationservice/cmd

FROM alpine:latest

WORKDIR /grpc/
EXPOSE 8080

COPY --from=build /grpc/main /grpc/main

RUN chmod +x /grpc/main

ENV PORT=50501
CMD ["/grpc/main", "--config=./config.yaml"]
