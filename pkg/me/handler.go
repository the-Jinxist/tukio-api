package me

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/the-Jinxist/tukio-api/pkg"
	"github.com/thedevsaddam/renderer"
)

var (
	rnd = renderer.New()
	v   = validator.New(validator.WithRequiredStructEnabled())
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

func (h handler) getUserProfile(w http.ResponseWriter, r *http.Request) {

	userID := chi.URLParam(r, "user_id")
	_, err := uuid.FromString(userID)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid user id",
			Status:  "error",
		})
		return
	}

	profile, err := h.svc.getUserProfile(r.Context(), userID)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "user profile retrieved successfully",
		Status:  "success",
		Data:    profile,
	})

}

func (h handler) update(w http.ResponseWriter, r *http.Request) {
	var req updateProfileReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})
		return
	}

	err = v.Struct(req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})
		return
	}

	err = h.svc.update(r.Context(), req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "profile updated successfully",
		Status:  "success",
	})
}
