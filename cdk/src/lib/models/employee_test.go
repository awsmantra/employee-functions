package models

import (
	"allergies/lib/dtos"
	"testing"
	"time"
)

func TestAllergy_GetPK(t *testing.T) {
	medication := &Allergies{
		PK: "TENANT#697190",
	}

	if medication.PK != "TENANT#697190" {
		t.Fail()
	}
}

func TestNewAllergyPK(t *testing.T) {
	newPK := NewAllergyPK(697190)

	if newPK != "TENANT#697190" {
		t.Fail()
	}

}

func TestNewAllergySK(t *testing.T) {
	newSK := NewAllergySK(78901, "food", "294F0Ad8RiEnHrsXFfTtZqZswkR")

	if newSK != "PATIENTCHARTID#78901#TYPE#food#294F0Ad8RiEnHrsXFfTtZqZswkR" {
		t.Fail()
	}
}

func TestNewAllergy(t *testing.T) {
	allergyDto := dtos.AllergiesDto{
		AllergyType:    "food",
		AllergyName:    "Eggs",
		AllergyId:      "294F0Ad8RiEnHrsXFfTtZqZswkR",
		TenantId:       78901,
		PatientChartId: 12345,
		Reaction:       "Rashes",
		CreatedBy:      "manuel@manuel.com",
		CreatedTs:      time.Now(),
	}
	allergyClaim := Claim{
		TenantId: 78901,
		Username: "manuel",
		UserId:   "manuel@manuel.com",
	}
	newAllergy := NewAllergy(&allergyDto, &allergyClaim, 12345, 78901)

	if newAllergy.AllergyName != "Eggs" {
		t.Fail()
	}
	if newAllergy.AllergyType != "food" {
		t.Fail()
	}
	if newAllergy.PatientChartId != 12345 {
		t.Fail()
	}
}

func TestEditedAllergy(t *testing.T) {
	allergy := Allergies{
		AllergyType:      "food",
		AllergyName:      "Eggs",
		AllergyId:        "294F0Ad8RiEnHrsXFfTtZqZswkR",
		TenantId:         78901,
		PatientChartId:   12345,
		Reaction:         "Rashes",
		ReasonForChange:  "By Mistake",
		ModifiedBy:       "uj@uj.com",
		ModifiedDatetime: time.Now(),
		PK:               "TENANT#78901",
		SK:               "PATIENTCHARTID#12345#TYPE#FOOD#294F0Ad8RiEnHrsXFfTtZqZswkR",
	}
	dto := dtos.AllergyReasonForChangeDto{
		ReasonForChange: "By Mistake",
		AllergyType:     "food",
		AllergyName:     "miscinline",
		Reaction:        "rashes",
	}
	allergyClaim := Claim{
		TenantId: 78901,
		Username: "uj",
		UserId:   "uj@uj.com",
	}
	newAllergy, _ := EditedAllergy(&dto, &allergy, &allergyClaim)
	if newAllergy.AllergyName != "miscinline" {
		t.Fail()
	}
	if newAllergy.AllergyType != "food" {
		t.Fail()
	}
	if newAllergy.PatientChartId != 12345 {
		t.Fail()
	}
}

func TestSetDeleteAllergy(t *testing.T) {
	allergy := Allergies{
		AllergyType:     "food",
		AllergyName:     "Eggs",
		AllergyId:       "294F0Ad8RiEnHrsXFfTtZqZswkR",
		TenantId:        78901,
		PatientChartId:  12345,
		Reaction:        "Rashes",
		CreatedBy:       "manuel@manuel.com",
		CreatedDatetime: time.Now(),
		IsDeleted:       false,
		PK:              "TENANT#78901",
		SK:              "PATIENTCHARTID#12345#TYPE#FOOD#294F0Ad8RiEnHrsXFfTtZqZswkR",
	}
	allergyReasonForDeleteDto := dtos.AllergyReasonForDeleteDto{
		ReasonForDelete: "Update medication",
	}

	allergyClaim := Claim{
		TenantId: 78901,
		Username: "manuel",
		UserId:   "manuel@manuel.com",
	}
	newAllergy, _ := SetDeletedAllergy(&allergyReasonForDeleteDto, &allergy, &allergyClaim)

	if newAllergy.PK != "TENANT#78901" {
		t.Fail()
	}
	if newAllergy.AllergyType != "food" {
		t.Fail()
	}
	if newAllergy.PatientChartId != 12345 {
		t.Fail()
	}
}

func TestNewHistoryMetaDataSuccess(t *testing.T) {
	allergy := Allergies{
		AllergyType:     "food",
		AllergyName:     "Eggs",
		AllergyId:       "294F0Ad8RiEnHrsXFfTtZqZswkR",
		TenantId:        78901,
		PatientChartId:  12345,
		Reaction:        "Rashes",
		CreatedBy:       "manuel@manuel.com",
		CreatedDatetime: time.Now(),
		IsDeleted:       false,
		PK:              "TENANT#78901",
		SK:              "PATIENTCHARTID#12345#TYPE#FOOD#294F0Ad8RiEnHrsXFfTtZqZswkR",
		HistoryMetaData: "[{\"PK\":\"TENANT#697195\",\"SK\":\"PATIENTCHARTID#78906#TYPE#food#29Lg3xS40PfTN2bLM3SvtxnCvFj\",\"AllergyType\":\"food\",\"AllergyName\":\"Eggs\",\"SortId\":0,\"AllergyId\":\"29Lg3xS40PfTN2bLM3SvtxnCvFj\",\"TenantId\":697195,\"PatientChartId\":78906,\"Reaction\":\"Rashes\",\"CreatedBy\":\"manuel@manuel.com\",\"ModifyBy\":\"\",\"DeletedBy\":\"\",\"CreatedTs\":\"2022-05-18T18:13:04.9339826Z\",\"ModifyTs\":null,\"ReasonForChange\":\"\",\"ReasonForDelete\":\"\",\"IsDelete\":false,\"IsEdited\":false,\"HistoryMetaData\":\"\"}]",
	}

	newHistoryMetaData, _ := NewHistoryMetaData(allergy)
	if newHistoryMetaData == "" {
		t.Fail()
	}
}

func TestNewHistoryMetaDataFailure(t *testing.T) {
	allergy := Allergies{
		AllergyType:     "food",
		AllergyName:     "Eggs",
		AllergyId:       "294F0Ad8RiEnHrsXFfTtZqZswkR",
		TenantId:        78901,
		PatientChartId:  12345,
		Reaction:        "Rashes",
		CreatedBy:       "manuel@manuel.com",
		CreatedDatetime: time.Now(),
		IsDeleted:       false,
		PK:              "TENANT#78901",
		SK:              "PATIENTCHARTID#12345#TYPE#FOOD#294F0Ad8RiEnHrsXFfTtZqZswkR",
		HistoryMetaData: "[{\"PK\"NTCHARTID#78906#TYPE#food#29Lg3xS40PfTN2bLM3SvtxnCvFj\",\"AllergyType\":\"food\",\"AllergyName\":\"Eggs\",\"SortId\":0,\"AllergyId\":\"29Lg3xS40PfTN2bLM3SvtxnCvFj\",\"TenantId\":697195,\"PatientChartId\":78906,\"Reaction\":\"Rashes\",\"CreatedBy\":\"manuel@manuel.com\",\"ModifyBy\":\"\",\"DeletedBy\":\"\",\"CreatedTs\":\"2022-05-18T18:13:04.9339826Z\",\"ModifyTs\":null,\"ReasonForChange\":\"\",\"ReasonForDelete\":\"\",\"IsDelete\":false,\"IsEdited\":false,\"HistoryMetaData\":\"\"}]",
	}

	_, err := NewHistoryMetaData(allergy)
	if err != nil {
		t.Fail()
	}
}

func TestMapAllergiesToDto(t *testing.T) {

	//Arrange
	allergy := &Allergies{
		TenantId:        697190,
		PatientChartId:  98765,
		AllergyId:       "294F0Ad8RiEnHrsXFfTtZqZswkR",
		AllergyType:     "drug",
		AllergyName:     "Aspirin",
		CreatedBy:       "Rashes",
		CreatedDatetime: time.Now(),
	}

	//Act
	dto := MapAllergiesToDto(allergy)

	//Assert
	if dto.AllergyId != "294F0Ad8RiEnHrsXFfTtZqZswkR" {
		t.Fail()
	}
}

func TestAllergyType_GetPK(t *testing.T) {

	allergyType := &AllergyType{
		PK: "TENANT#697190",
	}

	if allergyType.PK != "TENANT#697190" {
		t.Fail()
	}
}
