package registry

import "github.com/gin-gonic/gin"

type RouteContract interface {
	Endpoints(e *gin.Engine)
}

var routes []RouteFactory

type RouteFactory func() RouteContract

func RegisterRouter(route RouteFactory) {
	routes = append(routes, route)
}

func LoadRouter() []RouteFactory {
	return routes
}
