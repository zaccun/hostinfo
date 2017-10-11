FROM golang:1.9.1-alpine3.6
WORKDIR /go/src/app
COPY . .
RUN go build hi.go
ENTRYPOINT [ "./hi" ]

