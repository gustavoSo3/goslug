FROM golang:1.27-alpine

WORKDIR /goslug

COPY . .

RUN go mod tidy

RUN mkdir ./bin

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main ./cmd/goslug/main.go


EXPOSE 8080

CMD ["./bin/main"]

