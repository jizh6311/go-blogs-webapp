package main

import (
	"go-blogs-webapp/main/config"

	"go-blogs-webapp/main/dao"

	"go-blogs-webapp/main/handlers"

	"github.com/labstack/echo"
)

func init() {
	var config config.Config
	var blogsDAO dao.BlogsDAO
	config.Read()
	blogsDAO.Server = config.Server
	blogsDAO.Database = config.Database
	blogsDAO.Connect()
}

func main() {
	e := echo.New()

	e.File("/", "main/public/index.html")

	e.GET("/blogs/:username", handlers.GetBlogs)
	e.POST("/blogs", handlers.PostBlogs)

	//TODO: Add more APIs for blogs
	// e.PUT("/blogs", handlers.PutBlogs(db))
	// e.DELETE("/blogs/:id", handlers.DeleteBlogs(db))

	e.Logger.Fatal(e.Start(":8000"))
}
