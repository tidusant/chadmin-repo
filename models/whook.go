package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Whook struct {
	ID bson.ObjectId `bson:"_id,omitempty"`

	Name    string `bson:"name"`
	Action  string `bson:"action"`
	ShopID  string `bson:"shopid"`
	Created int    `bson:"created"`
	Data    string `bson:"data"`
	Status  int32  `bson:"status"`
}
type FBWhook struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Method      string        `bson:"method"`
	ContentType string        `bson:"contenttype"`
	URL         string        `bson:"url"`
	Created     time.Time     `bson:"created"`
	Data        string        `bson:"data"`
}
