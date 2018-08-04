package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Page ...
type Page struct {
	ID          bson.ObjectId        `bson:"_id,omitempty"`
	Code        string               `bson:"code"`
	UserID      string               `bson:"userid"`
	ShopID      string               `bson:"shopid"`
	Langs       map[string]*PageLang `bson:"langs"`
	Created     time.Time            `bson:"created"`
	Modified    time.Time            `bson:"modified"`
	Publish     bool                 `bson:"publish"`
	AltPagename string
}

//NewsLang ...
type PageLang struct {
	Title       string `bson:"title"`
	Slug        string `bson:"slug"`
	Content     string `bson:"content"`
	Description string `bson:"description"`
	Avatar      string `bson:"avatar"`
	Viewed      int    `bson:"viewed"`
}
