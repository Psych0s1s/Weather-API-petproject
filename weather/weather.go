package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherData struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Name string `json:"name"`
}

type Client struct {
	APIKey string
}

func (c *Client) FetchWeather(cityName string) (*WeatherData, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric&lang=ru", baseURL, cityName, c.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
