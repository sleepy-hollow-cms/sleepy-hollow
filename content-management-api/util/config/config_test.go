package config_test

import (
	"content-management-api/util/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

		// set environment variables to override
		os.Setenv("SLEEPY_HOLLOW_LOG_LEVEL", "info")
		os.Setenv("SLEEPY_HOLLOW_MONGODB_PASSWORD", "password2")
		t.Cleanup(func() {
			os.Unsetenv("SLEEPY_HOLLOW_LOG_LEVEL")
			os.Unsetenv("SLEEPY_HOLLOW_MONGODB_PASSWORD")
		})

		config.Conf.Load()
		actual := config.Conf

		assert.Equal(t, expected, actual)
	})
}
