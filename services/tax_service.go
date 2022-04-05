package services

import (
	"ols/helpers"
	"ols/models"
	res "ols/resources"
	req "ols/resources/requests"

	"gorm.io/gorm"

	"net/http"
)

type TaxService struct {
	Tax models.Tax
}

func (ts *TaxService) db() *gorm.DB {
	return models.GetDB()
}

func (ts *TaxService) CreateTax(input req.CreateTaxRequest) (tax *models.Tax, responseError *helpers.ResponseError) {
	// get item data where id = input.ID
	var item *models.Item
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

	// begin a transaction
	tx := ts.db().Begin()

	tax = &ts.Tax

	tax.ItemId = input.ItemId
	tax.Name = input.TaxName
	tax.Rate = input.TaxRate

	result = tx.Create(&tax)
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
	response["data"] = tax
	return tax, nil
}

func (ts *TaxService) UpdateTax(input req.TaxUpdateRequest) (tax *models.Tax, responseError *helpers.ResponseError) {
	// get tax data where id = input.ID and itema_id = input.ItemId
	result := ts.db().
		Where("id = ?", input.TaxId).
		Where("item_id = ?", input.ItemId).
		Find(&tax)
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

	// update data tax in table tax
	result = tx.Model(&models.Tax{}).
		Where("id = ?", input.TaxId).
		Where("item_id = ?", input.ItemId).
		Updates(map[string]interface{}{
			"name": input.TaxName,
			"rate": input.TaxRate,
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

	result = ts.db().Where("id = ?", input.TaxId).Find(&tax) // get new data Tax
	return tax, nil
}

func (ts *TaxService) DeleteTax(input req.TaxDeleteRequest) (responseError *helpers.ResponseError) {
	var tax TaxService

	// get tax data where id = input.ID and itema_id = input.ItemId
	findResult := ts.db().Model(&models.Tax{})

	if input.ItemId > 0 {
		findResult = findResult.Where("item_id = ?", input.ItemId)
	}

	findResult = findResult.Where("id = ?", input.TaxId).Find(&tax)
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

	// delete tax data where id = input.TaxId
	result := ts.db().Model(&models.Tax{})

	if input.ItemId > 0 {
		result = result.Where("item_id = ?", input.ItemId)
	}

	result = result.Debug().Where("id = ?", input.TaxId).Delete(&tax)
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
