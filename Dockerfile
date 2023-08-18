FROM golang:latest

SHELL ["/bin/bash", "-c"]

RUN go get github.com/michimani/gotwi
RUN go get github.com/joho/godotenv