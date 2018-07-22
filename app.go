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
	//GET all the blogs by username
	e.GET("/blogs/:username", handlers.GetBlogs)

	//GET one blog by bson id
	e.GET("/blog/:id", handlers.GetBlogById)

	e.POST("/blogs", handlers.PostBlogs)

	//TODO: Add more APIs for blogs
	// e.PUT("/blogs", handlers.PutBlogs(db))
	// e.DELETE("/blogs/:id", handlers.DeleteBlogs(db))

	//Use npm and go CLI to start frontend and backend services seperately.
	//Then use proxy in package.json
	e.Logger.Fatal(e.Start(":8081"))
}
