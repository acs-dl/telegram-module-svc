FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/distributed_lab/knox/knox-fork
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/knox /go/src/gitlab.com/distributed_lab/knox/knox-fork/cmd/dev_server

FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/knox /usr/local/bin/knox
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["knox"]
