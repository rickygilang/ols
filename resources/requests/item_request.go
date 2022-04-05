package resources_requests

type CreateItemRequest struct {
	ItemName string `json:"name" binding:"required"`
}

type ItemUpdateRequest struct {
	ItemId   int    `json:"id" binding:"required"`
	ItemName string `json:"name" binding:"required"`
}

type ItemDeleteRequest struct {
	ItemId int `json:"id" binding:"required"`
}
