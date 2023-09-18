package rest

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/db"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/logger"
	"go.uber.org/zap"
)

func getCommonDependency() *common.RestDependency {
	// load config
	cfg := config.LoadConfig(".env")

	var err error

	// logger
	zapLogger := logger.InitLogger(cfg)

	// database
	dbpool, err := db.NewPostgreSQL(cfg.Postgres, cfg.App.Env)
	if err != nil {
		zapLogger.Fatal("Failed to start, error connect to DB Postgre",
			zap.String("url", err.Error()),
		)
		return nil
	}

	comm := common.RestDependency{
		Config:     cfg,
		PostgreSQL: dbpool,
		Logger:     *zapLogger,
	}

	return &comm
}
