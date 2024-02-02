package service

import (
	"context"
	"dwh/internal/domain"
	"dwh/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Service struct {
	repo *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{repo: r}
}

func (svc *Service) Build(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := &domain.Database{}

	body, err := r.GetBody()
	if err != nil {
		http.Error(w, "can't get body", http.StatusInternalServerError)
		return
	}
	defer func() {
		err = body.Close()
		fmt.Fprintf(os.Stderr, err.Error())
	}()

	var bytes []byte
	_, err = body.Read(bytes)
	if err != nil {
		http.Error(w, "can't read body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, db)
	if err != nil {
		http.Error(w, "can't unmarshal body", http.StatusInternalServerError)
		return
	}

	sql := db.ToSql()
	err = svc.repo.ExecSql(ctx, sql)
	if err != nil {
		http.Error(w, "Build: "+err.Error(), http.StatusInternalServerError)
	}
}

func (svc *Service) Migrate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	routes := &domain.Routes{}

	body, err := r.GetBody()
	if err != nil {
		http.Error(w, "can't get body", http.StatusInternalServerError)
		return
	}
	defer func() {
		err = body.Close()
		fmt.Fprintf(os.Stderr, err.Error())
	}()

	var bytes []byte
	_, err = body.Read(bytes)
	if err != nil {
		http.Error(w, "can't read body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(bytes, routes)
	if err != nil {
		http.Error(w, "can't unmarshal body", http.StatusInternalServerError)
		return
	}

	sql := routes.ToSql()
	err = svc.repo.ExecSql(ctx, sql)
	if err != nil {
		http.Error(w, "Build: "+err.Error(), http.StatusInternalServerError)
	}
}
