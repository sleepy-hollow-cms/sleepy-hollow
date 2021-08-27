package config_test

import (
	"os"
	"testing"

	"github.com/sleepy-hollow-cms/sleepy-hollow/content-management-api/util/config"

	"github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestConfig(t *testing.T) {

	//TODO: Delete .conf/ directory after implement reading config from path specified by user
	// Now, this test is depending on .conf/config.yaml
	t.Run("外部ファイルと環境変数からコンフィグをロードできる", func(t *testing.T) {
		expected := &config.Config{
			Log: config.Log{
				Encoding: "json",
				Output:   "stdout",
				Level:    "info",
			},
			Server: config.Server{
				Port: 3000,
			},
			MongoDB: config.MongoDB{
				User:     "root",
				Password: "password2",
				Host:     "mongo",
				Port:     27017,
			},
		}

		// create temp config file
		c := &config.Config{
			Log: config.Log{
				Encoding: "json",
				Output:   "stdout",
				Level:    "debug",
			},
			Server: config.Server{
				Port: 3000,
			},
			MongoDB: config.MongoDB{
				User:     "root",
				Password: "password",
				Host:     "mongo",
				Port:     27017,
			},
		}
		b, _ := yaml.Marshal(c)
		filet.File(t, "config.yaml", string(b))
		os.Setenv("SLEEPY_HOLLOW_CONFIG", "./config.yaml")

		// set environment variables to override
		os.Setenv("SLEEPY_HOLLOW_LOG_LEVEL", "info")
		os.Setenv("SLEEPY_HOLLOW_MONGODB_PASSWORD", "password2")

		t.Cleanup(func() {
			os.Unsetenv("SLEEPY_HOLLOW_LOG_LEVEL")
			os.Unsetenv("SLEEPY_HOLLOW_MONGODB_PASSWORD")
			os.Unsetenv("SLEEPY_HOLLOW_CONFIG")
			filet.CleanUp(t)
		})

		config.Conf.Load()
		actual := config.Conf

		assert.Equal(t, expected, actual)
	})

	t.Run("外部ファイルの読み込みに失敗したとき環境変数からコンフィグをロードできる", func(t *testing.T) {
		expected := &config.Config{
			Log: config.Log{
				Encoding: "json",
				Output:   "stdout",
				Level:    "info",
			},
			Server: config.Server{
				Port: 3000,
			},
			MongoDB: config.MongoDB{
				User:     "root",
				Password: "password2",
				Host:     "mongo",
				Port:     27017,
			},
		}

		os.Setenv("SLEEPY_HOLLOW_CONFIG", "./config.yaml")

		// set environment variables to override
		os.Setenv("SLEEPY_HOLLOW_LOG_LEVEL", "info")
		os.Setenv("SLEEPY_HOLLOW_MONGODB_PASSWORD", "password2")

		t.Cleanup(func() {
			os.Unsetenv("SLEEPY_HOLLOW_LOG_LEVEL")
			os.Unsetenv("SLEEPY_HOLLOW_MONGODB_PASSWORD")
			os.Unsetenv("SLEEPY_HOLLOW_CONFIG")
		})

		config.Conf.Load()
		actual := config.Conf

		assert.Equal(t, expected, actual)
	})
}
