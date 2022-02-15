package main

import (
	"github.com/gin-gonic/gin"
	"go-library/config"
	"go-library/registry"
	"strconv"

	_ "go-library/api/author/gateway/route"
	_ "go-library/api/category/gateway/route"
)

func main() {
	e := gin.Default()
	for _, router := range registry.LoadRouter() {
		router().Endpoints(e)
	}
	e.Run(":" + strconv.Itoa(config.Conf.AppConfig.Port))
}
