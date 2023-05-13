package app

import "net/http"

func (m *ManagerService) ListConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("implement list config"))
}
