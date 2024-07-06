package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
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

func (h handler) login(w http.ResponseWriter, r *http.Request) {
	var req loginrReq

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

	tokenStr, err := h.svc.login(r.Context(), req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request: %s", err.Error()),
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusAccepted, pkg.DataResponse{
		Status:  "success",
		Message: "Login successful!",
		Data:    tokenStr,
	})
}
