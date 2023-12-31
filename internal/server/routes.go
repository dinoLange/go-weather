package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.BaseHandler)
	r.Get("/weatherforecast", s.weatherForecastHandler)
	r.Get("/currentweather", s.currentWeatherHandler)

	return r
}
