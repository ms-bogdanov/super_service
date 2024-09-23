package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"super_service/internal/model"
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
	msg, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	var takeRequest model.TakeBookRequest
	err = json.Unmarshal(msg, &takeRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = c.svc.TakeBook(takeRequest.UserID, takeRequest.BookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func (c Controller) ReturnBooks(w http.ResponseWriter, r *http.Request) {
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

	err = c.svc.ReturnBook(uID, bID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
