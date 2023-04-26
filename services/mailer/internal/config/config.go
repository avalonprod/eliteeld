package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	defaultHTTPPort           = "5001"
	defaultHTTPWriteTimeout   = 10 * time.Second
	defaultHTTPReadTimeout    = 10 * time.Second
	defaultHTTPMaxHeaderBytes = 1
)

type (
	Config struct {
		HTTP  HTTPConfig
		SMTP  SMTPConfig
		Email EmailConfig
	}

	HTTPConfig struct {
		Host           string        `mapstructure:"host"`
		Port           string        `mapstructure:"port"`
		ReadTimeout    time.Duration `mapstructure:"readTimeout"`
		WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderBytes int           `mapstructure:"maxHeaderBytes"`
	}

	EmailSubjects struct {
		CompanyRegistration string `mapstructure:"companyRegistration"`
	}

	EmailConfig struct {
		Templates EmailTemplates
		Subjects  EmailSubjects
	}

	EmailTemplates struct {
		CompanyRegistrationTemplate string `mapstructure:"companyRegistration"`
	}

	SMTPConfig struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		From     string `mapstructure:"from"`
		Password string
	}
)

func Init(configsDir string) (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	SetDefault()
	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}
	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}
	setFromEnv(&cfg)
	return &cfg, nil
}

func parseConfigFile(ConfigsDir string) error {
	viper.AddConfigPath(ConfigsDir)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("smtp", &cfg.SMTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("email.templates", &cfg.Email.Templates); err != nil {
		return err
	}
	if err := viper.UnmarshalKey("email.subjects", &cfg.Email.Subjects); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	cfg.SMTP.Password = os.Getenv("SMTP_PASSWORD")
}

func SetDefault() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderBytes", defaultHTTPMaxHeaderBytes)
	viper.SetDefault("http.writeTimeout", defaultHTTPWriteTimeout)
	viper.SetDefault("http.readTimeout", defaultHTTPReadTimeout)
}
