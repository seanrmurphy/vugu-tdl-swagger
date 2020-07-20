package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/seanrmurphy/go-fullstack/backend/model"
	"github.com/seanrmurphy/go-fullstack/backend/util"
)

// GetTodos gets an array of todos and returns them
func GetTodos() (tarray []model.Todo, e error) {

	tableName := "Todos"

	// Create the dynamo client object
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	proj := expression.NamesList(expression.Name("ID"), expression.Name("Title"), expression.Name("Completed"), expression.Name("CreationDate"))

	//expr, err := expression.NewBuilder().Build()
	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		log.Println("Got error building expression:", err.Error())
		e = err
		return
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		log.Println("Query API call failed:", err.Error())
		e = err
		return
	}

	for _, i := range result.Items {
		t := model.Todo{}

		err = dynamodbattribute.UnmarshalMap(i, &t)

		if err != nil {
			log.Println("Got error unmarshalling:", err.Error())
			e = err
			return
		} else {
			tarray = append(tarray, t)
		}
	}

	return
}

// HandleRequest obtains the set of todos from dynamodb
func HandleRequest(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// TODO(murp): add some error handling here
	tarray, _ := GetTodos()

	tbody, _ := json.Marshal(tarray)
	return util.CreateResponseWithCors(http.StatusOK, string(tbody)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
