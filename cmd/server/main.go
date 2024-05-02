package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/xyproto/musicsite/pkg/handlers"
)

func main() {
    r := mux.NewRouter()

    // Define route handlers using the mux router
    r.HandleFunc("/api/login", handlers.LoginHandler).Methods("POST")
    r.HandleFunc("/api/register", handlers.RegisterHandler).Methods("POST")
    r.HandleFunc("/api/user", handlers.UserHandler).Methods("GET", "PUT", "DELETE")
    r.HandleFunc("/api/bands", handlers.BandsHandler).Methods("GET", "POST", "PUT", "DELETE")

    // Serve static files
    r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

    log.Println("Server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}
