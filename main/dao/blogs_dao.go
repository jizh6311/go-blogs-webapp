package dao

import (
	"bytes"
	"errors"
	"fmt"
	"go-blogs-webapp/main/models"
	"io"
	"log"
	"mime/multipart"
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

func (m *BlogsDAO) Insert(blog *models.BlogJson, file *multipart.FileHeader) error {
	blog.PostDate = time.Now()
	uploadErr := UploadImage(blog, file)
	if uploadErr != nil {
		return uploadErr
	}

	blogMeta := models.BlogMeta{ID: blog.ID, Username: blog.Username,
		Description: blog.Description, PostDate: blog.PostDate}
	err := db.C(COLLECTION).Insert(&blogMeta)
	return err
}

func UploadImage(blog *models.BlogJson, multipartFile *multipart.FileHeader) error {
	file, createErr := db.GridFS("fs").Create(blog.ID.Hex() + ".jpg")
	if createErr != nil {
		fmt.Println("Create error message: " + createErr.Error())
		return createErr
	}

	src, err := multipartFile.Open()
	if err != nil {
		fmt.Println("Open error message: " + err.Error())
		return err
	}
	defer src.Close()

	readBuffer := bytes.NewBuffer(nil)
	if _, readErr := io.Copy(readBuffer, src); readErr != nil {
		fmt.Println("Read error message: " + readErr.Error())
		return readErr
	}

	_, writeErr := file.Write(readBuffer.Bytes())
	if writeErr != nil {
		fmt.Println("Write error message: " + writeErr.Error())
		return writeErr
	}

	closeErr := file.Close()
	if closeErr != nil {
		fmt.Println("Close error message: " + closeErr.Error())
		return closeErr
	}

	return nil
}

func (m *BlogsDAO) Find(username string) (error, []*models.BlogMeta) {
	blogs := []*models.BlogMeta{}
	err := db.C(COLLECTION).Find(bson.M{"Username": username}).All(&blogs)
	return err, blogs
}

func (m *BlogsDAO) RetriveImage(ID string) (error, string) {
	//TODO: Able to open multiple types of files
	file, findErr := db.GridFS("fs").Open(ID + ".jpg")
	if findErr != nil {
		return findErr, ""
	}

	byteArray := make([]byte, file.Size())
	_, readErr := file.Read(byteArray)
	if readErr != nil {
		return readErr, ""
	}
	closeErr := file.Close()
	if closeErr != nil {
		return closeErr, ""
	}

	return nil, string(byteArray)
}

func (m *BlogsDAO) Mock_Success_Insert(blog *models.BlogJson) error {
	return nil
}

func (m *BlogsDAO) Mock_Failure_Insert(blog *models.BlogJson) error {
	return errors.New("Insertion failed")
}
