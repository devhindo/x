FROM golang:latest

RUN github.com/michimani/gotwi
RUN go get github.com/joho/godotenv