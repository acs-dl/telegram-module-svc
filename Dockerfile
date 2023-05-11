FROM golang:1.19-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/acs-dl/telegram-module-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/telegram-module /go/src/github.com/acs-dl/telegram-module-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/telegram-module /usr/local/bin/telegram-module
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["telegram-module"]
