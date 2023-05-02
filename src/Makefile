build:
	go mod download
	env CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main handler/main.go

deploy:
	serverless deploy
