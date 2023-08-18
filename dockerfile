FROM golang:latest

RUN go get github.com/michimani/gotwi
RUN go get github.com/joho/godotenv