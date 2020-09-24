package controllers

import (
	"customer/dbconnection"
	"customer/models"
	"customer/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// InsertCustomer details
func InsertCustomer(c *gin.Context) {

	customer := models.Customer{}

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": "fail",
			"data":   "invalid request body; " + err.Error(),
		})
		return
	}

	customer.CreatedDate = time.Now().UTC()
	db := dbconnection.Get()

	if _, err := db.Model(&customer).Insert(); err == nil {
		c.JSON(http.StatusCreated, gin.H{
			"status": "success",
			"data": gin.H{
				"customer": customer,
			},
		})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"data":   err.Error(),
		})
	}
}

// GetCustomer Details by legalEntityID
func GetCustomer(c *gin.Context) {

	legalEntityID := utils.ParamID(c, "legalEntityID")

	customer := models.Customer{}
	db := dbconnection.Get()

	err := db.Model(&customer).Where("legal_entity_id = ?", legalEntityID).Select()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"data": gin.H{
				"customer": customer,
			},
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status": "fail",
		"data": gin.H{
			"customer": nil,
		},
	})

}
