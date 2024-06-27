package main

import (
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

    // Construct absolute path to static files directory
    currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    staticDir := filepath.Join(currentDir, "../..", "static") // Adjust the relative path as needed

    // Serve static files
    fs := http.FileServer(http.Dir(staticDir))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    addr := "0.0.0.0:" + port
    log.Printf("Starting server on %s", addr)
    log.Fatal(http.ListenAndServe(addr, r))
}

