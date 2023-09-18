package rest

import "math"

type ResponsePaging struct {
	CurrentPage  int   `json:"current_page" example:"1"`
	PageSize     int   `json:"page_size" example:"50"`
	FirstPage    int   `json:"first_page" example:"1"`
	LastPage     int   `json:"total_page" example:"1"`
	TotalRecords int64 `json:"total_count" example:"1"`
}

func CalculatePaging(totalRecords int64, page, pageSize int) ResponsePaging {
	if totalRecords == 0 {
		return ResponsePaging{
			CurrentPage: 1,
			FirstPage:   1,
			LastPage:    1,
		}
	}

	return ResponsePaging{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}
