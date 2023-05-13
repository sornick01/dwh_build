package app

import "net/http"

func (m *ManagerService) Stop(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("implement stop"))
}
