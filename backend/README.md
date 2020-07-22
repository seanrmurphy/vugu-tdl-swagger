# Sample backend for Todo application (Secure)

This is a backend that provides as a set of AWS Lambda functions written in Go.
The basic idea is that each lambda function maps to REST API endpoint and hence
it is necessary to configure this via the API Gateway.

# Prerequisites

It is assumed that the AWS client is installed, authenticated and that the
associated user has many privileges. A comprehensive list of required
privileges is not included here, but it is necessary to be able to create and
trigger lambda functions, create and access dynamodb tables, create API Gateway
state and create and interact with AWS Cognito.

The scripts make use of `jq` - please ensure that this is installed.

# Setting up the application

Setting up the application requires the following:
- create a new role with appropriate trusts and policies
- create a dynamodb database
- upload lambda functions to AWS Lambda
- create AWS Cognito configuration (a user pool, client and a test user)
- create and configure a new REST API Gateway using the provided Swagger definition

Each of the following are dealt with in the subsections below.

It is worth noting that the automation scripts are not very robust right now and
perform minimal state detemination and error checking.

## Create a new role

It is necessary to have a role which can trigger lambda functions and access
DynamoDB; also, it is necessary that this role be trusted by AWS Lambda services
and AWS API Gateway.

A script is provided to create a new role with these permissions. The name of
the role is `go-fullstack-test-role`.

    ./create-role.sh

The remainder of the scripts require information pertaining to the created role;
hence it is necessary to set an environment variable as indicated by the script
output.

## Create the database

Before testing the REST API, it is necessary to create a database; this simple
application uses very simple dynamodb. A script is provided to set up the database

    ./create-db.sh

## Uploading the lambda functions

Before uploading the lambda functions, it is necessary to ensure a set of go
modules are installed (unfortunately, I have not integrated all of this with
go modules as yet). Run the `get_go_dependencies.sh` script to get the necesary
dependencies to build the lambda functions.

AWS supports two specific operations relating to creation/modification of
lambda functions - upload and update - the former is intended for creating a
lambda function from scratch; the latter is intended for updating the
executable associated with a lambda function without updating other aspects (eg
permissions and roles, name, RAM requirements etc).

A set of scripts is provided which supports both modes; these scripts compile
each of the lambda functions and upload or update them.

A single script which runs all scripts is provided - to compile and upload all
scripts, use the following:

    cd lambda; ./upload-all-functions.sh

If the functions are changed, they can be updated as follows:

    cd lambda; ./upload-all-functions.sh -u

## Creating the Cognito User Pool and Client

Creating the user pool and client is straightforward, using the script provided:

    ./create-cognito-configuration.sh

## Creating the API gateway bindings

Once the Lambda functions are uploaded and the cognito user pools and client
generated, it is necessary to bind them to the API Gateway. This involves
modifying the provided swagger template to include user pool information and then
provide it to the API Gateway where the appropriate state can be created.

First, copy the `swagger-secure.yaml.tmpl` file to `swagger-secure.yaml`. Modify
this file to include the cognito user pool that should be used to perform
authentication. This is in the `providerARNs` field in the security definitions
in the swagger file.

Once this has been done, it is possible to create the bindings using the script
provided.

A script is provided which supports this.

    ./create-rest-api-secure.sh

## Testing

At present, testing is not so straightforward as it requires generation of a valid
access token.
