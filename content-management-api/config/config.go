package config

import (
	"github.com/creasty/defaults"
	"github.com/spf13/viper"
	"strings"
)

const (
	prefix string = "SLEEPY_HOLLOW"
)

var (
	Config *config
)

type config struct {
	Log     log     `yaml:"log"`
	Server  server  `yaml:"server"`
	MongoDB mongoDB `yaml:"mongodb"`
}

type log struct {
	Encoding string `yaml:"encoding" default:"json"`
	Output   string `yaml:"output" default:"stdout"`
}

type server struct {
	Port int `yaml:"port" default:"3000"`
}

type mongoDB struct {
	User     string `yaml:"user" default:"root"`
	Password string `yaml:"password" default:"password"`
	Host     string `yaml:"host" default:"mongo"`
	Port     int    `yaml:"port" default:"27017"`
}

func init() {
	Config = &config{
		Log:     log{},
		Server:  server{},
		MongoDB: mongoDB{},
	}
}

func (c *config) Load() error {
	defaults.Set(c)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		// TODO: output log when cannot load config file
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return err
	}

	Config = c
	return nil
}
