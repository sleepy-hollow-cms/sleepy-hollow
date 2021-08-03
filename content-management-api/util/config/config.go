package config

import (
	"strings"

	"github.com/creasty/defaults"
	"github.com/spf13/viper"
)

const (
	prefix            string = "SLEEPY_HOLLOW"
	configFileKey     string = "config"
	configDefaultPath string = "/etc/sleepy-hollow/config.yaml"
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
	viper.SetDefault(configFileKey, configDefaultPath)

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(viper.GetString(configFileKey))

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return err
	}

	return nil
}
