package http

import (
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
	userStorageAPI "wbHTTPServer/storage-service/api/user"
)

var servicePath = "/user"

type UserHTTPService struct {
	httpMux *httpRouter.HttpRouter

	userStorageApi userStorageAPI.UserStorageServiceClient
}

func NewUserHTTPService(httpMux *httpRouter.HttpRouter, userStorageApi userStorageAPI.UserStorageServiceClient) *UserHTTPService {
	return &UserHTTPService{httpMux: httpMux, userStorageApi: userStorageApi}
}

func (s *UserHTTPService) SetRoutes() {

	routes := []httpRouter.RouteHandler{
		{Route: handleCreateRoute, HandlerFunc: s.handleCreate},
		{Route: handleGetRoute, HandlerFunc: s.handleGet},
		{Route: handleDeleteRoute, HandlerFunc: s.handleDelete},
		{Route: handleUpdateRoute, HandlerFunc: s.handleUpdate},
	}

	s.httpMux.SetRoutes(routes)
}
