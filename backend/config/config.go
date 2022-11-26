package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Configs struct {
	JWT struct {
		TTLHour int    `yaml:"ttl"`
		Secret  string `yaml:"secret"`
	} `yaml:"jwt"`

	PG struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       string `yaml:"db"`
	} `yaml:"pg"`

	Bcrypt struct {
		Salt int `yaml:"salt"`
	} `yaml:"bcrypt"`

	App struct {
		Port int `yaml:"port"`
	} `yaml:"app"`
}

func Default() *Configs {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Configs{}

	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return config
}
