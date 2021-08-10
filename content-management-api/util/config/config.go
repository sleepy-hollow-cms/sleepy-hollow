package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	prefix            string = "SLEEPY_HOLLOW"
	configFileKey     string = "config"
	configDefaultPath string = "/etc/sleepy-hollow/config.yaml"
)

var (
	Conf                *Config
	configDefaultValues map[string]interface{}
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
	Host     string `yaml:"host" default:"localhost"`
	Port     int    `yaml:"port" default:"27017"`
}

func init() {
	Conf = &Config{
		Log:     Log{},
		Server:  Server{},
		MongoDB: MongoDB{},
	}

	configDefaultValues = map[string]interface{}{
		"LOG.ENCODING":     "json",
		"LOG.OUTPUT":       "stdout",
		"LOG.LEVEL":        "debug",
		"SERVER.PORT":      3000,
		"MONGODB.USER":     "root",
		"MONGODB.PASSWORD": "password",
		"MONGODB.HOST":     "localhost",
		"MONGODB.PORT":     27017,
	}
}

func (c *Config) Load() error {
	viper.SetDefault(configFileKey, configDefaultPath)

	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(viper.GetString(configFileKey))

	// set default values
	for k, v := range configDefaultValues {
		viper.SetDefault(k, v)
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Print("Config file is not found")
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return err
	}

	return nil
}
