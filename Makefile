build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/handler handler/main.go

deploy: build
	sls deploy --verbose --stage $(STAGE)