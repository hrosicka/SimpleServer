package main

// Structure for data expected in POST requests
type InputData struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
}

// Structure for JSON response
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
