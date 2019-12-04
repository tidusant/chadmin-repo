package models

import (
	"gopkg.in/mgo.v2/bson"
)

type CHImage struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Uid       string        `bson:"uid"`
	Shopid    string        `bson:"shopid"`
	Albumname string        `bson:"albumname"`
	AppName   string        `bson:"appname"`
	Filename  string        `bson:"filename"`
	Created   int64         `bson:"created"`
}
