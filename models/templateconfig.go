package models

import (
	"gopkg.in/mgo.v2/bson"
)

//News ...
type TemplateConfig struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	TemplateCode string        `bson:"templatecode"`
	ShopID       string        `bson:"shopid"`
	Key          string        `bson:"key"`
	Type         string        `bson:"type"`
	Value        string        `bson:"value"`
}
