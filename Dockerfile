FROM golang:alpine

RUN apk update && \
    apk add bash

WORKDIR /workspace

COPY . /workspace

RUN go install -v && mkdir /home/test

ENV HOME=/home/test

ENTRYPOINT [ "/go/bin/game-manager-go" ]
