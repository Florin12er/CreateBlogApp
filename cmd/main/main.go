package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"github.com/florin12er/GoBlogApp/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBlogRoutes(r)

	// Serve static files
	fs := http.FileServer(http.Dir("../../static"))
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", fs))

	// Register the template handler for the root route
	r.HandleFunc("/{page}", serveTemplate).Methods("GET")
	r.HandleFunc("/", serveTemplate).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	// Correctly determine template paths
	layoutPath := filepath.Join("../../templates", "layout.html")
	pagePath := filepath.Join("../../templates", r.URL.Path)

	// Ensure the page path ends with ".html"
	if filepath.Ext(pagePath) == "" {
		pagePath += ".html"
	}

	// Check if the requested template exists
	if _, err := os.Stat(pagePath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}

	// Parse the layout and page templates together
	tmpl, err := template.ParseFiles(layoutPath, pagePath)
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Example data to pass to the template
	data := map[string]interface{}{
		"Title": "My Blog",
	}

	// Execute the template with the data
	if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
