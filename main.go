package main

import (
	"log"
	"net/http"
)

// Declare db as a global variable
var db *Database // Assuming Database is the type of your database connection

func main() {
	// Initialize database connection
	err := initDB()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	// Defer the closing of the database connection
	defer db.Close()

	// Set up HTTP request handlers
	http.HandleFunc("/asset", handleGetAsset)
	http.HandleFunc("/order", handlePlaceOrder)

	// Start server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
