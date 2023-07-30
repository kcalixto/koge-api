build-lambda-function:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(HANDLER) $(DIR)/$(HANDLER).go
	mv $(HANDLER) $(ARTIFACTS_DIR)/

build-KogeApiFunction: HANDLER=main
build-KogeApiFunction: DIR=handler
build-KogeApiFunction: build-lambda-function