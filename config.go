package main

import (
	"github.com/BurntSushi/toml"
)

type mainConfig struct {
	Path   string
	Widths []int
}

func ReadConfig() mainConfig {
	log.Debug("Initializing configuration")
	var conf mainConfig
	if _, err := toml.DecodeFile("config.cfg", &conf); err != nil {
		log.Critical("Could not read configuration: ", err)
	}
	return conf
}
