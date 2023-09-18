package dto

import (
	"net/url"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/rest"
)

type Filter struct {
	Name          string
	Search        string
	Channel       string
	ShipTime      string
	Status        string
	PaymentMethod string
	PaymentStatus string
}

func GetFilterFromURL(urlValues url.Values) Filter {

	return Filter{
		Search:        rest.GetString(urlValues, "search", ""),
		Name:          rest.GetString(urlValues, "name", ""),
		Channel:       rest.GetString(urlValues, "is_customer", ""),
		ShipTime:      rest.GetString(urlValues, "products", ""),
		Status:        rest.GetString(urlValues, "point_name", ""),
		PaymentMethod: rest.GetString(urlValues, "point_uuid", ""),
		PaymentStatus: rest.GetString(urlValues, "created_by", ""),
	}
}
