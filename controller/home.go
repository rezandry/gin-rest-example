package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetHome(c *gin.Context) {
	var res Response
	fmt.Println(c.Params.ByName("id"))
	fmt.Println(c.Request.URL.Query())
	res.Code = "1"
	res.Message = "Success From Home"
	c.JSON(200, res)
}

func Login(c *gin.Context) {
	var res Response
	var login LoginStruct
	c.Bind(&login)
	fmt.Println(login)
	res.Code = "1"
	res.Message = "Success from Public"
	c.JSON(200, res)
}
