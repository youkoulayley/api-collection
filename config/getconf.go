package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/youkoulayley/api-collection/models"
)

// GetConf get the config file
func GetConf() *models.Conf {
	file, err := os.Open("config/conf.json")
	if err != nil {
		fmt.Printf("Can't open file : %s", err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := models.Conf{}

	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}

	return &configuration
}
