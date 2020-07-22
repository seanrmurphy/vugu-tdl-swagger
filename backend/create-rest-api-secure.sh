#! /usr/bin/env bash

# assumes AWS CLI v2
#
# this really is stretching bash...I guess this is what the serverless framework
# really helps with...

SWAGGER_FILE=swagger-secure.yaml

create_rest_api() {
  RESTAPIID=$(aws apigateway import-rest-api --cli-binary-format raw-in-base64-out --body file://$PWD/$SWAGGER_FILE --profile $PROFILE | jq -r '.id')
  printf "REST API ID = $RESTAPIID\n"
}

get_resource_id() {
	RESOURCE=$1
  RESOURCEID=$(aws apigateway get-resources --rest-api-id $RESTAPIID --profile "$PROFILE" | jq -r ".items[] | select( .path == \"$RESOURCE\") | .id")
  printf "TODO RESOURCE ID = $RESOURCEID \n"
}

get_function_arn() {
  FUNCTIONARN=$(aws lambda get-function --function-name $1 --profile $PROFILE | jq -r '.Configuration.FunctionArn')
  URI="arn:aws:apigateway:$REGION:lambda:path/2015-03-31/functions/$FUNCTIONARN/invocations"
  printf "URI = $URI \n"
}

put_integration() {
  aws apigateway put-integration --rest-api-id $RESTAPIID --resource-id $1 --http-method $2 --type AWS_PROXY --integration-http-method POST --uri $3 --profile $PROFILE --credentials $ROLE
}

check_env_vars() {
	if [ -z "$GOFULLSTACKPROFILE" ]
	then
		echo "Environment variable GOFULLSTACKPROFILE not defined...exiting..."
		exit
	fi

	if [ -z "$GOFULLSTACKROLE" ]
	then
		echo "Environment variable GOFULLSTACKROLE not defined...exiting..."
		exit
	fi
}

check_env_vars

PROFILE="$GOFULLSTACKPROFILE"
ROLE="$GOFULLSTACKROLE"

REGION=eu-west-2

create_rest_api
get_resource_id "/todo"

get_function_arn list-todos
put_integration $RESOURCEID "GET" $URI

get_function_arn create-todo
put_integration $RESOURCEID "POST" $URI

get_resource_id "/todo/{todoid}"

get_function_arn get-todo
put_integration $RESOURCEID "GET" $URI

get_function_arn delete-todo
put_integration $RESOURCEID "DELETE" $URI

get_function_arn update-todo
put_integration $RESOURCEID "PUT" $URI

aws apigateway create-deployment --rest-api-id $RESTAPIID --stage-name prod --stage-description prod --profile $PROFILE

echo
echo "To run the basic tests you need to set and environment variable as follows:"
echo
echo "For bash"
echo "export RESTAPI=\"https://$RESTAPIID.execute-api.eu-west-1.amazonaws.com/prod\""
echo
echo "For fish"
echo "set -x RESTAPI \"https://$RESTAPIID.execute-api.eu-west-1.amazonaws.com/prod\""
echo
echo "*** Please don't forget to limit access to your API when done; you can delete it completely using the following"
echo
echo "    ./delete-rest-api.sh"
echo
echo "Note that does not remote lambda and dynamodb..."


