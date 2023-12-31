package service

import (
	"go-weather/internal/view"
)

type CurrentWeather struct {
	StationID string
	Name      string
	Lat       float64
	Lon       float64
	Ele       int
	Data      []struct {
		DateTime             string
		Temp                 float64
		TempMean             float64
		TempMin              float64
		TempMax              float64
		DewpointMean         float64
		PrecSum              float64
		Temp5CmMin           float64
		WindSpeedMean        float64
		WindSpeedMax         float64
		WindGustMax          float64
		WindDirectionMean    float64
		HumidityRelativeMean float64
	} `json:"data"`
}

func (currentWeather *CurrentWeather) GenerateTemperatureMax() []view.RowValue {
	items := make([]view.RowValue, 0)
	for _, data := range currentWeather.Data {
		items = append(items, view.RowValue{Time: data.DateTime, Value: data.TempMax})
		//	items = append(items, data.TimeOfDay.Morning.TempMax)
		//	items = append(items, data.TimeOfDay.Afternoon.TempMax)
		//	items = append(items, data.TimeOfDay.Evening.TempMax)
	}
	return items
}
