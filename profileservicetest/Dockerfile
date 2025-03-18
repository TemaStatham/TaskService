FROM golang:alpine AS build

WORKDIR /grpc/

COPY go.mod go.sum ./
RUN go mod download

COPY proto/profile/ proto/profile/
COPY server/ server/

RUN go build -v -o main ./server

FROM alpine:latest

WORKDIR /grpc/

COPY --from=build /grpc/main .

ENV PORT=50501
ENTRYPOINT ["/grpc/main"]
