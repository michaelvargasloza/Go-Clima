package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func main() {
	apiKey := "bd5e378503939ddaee76f12ad7a97608" // API key de OpenWeatherMap
	city := "La Paz"                             // Ciudad de la que desea obtener el clima
	lang := "es"                                 // Configura el idioma a español

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&lang=%s", city, apiKey, lang)

	response, err := http.Get(url)
	//nil == null
	if err != nil {
		fmt.Printf("Error al hacer la solicitud: %s\n", err)
		return
	}
	defer response.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Printf("Error al decodificar la respuesta JSON: %s\n", err)
		return
	}

	tempCelsius := data.Main.Temp - 273.15
	description := data.Weather[0].Description

	fmt.Printf("Clima en %s:\n", city)
	fmt.Printf("Temperatura: %.2f°C\n", tempCelsius)
	fmt.Printf("Descripción: %s\n", description)
}
