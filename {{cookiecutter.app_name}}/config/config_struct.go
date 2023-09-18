package config

import "time"

type App struct {
	Env      string
	Host     string
	Port     int
	Name     string
	Timezone string
	Secret   string
	Version  string
}

type Cookies struct {
	SSODomain    string
	SSOURL       string
	AccessToken  string // token
	RefreshToken string // refresh_token
}

type Postgres struct {
	Url             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

type JWT struct {
	Key         string
	StaticToken string
}

type Telemetry struct {
	URL      string
	Key      string
	Insecure bool
}
