// Why: config loading gives the app one typed place for environment-driven settings.
// What to do: add new environment variables here with safe local defaults.
package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppEnv         string `mapstructure:"APP_ENV"`
	AppPort        string `mapstructure:"APP_PORT"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBName         string `mapstructure:"DB_NAME"`
	DBSSLMode      string `mapstructure:"DB_SSLMODE"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JWTExpireHours int    `mapstructure:"JWT_EXPIRE_HOURS"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	viper.SetDefault("APP_NAME", "example-gin")
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_NAME", "example_gin")
	viper.SetDefault("DB_SSLMODE", "disable")
	viper.SetDefault("JWT_SECRET", "change-me-in-production")
	viper.SetDefault("JWT_EXPIRE_HOURS", 24)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &cfg, nil
}
