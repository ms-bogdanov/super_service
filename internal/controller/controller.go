package controller

import (
	"net/http"
	"strconv"
	"super_service/internal/service"
)

type Controller struct {
	svc *service.Service
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (c Controller) TakeBook(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	bookID := r.URL.Query().Get("id")

	uID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	bID, err := strconv.Atoi(bookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = c.svc.TakeBook(uID, bID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
