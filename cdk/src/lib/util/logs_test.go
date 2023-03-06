package util

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogs_Debug(t *testing.T) {
	//Arrange
	logger := logrus.New()

	//Act
	SetLogLevel(logger, "debug")

	//Assert
	if logger.GetLevel().String() != "debug" {
		t.Fail()
	}
}

func TestLogs_Error(t *testing.T) {
	//Arrange
	logger := logrus.New()

	//Act
	SetLogLevel(logger, "error")

	//Assert
	if logger.GetLevel().String() != "error" {
		t.Fail()
	}
}

func TestLogs_Info(t *testing.T) {
	//Arrange
	logger := logrus.New()

	//Act
	SetLogLevel(logger, "info")

	//Assert
	if logger.GetLevel().String() != "info" {
		t.Fail()
	}
}

func TestLogs_Null_Level(t *testing.T) {
	//Arrange
	logger := logrus.New()

	//Act
	SetLogLevel(logger, "")

	//Assert
	if logger.GetLevel().String() != "debug" {
		t.Fail()
	}
}
