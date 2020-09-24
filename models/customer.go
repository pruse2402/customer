package models

import (
	"time"
)

type Customer struct {
	LegalEntityID           int       `json:"legalEntityID"`
	BankruptcyIndicatorFlag bool      `json:"bankruptcyIndicatorFlag"`
	CompanyName             string    `json:"companyName" sql:",notnull"`
	FirstName               string    `json:"firstName" sql:",notnull"`
	LastName                string    `json:"lastName"`
	LegalEntityStage        string    `json:"legalEntityStage"`
	LegalEntityType         string    `json:"legalEntityType"`
	CreatedDate             time.Time `json:"createdDate"`
	DateOfBirth             time.Time `json:"dateOfBirth"`
}
