#! /usr/bin/env bash

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi
}

check_env_vars

# Get the ARN of the user pool
POOLARN=$(aws cognito-idp list-user-pools --max-results 10 --profile $GOFULLSTACKPROFILE | jq -r '.UserPools[] | select(.Name=="todo_api") | .Id')

# Delete the user pool
echo 'Deleting user pool domain'
aws cognito-idp delete-user-pool-domain \
--domain todo-api-client \
--user-pool-id $POOLARN \
--profile $GOFULLSTACKPROFILE

# Delete the user pool
echo 'Deleting user pool todo_api (and associated state)'
aws cognito-idp delete-user-pool \
--user-pool-id $POOLARN \
--profile $GOFULLSTACKPROFILE



