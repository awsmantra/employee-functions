package main

import (
	"context"
	"employee-functions/lib/clients"
	"employee-functions/lib/data"
	"employee-functions/lib/models"
	"employee-functions/lib/util"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var (
	logger             *logrus.Logger
	employeeRepository data.EmployeeRepository
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get id from PathParameter
	if idStr, ok := request.PathParameters["id"]; ok {
		id, _ := strconv.Atoi(idStr)
		emp, err := employeeRepository.GetEmployee(ctx, id)

		if err != nil {
			return util.NewResponse(logger, request, "", 500, errors.New("internal server error. unable to process request"))
		}

		//if employee not found in database
		if emp.PK == "" {
			return util.NewResponse(logger, request, "resource not found", 404, err)
		}

		//Return employeeDto
		newDto := models.NewEmployeeDtoFromEmployee(*emp)
		return util.NewResponse(logger, request, newDto, 200, nil)

	}
	return util.NewResponse(logger, request, "POST", 400, errors.New("missing path parameter"))
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

	logger.Info("employee-get-by-id function initialized successfully")
}
