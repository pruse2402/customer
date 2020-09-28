package main

import (
	"customer/config"
	dbcon "customer/dbconnection"
	"customer/dbscripts"
	"customer/routes"
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Commandline flags to specify
	var configFile = flag.String("conf", "", "configuration file(mandatory)")
	flag.Parse()
	if flag.NFlag() != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parsing configuration
	if err := config.Parse(*configFile); err != nil {
		log.Fatalln("ERROR: ", err)
	}

	dbcon.Connect()
	defer dbcon.Close()

	dbscripts.InitDB()
	log.Info("Debug mode is ON")

	r := gin.Default()
	routes.InitRoutes(r)

	r.Run(fmt.Sprintf(":%d", config.Cfg.Port))
}
