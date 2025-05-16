package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello?name=TestUser", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"Hello TestUser!"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestDataHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/data?id=testID&type=testType", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(dataHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"Data retrieved.","data":{"id":"testID","type":"testType"}}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSubmitHandler(t *testing.T) {
	var jsonData = []byte(`{"name": "Test", "surname": "User", "email": "test@example.com"}`)
	req, err := http.NewRequest("POST", "/submit", strings.NewReader(string(jsonData)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(submitHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	expectedMessagePrefix := "Data received for user Test User"
	if !strings.HasPrefix(response.Message, expectedMessagePrefix) {
		t.Errorf("handler returned unexpected message: got %v, want prefix %v", response.Message, expectedMessagePrefix)
	}

	expectedData := map[string]interface{}{
		"name":    "Test",
		"surname": "User",
		"email":   "test@example.com",
	}

	responseDataMap, ok := response.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("response.Data is not a map: %v", response.Data)
	}

	if responseDataMap["Name"] != expectedData["Name"] ||
		responseDataMap["Surname"] != expectedData["Surname"] ||
		responseDataMap["Email"] != expectedData["Email"] {
		t.Errorf("handler returned unexpected data: got %v, want %v", responseDataMap, expectedData)
	}
}
