package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
		TimeZone string `yaml:"TimeZone"`
	} `yaml:"database"`
}

func ReadConfig() *Config {
	filename := "config/config.yaml"
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		panic(err)
	}
	return c
}
