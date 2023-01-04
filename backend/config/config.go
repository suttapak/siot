package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Configs struct {
	JWT struct {
		TTLHour int    `yaml:"ttl"`
		Secret  string `yaml:"secret"`
	} `yaml:"jwt"`

	PG struct {
		Username string `yaml:"username" `
		Password string `yaml:"password" `
		Host     string `yaml:"host"`
		Port     int    `yaml:"port" `
		DB       string `yaml:"db"`
	} `yaml:"pg"`

	Bcrypt struct {
		Salt int `yaml:"salt" `
	} `yaml:"bcrypt"`

	App struct {
		Port int `yaml:"port" `
	} `yaml:"app"`

	Mqtt struct {
		Port     int    `yaml:"port" `
		Broker   string `yaml:"broker" `
		Username string `yaml:"username"`
		Password string `yaml:"password" `
	}
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

	Config := Configs{}

	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
	fmt.Println(Config)

	return &Config
}
