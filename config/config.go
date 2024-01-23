package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

const confFile = "./config/config.yaml"

type Config struct {
	HTTP     struct{ Port int }
	GRCP     struct{ Port int }
	POSTGRES struct {
		Host     string
		User     string
		Password string
		Dbname   string
		Port     int
	}
}

func NewConfigStruct() (*Config, error) {
	con := &Config{}
	file, err := os.Open(confFile)
	if err != nil {
		return nil, fmt.Errorf("Failed open config file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Failed reading file :%v", err)
	}
	err = yaml.Unmarshal(read, con)
	if err != nil {
		return nil, fmt.Errorf("Yaml unmarshall error :%v", err)
	}
	return con, nil
}
