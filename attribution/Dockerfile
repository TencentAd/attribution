FROM golang:latest

ENV TZ "PRC"

WORKDIR /app
ADD . .

RUN go install -v attribution_server.go

ENTRYPOINT ["attribution_server"]
