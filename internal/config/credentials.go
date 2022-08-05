package config

import (
	"io/ioutil"

	"github.com/BRO3886/gtasks/internal/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/tasks/v1"
)

// ReadCredentials reads the config.json file
func ReadCredentials() *oauth2.Config {
	folderPath := GetInstallLocation()
	b, err := ioutil.ReadFile(folderPath + "/config.json")
	if err != nil {
		utils.ErrorP("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, tasks.TasksScope)
	if err != nil {
		utils.ErrorP("Unable to parse client secret file to config: %v", err)
	}
	return config
}
