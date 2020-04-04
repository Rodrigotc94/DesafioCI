FROM golang:alpine

ADD . /go/src/DesafioCI

RUN go install DesafioCI

ENTRYPOINT /go/bin/DesafioCI 