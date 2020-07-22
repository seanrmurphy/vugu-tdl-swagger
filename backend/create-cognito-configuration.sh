#! /usr/bin/env bash

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi
}

check_env_vars

# Create the user pool
echo 'Creating user pool todo_api'
aws cognito-idp create-user-pool \
--pool-name todo_api \
--profile $GOFULLSTACKPROFILE \
--no-paginate

# Get the ARN of the user pool
POOLARN=$(aws cognito-idp list-user-pools --max-results 10 --profile $GOFULLSTACKPROFILE | jq -r '.UserPools[] | select(.Name=="todo_api") | .Id')

# Create a client
echo 'Creating user pool client todo_api_client'
aws cognito-idp create-user-pool-client \
--user-pool-id $POOLARN \
--client-name todo_api_client \
--no-generate-secret \
--supported-identity-providers "COGNITO" \
--callback-urls "http://localhost:8844" \
--allowed-o-auth-flows "code" \
--allowed-o-auth-scopes "email" "openid" "profile" "aws.cognito.signin.user.admin" \
--allowed-o-auth-flows-user-pool-client \
--explicit-auth-flows "ALLOW_USER_SRP_AUTH" "ALLOW_REFRESH_TOKEN_AUTH" "ALLOW_CUSTOM_AUTH" \
--profile $GOFULLSTACKPROFILE

# Create a domain
echo 'Creating user pool domain'
aws cognito-idp create-user-pool-domain \
--user-pool-id $POOLARN \
--domain todo-api-client \
--profile $GOFULLSTACKPROFILE

# Create a User
echo 'Creating test user in the todo_api user pool'
aws cognito-idp admin-create-user \
--user-pool-id $POOLARN \
--username testuser \
--user-attributes Name=email,Value=testuser@test.org Name=email_verified,Value=true \
--message-action "SUPPRESS" \
--temporary-password "Passw0rd!" \
--profile $GOFULLSTACKPROFILE

echo
echo 'Cognito configuration created'
echo 'Note that it can take 15 minutes before the domain is active and usable'
echo
echo 'ARN for Cognito user pool is'
echo "    $POOLARN"
echo

