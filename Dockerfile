FROM golang:alpine as builder

ENV GO111MODULE=off

WORKDIR /app
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

ADD main.go /app/main.go

RUN GOOS=linux GOARCH=amd64 go build -tags netgo -o go-word-counter .

FROM scratch
COPY --from=builder /app/go-word-counter .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["./go-word-counter"]
