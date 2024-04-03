package weather

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

var apiKey string

func init() {
	apiKey = os.Getenv("OPENWEATHERMAP_API_KEY")
	if apiKey == "" {
		log.Fatal("Отсутствует ключ API") //Установите ключ в переменные окружения командой set (для windows) или export (для MacOS, Linux) OPENWEATHERMAP_API_KEY=ваш_api_ключ
	}
}

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

func FetchWeather(cityName string) (*WeatherData, error) {
	requestURL := fmt.Sprintf("%s?q=%s&appid=%s&units=metric&lang=ru", baseURL, cityName, apiKey)
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Ошибка API: статус %v", resp.Status)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
