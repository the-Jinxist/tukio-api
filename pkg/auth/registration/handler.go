package registration

import (
	"encoding/json"
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

func (h handler) register(w http.ResponseWriter, r *http.Request) {
	var req registerUserReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	err = v.Struct(req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	err = h.svc.registerUser(r.Context(), req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusAccepted, pkg.GenericResponse{
		Status:  "success",
		Message: "Please check your email for the otp to complete your registration",
	})
}
