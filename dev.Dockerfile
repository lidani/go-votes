FROM golang:1.16-alpine

WORKDIR /app

EXPOSE 9090

ENTRYPOINT [ "go", "run", "./cmd/api/main.go" ]