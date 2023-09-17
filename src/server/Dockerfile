FROM golang:1.20

RUN mkdir /server

COPY . /server

WORKDIR /server

RUN go build -o main .

EXPOSE 8080

CMD ["/server/main"]

# cmd : docker build -t devhindo/xcli:v1.0 . && docker container run -it --name xcli-server -p 3000:3000 devhindo/xcli:v1.0