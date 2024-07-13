package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort        string
	AgentService    *AgentServiceConfig
	SettingsService *SettingsServiceConfig
	VisitorService  *VisitorServiceConfig
}

type AgentServiceConfig struct {
	Host string
	Port string
}

type SettingsServiceConfig struct {
	Host string
	Port string
}

type VisitorServiceConfig struct {
	Host string
	Port string
}

var Conf *Config

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	viper.BindEnv("HTTP_PORT")
	viper.BindEnv("AGENTS_SERVICE_HOST")
	viper.BindEnv("AGENTS_SERVICE_PORT")

	viper.BindEnv("SETTINGS_SERVICE_HOST")
	viper.BindEnv("SETTINGS_SERVICE_PORT")

	Conf = &Config{
		HttpPort: viper.GetString("HTTP_PORT"),
		AgentService: &AgentServiceConfig{
			Host: viper.GetString("AGENTS_SERVICE_HOST"),
			Port: viper.GetString("AGENTS_SERVICE_PORT"),
		},
		SettingsService: &SettingsServiceConfig{
			Host: viper.GetString("SETTINGS_SERVICE_HOST"),
			Port: viper.GetString("SETTINGS_SERVICE_PORT"),
		},
		VisitorService: &VisitorServiceConfig{
			Host: viper.GetString("VISITOR_SERVICE_HOST"),
			Port: viper.GetString("VISITOR_SERVICE_PORT"),
		},
	}
}
