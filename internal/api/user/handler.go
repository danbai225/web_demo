package user

import (
	"web_demo/internal/services"
	"web_demo/internal/services/user"
)

type handler struct {
	userServer user.Service
}

func New() *handler {
	return &handler{
		userServer: services.SeverM.Get("user").(user.Service),
	}
}

func (h *handler) i() {}
