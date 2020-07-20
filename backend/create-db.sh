#! /usr/bin/env bash

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi
}

check_env_vars

# note that with dynamodb you only specify essential attributes at db creation time...
aws dynamodb create-table --table-name Todos \
--attribute-definitions AttributeName=ID,AttributeType=B  \
--key-schema AttributeName=ID,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
--profile $GOFULLSTACKPROFILE
