FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLE=0

RUN apt update && apt install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]
