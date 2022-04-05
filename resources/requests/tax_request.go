package resources_requests

type CreateTaxRequest struct {
	ItemId  int    `json:"item_id" binding:"required"`
	TaxName string `json:"name" binding:"required"`
	TaxRate int    `json:"rate" binding:"required"`
}

type TaxUpdateRequest struct {
	TaxId   int    `json:"id" binding:"required"`
	ItemId  int    `json:"item_id" binding:"required"`
	TaxName string `json:"name" binding:"required"`
	TaxRate int    `json:"rate" binding:"required"`
}

type TaxDeleteRequest struct {
	TaxId  int `json:"id" binding:"required"`
	ItemId int `json:"item_id"`
}
