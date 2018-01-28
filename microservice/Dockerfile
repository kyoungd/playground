FROM golang:latest
MAINTAINER Young Kwon

ENV SOURCES /go/src/github.com/young/Cloud-Native-Go/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT Cloud-Native-Go
