FROM golang:1.14.1 as build-img
LABEL maintainer "abishekkumar92@gmail.com"

ENV GO111MODULE=off

RUN apk update && apk add git && mkdir /home/com
COPY ${PWD} /go/src/github.com/kumarabd/go-http/
WORKDIR /go/src/github.com/kumarabd/go-http/
RUN git rev-parse HEAD > /home/com/version
# RUN git describe --tags `git rev-list --tags --max-count=1` >> /home/com/version

RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /home/com/service main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=build-img /home/com/** /home/
WORKDIR /home/
EXPOSE 80 5001
CMD ["sh","-c","./service"]