package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MySQL      MySQL      `env-prefix:"MYSQL_"`
	HttpRouter HttpRouter `env-prefix:"HTTP_ROUTER_"`
}

type MySQL struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"3306"`
	DBName   string `env:"DBNAME" envDefault:"mysql"`
	User     string `env:"USER" envDefault:"root"`
	Password string `env:"PASSWORD" envDefault:""`
	MaxIdleConns    int    `env:"MAX_IDLE_CONNS" envDefault:"10"`
	MaxOpenConns    int    `env:"MAX_OPEN_CONNS" envDefault:"100"`
	ConnMaxLifetime int    `env:"CONN_MAX_LIFETIME" envDefault:"10"`
}

type HttpRouter struct {
	Port int `env:"PORT" envDefault:"8080"`
}

var once sync.Once

func MustLoad() Config {
	cfg := Config{}

	once.Do(func() {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			log.Fatalf("Error loading environment variables: %v", err)
		}
	})

	return cfg
}
