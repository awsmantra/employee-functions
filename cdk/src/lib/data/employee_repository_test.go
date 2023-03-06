package data

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	allergiesRepository EmployeeRepository
)

type MockEmployeeDao struct {
	dynamodbiface.DynamoDBAPI
	TestSuccess        bool
	TestEmptyResultSet bool
	TestSuccessEmpty   bool
	TestAllergySuccess bool
}
