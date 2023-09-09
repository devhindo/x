FROM golang:latest

SHELL ["/bin/bash", "-c"]

RUN mkdir /server

COPY . /server

WORKDIR /server

RUN go mod download

# expose port 3000

EXPOSE 3000

CMD ["go", "run", "main.go"]

# cmd : docker build -t devhindo/xcli:v1.0 . && docker container run -it --name xcli-server -p 3000:3000 devhindo/xcli:v1.0