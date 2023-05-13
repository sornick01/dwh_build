package app

import "net/http"

func (m *ManagerService) GetConfig(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("implement get_config"))
}
