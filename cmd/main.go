package cmd

import (
	"github.com/go-chi/chi"
	"net/http"
	"super_service/config"
	"super_service/internal/controller"
	"super_service/internal/repository"
	"super_service/internal/service"
)

func main() {
	repo := repository.NewUserStorage(config.NewConfig())
	svc := service.NewService(repo)
	ctl := controller.NewController(svc)

	r := chi.NewRouter()
	r.Get("/", ctl.TakeBook)

	http.ListenAndServe(":8080", r)
}
