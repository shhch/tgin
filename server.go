package main

import (
	"github.com/gin-gonic/gin"
	"src"
)

func main() {

	//url := map[string]string{
	//	"a": "asd",
	//	"b":"zxc",
	//}

	src.checkBuild()

	r := gin.Default()
	r.GET("/check_build")
	r.Run("127.0.0.1:8089")
}
