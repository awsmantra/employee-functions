package clients

import (
	"testing"
)

func TestNewDynamoDBClient(t *testing.T) {

	//Act
	dbClient := NewDynamoDBClient()

	//Assert
	if dbClient == nil {
		t.Fail()
	}
}
