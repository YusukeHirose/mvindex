FROM golang:1.14-alpine

EXPOSE 8080

ENV GO111MODULE="on"

WORKDIR /var/app/movie

RUN apk update \
    add --no-cache \ 
        alpine-sdk \
        git && \
        go get github.com/pilu/fresh
    
CMD [ "fresh" ]



