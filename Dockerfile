FROM ubuntu:14.04

RUN apt-get update
RUN apt-get -y install golang git mercurial build-essential
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN go get github.com/codegangsta/gin
RUN go get bitbucket.org/liamstask/goose/cmd/goose

WORKDIR /go/src/github.com/hackedu/hackedu
ADD . /go/src/github.com/hackedu/hackedu
RUN go get
