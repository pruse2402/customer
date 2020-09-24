package main

import (
	"customer/config"
	dbcon "customer/dbconnection"
	"customer/dbscripts"
	"customer/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	dbcon.Connect()
	defer dbcon.Close()

	dbscripts.InitDB()
	log.Info("Debug mode is ON")

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(fmt.Sprintf(":%d", config.Port))
}
