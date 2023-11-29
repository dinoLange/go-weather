package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func TryApi() (*Forecast, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	apiUrl := "https://api.kachelmannwetter.com/v02/"
	req, err := http.NewRequest("GET", apiUrl+"forecast/53.868525/10.682954/3day", nil)
	if err != nil {
		fmt.Println("Error calling endpoint:", err)
		return nil, err
	}
	req.Header.Add("X-API-Key", "CY3jLwIBWWeNml_Rh2pWxI7fRVpa9cR_swbt-3INMhDS4OS2hLA_FXcqUiCXG8VBLg0GViyOXKiTw2kiR-MkDeAMVmIXRJkJcXeNJgRkibgrEfoTJ_9vvkSBrroe7kte")

	response, err := client.Do(req)

	if err != nil {
		fmt.Println("Error calling endpoint:", err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var forecastData Forecast

	err = json.Unmarshal(body, &forecastData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}
	fmt.Println(forecastData.Resolution)

	return &forecastData, nil
}
