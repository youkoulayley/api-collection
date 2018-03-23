package bootstrap

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/youkoulayley/api-collection/models"
)

// GetConf get the config file
func GetConf(confPath string, c string) *models.Conf {
	log.Info("Conf - Path defined at ", confPath+c)
	file, err := os.Open(confPath + c)
	if err != nil {
		// If the conf file is not loaded, we exit the app
		log.Fatal("Conf - Failed to open file : ", err.Error())
	} else {
		log.Debug("Conf - Open")
	}
	defer file.Close()
	defer log.Debug("Conf - Closed")

	// Decode the conf file and initiate the model
	decoder := json.NewDecoder(file)
	configuration := models.Conf{}

	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("Conf - Failed to decode conf file : ", err.Error())
	}

	return &configuration
}
