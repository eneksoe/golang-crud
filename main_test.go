package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"wallester_test/controllers"
	"wallester_test/database"
	"wallester_test/entities"
)

func TestGetAll(t *testing.T) {
	LoadAppConfig()
	database.Connect(AppConfig.ConnectionString)

	req, err := http.NewRequest("GET", "customer/getAll", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var response []entities.Customer
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Errorf("got invalid response, expected list of customers, got: %v", rr.Body.String())
	}

	if len(response) < 1 {
		t.Errorf("expected at least 1 customer, got %v", len(response))
	}

	for _, customer := range response {
		if customer.ID == 0 {
			t.Errorf("expected customer id %d to  have a source path, was empty", customer.ID)
		}
	}
}
