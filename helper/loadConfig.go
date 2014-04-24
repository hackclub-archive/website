package helper

import (
	"log"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
)

var config *yaml.File

func init() {
	var err error

	config, err = yaml.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal("Error loading config", err)
	}
}

func GetConfig(param string) string {
	val, err := config.Get(param)
	if err != nil || val == "" {
		val = os.Getenv(param)
	}

	return val
}
