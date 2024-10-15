package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gustavo-villar/go-weather-tracker/internal/model"
	"github.com/gustavo-villar/go-weather-tracker/internal/service"
)

func HandleGetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	// Check if empty
	if cep == "" {
		http.Error(w, "cep is required", http.StatusBadRequest)
		return
	}

	// Check if valid
	if !service.IsValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Get ViaCep Object
	viaCepObj, err := service.GetLocationByCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	// Get Location from ViaCepObj
	location := viaCepObj.Localidade + ", " + viaCepObj.Uf
	// Construct location string and escape it for URL use
	escapedLocation := url.QueryEscape(location)

	// Get WeatherAPI Object
	weatherApiObj, err := service.GetTemperature(escapedLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Create response object
	response := model.Response{
		TempC: weatherApiObj.Current.TempC,
		TempF: service.ConvertTemperatureToFahrenheit(weatherApiObj.Current.TempC),
		TempK: service.ConvertTemperatureToKelvin(weatherApiObj.Current.TempC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
