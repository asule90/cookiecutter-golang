package common

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// RestDependency struct for all required objects
// commonly used by rest api applications
type RestDependency struct {
	Config     *config.Provider
	PostgreSQL *gorm.DB
	Logger     zap.Logger
}
