package main

import (
	"encoding/json"
	"net/http"
)

// Task struct defines the car structure
type Car struct {
	ID    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var cars []Car
var nextID = 1

func CreateCar(w http.ResponseWriter, r *http.Request) {
	// Create a new Car
	var newCar Car

	// Decode the request body into the newCar struct
	json.NewDecoder(r.Body).Decode(&newCar)

	// Assign an ID to the car
	newCar.ID = nextID
	nextID++

	// Add the car to the slice
	cars = append(cars, newCar)

	// Send the created car as the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCar)
}

// GetCars retrieves all cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func main() {
	// Handle /cars route for creating cars
	http.HandleFunc("/cars", CreateCar)

	// Start the server
	http.ListenAndServe(":8080", nil)

}
