package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	Redis    RedisConfig    `mapstructure:"redis"`
	JWT      JWTConfig      `mapstructure:"jwt"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // debug / release
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"expiration"` // hours
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DBName, d.SSLMode,
	)
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 环境变量覆盖
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	// viper.Unmarshal 不触发 AutomaticEnv 对已有值的 key 查找，需要显式绑定
	_ = viper.BindEnv("database.host")
	_ = viper.BindEnv("database.port")
	_ = viper.BindEnv("database.user")
	_ = viper.BindEnv("database.password")
	_ = viper.BindEnv("database.dbname")
	_ = viper.BindEnv("database.sslmode")
	_ = viper.BindEnv("redis.host")
	_ = viper.BindEnv("redis.port")
	_ = viper.BindEnv("redis.password")
	_ = viper.BindEnv("redis.db")
	_ = viper.BindEnv("jwt.secret")
	_ = viper.BindEnv("jwt.expiration")
	_ = viper.BindEnv("server.port")
	_ = viper.BindEnv("server.mode")

	// 默认值
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.password", "postgres")
	viper.SetDefault("database.dbname", "student_admin")
	viper.SetDefault("database.sslmode", "disable")
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("jwt.secret", "change-me-in-production")
	viper.SetDefault("jwt.expiration", 24)

	// 如果有配置文件则读取，没有则用默认值+环境变量
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	// 环境变量直接覆盖，不依赖 Viper 的 AutomaticEnv/BindEnv 机制
	if v := os.Getenv("APP_SERVER_PORT"); v != "" {
		cfg.Server.Port = v
	}
	if v := os.Getenv("APP_SERVER_MODE"); v != "" {
		cfg.Server.Mode = v
	}
	if v := os.Getenv("APP_DATABASE_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("APP_DATABASE_PORT"); v != "" {
		cfg.Database.Port = v
	}
	if v := os.Getenv("APP_DATABASE_USER"); v != "" {
		cfg.Database.User = v
	}
	if v := os.Getenv("APP_DATABASE_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("APP_DATABASE_DBNAME"); v != "" {
		cfg.Database.DBName = v
	}
	if v := os.Getenv("APP_DATABASE_SSLMODE"); v != "" {
		cfg.Database.SSLMode = v
	}
	if v := os.Getenv("APP_REDIS_HOST"); v != "" {
		cfg.Redis.Host = v
	}
	if v := os.Getenv("APP_REDIS_PORT"); v != "" {
		cfg.Redis.Port = v
	}
	if v := os.Getenv("APP_REDIS_PASSWORD"); v != "" {
		cfg.Redis.Password = v
	}
	if v := os.Getenv("APP_REDIS_DB"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.Redis.DB = n
		}
	}
	if v := os.Getenv("APP_JWT_SECRET"); v != "" {
		cfg.JWT.Secret = v
	}
	if v := os.Getenv("APP_JWT_EXPIRATION"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.JWT.Expiration = n
		}
	}

	// 生产环境必须设置 JWT secret
	if cfg.Server.Mode == "release" && cfg.JWT.Secret == "change-me-in-production" {
		fmt.Fprintln(os.Stderr, "WARNING: JWT secret is not set in production!")
	}

	return &cfg, nil
}
