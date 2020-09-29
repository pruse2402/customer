package dbconnection

import (
	"customer/config"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

//Connect database
func Connect() {
	dbCon := pg.Connect(&pg.Options{
		Addr:     config.Cfg.Database.Address,
		User:     config.Cfg.Database.Username,
		Password: config.Cfg.Database.Password,
		Database: config.Cfg.Database.Name,
	})

	db = dbCon
	log.Printf("Connected successfully")

	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
		log.Fatalf("Unable to connect Postgres Database: %v\n", err)
		os.Exit(1)
	}
}

//Get db connection
func Get() *pg.DB {
	return db
}

//Close db connection
func Close() {
	err := db.Close()

	if err != nil {
		log.Printf("Closing DB err", err)
	}
	log.Printf("DB closed")
}
