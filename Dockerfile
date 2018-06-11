FROM golang:latest

WORKDIR /go/src/github.com/makki0205/ogp-cache

ADD ./ ./

RUN ls -la

EXPOSE 8080

ENTRYPOINT ["go","run","main.go"]
