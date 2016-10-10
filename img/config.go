package img

import (
	"errors"
	"github.com/BurntSushi/toml"
)

func ReadConfig(fname string) (ImgConfig, error) {
	var conf ImgConfig
	if _, err := toml.DecodeFile(fname, &conf); err != nil {
		return ImgConfig{}, errors.New("Malformed configuration file")
	}
	return conf, nil
}
