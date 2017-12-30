FROM golang:1.9.2-stretch

ENV GOPATH $GOPATH:/go
ENV GOBIN /go/bin

RUN apt-get update && \
    apt-get upgrade -y

ADD . /go/src/Repo-watcher
WORKDIR /go/src/Repo-watcher/src

CMD go run main.go
