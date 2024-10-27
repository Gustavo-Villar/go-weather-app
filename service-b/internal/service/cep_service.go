package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gustavo-villar/go-weather-tracker/service-b/internal/model"
)

// IsValidCEP checks if the given CEP is valid (8 digits and all numeric).
func IsValidCEP(cep string) bool {
	// Remove any non-digit characters (like hyphens)
	cep = regexp.MustCompile(`\D`).ReplaceAllString(cep, "")

	// Check if the length is exactly 8 and all characters are digits
	return len(cep) == 8 && regexp.MustCompile(`^\d{8}$`).MatchString(cep)
}

// GetLocationByCEP fetches location details by CEP.
func GetLocationByCEP(cep string) (model.ViaCEPResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return model.ViaCEPResponse{}, fmt.Errorf("failed to connect to ViaCEP API: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return model.ViaCEPResponse{}, fmt.Errorf("error fetching data: status code %d", resp.StatusCode)
	}

	var result model.ViaCEPResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return model.ViaCEPResponse{}, fmt.Errorf("failed to decode response: %v", err)
	}

	// Check if the result contains an error
	if result.Localidade == "" {
		return model.ViaCEPResponse{}, fmt.Errorf("can not find zipcode")
	}

	return result, nil
}
