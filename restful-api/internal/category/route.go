package category

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Handler Handler
}

func NewRoute(handler Handler) *Route {
	return &Route{
		Handler: handler,
	}
}

func (route *Route) RegisterRoutes(router *httprouter.Router) {
	router.GET("/categories", route.Handler.FindAll)
	router.GET("/categories/:categoryId", route.Handler.FindById)
	router.POST("/categories", route.Handler.Create)
	router.PUT("/categories/:categoryId", route.Handler.Update)
	router.DELETE("/categories/:categoryId", route.Handler.Delete)
}
