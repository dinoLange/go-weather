package service

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type LineChartData struct {
	Columns string
	Rows    string
	Title   string
	Width   string
	Height  string
}

func RenderLineChart(responseWriter http.ResponseWriter) {
	forecast, err := LoadThreeDayForeCast()
	maxTemps := forecast.generateTemperatureMax()
	var temperatures [][]string

	for _, row := range maxTemps {
		temperatures = append(temperatures, row)
	}
	fmt.Println(stringify2DSlice(temperatures))
	data := LineChartData{
		Columns: "number: Datum, number: Temperatur",
		Rows:    stringify2DSlice(temperatures),
		Title:   "3 Tage Wetter Lübeck ",
		Width:   "500",
	}
	err = templates.ExecuteTemplate(responseWriter, "lineChart.html", data)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}

// stringify2DSlice converts a 2D slice to its string representation
func stringify2DSlice(slice [][]string) string {
	var builder strings.Builder

	// Opening bracket
	builder.WriteString("[\n")

	for i, row := range slice {
		builder.WriteString("  [")
		for j, element := range row {
			// Format each element with a comma
			if j > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(element)
		}
		// Closing bracket for the row
		builder.WriteString("]")
		// Add a comma after each row except the last one
		if i < len(slice)-1 {
			builder.WriteString(",")
		}
		builder.WriteString("\n")
	}

	// Closing bracket
	builder.WriteString("]")

	return builder.String()
}

func (forecast *Forecast) generateTemperatureMax() [][]string {
	items := make([][]string, 0)
	for i, data := range forecast.Data {
		items = append(items, []string{fmt.Sprintf("%v", i), fmt.Sprintf("%g", data.TempMax)})
		//	items = append(items, data.TimeOfDay.Morning.TempMax)
		//	items = append(items, data.TimeOfDay.Afternoon.TempMax)
		//	items = append(items, data.TimeOfDay.Evening.TempMax)
	}
	return items
}
