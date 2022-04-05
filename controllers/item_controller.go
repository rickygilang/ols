package controllers

import (
	"net/http"
	"ols/helpers"
	res "ols/resources"
	req "ols/resources/requests"
	resp "ols/resources/responses"
	s "ols/services"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var itemService s.ItemService
	var input req.CreateItemRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	//call service
	itemDetail, responseError := itemService.CreateItem(input)

	if responseError != nil {
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else {
		response = res.Message(false, http.StatusOK, nil)
		response["data"] = resp.FormatItemResponse(*itemDetail)
	}

	//return response using api helper
	res.Respond(c, c.Writer, response)

}

func UpdateItem(c *gin.Context) {
	var itemService s.ItemService
	var input req.ItemUpdateRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	ItemDetail, responseError := itemService.UpdateItem(input) //call service

	if responseError != nil { // check generalError is empty or not, if not nill return detailError
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else { // return success status and data Item detail
		response = res.Message(false, http.StatusOK, nil)
		formatter := resp.FormatItemResponse(*ItemDetail)
		response["data"] = formatter
	}

	res.Respond(c, c.Writer, response)
}

func DeleteItem(c *gin.Context) {
	var itemService s.ItemService
	var input req.ItemDeleteRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	responseError := itemService.DeleteItem(input) //call service

	if responseError != nil { // check generalError is empty or not, if not nill return detailError
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else { // return success status and data Item detail
		response = res.Message(false, http.StatusOK, nil)
		response["data"] = resp.FormatItemSuccessResponse(true)
	}

	res.Respond(c, c.Writer, response)
}

// Func for get list
func GetListItem(c *gin.Context) {
	var itemService s.ItemService
	var response map[string]interface{}

	// call service
	items, responseError := itemService.GetListItem()
	if responseError != nil { // check generalError is empty or not, if not nill return detailError
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else { // return success status and data Item detail
		response = res.Message(false, http.StatusOK, nil)
		response["data"] = resp.FormatGetListItemResponse(items)
	}

	res.Respond(c, c.Writer, response)
}
