package infrastucture

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MongoConfig struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"dbName"`
	} `yaml:"mongodb"`
}

func processConfigError(err error) {
	log.Println("Problem with reading config file.")
	log.Fatal(err)
}

func ReadConfig(configPath string) Config {
	f, err := os.Open(configPath)
	if err != nil {
		processConfigError(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processConfigError(err)
	}
	return cfg
}
