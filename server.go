package main

import (
	"github.com/gin-gonic/gin"
	"server/src"
)



func main() {

	//url := map[string]string{
	//	"a": "asd",
	//	"b":"zxc",
	//}

	r := gin.Default()
	r.GET("/check_build", src.check_build)
	r.Run("127.0.0.1:8089")
}
