package routers

import (
	c "ols/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter function will perform all route operations
func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	routeTax := r.Group("/tax")
	{
		routeTax.POST("/create", c.CreateTax)
		routeTax.POST("/update", c.UpdateTax)
		routeTax.POST("/delete", c.DeleteTax)
	}

	routeItem := r.Group("/item")
	{
		routeItem.POST("/create", c.CreateItem)
		routeItem.POST("/update", c.UpdateItem)
		routeItem.POST("/delete", c.DeleteItem)
		routeItem.POST("/get_list", c.GetListItem)
	}

	return r

}
