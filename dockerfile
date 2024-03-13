# Dockerfile to build and run the fetch program

FROM golang:latest

WORKDIR /app

COPY . .

RUN go get github.com/snoopy910/fetch-save-url
RUN go build fetch.go

CMD ["./fetch"]
