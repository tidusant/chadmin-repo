package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Shop struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	Users   []ShopUser    `bson:"users"`
	Name    string        `bson:"name"`
	Phone   string        `bson:"phone"`
	Created time.Time     `bson:"created"`
	Albums  []ShopAlbum   `bson:"albums"`
	Theme   string        `bson:"theme"`
	Config  ShopConfigs   `bson:"config"`
}

type ShopConfigs struct {
	Title       string   `bson:"title"`
	Description string   `bson:"description"`
	Avatar      string   `bson:"avatar"`
	Multilang   bool     `bson:"multilang"`
	Langs       []string `bson:"langs"`
	Defaultlang string   `bson:"defaultlang"`
	CurrentLang string   `bson:"currentlang"`
}

type ShopUser struct {
	Id    string `bson:"userid"`
	Level string `bson:"level"`
}
type ShopAlbum struct {
	Slug    string    `bson:"slug"`
	Name    string    `bson:"name"`
	UserId  string    `bson:"userid"`
	Created time.Time `bson:"created"`
}
