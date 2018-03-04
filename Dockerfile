FROM golang:1.10.0

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main .
