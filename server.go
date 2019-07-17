package main

import (
	"github.com/gin-gonic/gin"
	"tgin/tool"
)


func main() {

	r := gin.Default()
	r.GET("/checkbuild", func(c *gin.Context) {
		ret, err := tool.CheckBuild()
		if err != nil {
			c.JSON(200,gin.H{
				"msg":err.Error(),
			})
		}
		c.JSON(200,gin.H{
			"result":ret,
		})
	})
	r.GET("/build", func(c *gin.Context) {
		param := tool.Param{}
		param.Password := c.Param("password")
	})

	r.Run("127.0.0.1:8089")
}
