FROM ubuntu:13.10

RUN apt-get update
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install golang git mercurial build-essential
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN go get github.com/codegangsta/gin
RUN go get bitbucket.org/liamstask/goose/cmd/goose

WORKDIR /go/src/github.com/hackedu/backend
ADD . /go/src/github.com/hackedu/backend
RUN go get
