FROM golang:1.14-alpine

COPY ./ /go

WORKDIR /go

ENV GO111MODULE="on"

RUN apk add --no-cache \
        alpine-sdk \
        git && \
        go get github.com/pilu/fresh
    
CMD [ "fresh" ]



