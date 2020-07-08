FROM golang:1.14-alpine

COPY ./ /var/app

WORKDIR /var/app

ENV GO111MODULE="on"

RUN apk add --no-cache \
        alpine-sdk \
        git && \
        go get github.com/pilu/fresh
    
CMD [ "fresh" ]



