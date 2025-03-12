package config

import (
	"github.com/BurntSushi/toml"
	"math/rand"
	"strconv"
	"time"
)

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

func CodeGen() string {
	rand.Seed(time.Now().UnixNano())
	code := ""
	for i := 0; i < 6; i++ {
		code += strconv.Itoa(rand.Intn(9))
	}

	return code
}
