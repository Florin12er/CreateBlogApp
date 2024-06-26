package controllers

import (
	"fmt"
	"github.com/florin12er/GoBlogApp/pkg/models"
	"github.com/florin12er/GoBlogApp/pkg/utils"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
)

var NewBlog models.Blog

func ShowAllBlogs(w http.ResponseWriter, r *http.Request) {
	NewBlogs := models.GetAllBlogs()
	layoutPath, err := utils.GetTemplatePath("layout.html")
	tmplPath, err := utils.GetTemplatePath("index.html")

	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, NewBlogs)

}
func GetBlog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    layoutPath, err := utils.GetTemplatePath("layout.html")
    tmplPath, err := utils.GetTemplatePath("show.html")
	blogId := vars["blogId"]
	ID, err := strconv.ParseInt(blogId, 0, 0)
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
	tmpl.Execute(w, blogDetails)
}
func GoToCreateBlog(w http.ResponseWriter, r *http.Request) {
    layoutPath, err := utils.GetTemplatePath("layout.html")
    tmplPath, err := utils.GetTemplatePath("create.html")
	NewBlogs := models.GetAllBlogs()
	tmpl, err := template.ParseFiles(layoutPath, tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, NewBlogs)
}
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		CreateBlog := &models.Blog{}
		CreateBlog.Name = r.FormValue("name")
		CreateBlog.Author = r.FormValue("author")
		CreateBlog.Links = r.FormValue("links")
		CreateBlog.Content = r.FormValue("publication")
		var _ = CreateBlog.CreateBlog()
		http.Redirect(w, r, "/blog", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		vars := mux.Vars(r)
		blogId := vars["blogId"]
		ID, err := strconv.ParseInt(blogId, 10, 64)
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		// Call your delete function from models package
		err = models.DeleteBlog(ID)
		if err != nil {
			http.Error(w, "Failed to delete book", http.StatusInternalServerError)
			return
		}

		// Respond with success message
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Blog deleted successfully"))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBlog = &models.Blog{}
	utils.ParseBody(r, updateBlog)
	vars := mux.Vars(r)
	blogId := vars["blogId"]
	ID, err := strconv.ParseInt(blogId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	blogDetails, db := models.GetBlogsById(ID)
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
	db.Save(&blogDetails)
	http.Redirect(w, r, fmt.Sprintf("/book/%d", blogDetails.ID), http.StatusSeeOther)
}
