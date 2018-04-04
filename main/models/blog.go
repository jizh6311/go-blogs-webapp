package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Blog struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"Username" json:"Username"`
	Image       string        `bson:"Image" json:"Image"`
	Description string        `bson:"Description" json:"Description"`
}
