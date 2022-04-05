package services

import (
	"ols/helpers"
	"ols/models"
	res "ols/resources"
	req "ols/resources/requests"

	"gorm.io/gorm"

	"net/http"
)

type ItemService struct {
	Item models.Item
}

func (ts *ItemService) db() *gorm.DB {
	return models.GetDB()
}

func (ts *ItemService) CreateItem(input req.CreateItemRequest) (item *models.Item, responseError *helpers.ResponseError) {

	// begin a transaction
	tx := ts.db().Begin()

	item = &ts.Item

	item.Name = input.ItemName

	result := tx.Create(&item)
	if result.Error != nil {
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail insert to database",
			ErrorDetail:      result.Error,
		}
		return nil, responseError
	}

	tx.Commit()
	response := res.Message(false, http.StatusOK, nil)
	response["data"] = item
	return item, nil
}

func (ts *ItemService) UpdateItem(input req.ItemUpdateRequest) (item *models.Item, responseError *helpers.ResponseError) {
	// get Item data where id = input.ID
	result := ts.db().Where("id = ?", input.ItemId).Find(&item)
	if result.Error != nil { // check error is empty or not
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail get data from database",
			ErrorDetail:      result.Error,
		}
		return nil, responseError
	}

	if result.RowsAffected < 1 { // handle record not found
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusBadRequest,
			ErrorMessage:     "Data not found",
			ErrorDetail:      nil,
		}
		return nil, responseError
	}

	tx := ts.db().Begin() // begin a transaction

	// update data student in table Item
	result = tx.Model(&models.Item{}).
		Where("id = ?", input.ItemId).
		Updates(map[string]interface{}{
			"name": input.ItemName,
		})
	if result.Error != nil {
		tx.Rollback()
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail update database",
			ErrorDetail:      result.Error,
		}
		return nil, responseError
	}

	tx.Commit()

	result = ts.db().Where("id = ?", input.ItemId).Find(&item) // get new data Item
	return item, nil
}

func (ts *ItemService) DeleteItem(input req.ItemDeleteRequest) (responseError *helpers.ResponseError) {
	var item *models.Item

	// get Item data where id = input.ID
	findResult := ts.db().Where("id = ?", input.ItemId).Find(&item)
	if findResult.Error != nil { // check error is empty or not
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail get data from database",
			ErrorDetail:      findResult.Error,
		}
		return responseError
	}

	if findResult.RowsAffected < 1 { // handle record not found
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusBadRequest,
			ErrorMessage:     "Data not found",
			ErrorDetail:      nil,
		}
		return responseError
	}

	// delete Item data where id = input.ItemId
	// result := ts.db().Where("id = ?", input.ItemId).Delete(&item)
	result := ts.db().Select("Taxes").Where("id = ?", input.ItemId).Delete(&item)
	if result.Error != nil { // check error is empty or not
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail get data from database",
			ErrorDetail:      result.Error,
		}
		return responseError
	}

	return nil
}

func (ts *ItemService) GetListItem() (items []models.Item, responseError *helpers.ResponseError) {
	result := ts.db().
		Preload("Taxes").
		Find(&items)
	if result.Error != nil { // check error is empty or not
		responseError = &helpers.ResponseError{
			HttpResponseCode: http.StatusInternalServerError,
			ErrorMessage:     "Fail get data from database",
			ErrorDetail:      result.Error,
		}
		return nil, responseError
	}

	return items, nil
}
