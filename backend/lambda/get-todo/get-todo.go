package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"github.com/seanrmurphy/go-fullstack/backend/model"
	"github.com/seanrmurphy/go-fullstack/backend/util"
)

// GetTodo gets a todo with the specified id; returns an error if this does not
// exist
func GetTodo(id uuid.UUID) (t model.Todo, e error) {

	tableName := "Todos"

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	uuidBinary, _ := id.MarshalBinary()
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				B: uuidBinary,
			},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
		e = err
		return
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &t)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return
}

// HandleRequest extracts the id of the record, passes it on to the get function,
// marshalls it and returns it
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idString := req.PathParameters["todoid"]
	if idString == "" {
		return util.CreateResponseWithCors(http.StatusInternalServerError, "No ID provided"), nil
	}

	id, _ := uuid.Parse(idString)
	t, _ := GetTodo(id)
	// TODO(murp): add error checking here

	tbody, _ := json.Marshal(t)
	return util.CreateResponseWithCors(http.StatusOK, string(tbody)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
