package controllers

import (
	"bytes"
	"customer/models"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	logrus.FieldLogger
}

func TestGetCustomer(t *testing.T) {

	tests := []struct {
		name           string
		expectedStatus int
		path           string
		status         string
		legalEntityID  string
	}{
		{
			name:           "customerGetApi",
			expectedStatus: 200,
			status:         "success",
			path:           "http://localhost:8081/v1",
			legalEntityID:  "6",
		},
		{
			name:           "customerGetApi",
			expectedStatus: 404,
			status:         "fail",
			path:           "http://localhost:8081/v1",
			legalEntityID:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			res, _ := http.Get(tt.path + "/customer/" + tt.legalEntityID)
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("status unmached: got %d, expectes %d", res.StatusCode,
					tt.expectedStatus)
			}
		})
	}

}

func TestInsertCustomer(t *testing.T) {
	type fields struct {
		FieldLogger logrus.FieldLogger
		Body        models.Customer
	}

	tests := []struct {
		name           string
		expectedStatus int
		status         string
		path           string
		fields         fields
	}{
		{
			name:           "customerInsertApi",
			expectedStatus: 200,
			path:           "http://localhost:8081/v1",
			status:         "success",
			fields: fields{
				Body: models.Customer{
					LegalEntityID:           1,
					BankruptcyIndicatorFlag: true,
					CompanyName:             "xcompany",
					FirstName:               "sandeep",
					LastName:                "Dev",
					LegalEntityStage:        "sandeep",
					LegalEntityType:         "sandeep",
				},
			},
		},
		{
			name:           "customerInsertApi",
			expectedStatus: 400,
			path:           "http://localhost:8081/v1",
			status:         "fail",
			fields: fields{
				Body: models.Customer{
					LegalEntityID:           1,
					BankruptcyIndicatorFlag: true,
					CompanyName:             "xcompany",
					FirstName:               "sandeep",
					LastName:                "Dev",
					LegalEntityStage:        "sandeep",
					LegalEntityType:         "sandeep",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(tt.fields.Body)

			res, _ := http.Post(tt.path+"/customer", "", b)
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("status unmached: got %d, expectes %d", res.StatusCode,
					tt.expectedStatus)
			}
		})
	}
}

func TestUpdateCustomer(t *testing.T) {
	type fields struct {
		FieldLogger logrus.FieldLogger
		Body        models.Customer
	}

	tests := []struct {
		name           string
		expectedStatus int
		path           string
		fields         fields
		status         string
		legalEntityID  string
	}{
		{
			name:           "customerUpdateApi",
			expectedStatus: 200,
			path:           "http://localhost:8081/v1",
			status:         "success",
			legalEntityID:  "7",
			fields: fields{
				Body: models.Customer{
					LegalEntityID:           6,
					BankruptcyIndicatorFlag: true,
					CompanyName:             "xcompany",
					FirstName:               "sandeep",
				},
			},
		},
		{
			name:           "customerUpdateApi",
			expectedStatus: 400,
			path:           "http://localhost:8081/v1",
			status:         "fail",
			legalEntityID:  "5",
			fields: fields{
				Body: models.Customer{
					LastName:         "Dev",
					LegalEntityStage: "sandeep",
					LegalEntityType:  "sandeep",
				},
			},
		},
		{
			name:           "customerUpdateApi",
			expectedStatus: 404,
			path:           "http://localhost:8081/v1",
			status:         "fail",
			legalEntityID:  "4",
			fields: fields{
				Body: models.Customer{
					LastName:         "Dev",
					LegalEntityStage: "sandeep",
					LegalEntityType:  "sandeep",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(tt.fields.Body)

			res, _ := http.Post(tt.path+"/customer/"+tt.legalEntityID, "", b)
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("status unmached: got %d, expectes %d", res.StatusCode,
					tt.expectedStatus)
			}
		})
	}
}

func TestRemoveCustomer(t *testing.T) {

	tests := []struct {
		name           string
		expectedStatus int
		path           string
		status         string
		legalEntityID  string
	}{
		{
			name:           "customerRemoveApi",
			expectedStatus: 200,
			path:           "http://localhost:8081/v1",
			status:         "success",
			legalEntityID:  "10",
		},
		{
			name:           "customerRemoveApi",
			expectedStatus: 404,
			path:           "http://localhost:8081/v1",
			status:         "error",
			legalEntityID:  "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(tt.path + "/customer/" + tt.legalEntityID)
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("status unmached: got %d, expectes %d", res.StatusCode,
					tt.expectedStatus)
			}
		})
	}
}

func TestSearchCustomer(t *testing.T) {
	type fields struct {
		FieldLogger logrus.FieldLogger
		Body        models.Customer
	}

	tests := []struct {
		name           string
		expectedStatus int
		status         string
		path           string
		fields         fields
	}{
		{
			name:           "customerSearchApi",
			expectedStatus: 200,
			path:           "http://localhost:8081/v1",
			status:         "success",
			fields: fields{
				Body: models.Customer{
					LegalEntityID:           1,
					BankruptcyIndicatorFlag: true,
					CompanyName:             "xcompany",
					FirstName:               "sandeep",
					LastName:                "Dev",
					LegalEntityStage:        "sandeep",
					LegalEntityType:         "sandeep",
				},
			},
		},
		{
			name:           "customerSearchApi",
			expectedStatus: 400,
			path:           "http://localhost:8081/v1",
			status:         "fail",
			fields: fields{
				Body: models.Customer{},
			},
		},
		{
			name:           "customerSearchApi",
			expectedStatus: 500,
			path:           "http://localhost:8081/v1",
			status:         "fail",
			fields: fields{
				Body: models.Customer{
					LegalEntityID:           11,
					BankruptcyIndicatorFlag: true,
					CompanyName:             "xcompany",
					FirstName:               "sandeep",
					LastName:                "Dev",
					LegalEntityStage:        "sandeep",
					LegalEntityType:         "sandeep",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := new(bytes.Buffer)
			json.NewEncoder(b).Encode(tt.fields.Body)

			res, _ := http.Post(tt.path+"/searchCustomer", "", b)
			if res.StatusCode != tt.expectedStatus {
				t.Errorf("status unmached: got %d, expectes %d", res.StatusCode,
					tt.expectedStatus)
			}
		})
	}
}
