package validator

import (
	"employee-functions/lib/dtos"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate_Failure(t *testing.T) {

	dto := &dtos.EmployeeDto{}

	err := Valid(dto)

	assert.Equal(t, fmt.Sprint(err), "EmployeeId is a required field", "dto validation got failed")
}

func TestValidate_Success(t *testing.T) {

	dto := &dtos.EmployeeDto{}

	err := Valid(dto)

	if err != nil {
		t.Fail()
	}
}
