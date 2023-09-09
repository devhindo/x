FROM golang:latest

SHELL ["/bin/bash", "-c"]

RUN mkdir /x

COPY ./src/cli /x

WORKDIR /x

RUN go get github.com/michimani/gotwi
RUN go get github.com/joho/godotenv
RUN go get golang.ngrok.com/ngrok
