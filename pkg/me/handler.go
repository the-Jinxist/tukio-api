package me

import (
	"net/http"

	"github.com/the-Jinxist/tukio-api/pkg"
	"github.com/thedevsaddam/renderer"
)

var (
	rnd = renderer.New()
)

type handler struct {
	svc service
}

func NewHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) get(w http.ResponseWriter, r *http.Request) {

	profile, err := h.svc.get(r.Context())
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "profile retrieved successfully",
		Status:  "success",
		Data:    profile,
	})

}
