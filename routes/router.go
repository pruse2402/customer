package routes

import (
	"customer/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	authOnlyRoute := r.Group("/v1")

	// APIs for Customer
	authOnlyRoute.POST("/customer", controllers.InsertCustomer) // CREATE new
}
