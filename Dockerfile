FROM golang:latest

RUN mkdir /server

ADD . /server/

WORKDIR /server

RUN go build -o main .

CMD ["/server/main"]
