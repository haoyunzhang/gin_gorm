package main

import (
	"github.com/gin-gonic/gin"
	"hfga.com.cn/ginorm/router"
)


func main() {

	r := gin.Default()
	router.Route(r)
	r.Run()
}
