package controllers

import (
	"github.com/gin-gonic/gin"
)

type QueryParams struct {
	CompanyName string
	FirstName   string
	LastName    string
}

// ListQueryParams which the arams struct
func ListQueryParams(c *gin.Context) (QueryParams, error) {
	query := c.Request.URL.Query()
	queryParams := QueryParams{
		CompanyName: query.Get("companyName"),
		FirstName:   query.Get("firstName"),
		LastName:    query.Get("lastName"),
	}

	return queryParams, nil
}
