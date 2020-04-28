.PHONY: run down bash test clean

help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run:		## Start the node container
	@docker-compose up

.PHONY: build
build:		## Build the services
	@docker-compose build
down:		## Stop the node container
	@docker-compose down

.PHONY: swagger
swagger:	## (Re)generate the swagger documentation
	@docker-compose exec -e GO111MODULE=off go swagger generate spec -o ./swagger.yaml --scan-models

.PHONY: client
client: 	## (Re)generate the client-api package
	@docker-compose exec go swagger generate client -f ./swagger.yaml -A product-api -t ./cmd

.PHONY: protos
protos:		## (Re)generate the code for the .proto files
	@docker-compose exec go protoc -I grpc/protos/ grpc/protos/currency.proto --go_out=grpc/protos/currency

bash:		## Open a new interactive bash in the go container
	@docker-compose exec go /bin/bash

test:		## Execute test

clean:		## Clean all the data created
