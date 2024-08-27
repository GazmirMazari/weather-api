FROM golang:1.19.2-alpine as builder

COPY . /app
WORKDIR /app

RUN go get ./...
RUN cat /etc/resolv.conf

WORKDIR /app/cmd/svr

RUN go build -o app

FROM alpine:3.18

EXPOSE 8080
COPY --from=builder /app/cmd/svr/app .
COPY --from=builder /app/cmd/svr/config.yaml .


CMD ["./app"]