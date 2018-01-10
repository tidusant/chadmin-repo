package models

import (
	"gopkg.in/mgo.v2/bson"
)

type BuildScript struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	ShopID       string        `bson:"shopid"`
	Domain       string        `bson:"domain"`
	TemplateCode string        `bson:"templatecode"`
	ObjectID     string        `bson:"objectid"`
	Collection   string        `bson:"collection"`
	Status       int           `bson:"status"` //0: news, 1: building, 2: finish
	Created      int64         `bson:"created"`
	Modified     int64         `bson:"modified"`
	Lang         string        `bson:"lang"`
}
