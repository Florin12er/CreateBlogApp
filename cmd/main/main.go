package main

import (
	"log"
	"net/http"

	"github.com/florin12er/GoBlogApp/pkg/routes"
	"github.com/gorilla/mux"
)
func main() {
    r := mux.NewRouter()
    routes.RegisterBlogRoutes(r)
    http.Handle("/",r)
    addr := "localhost:3000"
    log.Printf("Starting server on %s", addr)
    log.Fatal(http.ListenAndServe(addr,r))
}
