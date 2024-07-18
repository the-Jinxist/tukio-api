package events

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid"
	"github.com/nfnt/resize"
	"github.com/the-Jinxist/tukio-api/internal/media"
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
	params := getQueryParams(q)

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
			"events":     append([]EventResponse{}, event...),
			"pagination": responseParams,
		},
	})

}

func (h handler) listUserEvents(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	params := getQueryParams(q)

	event, responseParams, err := h.svc.listUserEvents(r.Context(), params)
	if err != nil {
		rnd.JSON(w, http.StatusInternalServerError, pkg.GenericResponse{
			Message: "sorry an error occurred.",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusOK, pkg.DataResponse{
		Message: "user events retrieved successfully",
		Status:  "success",
		Data: map[string]any{
			"events":     append([]EventResponse{}, event...),
			"pagination": responseParams,
		},
	})

}

func (h handler) create(w http.ResponseWriter, r *http.Request) {
	var req createEventParams

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

	err = h.svc.create(r.Context(), req)
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: "invalid request",
			Status:  "error",
		})
		return
	}

	rnd.JSON(w, http.StatusAccepted, pkg.GenericResponse{
		Status:  "success",
		Message: "Event created successfully!",
	})
}

func (h handler) uploadImage(w http.ResponseWriter, r *http.Request) {
	file, fH, err := r.FormFile("file")
	if err != nil {
		rnd.JSON(w, http.StatusBadRequest, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})
		return
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		rnd.JSON(w, http.StatusUnsupportedMediaType, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})

		return
	}

	fileExt := strings.Split(fH.Filename, ".")[1]
	imageID := uuid.Must(uuid.NewV7())

	fileName := fmt.Sprintf("%s.%s", imageID.String(), fileExt)

	m := resize.Resize(1000, 0, img, resize.Lanczos3)
	out, err := os.Create(fileName)
	if err != nil {
		rnd.JSON(w, http.StatusUnsupportedMediaType, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})
		return
	}
	defer out.Close()

	err = jpeg.Encode(out, m, nil)
	if err != nil {

		rnd.JSON(w, http.StatusUnsupportedMediaType, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})

		return
	}

	meta := media.ImageMeta{
		AssociatedID: imageID,
		ImageType:    "event_poster",
	}

	imageUrl, err := media.UploadImage(r.Context(), fileName, meta)
	if err != nil {

		rnd.JSON(w, http.StatusUnsupportedMediaType, pkg.GenericResponse{
			Message: fmt.Sprintf("invalid request; %s", err.Error()),
			Status:  "error",
		})

		return
	}

	rnd.JSON(w, http.StatusUnsupportedMediaType, pkg.DataResponse{
		Message: "image uploaded successfully",
		Status:  "success",
		Data: map[string]any{
			"image_url": imageUrl,
		},
	})
}

func getQueryParams(q url.Values) queryParams {
	limitStr, cursor := q.Get("limit"), q.Get("cursor")

	limit, _ := strconv.Atoi(limitStr)
	if limit == 0 {
		limit = 20
	}

	params := queryParams{
		limit:  limit,
		cursor: cursor,
	}
	return params
}
