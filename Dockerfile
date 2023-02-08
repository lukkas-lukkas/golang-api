FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLE=0

CMD ["tail", "-f", "/dev/null"]
