package route

import (
	"github.com/gin-gonic/gin"
	"go-library/api/author"
	"go-library/api/author/gateway/handler"
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
	//e.Group("/author").
	//	GET("/", svc.Handler.GetList).
	//	GET("/:id", svc.Handler.Get).
	//	POST("/", svc.Handler.Store).
	//	PUT("/:id", svc.Handler.Update).
	//	DELETE("/:id", svc.Handler.Delete)

	a := e.Group("/author")
	a.GET("/", svc.Handler.GetList)
	a.GET("/:id", svc.Handler.Get)
	a.POST("/", svc.Handler.Store)
	a.PUT("/:id", svc.Handler.Update)
	a.DELETE("/:id", svc.Handler.Delete)
}

func New() registry.RouteContract {
	return &service{
		Handler: author.ApiHandler(),
	}
}
