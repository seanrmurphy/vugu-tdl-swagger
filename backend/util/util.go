package util

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type GenericReturnMessage struct {
	Message string
}

func CreateResponseWithCors(s int, b string) events.APIGatewayProxyResponse {

	headers := make(map[string]string)
	headers["Access-Control-Allow-Origin"] = "*"
	return events.APIGatewayProxyResponse{
		StatusCode: s,
		Headers:    headers,
		Body:       b,
	}
}

func CreateResponse(code int, s string) events.APIGatewayProxyResponse {
	m := GenericReturnMessage{
		Message: s,
	}
	b, _ := json.Marshal(&m)
	return CreateResponseWithCors(code, string(b))
}
