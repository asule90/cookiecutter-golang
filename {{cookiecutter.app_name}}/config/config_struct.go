package config

import "time"

type App struct {
	Env     string
	Host    string
	Port    int
	Name    string
	Secret  string
	Version string
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