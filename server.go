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
		tool.P.Password = c.Param("password")
		tool.P.Client = c.Param("client")
		tool.P.Ip = c.Param("ip")
		tool.P.Swoole = c.Param("swoole")
		tool.P.Platform = c.Param("platform")
		tool.P.Datrixlog = c.Param("datrixlog")
		tool.P.Web = c.Param("web")


	})

	r.Run("127.0.0.1:8089")
}
