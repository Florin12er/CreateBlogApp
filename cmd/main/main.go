package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"html/template"

	"github.com/florin12er/GoBlogApp/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBlogRoutes(r)

	// Serve static files
	fs := http.FileServer(http.Dir("../../static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type header explicitly for CSS files
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}
		fs.ServeHTTP(w, r)
	})))

	// Register the template handler for the root route
	r.HandleFunc("/{page}", serveTemplate).Methods("GET")
	r.HandleFunc("/", serveTemplate).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	addr := "0.0.0.0:" + port
	log.Printf("Starting server on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", r.URL.Path)

	// Append ".html" if no extension is present
	if filepath.Ext(fp) == "" {
		fp += ".html"
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"Title": "My Blog",
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

