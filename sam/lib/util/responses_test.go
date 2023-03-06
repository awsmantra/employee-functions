package util

import (
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/sirupsen/logrus"
	"testing"
)

type tests struct {
	request  events.APIGatewayProxyRequest
	response events.APIGatewayProxyResponse
	err      error
}

func TestNewResponseSuccess(t *testing.T) {

	logger := logrus.New()

	test := tests{
		request:  events.APIGatewayProxyRequest{},
		response: events.APIGatewayProxyResponse{},
		err:      nil,
	}

	response, _ := NewResponse(logger, test.request, "success", 200, nil)

	if response.Body == "success" {
		t.Fail()
	}
}

func TestNewResponseFail(t *testing.T) {

	logger := logrus.New()

	test := tests{
		request:  events.APIGatewayProxyRequest{},
		response: events.APIGatewayProxyResponse{},
		err:      nil,
	}

	response, _ := NewResponse(logger, test.request, "Fail", 400, errors.New("bad request"))

	if response.StatusCode != 400 {
		t.Fail()
	}
}
