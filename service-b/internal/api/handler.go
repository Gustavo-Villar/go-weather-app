package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gustavo-villar/go-weather-tracker/service-b/internal/model"
	"github.com/gustavo-villar/go-weather-tracker/service-b/internal/service"
	"go.opentelemetry.io/otel"
)

func HandleGetWeather(w http.ResponseWriter, r *http.Request) {
	ctx, commandSpan := otel.GetTracerProvider().Tracer("weather").Start(r.Context(), "weather-command")
	defer commandSpan.End()

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
	ctx, zipcodeQuerySpan := otel.GetTracerProvider().Tracer("weather").Start(ctx, "weather-zipcode-query")
	viaCepObj, err := service.GetLocationByCEP(cep)
	zipcodeQuerySpan.End()
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	// Get Location from ViaCepObj
	location := viaCepObj.Localidade + ", " + viaCepObj.Uf
	// Construct location string and escape it for URL use
	escapedLocation := url.QueryEscape(location)

	// Get WeatherAPI Object
	_, weatherQuerySpan := otel.GetTracerProvider().Tracer("weather").Start(ctx, "weather-query")
	weatherApiObj, err := service.GetTemperature(escapedLocation)
	weatherQuerySpan.End()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Create response object
	response := model.Response{
		City:  viaCepObj.Localidade,
		TempC: weatherApiObj.Current.TempC,
		TempF: service.ConvertTemperatureToFahrenheit(weatherApiObj.Current.TempC),
		TempK: service.ConvertTemperatureToKelvin(weatherApiObj.Current.TempC),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
	}
}
