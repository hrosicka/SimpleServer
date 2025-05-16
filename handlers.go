package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const dataFile = "data.jsonl"

// Handler for the /hello path (GET)
// EXAMPLE: http://localhost:8080/hello?name=Hanka
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	response := Response{
		Message: fmt.Sprintf("Hello %s!", name),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for the /data path (GET)
// EXAMPLE: http://localhost:8080/data?id=123456&type=user
func dataHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	dataType := r.URL.Query().Get("type")
	response := Response{
		Message: "Data retrieved.",
		Data: map[string]string{
			"id":   id,
			"type": dataType,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler for the /submit path (POST)
// EXAMPLE: curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "surname": "Blue", "email": "alice.blue@example.com"}' http://localhost:8080/submit
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected a POST request.", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body.", http.StatusBadRequest)
		log.Println("Error reading body:", err)
		return
	}
	defer r.Body.Close()

	var input InputData
	err = json.Unmarshal(body, &input)
	if err != nil {
		http.Error(w, "Invalid JSON format.", http.StatusBadRequest)
		log.Println("Error parsing JSON:", err)
		return
	}

	response := Response{
		Message: fmt.Sprintf("Data received for user %s %s (%s).", input.Name, input.Surname, input.Email),
		Data:    input,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	// Save data to file
	err = saveDataToFile(input)
	if err != nil {
		log.Println("Error saving data to file:", err)
	}
}

func saveDataToFile(data InputData) error {
	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	if err != nil {
		return err
	}
	return nil
}
