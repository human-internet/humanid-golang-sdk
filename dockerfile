FROM golang:1.14 AS builder

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN go get github.com/pilu/fresh


RUN mkdir -p /go/src/github.com/bluenumberfoundation/humanid-golang-sdk/
WORKDIR /go/src/github.com/bluenumberfoundation/humanid-golang-sdk/

ADD . /go/src/github.com/bluenumberfoundation/humanid-golang-sdk/

RUN dep ensure
