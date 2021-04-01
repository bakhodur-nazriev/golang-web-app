package main

import (
	"fmt"
	"github.com/bakhodur-nazriev/modernWebApp/pkg/config"
	"github.com/bakhodur-nazriev/modernWebApp/pkg/handlers"
	"github.com/bakhodur-nazriev/modernWebApp/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port :8080")
	http.ListenAndServe(":8080", nil)
}
