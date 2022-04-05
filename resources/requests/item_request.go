package resources_requests

import "ols/models"

type CreateItemRequest struct {
	ItemName string       `json:"name" binding:"required"`
	Taxes    []models.Tax `json:"taxes,omitempty"`
}

type ItemUpdateRequest struct {
	ItemId   int          `json:"id" binding:"required"`
	ItemName string       `json:"name" binding:"required"`
	Taxes    []models.Tax `json:"taxes,omitempty"`
}

type ItemDeleteRequest struct {
	ItemId int `json:"id" binding:"required"`
}
