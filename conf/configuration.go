package conf

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	Port            int
	Database        string
	Collection      string
	DatabasePort    int
	DatabaseAddress string
}

func New() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("conf.json", &configuration)
	if err != nil {
		log.Fatal(err)
	}
	return configuration
}
