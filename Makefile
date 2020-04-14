.PHONY: run down bash test clean

help:		## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run:		## Start the node container
	@docker-compose up -d
down:		## Stop the node container
	@docker-compose down

bash:		## Open a new interactive bash in the go container
	@docker-compose exec go /bin/bash

test:		## Execute test

clean:		## Clean all the data created
