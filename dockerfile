FROM scratch
COPY /src/hi /app
WORKDIR /app
ENTRYPOINT [ "/app/hi" ]
