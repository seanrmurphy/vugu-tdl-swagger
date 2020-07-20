# Sample backend for Todo application

This is a backend that provides as a set of AWS Lambda functions written in Go.
The basic idea is that each lambda function maps to REST API endpoint and hence
it is necessary to configure this via the API Gateway.

# Setting up the application

Setting up the application requires the following:
- create a new role with appropriate trusts and policies
- upload lambda functions to AWS Lambda
- create a database
- create and configure a new REST API Gateway

Each of the following are dealt with in the subsections below.

It is worth noting that the automation scripts are not very robust right now and
do not perform error checking or state determination prior to performing their
actions.

## Create a new role

It is necessary to have a role which can trigger lambda functions and access
DynamoDB; also, it is necessary that this role by trusted by AWS Lambda services
and AWS API Gateway.

A script is provided to create a new role with these permissions. The name of
the role is `go-fullstack-test-role`.

    ./create-role.sh

The remainder of the scripts require information pertaining to the created role;
hence it is necessary to set an environment variable as indicated by the script.

## Uploading the lambda functions

AWS supports two operations relating to lambda functions - upload and update -
the latter is intended for updating the executable associated with a lambda
function without updating other aspects (eg permissions and roles, name, RAM
requirements etc).

A set of scripts is provided which supports both modes; these scripts compile
each of the lambda functions and upload or update them.

A single script which runs all scripts is provided - to compile and upload all
scripts, use the following:

    cd lambda; ./upload-all-functions.sh

If the functions are changed, they can be updated as follows:

    cd lambda; ./upload-all-functions.sh -u

## Creating the API gateway bindings

Once the Lambda functions are uploaded, it is necessary to bind them to the
API Gateway. This involves creating the REST API based on the provided `swagger.yaml`
file and binding the endpoints to the appropriate lambda functions.

A script is provided which supports this.

    ./create-rest-api.sh

As with the role, it is necessary to set an environment variable to test the
API - follow the instructions provided by the script.

## Create the database

Before testing the REST API, it is necessary to create a database; this simple
application uses very simple dynamodb. A script is provided to set up the database

    ./create-db.sh

# Initial testing

Ensure the db is configured; there is a create-db.sh helper script to support this.

Go through the following steps:
- Create todo

    `./test/create-todo-test.sh`

- List all todos

    `./test/list-todos-test.sh`

- Get specific todo

    `./test/get-todo-test.sh`

- Update specific todo

    `./test/update-todo-test.sh`

- Delete todo

    `./test/delete-todo-test.sh`

