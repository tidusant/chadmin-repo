package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//News ...
type News struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Code string        `bson:"code"`

	UserID   string               `bson:"userid"`
	ShopID   string               `bson:"shopid"`
	CatID    string               `bson:"catid"`
	Langs    map[string]*PageLang `bson:"langs"`
	Status   string               `bson:"status"`
	Created  time.Time            `bson:"created"`
	Modified time.Time            `bson:"modified"`
	Publish  bool                 `bson:"publish"`
	Home     bool                 `bson:"home"`
	Feature  bool                 `bson:"feature"`

	LangLinks []LangLink `bson:"langlinks"`
}

//NewsCat ...
type NewsCat struct {
	ID        bson.ObjectId        `bson:"_id,omitempty"`
	Code      string               `bson:"code"`
	UserId    string               `bson:"userid"`
	ShopId    string               `bson:"shopid"`
	Created   time.Time            `bson:"created"`
	Langs     map[string]*PageLang `bson:"langs"`
	LangLinks []LangLink           `bson:"langlinks"`
}
