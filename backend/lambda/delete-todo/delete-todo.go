package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/google/uuid"

	"github.com/seanrmurphy/go-fullstack/backend/util"
)

// DeleteTodo deletes a todo specified with a uuid. In the case that this does
// not exist an error is generated.
func DeleteTodo(id uuid.UUID) (err error) {

	tableName := "Todos"

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	uuidBinary, _ := id.MarshalBinary()
	var resp *dynamodb.DeleteItemOutput
	resp, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				B: uuidBinary,
			},
		},
		ReturnValues: aws.String("ALL_OLD"),
	})

	// unlikely that an error occurs...
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("Unable to delete item with given ID")
		return
	}

	// check if the return value returned a sensible value
	if _, ok := resp.Attributes["ID"]; !ok {
		// not botherig to confirm ID is correct here; this should be done within
		// dynamodb
		err = errors.New("Item with given ID not found")
	}
	return
}

// HandleRequest performs some basic validation on the input id, if valid sends
// to the delete function and generates a return JSON string which is human readabled
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	idString := req.PathParameters["todoid"]
	if idString == "" {
		e := util.CreateResponse(http.StatusInternalServerError, "No valid ID provided")
		return e, nil
	}

	id, _ := uuid.Parse(idString)
	err := DeleteTodo(id)

	if err != nil {
		e := util.CreateResponse(http.StatusNotFound, "No object with given ID found")
		return e, nil
	}

	//return fmt.Sprintf("Hello %s!", name.Name), nil
	e := util.CreateResponse(http.StatusOK, "Record deleted")
	return e, nil
}

// main starts the lambda function
func main() {
	lambda.Start(HandleRequest)
}
