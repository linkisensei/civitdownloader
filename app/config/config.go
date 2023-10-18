package config

import (
	"log"

	"github.com/spf13/viper"
)

const INSTALLATION_PATH string = "installation_path"

// var viperInstance *viper.Viper

// func getViperInstance() *viper.Viper {
// 	if viperInstance == nil {

// 	}

// 	return viperInstance
// }

type AppConfig struct {
	Viper *viper.Viper
}

func (c *AppConfig) Init() error {
	if c.Viper != nil {
		return nil
	}

	c.Viper = viper.New()

	c.Viper.SetConfigName("config")
	c.Viper.SetConfigType("json")
	c.Viper.AddConfigPath(".")

	c.Viper.SetDefault(INSTALLATION_PATH, "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil
		} else {
			log.Fatal(err.Error())
			return err
		}
	}

	return nil
}

func (c *AppConfig) Set(key string, value interface{}) error {
	c.Init()
	c.Viper.Set(key, value)
	return c.Viper.SafeWriteConfig()
}

func (c *AppConfig) Get(key string) interface{} {
	c.Init()
	return c.Viper.Get(key)
}

func (c *AppConfig) GetString(key string) string {
	c.Init()
	log.Println(c.Viper.Get(key))
	return c.Viper.GetString(key)
}

func NewConfig() AppConfig {
	config := AppConfig{}
	config.Init()
	return config
}

var Config AppConfig = NewConfig()
