package service

import (
	"go-weather/internal/view"
	"time"
)

type CustomTime time.Time

type Forecast struct {
	Lat           float64   `json:"lat"`
	Lon           float64   `json:"lon"`
	Alt           int       `json:"alt"`
	Resolution    string    `json:"resolution"`
	TimeZone      string    `json:"timeZone"`
	SystemOfUnits string    `json:"systemOfUnits"`
	Run           time.Time `json:"run"`
	Data          []struct {
		DateTime      CustomTime `json:"dateTime"`
		DayName       string     `json:"dayName"`
		TempMax       float64    `json:"tempMax"`
		TempMin       float64    `json:"tempMin"`
		PrecCurrent   float64    `json:"precCurrent"`
		WindGust      float64    `json:"windGust"`
		WindSpeed     float64    `json:"windSpeed"`
		WindDirection int        `json:"windDirection"`
		SunHours      float64    `json:"sunHours"`
		CloudCoverage int        `json:"cloudCoverage"`
		WeatherSymbol string     `json:"weatherSymbol"`
		Risks         []any      `json:"risks"`
		TimeOfDay     struct {
			Night struct {
				TimeOfDay     string  `json:"timeOfDay"`
				IsDay         bool    `json:"isDay"`
				TempMax       float64 `json:"tempMax"`
				TempMin       float64 `json:"tempMin"`
				PrecCurrent   float64 `json:"precCurrent"`
				WindGust      float64 `json:"windGust"`
				WindSpeed     float64 `json:"windSpeed"`
				WindDirection int     `json:"windDirection"`
				WeatherSymbol string  `json:"weatherSymbol"`
				Risks         []any   `json:"risks"`
			} `json:"night"`
			Morning struct {
				TimeOfDay     string  `json:"timeOfDay"`
				IsDay         bool    `json:"isDay"`
				TempMax       float64 `json:"tempMax"`
				TempMin       float64 `json:"tempMin"`
				PrecCurrent   float64 `json:"precCurrent"`
				WindGust      float64 `json:"windGust"`
				WindSpeed     float64 `json:"windSpeed"`
				WindDirection int     `json:"windDirection"`
				WeatherSymbol string  `json:"weatherSymbol"`
				Risks         []any   `json:"risks"`
			} `json:"morning"`
			Afternoon struct {
				TimeOfDay     string  `json:"timeOfDay"`
				IsDay         bool    `json:"isDay"`
				TempMax       float64 `json:"tempMax"`
				TempMin       float64 `json:"tempMin"`
				PrecCurrent   float64 `json:"precCurrent"`
				WindGust      float64 `json:"windGust"`
				WindSpeed     float64 `json:"windSpeed"`
				WindDirection int     `json:"windDirection"`
				WeatherSymbol string  `json:"weatherSymbol"`
				Risks         []any   `json:"risks"`
			} `json:"afternoon"`
			Evening struct {
				TimeOfDay     string  `json:"timeOfDay"`
				IsDay         bool    `json:"isDay"`
				TempMax       float64 `json:"tempMax"`
				TempMin       float64 `json:"tempMin"`
				PrecCurrent   float64 `json:"precCurrent"`
				WindGust      float64 `json:"windGust"`
				WindSpeed     float64 `json:"windSpeed"`
				WindDirection int     `json:"windDirection"`
				WeatherSymbol string  `json:"weatherSymbol"`
				Risks         []any   `json:"risks"`
			} `json:"evening"`
		} `json:"timeOfDay"`
	} `json:"data"`
}

func (forecast *Forecast) GenerateTemperatureMax() []view.RowValue {
	items := make([]view.RowValue, 0)
	for _, data := range forecast.Data {
		items = append(items, view.RowValue{Time: time.Time(data.DateTime).Add(time.Hour * 2).String(), Value: data.TimeOfDay.Night.TempMax})
		items = append(items, view.RowValue{Time: time.Time(data.DateTime).Add(time.Hour * 8).String(), Value: data.TimeOfDay.Morning.TempMax})
		items = append(items, view.RowValue{Time: time.Time(data.DateTime).Add(time.Hour * 14).String(), Value: data.TimeOfDay.Afternoon.TempMax})
		items = append(items, view.RowValue{Time: time.Time(data.DateTime).Add(time.Hour * 20).String(), Value: data.TimeOfDay.Evening.TempMax})
	}
	return items
}
