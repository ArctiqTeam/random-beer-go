FROM golang:1.12.0-alpine3.9
RUN go version

COPY . /go/src/github.com/arctiqteam/random-beer-go
WORKDIR /go/src/github.com/arctiqteam/random-beer-go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .
EXPOSE 8080
CMD ["/go/src/github.com/arctiqteam/random-beer-go/app"]
