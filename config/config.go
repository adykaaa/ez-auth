package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	LogLevel          string `mapstructure:"LOG_LEVEL"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	DBConnString      string `mapstructure:"DB_CONN_STRING"`
	RedisConnString   string `mapstructure:"REDIS_CONN_STRING"`

	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

type GoogleConfig struct {
	RedirectURL  string `mapstructure:"GOOGLE_REDIRECT_URL"`
	ClientID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	ClientSecret string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	Scopes       string `mapstructure:"GOOGLE_SCOPES"`
}

type FacebookConfig struct {
	RedirectURL  string `mapstructure:"FACEBOOK_REDIRECT_URL"`
	ClientID     string `mapstructure:"FACEBOOK_CLIENT_ID"`
	ClientSecret string `mapstructure:"FACEBOOK_CLIENT_SECRET"`
	Scopes       string `mapstructure:"FACEBOOK_SCOPES"`
}

type GithubConfig struct {
	RedirectURL  string `mapstructure:"GITHUB_REDIRECT_URL"`
	ClientID     string `mapstructure:"GITHUB_CLIENT_ID"`
	ClientSecret string `mapstructure:"GITHUB_CLIENT_SECRET"`
	Scopes       string `mapstructure:"GITHUB_SCOPES"`
}

// Load loads the environment variables into the config struct
func Load(path string) (config AppConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper couldn't read in the config file. %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper could not unmarshal the configuration. %v", err)
	}
	return
}
