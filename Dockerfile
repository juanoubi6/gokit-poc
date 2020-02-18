FROM golang:1.11

WORKDIR /gokit-poc

ADD . /gokit-poc

RUN go mod download

RUN go get github.com/pilu/fresh

RUN go build

EXPOSE 8080

CMD ["./gokit-poc"]