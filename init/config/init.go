package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"
)

const Version = "0.0.1"

var cfg = new(config)

func Init(file string) {
	byt, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Sprintf("Load config error, file: [%s], err: [%v]", file, err))
	}

	if err = yaml.Unmarshal(byt, cfg); err != nil {
		panic(fmt.Sprintf("Read config error, file: [%s], err: [%v]", file, err))
	}
}

func GetConfig() *config {
	return cfg
}
