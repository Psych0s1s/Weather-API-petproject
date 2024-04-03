package main

import (
	"Weather-API-petproject/weather"
	"fmt"
	"log"
	"net/url"
)

func main() {
	fmt.Println("Введите название города на русском языке:")
	var cityName string
	if _, err := fmt.Scanln(&cityName); err != nil {
		log.Fatal("Ошибка при чтении названия города:", err)
	}

	encodedCityName := url.QueryEscape(cityName)

	weatherData, err := weather.FetchWeather(encodedCityName)
	if err != nil {
		fmt.Println("Ошибка при получении данных о погоде:", err)
		return
	}

	if len(weatherData.Weather) == 0 {
		fmt.Println("Нет данных о погоде для указанного города.")
		return
	}

	fmt.Printf("Текущая погода в городе %s:\n", weatherData.Name)
	fmt.Printf("Температура воздуха: %.1f°C\n", weatherData.Main.Temp)
	fmt.Printf("Относительная влажность: %d%%\n", weatherData.Main.Humidity)
	fmt.Printf("Скорость ветра: %.1fm/s\n", weatherData.Wind.Speed)
	fmt.Printf("Облачность: %s\n", weatherData.Weather[0].Description)
}
