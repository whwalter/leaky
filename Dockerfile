FROM golang:1.16-alpine
COPY . /go/src/leak
WORKDIR /go/src/leak
RUN go build -o /go/bin/leaky ./server/main.go
RUN go build -o /go/bin/leak-client ./client/main.go
EXPOSE 8080/tcp
ENTRYPOINT ["leaky"]
