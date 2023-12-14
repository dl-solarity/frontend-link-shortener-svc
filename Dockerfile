FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dl-solarity/frontend-link-shortener-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/frontend-link-shortener-svc /go/src/github.com/dl-solarity/frontend-link-shortener-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/frontend-link-shortener-svc /usr/local/bin/frontend-link-shortener-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["frontend-link-shortener-svc"]
