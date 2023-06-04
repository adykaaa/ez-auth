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
func Load(path string) (ac AppConfig, goc GoogleConfig, fc FacebookConfig, ghc GithubConfig, err error) {

	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("couldn't read in the config file. %v", err)
	}

	var appConfig AppConfig
	if err := viper.Unmarshal(&appConfig); err != nil {
		log.Fatalf("unable to unmarshal values to appConfig, %v", err)
	}

	var googleConfig GoogleConfig
	if err := viper.Unmarshal(&googleConfig); err != nil {
		log.Fatalf("unable to unmarshal values to googleConfig, %v", err)
	}

	var facebookConfig FacebookConfig
	if err := viper.Unmarshal(&facebookConfig); err != nil {
		log.Fatalf("unable to unmarshal values to facebookConfig, %v", err)
	}

	var githubConfig GithubConfig
	if err := viper.Unmarshal(&githubConfig); err != nil {
		log.Fatalf("unable to unmarshal values to githubConfig, %v", err)
	}

	viper.WatchConfig()

	return appConfig, googleConfig, facebookConfig, githubConfig, err
}
