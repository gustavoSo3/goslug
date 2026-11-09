FROM golang:1.27-alpine

WORKDIR /app

# Copy dependency files first for caching
COPY go.mod ./
RUN go mod download

# Now copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main ./cmd/goslug/main.go

EXPOSE 8080

CMD ["./bin/main"]