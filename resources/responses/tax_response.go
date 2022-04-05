package resources_response

import (
	"ols/models"
	"strconv"
)

type CreateTaxResponse struct {
	ID     int    `json:"id"`
	ItemId int    `json:"item_id"`
	Name   string `json:"name"`
	Rate   string `json:"rate"`
}

func FormatCreateTaxResponse(tax models.Tax) CreateTaxResponse {
	data := CreateTaxResponse{
		ID:     tax.Id,
		ItemId: tax.ItemId,
		Name:   tax.Name,
		Rate:   strconv.Itoa(tax.Rate) + "%",
	}

	return data
}

type TaxSuccessResponse struct {
	IsSuccess bool `json:"is_success"`
}

func FormatTaxSuccessResponse(IsSuccess bool) TaxSuccessResponse {
	data := TaxSuccessResponse{
		IsSuccess: IsSuccess,
	}

	return data
}
