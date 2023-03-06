package data

import (
	"context"
	"employee-functions/lib/models"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	log "github.com/sirupsen/logrus"
)

type EmployeeRepository interface {
	CreateEmployee(context.Context, *models.Employee) (*dynamodb.PutItemOutput, error)
	GetEmployee(context.Context, int) (*models.Employee, error)
	DeleteEmployee(context.Context, int) error
	UpdateEmployee(context.Context, *models.Employee)  error
}

type EmployeeDao struct {
	DB     dynamodbiface.DynamoDBAPI
	Logger *log.Logger
}

func (dao *EmployeeDao) GetEmployee(ctx context.Context, id int) (*models.Employee, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":PK": {
				S: aws.String("EMPLOYEE"),
			},
			":SK": {
				S: aws.String(fmt.Sprintf("ID#%d", id)),
			},
		},
		KeyConditionExpression: aws.String("PK = :PK AND SK = :SK"),
		TableName:              aws.String("Employee"),
	}
	result, err := dao.DB.QueryWithContext(ctx, input)

	if err != nil {
		dao.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("error while getting employee")
		return nil, err
	}

	emp := &models.Employee{}
	for _, item := range result.Items {
		_ = dynamodbattribute.UnmarshalMap(item, &emp)
	}
    return emp, nil
}

func (dao *EmployeeDao) CreateEmployee(ctx context.Context, employee *models.Employee) (*dynamodb.PutItemOutput, error) {
	marshalledEvent, err := dynamodbattribute.MarshalMap(employee)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      marshalledEvent,
		TableName: aws.String("Employee"),
	}

	out, err := dao.DB.PutItemWithContext(ctx, input)
	if err != nil {
		dao.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("error creating employee")
		return nil, err
	}

	return out, nil

}

func (dao *EmployeeDao) DeleteEmployee(ctx context.Context, id int) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String("EMPLOYEE"),
			},
			"SK": {
				S: aws.String(fmt.Sprintf("ID#%d", id)),
			},
		},
		TableName: aws.String("Employee"),
		ReturnValues: aws.String("ALL_OLD"),
	}

	result, err := dao.DB.DeleteItemWithContext(ctx, input)

	if err != nil {
		dao.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("error while getting Employee")
		return err
	}
	fmt.Println("returned values ", result) // check whether it returned null value for OLD_VALUE, if yes ,Employee doesn't exist
	return nil
}

func (dao *EmployeeDao) UpdateEmployee(ctx context.Context, employee *models.Employee)  error {
	marshalledEvent, err := dynamodbattribute.MarshalMap(employee)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      marshalledEvent,
		TableName: aws.String("Employee"),
	}

	//PutItem replaces any existing item that has the same key, use updateItem if you want to merge new attributes.
	_, err = dao.DB.PutItemWithContext(ctx, input)

	if err != nil {
		dao.Logger.WithFields(log.Fields{
			"err": err,
		}).Error("error while updating employee record")
		return err
	}
	return nil
}
