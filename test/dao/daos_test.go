package daosTest

import (
	"testing"

	"../../main/dao"
	"../../main/models"
	"gopkg.in/mgo.v2/bson"
)

func Test_Success_Insert(t *testing.T) {
	var blogsDAO dao.BlogsDAO
	blog := &models.Blog{ID: bson.NewObjectId()}
	if err := blogsDAO.Mock_Success_Insert(blog); err != nil {
		t.Errorf("Error thrown when inserting the data.")
	}
}

func Test_Failure_Insert(t *testing.T) {
	var blogsDAO dao.BlogsDAO
	var err error
	blog := &models.Blog{ID: bson.NewObjectId()}
	if err = blogsDAO.Mock_Failure_Insert(blog); err == nil {
		t.Errorf("The insertion should fail.")
	}

	if err.Error() != "Insertion failed" {
		t.Errorf("The error message is incorrect")
	}
}

//TODO: Add more unit test cases for blog_dao
