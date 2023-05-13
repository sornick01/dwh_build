package app

import "net/http"

func (m *ManagerService) AddConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("implement add_config"))
}
