package httpserver

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	file, err := os.Open("configs/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}
