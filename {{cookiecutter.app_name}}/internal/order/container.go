package order

import (
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/repo"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/service"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
)

type Container struct {
	Repo    repo.Repo
	Service service.Service
	Handler *Handler
}

func NewContainer(dep *common.RestDependency) *Container {
	repo := repo.NewPostgreRepo(dep)
	service := service.NewServiceImpl(dep, repo)
	handler := NewHandler(dep, service)

	return &Container{
		Repo:    repo,
		Service: service,
		Handler: handler,
	}
}
