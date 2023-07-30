build:
	sam build --config-file samconfig.toml --parameter-overrides "ParameterKey=Stage,ParameterValue=$(STAGE)"

deploy: 
	sam deploy --config-file samconfig.toml --parameter-overrides "ParameterKey=Stage,ParameterValue=$(STAGE)"