package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// Car struct defines the car structure
type Car struct {
	ID    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

// Slice to store cars
var cars []Car
var nextID = 1

// CreateCar creates a new car
func CreateCar(w http.ResponseWriter, r *http.Request) {
	var newCar Car
	json.NewDecoder(r.Body).Decode(&newCar)

	newCar.ID = nextID
	nextID++

	cars = append(cars, newCar)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newCar)
}

// GetCars retrieves all cars
func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// GetCarByID retrieves a car by its ID
func GetCarByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for _, car := range cars {
		if car.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// UpdateCar updates an existing car by ID
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	json.NewDecoder(r.Body).Decode(&updatedCar)

	for i, car := range cars {
		if car.ID == id {
			// Update car details
			cars[i].Make = updatedCar.Make
			cars[i].Model = updatedCar.Model
			cars[i].Year = updatedCar.Year

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(cars[i])
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

// DeleteCar deletes a car by ID
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	for i, car := range cars {
		if car.ID == id {
			// Remove car from slice
			cars = append(cars[:i], cars[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // 204 No Content
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func main() {
	// Route to create and get all cars
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreateCar(w, r)
		case http.MethodGet:
			GetCars(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Route for operations on a single car by ID
	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetCarByID(w, r)
		case http.MethodPut:
			UpdateCar(w, r)
		case http.MethodDelete:
			DeleteCar(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	http.ListenAndServe(":8000", nil)
}
