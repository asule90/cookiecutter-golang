package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

// NewPostgreSQL return a client connection handle to a Postgre server.
func NewPostgreSQL(option *config.Postgres, env string) (*gorm.DB, error) {
	logger := zapgorm2.New(zap.L())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	pgxDB, err := sql.Open("pgx", option.Url)
	if err != nil {
		return nil, err
	}
	pgxDB.SetMaxIdleConns(option.MaxIdleConns)
	pgxDB.SetMaxOpenConns(option.MaxOpenConns)
	pgxDB.SetConnMaxLifetime(option.ConnMaxLifetime)
	pgxDB.SetConnMaxIdleTime(option.ConnMaxIdleTime)

	db, err := gorm.Open(
		postgres.New(postgres.Config{Conn: pgxDB}),
		&gorm.Config{
			SkipDefaultTransaction: true,
		})
	if err != nil {
		log.Panic(err) // its okey to panic, if database missing apps should not be run
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return db, err
	}

	return db, nil
}
