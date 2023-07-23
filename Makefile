build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/handler handler/main.go

deploy: build
	sls deploy --verbose --stage prd