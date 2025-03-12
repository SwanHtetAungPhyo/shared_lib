package config

import "github.com/BurntSushi/toml"

type SMTPConfig struct {
	Protocol string `toml:"protocol"`
	Port     int    `toml:"port"`
	Email    string `toml:"email"`
	Password string `toml:"password"`
}

type Config struct {
	SMTP SMTPConfig `toml:"smtp"`
}

func LoadConfig(filePath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
