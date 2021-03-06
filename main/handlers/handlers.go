package handlers

import (
	"net/http"

	"go-blogs-webapp/main/dao"

	"go-blogs-webapp/main/models"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

var blogsDAO dao.BlogsDAO

func GetBlogs(c echo.Context) (err error) {
	username := c.Param("username")

	if err, _ := blogsDAO.Find(username); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	_, blogs := blogsDAO.Find(username)
	return c.JSON(http.StatusOK, blogs)
}

func GetBlogById(c echo.Context) (err error) {
	id := c.Param("id")
	err, fileBytes := blogsDAO.RetriveImage(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fileBytes)
}

func PostBlogs(c echo.Context) (err error) {
	blog := &models.BlogJson{ID: bson.NewObjectId()}

	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err = c.Bind(blog); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if err = blogsDAO.Insert(blog, file); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, blog)
}

func DeleteBlog(c echo.Context) (err error) {
	id := c.Param("id")
	err, deleteMessage := blogsDAO.Delete(id)

	if err != nil {
		return c.String(http.StatusInternalServerError, deleteMessage)
	}

	return c.String(http.StatusOK, deleteMessage)
}

func Mock_Get_Success_Blogs(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "testLogs")
}
