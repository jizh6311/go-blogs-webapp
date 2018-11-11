package main

import (
	"go-blogs-webapp/main/config"

	"go-blogs-webapp/main/dao"

	"go-blogs-webapp/main/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// ALlow origin cross
	e.Use(middleware.CORS())

	// APIs below:
	e.GET("/blogs/:username", handlers.GetBlogs)
	e.GET("/blog/:id", handlers.GetBlogById)
	e.POST("/blog", handlers.PostBlogs)
	e.DELETE("/blog/:id", handlers.DeleteBlog)

	//TODO: Add more APIs for blogs
	// e.PUT("/blogs", handlers.PutBlogs(db))

	//Use npm and go CLI to start frontend and backend services seperately.
	//Then use proxy in package.json
	e.Logger.Fatal(e.Start(":8081"))
}
