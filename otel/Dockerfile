# Build the host-info binary with no dependencies
FROM golang:1.24-alpine AS hi-build
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo hi.go
RUN chown 1000:1000 /go/src/app/hi ; chmod 500 /go/src/app/hi

# Package the binary in its own isolated container
FROM scratch
COPY --from=hi-build /go/src/app/hi /hi
LABEL org.opencontainers.image.ref.name=hostinfo
LABEL org.opencontainers.image.version=1.0.0
LABEL org.opencontainers.image.authors=rx-m
LABEL org.opencontainers.image.url=https://rx-m.com
USER 1000:1000
EXPOSE 9898
ENTRYPOINT [ "/hi" ]
