FROM golang:1.27-alpine

WORKDIR /goslug

COPY . .

RUN go mod tidy

RUN go build -o ./bin ./api/cmd/goslug/main.go

EXPOSE 8080

CMD ["./bin/main"]

