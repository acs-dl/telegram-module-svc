FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/distributed_lab/acs/telegram-module
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/telegram-module /go/src/gitlab.com/distributed_lab/acs/telegram-module


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/telegram-module /usr/local/bin/telegram-module
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["telegram-module"]
