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

func CreateTax(c *gin.Context) {
	var taxService s.TaxService
	var input req.CreateTaxRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	//call service
	taxDetail, responseError := taxService.CreateTax(input)

	if responseError != nil {
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else {
		response = res.Message(false, http.StatusOK, nil)
		response["data"] = resp.FormatCreateTaxResponse(*taxDetail)
	}

	//return response using api helper
	res.Respond(c, c.Writer, response)

}

func UpdateTax(c *gin.Context) {
	var TaxService s.TaxService
	var input req.TaxUpdateRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	TaxDetail, responseError := TaxService.UpdateTax(input) //call service

	if responseError != nil { // check generalError is empty or not, if not nill return detailError
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else { // return success status and data Tax detail
		response = res.Message(false, http.StatusOK, nil)
		formatter := resp.FormatCreateTaxResponse(*TaxDetail)
		response["data"] = formatter
	}

	res.Respond(c, c.Writer, response)
}

func DeleteTax(c *gin.Context) {
	var TaxService s.TaxService
	var input req.TaxDeleteRequest
	var response map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil { // check form validation
		response = res.Message(true, http.StatusBadRequest, "Invalid parameter")
		response["data"] = gin.H{"errors": helpers.FormatValidationError(err)}
		res.Respond(c, c.Writer, response)
		return
	}

	responseError := TaxService.DeleteTax(input) //call service

	if responseError != nil { // check generalError is empty or not, if not nill return detailError
		response = res.Message(true, responseError.HttpResponseCode, responseError.ErrorMessage)
		response["data"] = responseError.ErrorDetail
	} else { // return success status and data Tax detail
		response = res.Message(false, http.StatusOK, nil)
		response["data"] = resp.FormatItemSuccessResponse(true)
	}

	res.Respond(c, c.Writer, response)
}
