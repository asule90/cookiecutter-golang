package repo

import (
	"context"
	"time"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/dto"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/entities"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/internal/order/repo/models"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/common"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/errr"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/rest"
)

type postgreRepo struct {
	dep *common.RestDependency
}

func NewPostgreRepo(dep *common.RestDependency) *postgreRepo {
	return &postgreRepo{
		dep: dep,
	}
}

func (pr postgreRepo) Find(ctx context.Context, filter dto.Filter, page rest.RequestPaging) ([]entities.Order, *rest.ResponsePaging, error) {
	page.SetSortSafeList([]string{"name", "-name", "channel", "-channel", "ship_time", "-ship_time", "status", "-status", "payment_status", "-payment_status", "payment_method", "-payment_method"})
	if err := page.Validate(); err != nil {
		return nil, nil, errr.WrapF(400, errr.ErrDBSortFilter.Error(), err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var db_orders []models.Orders
	var totalRecord int64
	var query = pr.dep.PostgreSQL.WithContext(ctx).Model(&db_orders)

	if filter.Search != "" {
		query.Where(
			query.Where("name ILIKE ?", "%"+filter.Search+"%").
				Or("channel ILIKE ?", "%"+filter.Search+"%").
				Or("ship_time ILIKE ?", "%"+filter.Search+"%").
				Or("payment_method ILIKE ?", "%"+filter.Search+"%").
				Or("payment_status ILIKE ?", "%"+filter.Search+"%").
				Or("note ILIKE ?", "%"+filter.Search+"%").
				Or("promo ILIKE ?", "%"+filter.Search+"%").
				Or("created_by ILIKE ?", "%"+filter.Search+"%").
				Or("updated_by ILIKE ?", "%"+filter.Search+"%"),
		)
	}

	if filter.Name != "" {
		query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Channel != "" {
		query.Where("channel = ?", filter.Channel)
	}
	if filter.ShipTime != "" {
		query.Where("ship_time = ?", filter.ShipTime)
	}
	if filter.Status != "" {
		query.Where("status = ?", filter.Status)
	}
	if filter.PaymentStatus != "" {
		query.Where("payment_status = ?", filter.PaymentStatus)
	}
	if filter.PaymentMethod != "" {
		query.Where("payment_status = ?", filter.PaymentMethod)
	}

	result := query.
		Count(&totalRecord). // count must be before limit
		Order(page.SortColumnDirection()).
		Limit(page.Limit()).
		Offset(page.Offset()).
		Find(&db_orders)

	if result.Error != nil {
		return nil, nil, errr.Wrap(422, errr.ErrDBSortFilter)
	}

	paging := rest.CalculatePaging(totalRecord, page.Page, page.PageSize)

	// Transform model to entity
	results := make([]entities.Order, 0, len(db_orders))
	for _, model := range db_orders {
		results = append(results, model.ToEntity())
	}

	return results, &paging, nil
}
