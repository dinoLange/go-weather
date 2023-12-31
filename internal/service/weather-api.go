package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiUrl = "https://api.kachelmannwetter.com/v02/"

func LoadCurrentWeather() (*CurrentWeather, error) {
	body := loadBody("station/VH0568/observations/1h")

	var data CurrentWeather

	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &data, nil
}

func LoadThreeDayForeCast() (*Forecast, error) {
	body := loadBody("forecast/53.868525/10.682954/3day")

	var data Forecast

	err := json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return &data, nil
}

func loadBody(path string) []byte {
	req, err := http.NewRequest("GET", apiUrl+path, nil)
	if err != nil {
		fmt.Println("Error calling endpoint:", err)
		return nil
	}
	req.Header.Add("X-API-Key", "CY3jLwIBWWeNml_Rh2pWxI7fRVpa9cR_swbt-3INMhDS4OS2hLA_FXcqUiCXG8VBLg0GViyOXKiTw2kiR-MkDeAMVmIXRJkJcXeNJgRkibgrEfoTJ_9vvkSBrroe7kte")

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error calling endpoint:", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}
	return body
}
