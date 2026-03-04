package viper

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	var config = viper.New()
	assert.NotNil(t, config)
}

func TestViperJSON(t *testing.T) {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.NoError(t, err)
	assert.Equal(t, "localhost", config.GetString("database.host"))
}

func TestViperYAML(t *testing.T) {
	var config = viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	assert.NoError(t, err)
	assert.Equal(t, "localhost", config.GetString("database.host"))
}

func TestViperENV(t *testing.T) {
	var config = viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	// make sure the env var is present for the test
	_ = os.Setenv("FROM_ENV", "hello")

	err := config.ReadInConfig()
	assert.NoError(t, err)
	assert.Equal(t, "localhost", config.GetString("DATABASE_HOST"))
	assert.Equal(t, "hello", config.GetString("FROM_ENV"))
}
