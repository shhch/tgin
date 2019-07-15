package main

import (
	"github.com/gin-gonic/gin"
	"tgin/tool"
)

func main() {

	r := gin.Default()
	r.GET("/checkbuild", func(c *gin.Context) {
		ret := tool.CheckBuild()
		c.JSON(200,gin.H{
			"result":ret,
		})
	})

	r.GET("/build", func(c *gin.Context) {

	})

	r.Run("127.0.0.1:8089")
}
