package configs

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
}

func GetDatabaseConfig() (Config, error) {
	databaseConfig := Config{}
	data, _ := ioutil.ReadFile("./configs/database.yaml")
	err := yaml.Unmarshal(data, &databaseConfig)
	return databaseConfig, err
}
