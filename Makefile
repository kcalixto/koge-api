build:
	sam build --config-env $(STAGE)

deploy: 
	sam deploy --config-env $(STAGE)