package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ManagerService struct {
	server *http.Server
}

func NewManagerService(port int) *ManagerService {
	root := chi.NewRouter()

	managerService := &ManagerService{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: root,
		},
	}

	root.Get("/config", managerService.GetConfig)
	root.Get("/list-config", managerService.ListConfig)
	root.Post("/config", managerService.AddConfig)
	root.Post("/start", managerService.Start)
	root.Post("/stop", managerService.Stop)

	return managerService
}

func Run(port int) error {
	serv := NewManagerService(port)

	err := serv.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
