package route

import (
	"github.com/gin-gonic/gin"
	"go-library/api/category"
	"go-library/api/category/gateway/handler"
	"go-library/registry"
)

func init() {
	registry.RegisterRouter(New)
}

type Route struct {
	Engine *gin.Engine
}

type service struct {
	Handler handler.HandlerContract
}

func (svc *service) Endpoints(e *gin.Engine) {
	a := e.Group("/category")
	a.GET("/", svc.Handler.GetList)
	a.GET("/:id", svc.Handler.Get)
	a.POST("/", svc.Handler.Store)
	a.PUT("/:id", svc.Handler.Update)
	a.DELETE("/:id", svc.Handler.Delete)
}

func New() registry.RouteContract {
	return &service{
		Handler: category.ApiHandler(),
	}
}
