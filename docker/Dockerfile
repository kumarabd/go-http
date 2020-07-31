FROM registry.bookmyshow.org/alpine/golang:1.14.1 as build-img
LABEL maintainer "abishek.kumar@bookmyshow.com"

ENV GO111MODULE=off

RUN apk update && apk add git && mkdir /home/com
COPY ${PWD} /go/src/github.com/kumarabd/go-backend/
WORKDIR /go/src/github.com/kumarabd/go-backend/
RUN git rev-parse HEAD > /home/com/version
# RUN git describe --tags `git rev-list --tags --max-count=1` >> /home/com/version

RUN CGO_ENABLED=0 go build -a -ldflags "-s -w" -o /home/com/email main.go

FROM registry.bookmyshow.org/alpine:latest

RUN apk --no-cache add ca-certificates
COPY --from=build-img /home/com/** /home/
WORKDIR /home/
EXPOSE 80 5001
CMD ["sh","-c","./email"]