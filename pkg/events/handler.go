package events

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
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
	eID := chi.URLParam(r, "event_id")
	if _, err := uuid.FromString(eID); err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid event id",
			Status:  "error",
		})
		return
	}

	event, err := h.svc.get(r.Context(), eID)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, pkg.GenericResponse{
			Message: "sorry an error occurred.",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "event retrieved successfully",
		Status:  "success",
		Data:    event,
	})
}

func (h handler) list(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	limitStr, cursor := q.Get("limit"), q.Get("cursor")

	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 20
	}

	params := queryParams{
		limit:  limit,
		cursor: cursor,
	}

	event, responseParams, err := h.svc.rlist(r.Context(), params)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, pkg.GenericResponse{
			Message: "sorry an error occurred.",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "events retrieved successfully",
		Status:  "success",
		Data: map[string]any{
			"events":     event,
			"pagination": responseParams,
		},
	})

}
