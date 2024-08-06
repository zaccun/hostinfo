# Build the host-info binary with no dependencies
FROM golang:1.22-alpine AS hi-build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo hi.go
RUN chown 1000:1000 /go/src/app/hi ; chmod 500 /go/src/app/hi

# Package the binary in its own isolated container
FROM scratch
COPY --from=hi-build /go/src/app/hi /hi
USER 1000:1000
EXPOSE 9898
ENTRYPOINT [ "/hi" ]