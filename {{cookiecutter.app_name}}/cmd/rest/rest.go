/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package rest

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	{% if cookiecutter.add_rest_server == "chi" -%}"github.com/go-chi/chi/v5"{%- endif %}
	{% if cookiecutter.add_rest_server == "echo" -%}
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	{%- endif %}
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/xerrors"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/xrender"
)

var (
	ENV     string
	Address string
)

// RestCmd represents the rest command.
var RestCmd = &cobra.Command{
	Use:   "rest",
	Short: "rest api for supply service",
	Run: func(cmd *cobra.Command, args []string) {
		logger, err := zap.NewProduction()
		if err != nil {
			panic(xerrors.WrapErrorf(err, xerrors.ErrorCodeUnknown, "zap.NewProduction"))
		}

		{% if cookiecutter.add_rest_server == "chi" -%}
		route := chi.NewRouter()

		// Add your service here
		route.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("Content-Type", "text/plain")
			xrender.Response(rw, "ok", http.StatusOK)
		})

		srv := &http.Server{
			Handler:           route,
			Addr:              fmt.Sprintf(":%s", Address),
			ReadTimeout:       1 * time.Second,
			ReadHeaderTimeout: 1 * time.Second,
			WriteTimeout:      1 * time.Second,
			IdleTimeout:       1 * time.Second,
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			logger.Info("Listening on port :", zap.String("address", Address))
			if err := srv.ListenAndServe(); err != nil {
				if errors.Is(err, http.ErrServerClosed) {
					log.Fatal(err)
				}
			}
		}()

		sig := <-c
		logger.Info("shutting down: %+v", zap.Any("signal", sig))

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatal("error shutingdown server", zap.Error(err))
		}
		{%- endif %}
		{% if cookiecutter.add_rest_server == "echo" -%}
		// Setup
		e := echo.New()
		e.Logger.SetLevel(log.INFO)
		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "OK")
		})

		// Start server
		go func() {
			if err := e.Start(fmt.Sprintf(":%s", Address)); err != nil && err != http.ErrServerClosed {
				e.Logger.Fatal("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds. 
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
		{%- endif %}
	},
}
