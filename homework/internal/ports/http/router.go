package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"homework/internal/app"
)

const (
	defaultReadTimeout  = 30 * time.Second
	defaultWriteTimeout = 30 * time.Second
)

type Handler struct {
	service      app.Service
	fullAddress  string
	readTimeout  time.Duration
	writeTimeout time.Duration
}

type Config struct {
	Service      app.Service
	Port         string
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewHandler(config *Config) Handler {
	readTimeout := config.ReadTimeout
	if readTimeout < 1 {
		readTimeout = defaultReadTimeout
	}

	writeTimeout := config.WriteTimeout
	if writeTimeout < 1 {
		writeTimeout = defaultWriteTimeout
	}

	fullAddress := fmt.Sprintf("%s:%s", config.Host, config.Port)

	return Handler{
		service:      config.Service,
		writeTimeout: writeTimeout,
		readTimeout:  readTimeout,
		fullAddress:  fullAddress,
	}
}

func (h *Handler) NewServer() *http.Server {
	mux := chi.NewRouter()

	mux.Route("/", func(r chi.Router) {
		r.Post("/devices", h.createDevice)
		r.Get("/devices/{id}", h.getDevice)
		r.Delete("/devices/{id}", h.deleteDevice)
		r.Put("/devices", h.updateDevice)
	})

	return &http.Server{
		Addr:         h.fullAddress,
		Handler:      mux,
		ReadTimeout:  h.readTimeout,
		WriteTimeout: h.writeTimeout,
	}
}
