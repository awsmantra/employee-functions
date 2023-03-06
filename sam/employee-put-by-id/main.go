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

		//Unmarshal request object
		dto := &dtos.EmployeeDto{}
		err := json.Unmarshal([]byte(request.Body), dto)

		if err != nil {
			return util.NewResponse(logger, request, "", 500, errors.New("internal server error. Unable to process request"))
		}

		//validate request
		dto.Id = id
		err = validator.Valid(dto)
		if err != nil {
			return util.NewResponse(logger, request, "", 400, err)
		}

		//Create model and call repository to save employee object
		employee := models.NewEmployeeFromUpdateDto(dto)
		err = employeeRepository.UpdateEmployee(ctx, employee)

		return util.NewResponse(logger, request, "employee updated successfully", 200, nil)
	}
	return util.NewResponse(logger, request, "PUT", 400, errors.New("missing path parameter"))
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

	logger.Info("employee-put-by-id function initialized successfully")
}
