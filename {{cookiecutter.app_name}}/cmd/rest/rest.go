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
	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	{%- endif %}
	"github.com/spf13/cobra"
)

// RestCmd represents the rest command.
var RestCmd = &cobra.Command{
	Use:   "rest",
	Short: "rest api",
	Run: func(cmd *cobra.Command, args []string) {

		dep := getCommonDependency()
		defer dep.Logger.Sync()
		
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
			dep.Logger.Info("Listening on port :", zap.String("address", Address))
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
		// Set a default timeout for all requests to this Echo instance.
		e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: 10 * time.Second,
		}))

		e.Use(echozap.ZapLogger(&dep.Logger))

		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "OK")
		})

		// wiring dependency injection

		router(dep, e)

		// Start server
		go func() {
			if err := e.Start(fmt.Sprintf(":%d", dep.Config.App.Port)); err != nil && err != http.ErrServerClosed {
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
