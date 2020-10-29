FROM golang:latest

ENV TZ "PRC"

WORKDIR /app
ADD . .

RUN bash build/build.sh

ENTRYPOINT ["attribution_server"]
