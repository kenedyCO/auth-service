FROM golang:1.20.3-alpine AS builder

COPY . /github.com/kenedyCO/auth-service/source/
WORKDIR /github.com/kenedyCO/auth-service/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kenedyCO/auth-service/source/bin/crud_server .

CMD ["./crud_server"]