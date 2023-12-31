package server

import (
	"go-weather/internal/service"
	"go-weather/internal/view"
	"net/http"
)

func (s *Server) BaseHandler(w http.ResponseWriter, r *http.Request) {
	view.Base().Render(r.Context(), w)
}

func (s *Server) weatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	forecast, err := service.LoadThreeDayForeCast()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := forecast.GenerateTemperatureMax()
	columns := []view.Column{{Type: "date", Name: "Datum"}, {Type: "number", Name: "Temperatur"}}

	renderLineChart(w, r, data, columns, "Wettervorhersage Lübeck")
}

func (s *Server) currentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	currentWeather, err := service.LoadCurrentWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := currentWeather.GenerateTemperatureMax()
	columns := []view.Column{{Type: "date", Name: "Datum"}, {Type: "number", Name: "Temperatur"}}
	renderLineChart(w, r, data, columns, "Aktuelles Wetter Lübeck")
}

func renderLineChart(w http.ResponseWriter, r *http.Request, data []view.RowValue, columns []view.Column, title string) {
	view.LineChart(data, columns, title, 800, 500).Render(r.Context(), w)
}
