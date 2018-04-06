package handlers

import (
	"net/http"

	"../dao"
	"../models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func GetBlogs(c echo.Context) (err error) {
	username := c.Param("username")
	var blogsDAO dao.BlogsDAO
	if err, _ := blogsDAO.Find(username); err != nil {
		return err
	}

	_, blogs := blogsDAO.Find(username)
	return c.JSON(http.StatusOK, blogs)
}

func PostBlogs(c echo.Context) (err error) {
	blog := &models.Blog{ID: bson.NewObjectId()}

	if err = c.Bind(blog); err != nil {
		return
	}

	var blogsDAO dao.BlogsDAO
	if err = blogsDAO.Insert(blog); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, blog)
}

func Mock_Get_Success_Blogs(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "testLogs")
}
