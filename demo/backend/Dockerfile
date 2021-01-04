FROM golang:1.14 AS builder

RUN mkdir -p /go/src/github.com/bluenumberfoundation/humanid-golang-demo/
ADD . /go/src/github.com/bluenumberfoundation/humanid-golang-demo/
WORKDIR /go/src/github.com/bluenumberfoundation/humanid-golang-demo/

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN go get github.com/pilu/fresh

RUN dep ensure

RUN dep ensure -update github.com/bluenumberfoundation/humanid-golang-sdk

RUN go build -o main .
RUN chmod +x ./main

EXPOSE 4000
