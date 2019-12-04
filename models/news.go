package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//News ...
type News struct {
	ID       bson.ObjectId        `bson:"_id,omitempty"`
	Code     string               `bson:"code"`
	UserID   string               `bson:"userid"`
	ShopID   string               `bson:"shopid"`
	CatID    string               `bson:"catid"`
	Langs    map[string]*NewsLang `bson:"langs"`
	Status   string               `bson:"status"`
	Created  time.Time            `bson:"created"`
	Modified time.Time            `bson:"modified"`
	Publish  bool                 `bson:"publish"`
}

//NewsLang ...
type NewsLang struct {
	Title       string `bson:"title"`
	Slug        string `bson:"slug"`
	Content     string `bson:"content"`
	Description string `bson:"description"`
	Avatar      string `bson:"avatar"`
	Viewed      int    `bson:"viewed"`
}

//NewsCat ...
type NewsCat struct {
	ID      bson.ObjectId           `bson:"_id,omitempty"`
	Code    string                  `bson:"code"`
	UserId  string                  `bson:"userid"`
	ShopId  string                  `bson:"shopid"`
	Created time.Time               `bson:"created"`
	Langs   map[string]*NewsCatLang `bson:"langs"`
}

//NewsCatLang ...
type NewsCatLang struct {
	Slug        string `bson:"slug"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Content     string `bson:"content"`
}
