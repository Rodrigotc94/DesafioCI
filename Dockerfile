FROM golang:alpine
RUN apk add bash

ADD . /go/src/DesafioCI

RUN go install DesafioCI

ENTRYPOINT /go/bin/DesafioCI 