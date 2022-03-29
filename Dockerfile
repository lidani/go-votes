FROM golang:1.16-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o /main ./cmd/api/main.go

FROM alpine:3.15.2

WORKDIR /app

COPY --from=build /main .

EXPOSE 9090

ENTRYPOINT ["./main"]