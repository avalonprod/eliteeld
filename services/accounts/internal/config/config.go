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
	HTTP   HTTPConfig
	Mongo  MongoConfig
	Auth   AuthConfig
	Emails EmailsConfig
}

type (
	EmailsConfig struct {
		ApiUrlCompanyRegistration string `mapstructure:"apiUrlCompanyRegistration"`
		ApiUrlDriverRegistration  string `mapstructure:"apiUrlDriverRegistration"`
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
	AuthConfig struct {
		JWT          JWTConfig
		PasswordSalt string
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
		SigningKey      string
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

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("emailsService", &cfg.Emails); err != nil {
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
	cfg.Auth.PasswordSalt = os.Getenv("PASSWORD_SALT")
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

func SetDefault() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderBytes", defaultHTTPMaxHeaderBytes)
	viper.SetDefault("http.writeTimeout", defaultHTTPWriteTimeout)
	viper.SetDefault("http.readTimeout", defaultHTTPReadTimeout)

}
