// go test
package main

import (
	"encoding/json"
	"testing"
)

func TestInputDataJSON(t *testing.T) {
	input := InputData{
		Name:    "Test",
		Surname: "User",
		Email:   "test@example.com",
	}

	expectedJSON := `{"name":"Test","surname":"User","email":"test@example.com"}`

	actualJSONBytes, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal InputData: %v", err)
	}
	actualJSON := string(actualJSONBytes)

	if actualJSON != expectedJSON {
		t.Errorf("JSON marshal mismatch: got %v, want %v", actualJSON, expectedJSON)
	}

	var unmarshaled InputData
	err = json.Unmarshal([]byte(expectedJSON), &unmarshaled)
	if err != nil {
		t.Fatalf("failed to unmarshal InputData: %v", err)
	}

	if unmarshaled != input {
		t.Errorf("JSON unmarshal mismatch: got %+v, want %+v", unmarshaled, input)
	}
}

func TestResponseJSON(t *testing.T) {
	response := Response{
		Message: "Test message",
		Data: map[string]int{
			"id": 123,
		},
	}

	expectedJSON := `{"message":"Test message","data":{"id":123}}`

	actualJSONBytes, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("failed to marshal Response: %v", err)
	}
	actualJSON := string(actualJSONBytes)

	// Nemusíme porovnávat přesně řetězce, protože pořadí klíčů v mapě nemusí být stejné.
	// Pro jednoduchost to ale zde uděláme. Pro robustnější testování bychom mohli JSON parsovat a porovnávat struktury.
	if actualJSON != expectedJSON {
		t.Errorf("JSON marshal mismatch: got %v, want %v", actualJSON, expectedJSON)
	}

	var unmarshaled Response
	err = json.Unmarshal([]byte(expectedJSON), &unmarshaled)
	if err != nil {
		t.Fatalf("failed to unmarshal Response: %v", err)
	}

	// Pro porovnání Response s mapou v Data musíme provést hlubší kontrolu
	if unmarshaled.Message != response.Message {
		t.Errorf("Message mismatch after unmarshal: got %v, want %v", unmarshaled.Message, response.Message)
	}
	unmarshaledData, ok := unmarshaled.Data.(map[string]interface{})
	if !ok {
		t.Fatalf("unmarshaled Data is not a map: %v", unmarshaled.Data)
	}
	originalData, ok := response.Data.(map[string]int)
	if !ok {
		t.Fatalf("original Data is not a map[string]int: %v", response.Data)
	}
	if float64(originalData["id"]) != unmarshaledData["id"] { // JSON unmarshal číselné typy do float64
		t.Errorf("Data mismatch after unmarshal: got %v, want %v", unmarshaledData, originalData)
	}
}
