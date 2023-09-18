package rest

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
	"github.com/labstack/echo/v4"
)

func router(dep *common.RestDependency, server *echo.Echo) {
	orderContainer := order.NewContainer(dep)

	v1Routes := server.Group("/api/v1")

	oRoute := v1Routes.Group("/orders")
	oRoute.GET("", orderContainer.Handler.Get)
}
