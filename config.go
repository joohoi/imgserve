package main

import (
	"github.com/BurntSushi/toml"
	"imgserve/img"
)

func ReadConfig() img.ImgConfig {
	log.Debug("Initializing configuration")
	var conf img.ImgConfig
	if _, err := toml.DecodeFile("config.cfg", &conf); err != nil {
		log.Critical("Could not read configuration: ", err)
	}
	return conf
}
