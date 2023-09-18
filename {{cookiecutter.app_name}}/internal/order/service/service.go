package service

import (
	"context"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/dto"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/entities"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/rest"
)

type Service interface {
	Find(ctx context.Context, filter dto.Filter, page rest.RequestPaging) ([]entities.Order, *rest.ResponsePaging, error)
}
