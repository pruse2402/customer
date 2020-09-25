package models

import (
	"time"
)

type Customer struct {
	LegalEntityID           int64     `json:"legalEntityID"`
	BankruptcyIndicatorFlag bool      `json:"bankruptcyIndicatorFlag"`
	CompanyName             string    `json:"companyName" sql:",notnull"`
	FirstName               string    `json:"firstName" sql:",notnull"`
	LastName                string    `json:"lastName"`
	LegalEntityStage        string    `json:"legalEntityStage"`
	LegalEntityType         string    `json:"legalEntityType"`
	CreatedDate             time.Time `json:"createdDate"`
	DateOfBirth             time.Time `json:"dateOfBirth"`
}

type CustomerRequestBody struct {
	CompanyName   string `json:"companyName"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	LegalEntityID int64  `json:"legalEntityID"`
}
