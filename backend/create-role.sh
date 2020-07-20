#! /usr/bin/env bash

PROFILE=default
ROLENAME=go-fullstack-test-role

# note that with dynamodb you only specify essential attributes at db creation time...
aws iam create-role --role-name $ROLENAME --assume-role-policy-document file://role.json --profile $PROFILE

aws iam attach-role-policy --role-name $ROLENAME --policy-arn "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess" --profile $PROFILE
aws iam attach-role-policy --role-name $ROLENAME --policy-arn "arn:aws:iam::aws:policy/service-role/AWSLambdaDynamoDBExecutionRole" --profile $PROFILE
aws iam attach-role-policy --role-name $ROLENAME --policy-arn "arn:aws:iam::aws:policy/AWSLambdaExecute" --profile $PROFILE
aws iam attach-role-policy --role-name $ROLENAME --policy-arn "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole" --profile $PROFILE
aws iam attach-role-policy --role-name $ROLENAME --policy-arn "arn:aws:iam::aws:policy/AWSLambdaInvocation-DynamoDB" --profile $PROFILE

ROLEARN=$(aws iam get-role --role-name $ROLENAME --profile $PROFILE | jq -r '.Role.Arn')
echo
echo "To set up the lambda functions and the API gateway, please set the following environment variable"
echo
echo "For bash"
echo "export GOFULLSTACKROLE=\"$ROLEARN\""
echo
echo "For fish"
echo "set -x GOFULLSTACKROLE \"$ROLEARN\""

