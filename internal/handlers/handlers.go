package handlers

import "github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"

type Handlers struct {
	services *service.Service
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{services}
}
