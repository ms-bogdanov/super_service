package main

import (
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"net/http"
	"super_service/config"
	"super_service/internal/controller"
	"super_service/internal/repository"
	"super_service/internal/service"
)

func main() {
	cfg := config.NewConfig()

	repo := repository.NewUserStorage(cfg)
	svc := service.NewService(repo)
	ctl := controller.NewController(svc)

	svc.InitProject()

	r := chi.NewRouter()
	r.Post("/take", ctl.TakeBook)
	r.Post("/return", ctl.ReturnBooks)

	http.ListenAndServe(":8080", r)
}
