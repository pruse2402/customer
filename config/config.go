package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	DEV_MODE    bool   `json:"dev_mode"`
	STAG_MODE   bool   `json:"stag_mode"`
	PROD_MODE   bool   `json:"prod_mode"`
	APP_VERSION string `json:"app_version"`

	// SERVER CONFIG
	APP_URL     string `json:"app_url"`
	DB_NAME     string `json:"db_name"`
	DB_USERNAME string `json:"db_username"`
	DB_PASSWORD string `json:"db_password"`
	DB_Address  string `json:"db_address"`
	Port        int    `json:"port"`
}

var (
	Cfg  Config
	once sync.Once
)

// Parse parses the json configuration file
// And converting it into native type
func Parse(file string) error {
	once.Do(func() {
		// Reading the flags
		data, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("Error in ReadFile:", err)
		}
		if err := json.Unmarshal(data, &Cfg); err != nil {
			log.Println("Error in Unmarshal:", err)
		}
	})
	return nil
}
