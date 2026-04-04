FROM golang:1.26.1 as builder
WORKDIR /workspace


COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o go-server-pc main.go

FROM alpine:latest

RUN apk update
WORKDIR /
COPY --from=builder /workspace/go-server-pc .
ENTRYPOINT ["/go-server-pc"]