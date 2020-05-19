package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckHeader(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("In Check Header")
	if id != "333" {
		var res Response
		res.Code = "41"
		res.Message = "Unauthorize"
		c.JSON(200, res)
		c.Abort()
	}
	c.Next()

}
