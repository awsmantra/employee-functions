package util

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
)

func NewResponse(logger *logrus.Logger, request events.APIGatewayProxyRequest, response interface{}, status int, err error) (events.APIGatewayProxyResponse, error) {
	m, _ := json.Marshal(response)
	var body string

	if err != nil {
		body = err.Error()
	} else {
		body = string(m)
	}

	e := events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Method":      "OPTIONS,POST,GET,PUT,DELETE",
		},
	}

	logger.WithFields(logrus.Fields{
		"request":  request,
		"response": response,
	}).Tracef("Request and response")

	return e, nil

}
