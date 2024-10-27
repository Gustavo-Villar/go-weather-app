package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertTemperatureToKelvin(t *testing.T) {
	tests := []struct {
		name           string
		tempC          float64
		expectedKelvin float64
	}{
		{"Zero Celsius to Kelvin", 0.0, 273.15},
		{"Positive Celsius to Kelvin", 25.0, 298.15},
		{"Negative Celsius to Kelvin", -10.0, 263.15},
		{"High Celsius to Kelvin", 100.0, 373.15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertTemperatureToKelvin(tt.tempC)
			assert.Equal(t, tt.expectedKelvin, result)
		})
	}
}

func TestConvertTemperatureToFahrenheit(t *testing.T) {
	tests := []struct {
		name               string
		tempC              float64
		expectedFahrenheit float64
	}{
		{"Zero Celsius to Fahrenheit", 0.0, 32.0},
		{"Positive Celsius to Fahrenheit", 25.0, 77.0},
		{"Negative Celsius to Fahrenheit", -10.0, 14.0},
		{"High Celsius to Fahrenheit", 100.0, 212.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertTemperatureToFahrenheit(tt.tempC)
			assert.Equal(t, tt.expectedFahrenheit, result)
		})
	}
}
