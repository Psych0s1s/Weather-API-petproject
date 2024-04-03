package main

import (
	"Weather-API-petproject/weather"
	"fmt"
)

func main() {
	client := weather.Client{APIKey: "fe49090ccc4ec67af284dd84e74807fb"}
	weatherData, err := client.FetchWeather("Moscow")
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}

	fmt.Printf("Текущая погода в городе %s:\n", weatherData.Name)
	fmt.Printf("Температура воздуха: %.1f°C\n", weatherData.Main.Temp)
	fmt.Printf("Относительная влажность: %d%%\n", weatherData.Main.Humidity)
	fmt.Printf("Скорость ветра: %.1fm/s\n", weatherData.Wind.Speed)
	fmt.Printf("Облачность: %s\n", weatherData.Weather[0].Description)
}
