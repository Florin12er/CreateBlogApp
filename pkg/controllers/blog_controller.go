package controllers

import (
	"fmt"
	"github.com/florin12er/GoBlogApp/pkg/models"
	"github.com/florin12er/GoBlogApp/pkg/utils"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
    "encoding/json"
)

var NewBlog models.Blog

func ShowAllBlogs(w http.ResponseWriter, r *http.Request) {
	NewBlogs := models.GetAllBlogs()
	layoutPath, layoutErr := utils.GetTemplatePath("layout.html")
	tmplPath, tmplErr := utils.GetTemplatePath("index.html")

	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if layoutErr != nil || tmplErr != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, NewBlogs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	layoutPath, layoutErr := utils.GetTemplatePath("layout.html")
	tmplPath, tmplErr := utils.GetTemplatePath("show.html")

	if layoutErr != nil || tmplErr != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		return
	}

	blogId := vars["blogId"]
	ID, err := strconv.ParseInt(blogId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing")
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	blogDetails, _ := models.GetBlogsById(ID)
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, blogDetails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GoToCreateBlog(w http.ResponseWriter, r *http.Request) {
	layoutPath, layoutErr := utils.GetTemplatePath("layout.html")
	tmplPath, tmplErr := utils.GetTemplatePath("create.html")

	if layoutErr != nil || tmplErr != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		return
	}

	NewBlogs := models.GetAllBlogs()
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := tmpl.Execute(w, NewBlogs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		CreateBook := &models.Blog{}
		CreateBook.Name = r.FormValue("name")
		CreateBook.Author = r.FormValue("author")
		CreateBook.Content = r.FormValue("content")
		CreateBook.Links = r.FormValue("links")
		var _ = CreateBook.CreateBlog()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func GoToEditBlog(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blogId := vars["blogId"]
    ID, err := strconv.ParseInt(blogId, 0, 0)
    if err != nil {
        fmt.Println("error while parsing")
        http.Error(w, "Invalid blog ID", http.StatusBadRequest)
        return
    }

    blogDetails, _ := models.GetBlogsById(ID)
    if blogDetails == nil {
        http.Error(w, "Blog not found", http.StatusNotFound)
        return
    }

    layoutPath, layoutErr := utils.GetTemplatePath("layout.html")
    tmplPath, tmplErr := utils.GetTemplatePath("edit.html")

    if layoutErr != nil || tmplErr != nil {
        http.Error(w, "Error loading templates", http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles(layoutPath, tmplPath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    if err := tmpl.Execute(w, blogDetails); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	blogId := vars["blogId"]
	ID, err := strconv.ParseInt(blogId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	if err := models.DeleteBlog(ID); err != nil {
		http.Error(w, "Failed to delete blog", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Blog deleted successfully"))
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Blog struct
	var updateBlog models.Blog
	utils.ParseBody(r, &updateBlog)

	// Get the blog ID from URL parameters
	vars := mux.Vars(r)
	blogId := vars["blogId"]
	ID, err := strconv.ParseInt(blogId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	// Retrieve the existing blog from the database
	blogDetails, db := models.GetBlogsById(ID)
	if blogDetails == nil {
		http.Error(w, "Blog not found", http.StatusNotFound)
		return
	}

	// Update fields if they are not empty in the request
	if updateBlog.Name != "" {
		blogDetails.Name = updateBlog.Name
	}
	if updateBlog.Author != "" {
		blogDetails.Author = updateBlog.Author
	}
	if updateBlog.Links != "" {
		blogDetails.Links = updateBlog.Links
	}
	if updateBlog.Content != "" {
		blogDetails.Content = updateBlog.Content
	}

	// Save the updated blog back to the database
	db.Save(blogDetails)

	// Respond with updated blog details as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogDetails)
}

