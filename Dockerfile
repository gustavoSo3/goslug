FROM golang:1.27

WORKDIR /url-shortener

COPY . .

RUN go mod tidy

RUN go build -o ./bin ./api/cmd/url-shortener/main.go

EXPOSE 8080

CMD ["./bin/main"]

