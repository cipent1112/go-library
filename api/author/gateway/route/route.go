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
	e.Group("/author").
		GET("/", svc.Handler.GetList).
		GET("/:id", svc.Handler.Get).
		POST("/", svc.Handler.Store)
}

func New() registry.RouteContract {
	return &service{
		Handler: author.ApiHandler(),
	}
}
