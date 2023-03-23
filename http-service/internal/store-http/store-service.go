package http

import (
	httpRouter "wbHTTPServer/http-service/pkg/http-router"
	storeStorageAPI "wbHTTPServer/storage-service/api/store"
)

var servicePath = "/store"

type StoreHTTPService struct {
	httpMux *httpRouter.HttpRouter

	storeStorageApi storeStorageAPI.StoreStorageServiceClient
}

func NewStoreHTTPService(httpMux *httpRouter.HttpRouter, storeStorageApi storeStorageAPI.StoreStorageServiceClient) *StoreHTTPService {
	return &StoreHTTPService{httpMux: httpMux, storeStorageApi: storeStorageApi}
}

func (s *StoreHTTPService) SetRoutes() {

	routes := []httpRouter.RouteHandler{
		{Route: handleCreateRoute, HandlerFunc: s.handleCreate},
		{Route: handleGetRoute, HandlerFunc: s.handleGet},
		{Route: handleDeleteRoute, HandlerFunc: s.handleDelete},
		{Route: handleUpdateRoute, HandlerFunc: s.handleUpdate},
	}

	s.httpMux.SetRoutes(routes)
}
