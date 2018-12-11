package service

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Config holds the configuration values for the backend.
type Config struct {
	DBhost            string
	DBport            int64
	DBuser            string
	DBpassword        string
	DBname            string
	DBmaxopenconns    int
	DBconnmaxlifetime time.Duration
}

// NewConfig loads the config file into the Config struct.
func NewConfig() *Config {
	config := viper.New()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)
	config.AutomaticEnv()

	config.SetDefault("db.host", "localhost")
	config.SetDefault("db.port", 5432)
	config.SetDefault("db.maxopenconns", 15)
	config.SetDefault("db.connmaxlifetime", 15*time.Minute)
	config.SetDefault("db.port", 5432)

	return &Config{
		DBhost:            config.GetString("db.host"),
		DBport:            config.GetInt64("db.port"),
		DBuser:            config.GetString("db.user"),
		DBpassword:        config.GetString("db.password"),
		DBname:            config.GetString("db.name"),
		DBmaxopenconns:    config.GetInt("db.maxopenconns"),
		DBconnmaxlifetime: config.GetDuration("db.connmaxlifetime"),
	}
}
