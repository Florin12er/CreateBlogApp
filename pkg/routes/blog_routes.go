package routes

import (
    "github.com/florin12er/GoBlogApp/pkg/controllers"
    "github.com/gorilla/mux"
)

// RegisterBlogRoutes registers all routes related to blog operations
func RegisterBlogRoutes(router *mux.Router) {
    // GET all blogs
    router.HandleFunc("/", controllers.ShowAllBlogs).Methods("GET")

    // GET form to create a new blog
    router.HandleFunc("/blog/new", controllers.GoToCreateBlog).Methods("GET")

    // POST to create a new blog
    router.HandleFunc("/blog", controllers.CreateBlog).Methods("POST")

    // GET a specific blog by ID
    router.HandleFunc("/blog/{blogId}", controllers.GetBlog).Methods("GET")

    // DELETE a specific blog by ID
    router.HandleFunc("/blog/{blogId}", controllers.DeleteBlog).Methods("DELETE")

    // PUT to update a specific blog by ID
    router.HandleFunc("/blog/{blogId}", controllers.UpdateBlog).Methods("PUT")

    // GET form to edit a specific blog by ID
    router.HandleFunc("/blog/{blogId}/edit", controllers.GoToEditBlog).Methods("GET")
}

