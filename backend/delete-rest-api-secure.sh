#! /usr/bin/env bash

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi

}

check_env_vars

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

REGION=eu-west-1

# Does not work well if there ar  two APIs with the sane name..."
RESTAPIID=$(aws apigateway get-rest-apis --profile $PROFILE | jq -r '.items[] | select( .name == "Simple Todo API (Secure)") | .id')

echo "Running delete command on API $RESTAPIID ..."
aws apigateway delete-rest-api --rest-api-id $RESTAPIID --profile $PROFILE

