FROM golang:latest

SHELL ["/bin/bash", "-c"]

RUN mkdir /server

COPY . /server

WORKDIR /server

RUN go mod download

