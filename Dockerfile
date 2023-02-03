FROM golang:1.19

WORKDIR /application

COPY go.mod ./
COPY main.go ./

RUN go build -o /go-api-rest

EXPOSE 8080

CMD ["/go-api-rest"]