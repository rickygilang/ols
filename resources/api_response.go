package resource

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//ResponseData structure
type ResponseData struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Error            interface{} `json:"error"`
	ErrorMessage     interface{} `json:"error_message"`
	HttpResponseCode interface{} `json:"http_response_code"`
}

func Message(err bool, httpResponseCode int, errMessage interface{}) map[string]interface{} {
	return map[string]interface{}{
		"error":            err,
		"httpResponseCode": httpResponseCode,
		"errorMessage":     errMessage,
	}
}

//Respond returns basic response structure
func Respond(c *gin.Context, w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")

	meta := Meta{
		Error:            data["error"],
		HttpResponseCode: data["httpResponseCode"],
		ErrorMessage:     data["errorMessage"],
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data["data"],
	}

	c.JSON(data["httpResponseCode"].(int), jsonResponse)
}
