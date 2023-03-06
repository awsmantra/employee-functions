package clients

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient(region string) dynamodbiface.DynamoDBAPI {
	c := &aws.Config{
		Region: aws.String(region)}
	sess := session.Must(session.NewSession(c))
	svc := dynamodb.New(sess)
	return dynamodbiface.DynamoDBAPI(svc)
}
