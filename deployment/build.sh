rm ~/goslug/bin/main
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main ./cmd/goslug/main.go
