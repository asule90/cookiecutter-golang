package service

import (
	"context"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/dto"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/entities"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/repo"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/errr"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/rest"
)

type serviceImpl struct {
	dep  *common.RestDependency
	repo repo.Repo
}

func NewServiceImpl(dep *common.RestDependency, repo repo.Repo) *serviceImpl {
	return &serviceImpl{
		dep:  dep,
		repo: repo,
	}
}

func (s *serviceImpl) Find(ctx context.Context, filter dto.Filter, page rest.RequestPaging) ([]entities.Order, *rest.ResponsePaging, error) {

	// load all users
	orders, paging, err := s.repo.Find(ctx, filter, page)
	if err != nil {
		return nil, paging, errr.WrapF(400, "find users data:: %w", errr.ErrDBSortFilter)
	}

	return orders, paging, nil
}
