clear:
	@rm -rf main

build-main:
	@GOOS=linux GOARCH=amd64 go build hello.go

run-zip:
	zip -r function.zip hello

create-function:
	aws lambda update-function-code --function-name sns_notification --zip-file fileb://function.zip

invoke-test:
	aws lambda invoke --function-name sns_notification --cli-binary-format raw-in-base64-out --payload file://request.json response.json

run:
	@make clear
	@make build-main
	@make run-zip
	@make create-function
	@inoke-test