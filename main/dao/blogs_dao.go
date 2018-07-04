package dao

import (
	"errors"
	"go-blogs-webapp/main/models"
	"io/ioutil"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BlogsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "blogs"
)

// Establish a connection to database
func (m *BlogsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *BlogsDAO) Insert(blog *models.BlogJson) error {
	blog.PostDate = time.Now()
	uploadErr := UploadImage(blog)
	if uploadErr != nil {
		return uploadErr
	}

	blogMeta := models.BlogMeta{ID: blog.ID, Username: blog.Username,
		Description: blog.Description, PostDate: blog.PostDate}
	err := db.C(COLLECTION).Insert(&blogMeta)
	return err
}

func UploadImage(blog *models.BlogJson) error {
	file, createErr := db.GridFS("fs").Create(blog.ID.Hex() + ".jpg")
	if createErr != nil {
		return createErr
	}
	fileBytes, readErr := ioutil.ReadFile(blog.Image)
	if readErr != nil {
		return readErr
	}
	_, writeErr := file.Write(fileBytes)
	if writeErr != nil {
		return writeErr
	}

	closeErr := file.Close()
	if closeErr != nil {
		return writeErr
	}

	return nil
}

func (m *BlogsDAO) Mock_Success_Insert(blog *models.BlogJson) error {
	return nil
}

func (m *BlogsDAO) Mock_Failure_Insert(blog *models.BlogJson) error {
	return errors.New("Insertion failed")
}

func (m *BlogsDAO) Find(username string) (error, []*models.BlogMeta) {
	blogs := []*models.BlogMeta{}
	err := db.C(COLLECTION).Find(bson.M{"Username": username}).All(&blogs)
	return err, blogs
}

func RetriveImage(blog *models.BlogMeta) error {

	return nil
}
