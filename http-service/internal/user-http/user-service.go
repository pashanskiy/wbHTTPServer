package http

import (
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
)

var servicePath = "/user"

type UserHTTPService struct {
	httpMux *httpRouter.HttpRouter
}

func NewUserHTTPService(httpMux *httpRouter.HttpRouter) *UserHTTPService {
	return &UserHTTPService{httpMux: httpMux}
}

func (s *UserHTTPService) SetRoutes() {

	
	routes := []httpRouter.RouteHandler{
		{Route: handleCreateRoute, HandlerFunc: s.handleCreate},
	}

	s.httpMux.SetRoutes(routes)
}
