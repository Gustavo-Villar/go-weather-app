package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsValidCEP(t *testing.T) {
	valid := IsValidCEP("01001000")
	if !valid {
		t.Errorf("Expected valid CEP")
	}

	invalid := IsValidCEP("abc123")
	if invalid {
		t.Errorf("Expected invalid CEP")
	}
}

func TestGetLocationByCEP_Success(t *testing.T) {
	// Mock ViaCEP API response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"cep": "01001-000",
			"logradouro": "Praça da Sé",
			"bairro": "Sé",
			"localidade": "São Paulo",
			"uf": "SP"
		}`))
	}))
	defer server.Close()

	// Call the function with the mock server URL
	result, err := GetLocationByCEP("01001000")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Localidade != "São Paulo" {
		t.Errorf("Expected 'São Paulo', got %v", result.Localidade)
	}
}

func TestGetLocationByCEP_Invalid(t *testing.T) {
	// Mock ViaCEP API response for invalid CEP
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := GetLocationByCEP("invalidCEP")
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}
}
