package main

import (
	"context"
	"employee-functions/lib/clients"
	"employee-functions/lib/data"
	"employee-functions/lib/dtos"
	"employee-functions/lib/models"
	"employee-functions/lib/util"
	"employee-functions/lib/validator"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	logger             *logrus.Logger
	employeeRepository data.EmployeeRepository
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Unmarshal request object
	dto := &dtos.EmployeeDto{}
	err := json.Unmarshal([]byte(request.Body), dto)

	if err != nil {
		return util.NewResponse(logger, request, "", 500, errors.New("internal server error. unable to process request"))
	}

	//validate request
	err = validator.Valid(dto)
	if err != nil {
		return util.NewResponse(logger, request, "", 400, err)
	}

	//create model and call repository
	employee := models.NewEmployeeFromDto(dto)
	_, err = employeeRepository.CreateEmployee(ctx, employee)

	if err != nil {
		return util.NewResponse(logger, request, "internal server error. Unable to process request", 500, err)
	}

	return util.NewResponse(logger, request, "employee created successfully", 200, nil)
}

func main() {
	lambda.Start(Handler)
}

func init() {
	logger = logrus.New()
	util.SetLogLevel(logger, os.Getenv("LOG_LEVEL"))
	awsRegion := os.Getenv("REGION")
	dbClient := clients.NewDynamoDBClient(awsRegion)

	employeeRepository = &data.EmployeeDao{
		DB:     dbClient,
		Logger: logger,
	}

	logger.Info("employee-post function initialized successfully")
}
