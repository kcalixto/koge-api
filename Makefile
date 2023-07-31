include lambda.mk

build: build-KogeApiFunction
	sam build --config-env $(STAGE)

deploy: 
	sam deploy --config-env $(STAGE)