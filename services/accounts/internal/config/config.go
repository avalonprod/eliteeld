package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	defaultHTTPPort           = "8000"
	defaultHTTPWriteTimeout   = 10 * time.Second
	defaultHTTPReadTimeout    = 10 * time.Second
	defaultHTTPMaxHeaderBytes = 1
)

type Config struct {
	HTTP     HTTPConfig
	Mongo    MongoConfig
	Password PasswordConfig
	Emails   EmailsConfig
}

type (
	EmailsConfig struct {
		Url string `mapstructure:"url"`
	}
	HTTPConfig struct {
		Host           string        `mapstructure:"host"`
		Port           string        `mapstructure:"port"`
		ReadTimeout    time.Duration `mapstructure:"readTimeout"`
		WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderBytes int           `mapstructure:"maxHeaderBytes"`
	}
	MongoConfig struct {
		URL      string
		Database string
		Username string
		Password string
	}
	PasswordConfig struct {
		PasswordSalt string
	}
)

func Init(configDir string) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := parseConfigFile(configDir); err != nil {
		return nil, err
	}
	SetDefault()
	setFromEnv(&cfg)
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("emails", &cfg.Emails); err != nil {
		return err
	}
	return nil
}

func parseConfigFile(configsDir string) error {
	viper.AddConfigPath(configsDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func setFromEnv(cfg *Config) {
	cfg.Mongo.URL = os.Getenv("MONGODB_URL")
	cfg.Mongo.Username = os.Getenv("MONGODB_USERNAME")
	cfg.Mongo.Password = os.Getenv("MONGODB_PASSWORD")
	cfg.Mongo.Database = os.Getenv("MONGODB_DATABASE")
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
	cfg.Password.PasswordSalt = os.Getenv("PASSWORD_SALT")
}

func SetDefault() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderBytes", defaultHTTPMaxHeaderBytes)
	viper.SetDefault("http.writeTimeout", defaultHTTPWriteTimeout)
	viper.SetDefault("http.readTimeout", defaultHTTPReadTimeout)

}
