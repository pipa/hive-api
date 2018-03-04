FROM golang:1.10.0-alpine

WORKDIR /app
ADD . /app/

# RUN go build -o main .
CMD ["go","run","/app/Main.go"]
