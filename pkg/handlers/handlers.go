package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the root route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Musician Network API!")
}

// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	// Implement login logic here
	fmt.Fprintf(w, "Login endpoint hit")
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}
	// Implement registration logic here
	fmt.Fprintf(w, "Register endpoint hit")
}

// UserHandler handles operations on a user
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "User GET endpoint hit")
	case "PUT":
		fmt.Fprintf(w, "User PUT endpoint hit")
	case "DELETE":
		fmt.Fprintf(w, "User DELETE endpoint hit")
	default:
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
	}
}

// BandsHandler handles operations on bands
func BandsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Bands GET endpoint hit")
	case "POST":
		fmt.Fprintf(w, "Bands POST endpoint hit")
	case "PUT":
		fmt.Fprintf(w, "Bands PUT endpoint hit")
	case "DELETE":
		fmt.Fprintf(w, "Bands DELETE endpoint hit")
	default:
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
	}
}
