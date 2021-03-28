package main

import (
	"fmt"
	"github.com/bakhodur-nazriev/modernWebApp/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port :8000")
	http.ListenAndServe(":8000", nil)
}
