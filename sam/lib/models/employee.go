package models

import (
	"employee-functions/lib/dtos"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Employee struct {
	PK               string    `dynamodbav:"PK"`
	SK               string    `dynamodbav:"SK"`
	FirstName        string    `dynamodbav:"FirstName"`
	LastName         string    `dynamodbav:"LastName"`
	Age              int       `dynamodbav:"Age"`
	CreatedDatetime  time.Time `dynamodbav:"CreatedDatetime,omitempty"`
	ModifiedDatetime  time.Time `dynamodbav:"ModifiedDatetime,omitempty"`
}

type EmployeeList struct {
	Employees []Employee
}

func NewEmployeeFromDto(dto *dtos.EmployeeDto) *Employee {
	empId := getRandomNumber(1000,9999) //For simplicity using random number for employeeId
	fmt.Println("new employeeId ", empId)

	employee := Employee{
		PK:               "EMPLOYEE",
		SK:               fmt.Sprint("ID#", empId),
		FirstName:        dto.FirstName,
		LastName:         dto.LastName,
		Age:              dto.Age,
		CreatedDatetime:  time.Now(),
	}
	return &employee
}

func NewEmployeeDtoFromEmployee(employee Employee) *dtos.EmployeeDto {
	_id, err := strconv.Atoi(fmt.Sprint(employee.SK[3:])) //get the value after ID#

	if err != nil {
		return nil  //modified this code..
	}
	dto := dtos.EmployeeDto{
		Id:              _id,
		FirstName:       employee.FirstName,
		LastName:        employee.LastName,
		Age:             employee.Age,
		CreatedDatetime: employee.CreatedDatetime,
		ModifiedDatetime: employee.ModifiedDatetime,
	}
	return &dto
}


func NewEmployeeFromUpdateDto(dto *dtos.EmployeeDto) *Employee {

	employee := Employee{
		PK:               "EMPLOYEE",
		SK:               fmt.Sprint("ID#", dto.Id),
		FirstName:        dto.FirstName,
		LastName:         dto.LastName,
		Age:              dto.Age,
		CreatedDatetime:  dto.CreatedDatetime,
		ModifiedDatetime:  time.Now(),
	}
	return &employee
}


func getRandomNumber(low, hi int) int {
	fmt.Println("Current Time in UnixNano ", time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	return low + rand.Intn(hi-low)
}