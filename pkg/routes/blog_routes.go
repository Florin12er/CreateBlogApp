package routes

import (
	"github.com/florin12er/GoBlogApp/pkg/controllers"
	"github.com/gorilla/mux"
)

// RegisterBookStoreRoutes registers all routes related to bookstore operations
func RegisterBlogRoutes(router *mux.Router) {
	router.HandleFunc("/", controllers.ShowAllBlogs).Methods("GET")
    router.HandleFunc("/blog/new" , controllers.GoToCreateBlog).Methods("GET")
	router.HandleFunc("/blog", controllers.CreateBlog).Methods("POST")
	router.HandleFunc("/blog/{blogId}", controllers.GetBlog).Methods("GET")
	router.HandleFunc("/blog/{blogId}", controllers.DeleteBlog).Methods("DELETE")
}
