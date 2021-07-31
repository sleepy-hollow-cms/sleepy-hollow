package config

import (
	"strings"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

const (
	prefix string = "SLEEPY_HOLLOW"
)

var (
	Conf *Config
)

type Config struct {
	Log     Log     `yaml:"log"`
	Server  Server  `yaml:"server"`
	MongoDB MongoDB `yaml:"mongodb"`
}

type Log struct {
	Encoding string `yaml:"encoding" default:"json"`
	Output   string `yaml:"output" default:"stdout"`
	Level    string `yaml:"level" default:"debug"`
}

type Server struct {
	Port int `yaml:"port" default:"3000"`
}

type MongoDB struct {
	User     string `yaml:"user" default:"root"`
	Password string `yaml:"password" default:"password"`
	Host     string `yaml:"host" default:"mongo"`
	Port     int    `yaml:"port" default:"27017"`
}

func init() {
	Conf = &Config{
		Log:     Log{},
		Server:  Server{},
		MongoDB: MongoDB{},
	}
}

func (c *Config) Load() error {
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

	return nil
}
