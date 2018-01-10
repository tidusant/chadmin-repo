package models

import (
	"gopkg.in/mgo.v2/bson"
)

//News ...
type TemplateLang struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	TemplateCode string        `bson:"templatecode"`
	Lang         string        `bson:"lang"`
	ShopID       string        `bson:"shopid"`
	Key          string        `bson:"key"`
	Value        string        `bson:"value"`
}
