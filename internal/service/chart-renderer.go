package service

import (
	"fmt"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func RenderThreeDayForecast(responseWriter http.ResponseWriter) {
	forecast, err := LoadThreeDayForeCast()
	if err != nil {
		panic("oh no!")
	}
	htmlTemplate(responseWriter, forecast)

}

func (forecast *Forecast) generateTemperaturMaxLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, data := range forecast.Data {
		items = append(items, opts.LineData{Value: data.TimeOfDay.Night.TempMax})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Morning.TempMax})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Afternoon.TempMax})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Evening.TempMax})
	}
	return items
}

func (forecast *Forecast) generateTemperaturMinLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, data := range forecast.Data {
		items = append(items, opts.LineData{Value: data.TimeOfDay.Night.TempMin})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Morning.TempMin})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Afternoon.TempMin})
		items = append(items, opts.LineData{Value: data.TimeOfDay.Evening.TempMin})
	}
	fmt.Println("items", len(items))

	return items
}

func htmlTemplate(responseWriter http.ResponseWriter, forecast *Forecast) {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "3 Tage Vorhersage Lübeck",
	}))

	dayNames := []string{}
	for _, day := range forecast.Data {
		fmt.Println("day", day)

		dayNames = append(dayNames,
			day.DayName, "", "", "")
	}
	line.SetXAxis(dayNames).
		AddSeries("Temperatur Max", forecast.generateTemperaturMaxLineItems()).
		AddSeries("Temperatur Min", forecast.generateTemperaturMinLineItems())

	line.Render(responseWriter)
}
