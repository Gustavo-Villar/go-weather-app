package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gustavo-villar/go-weather-tracker/service-b/internal/model"
)

// GetTemperature fetches the current temperature for a given location.
func GetTemperature(location string) (model.WeatherAPIResponse, error) {
	apiKey := "4c03d5ea28454cb5be4233359241410"

	urlToCall := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, location)

	resp, err := http.Get(urlToCall)
	if err != nil {
		fmt.Println(err)
		return model.WeatherAPIResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.WeatherAPIResponse{}, fmt.Errorf("failed to fetch temperature, status code: %d", resp.StatusCode)
	}

	var result model.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.WeatherAPIResponse{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return result, nil
}

// Convert Temperature to Kelvin
func ConvertTemperatureToKelvin(tempC float64) float64 {
	return tempC + 273.15
}

// Convert Temperature to Fahrenheit
func ConvertTemperatureToFahrenheit(tempC float64) float64 {
	return (tempC * 9 / 5) + 32
}
