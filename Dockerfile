FROM golang:1.18
WORKDIR /go/src/simple-orm
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]