package logger

import (
	"time"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// import "github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/version"

func getEncoding(cfg *config.Provider) zapcore.EncoderConfig {
	var zapEncoderConf zapcore.EncoderConfig
	if cfg.App.Env == "production" {
		zapEncoderConf = zap.NewProductionEncoderConfig()
	} else {
		zapEncoderConf = zap.NewDevelopmentEncoderConfig()
	}

	// zapEncoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	zapEncoderConf.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		loc, _ := time.LoadLocation(cfg.App.Timezone)
		enc.AppendString(t.In(loc).Format("2006-01-02T15:04:05.000Z"))
	}

	return zapEncoderConf
}

func getConfig(cfg *config.Provider) zap.Config {

	config := zap.Config{
		Sampling:      nil,
		Encoding:      "json",
		EncoderConfig: getEncoding(cfg),
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}

	if cfg.App.Env == "production" {
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		config.Development = false
		config.DisableCaller = true
		config.DisableStacktrace = true
	} else {
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		config.Development = true
		config.DisableCaller = false
		config.DisableStacktrace = false
	}

	return config
}

func InitLogger(cfg *config.Provider) *zap.Logger {
	logg, _ := getConfig(cfg).Build()

	return logg

}
