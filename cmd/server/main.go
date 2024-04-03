package main

import (
	"Weather-API-petproject/weather"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		city := r.URL.Query().Get("city")
		if city == "" {
			http.Error(w, "Не указан параметр 'city'", http.StatusBadRequest)
			return
		}

		encodedCityName := url.QueryEscape(city)

		weatherData, err := weather.FetchWeather(encodedCityName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Ошибка при получении данных о погоде: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherData)

	})

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
