package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"

	"homework/internal/devices"
)

func (h *Handler) createDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		h.processError(w, "can not read request body", http.StatusBadRequest)
		return
	}

	var device devices.Device
	err = json.Unmarshal(buf, &device)
	if err != nil {
		h.processError(w, "can not unmarshal request body", http.StatusBadRequest)
		return
	}

	err = h.service.CreateDevice(&device)
	if err != nil {
		h.processError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		h.processError(w, "'id' is required param", http.StatusBadRequest)
		return
	}

	device, err := h.service.GetDevice(id)
	if err != nil {
		h.processError(w, err.Error(), http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(device)
	if err != nil {
		h.processError(w, "can not marshal device", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(buf)
}

func (h *Handler) deleteDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	id := chi.URLParam(r, "id")
	if len(id) == 0 {
		h.processError(w, "'id' is required param", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteDevice(id)
	if err != nil {
		h.processError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateDevice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	buf, err := io.ReadAll(r.Body)
	if err != nil {
		h.processError(w, "can not read request body", http.StatusBadRequest)
		return
	}

	var device devices.Device
	err = json.Unmarshal(buf, &device)
	if err != nil {
		h.processError(w, "can not unmarshal request body", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateDevice(&device)
	if err != nil {
		h.processError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type ErrorBody struct {
	Message string `json:"message"`
}

func (h *Handler) processError(w http.ResponseWriter, msg string, code int) {
	body := ErrorBody{
		Message: msg,
	}
	buf, _ := json.Marshal(body)

	w.WriteHeader(code)
	_, _ = w.Write(buf)
}
