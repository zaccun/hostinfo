FROM docker.io/python:3-alpine
WORKDIR /usr/src/hostinfo
COPY ./__main__.py .
RUN chown 1000:1000 __main__.py ; chmod 500 __main__.py
USER 1000
ENV PYTHONUNBUFFERED=1
ENTRYPOINT ["./__main__.py"]
CMD ["9898"]