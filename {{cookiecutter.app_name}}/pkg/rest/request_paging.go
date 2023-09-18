package rest

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.app_name}}/pkg/slices"
)

// RequestPaging used for pagination and sort on database
type RequestPaging struct {
	Page         int    `json:"page" validate:"gt=1,lte=1000000"`
	PageSize     int    `json:"page_size" validate:"gt=5,lte=100"`
	Sort         string `json:"sort" default:"geek"`
	sortSafelist []string
}

func GetRequestPaging(urlValues url.Values) RequestPaging {
	return RequestPaging{
		Page:     GetInt(urlValues, "page", 1),
		PageSize: GetInt(urlValues, "page_size", 20),
		Sort:     GetString(urlValues, "sort", ""),
	}
}

// SetSortSafeList do set sort safe list
func (f *RequestPaging) SetSortSafeList(sort []string) {
	f.sortSafelist = sort
}

func (f *RequestPaging) setDefault() {
	if f.Sort == "" {
		f.Sort = f.sortSafelist[0]
	}

	if f.Page == 0 {
		f.Page = 1
	}

	if f.PageSize == 0 {
		f.PageSize = 20
	}
}

// Validate do set default value for filter field and validate when
// user use the field
func (f *RequestPaging) Validate() error {
	f.setDefault()

	if f.Page < 1 || f.Page > 1000 {
		return errors.New("invalid page value")
	}

	if f.PageSize < 1 || f.PageSize > 100 {
		return errors.New("invalid page_size value")
	}

	if len(f.sortSafelist) != 0 {
		if !slices.In(f.Sort, f.sortSafelist) {
			return errors.New("invalid sort value")
		}
	}

	return nil
}

// Check that the client-provided Sort field matches one of the entries in our safelist
// and if it does, extract the column name from the Sort field by stripping the leading
// hyphen character (if one exists).
func (f *RequestPaging) sortColumn() string {
	for _, safeValue := range f.sortSafelist {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// return the sort direction ("ASC" or "DESC") depending on the prefix character of the
// Sort field.
func (f *RequestPaging) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

// SortColumnDirection return sort sql format
// ex: "id ASC"
func (f *RequestPaging) SortColumnDirection() string {
	return fmt.Sprintf("%s %s", f.sortColumn(), f.sortDirection())
}

// Limit return limit (size per page)
func (f *RequestPaging) Limit() int {
	return f.PageSize
}

// Offset return calculated offset
func (f *RequestPaging) Offset() int {
	return (f.Page - 1) * f.PageSize
}
