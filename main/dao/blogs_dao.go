package dao

import (
	"log"

	"../models"
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

func (m *BlogsDAO) Insert(blog *models.Blog) error {
	err := db.C(COLLECTION).Insert(blog)
	return err
}

func (m *BlogsDAO) Find(username string) (error, []*models.Blog) {
	blogs := []*models.Blog{}
	err := db.C(COLLECTION).Find(bson.M{"Username": username}).All(&blogs)
	return err, blogs
}
