package app

import "net/http"

func (m *ManagerService) Start(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("implement start"))
}
