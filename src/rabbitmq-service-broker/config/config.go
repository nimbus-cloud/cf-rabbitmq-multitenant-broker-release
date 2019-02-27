package config

import (
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
}

type ServiceConfig struct {
	UUID      string `yaml:"uuid"`
	Name      string `yaml:"name"`
	Shareable bool   `yaml:"shareable"`
	PlanUUID  string `yaml:"plan_uuid"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
}

type RabbitMQConfig struct {
	Hosts         []string            `yaml:"hosts"`
	Administrator RabbitMQCredentials `yaml:"administrator"`
}

type RabbitMQCredentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Read(path string) (Config, error) {
	configBytes, err := ioutil.ReadFile(filepath.FromSlash(path))
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err = yaml.Unmarshal(configBytes, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
