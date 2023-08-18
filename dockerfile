FROM golang:latest

RUN go get github.com/michimani/gotwi
RUN RUN go get github.com/joho/godotenv
