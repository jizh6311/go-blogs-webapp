package models

import "gopkg.in/mgo.v2/bson"
import "time"

type BlogMeta struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"Username" json:"Username"`
	Description string        `bson:"Description" json:"Description"`
	PostDate    time.Time     `bson:"PostDate" json:"PostDate"`
}
