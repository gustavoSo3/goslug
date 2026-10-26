rm bin/main
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/main ./api/cmd/goslug/main.go
