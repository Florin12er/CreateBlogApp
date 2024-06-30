package models

import (
	"fmt"
	"github.com/florin12er/GoBlogApp/pkg/config"
	"github.com/jinzhu/gorm"
	"strings"
)

var db *gorm.DB

type Blog struct {
	gorm.Model
	Name            string `gorm:"" json:"name"`
	Author          string `json:"author"`
	Date            string `json:"publication"`
	Tags            string `json:"tags"`
	Content         string `json:"content"`
	ShortContent    string `gorm:"-"` // Exclude from database schema, used only in Go code
}

func init() {
	config.Connect()
	fmt.Println("connected to postgres")
	db = config.GetDB()
	db.AutoMigrate(&Blog{})
}

func (b *Blog) CreateBlog() *Blog {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBlogs() []Blog {
	var Blogs []Blog
	db.Find(&Blogs)

	// Populate ShortContent for each blog entry
	for i := range Blogs {
		Blogs[i].ShortContent = truncateContent(Blogs[i].Content, 150) // Adjust length as needed
	}

	return Blogs
}

// Function to truncate content
func truncateContent(content string, length int) string {
	if len(content) > length {
		return strings.TrimSpace(content[:length]) + "..."
	}
	return content
}

func GetBlogsById(Id int64) (*Blog, *gorm.DB) {
	var getBlogs Blog
	db := db.Where("ID=?", Id).Find(&getBlogs)
	return &getBlogs, db
}

func DeleteBlog(ID int64) error {
	var blog Blog
	result := db.Where("ID=?", ID).Delete(&blog)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

