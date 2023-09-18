package order

import (
	"net/http"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/dto"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/service"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/rest"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Handler struct {
	dep     *common.RestDependency
	service service.Service
}

func NewHandler(dep *common.RestDependency, service service.Service) *Handler {
	return &Handler{
		dep:     dep,
		service: service,
	}
}

func (h *Handler) Get(c echo.Context) error {

	urlValues := c.QueryParams()
	requestPaging := rest.GetRequestPaging(urlValues)
	filter := dto.GetFilterFromURL(urlValues)

	orders, responsePaging, err := h.service.Find(c.Request().Context(), filter, requestPaging)

	if err != nil {
		h.dep.Logger.Error("Order Get", zap.String("err", err.Error()))

		statusCode, errMsg := rest.ParseErrHttp(err)

		return c.JSON(statusCode, rest.Envelope{
			Success: false,
			Message: errMsg,
		})
	}

	return c.JSON(http.StatusOK, rest.Envelope{
		Success: true,
		Message: "mengambil data pesanan berhasil",
		Data:    orders,
		Paging:  responsePaging,
	})
}
