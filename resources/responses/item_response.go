package resources_response

import (
	"ols/models"
	"strconv"
)

type ListTaxResponse struct {
	Id   int    ` json:"id"`
	Name string ` json:"name"`
	Rate string `json:"rate"`
}

type ItemResponse struct {
	ID    int          `json:"id"`
	Name  string       `json:"name"`
	Taxes []models.Tax `json:"taxes,omitempty"`
}

func FormatItemResponse(item models.Item) ItemResponse {
	data := ItemResponse{
		ID:    item.Id,
		Name:  item.Name,
		Taxes: item.Taxes,
	}

	return data
}

type ItemSuccessResponse struct {
	IsSuccess bool `json:"is_success"`
}

func FormatItemSuccessResponse(IsSuccess bool) ItemSuccessResponse {
	data := ItemSuccessResponse{
		IsSuccess: IsSuccess,
	}

	return data
}

type ListItemResponse struct {
	Id    int                `json:"id"`
	Name  string             `json:"name"`
	Taxes []*ListTaxResponse `json:"taxes,omitempty"`
}

func FormatGetListItemResponse(items []models.Item) (data []*ListItemResponse) {
	var listItem []*ListItemResponse

	for _, detailItem := range items {
		var listTax []*ListTaxResponse

		for _, detailTax := range detailItem.Taxes {
			listTax = append(
				listTax,
				&ListTaxResponse{
					Id:   detailTax.Id,
					Name: detailTax.Name,
					Rate: strconv.Itoa(detailTax.Rate) + "%",
				},
			)
		}

		listItem = append(
			listItem,
			&ListItemResponse{
				Id:    detailItem.Id,
				Name:  detailItem.Name,
				Taxes: listTax,
			},
		)
	}

	return listItem
}
