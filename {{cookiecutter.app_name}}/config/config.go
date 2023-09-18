package config

import (
	"fmt"
	"time"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/env"
	"github.com/joho/godotenv"
)

type Provider struct {
	App      *App
	Postgres *Postgres
	JWT      *JWT
}

func LoadConfig(file string) *Provider {
	err := godotenv.Overload(file)
	if err != nil {
		fmt.Println(".env not found (but expected on prod)")
	}

	app := &App{
		Env:      env.Get("APP_ENV", "local"),
		Host:     env.Get("APP_HOST", "localhost"),
		Port:     env.Get("APP_PORT", 8080),
		Name:     env.Get("APP_NAME", "golang project"),
		Timezone: env.Get("APP_TIMEZONE", "Asia/Jakarta"),
		Secret:   env.Get("APP_SECRET", ""),
		Version:  env.Get("APP_VERSION", ""),
	}

	postgres := &Postgres{
		Url:             env.Get("POSTGRES_URL", "postgres://postgres:postgres@localhost:5432/customer?sslmode=disable"),
		MaxIdleConns:    env.Get("POSTGRES_IDLE_CONNS", 10),
		MaxOpenConns:    env.Get("POSTGRES_MAX_OPEN_CONNS", 50),
		ConnMaxLifetime: env.GetDuration("POSTGRES_CONN_MAX_LIFETIME", 1*time.Hour),
		ConnMaxIdleTime: env.GetDuration("POSTGRES_CONN_MAX_IDLE_TIME", 20*time.Minute),
	}

	jwt := &JWT{
		Key:         env.Get("JWT_SECRET", ""),
		StaticToken: env.Get("STATIC_TOKEN", ""),
	}

	cfg := &Provider{
		App:      app,
		Postgres: postgres,
		JWT:      jwt,
	}

	return cfg
}
