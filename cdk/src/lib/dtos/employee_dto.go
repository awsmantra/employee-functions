package dtos

import "time"

type EmployeeDto struct {
	Id              int       `json:"Id"`
	FirstName       string    `json:"firstName" validate:"required"`
	LastName        string    `json:"lastName" validate:"required"`
	Age             int       `json:"age"  validate:"required"`
	CreatedDatetime time.Time `json:"createdDatetime,omitempty"`
	ModifiedDatetime time.Time `json:"modifiedDatetime,omitempty"`
}
