package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Shop struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Users       []ShopUser    `bson:"users"`
	Name        string        `bson:"name"`
	Phone       string        `bson:"phone"`
	Domain      string        `bson:"domain"`
	Description string        `bson:"description"`
	Status      int32         `bson:"status"`
	Created     time.Time     `bson:"created"`
	Albums      []ShopAlbum   `bson:"albums"`
	Theme       string        `bson:"theme"`
	Config      ShopConfigs   `bson:"config"`
}

type ShopConfigs struct {
	Title       string `bson:"title"`
	Description string `bson:"description"`
	Level       int    `bson:"level"`
	MaxUser     int    `bson:"maxuser"`
	MaxImage    int    `bson:"maximage"`
	MaxAlbum    int    `bson:"maxalbum"`
	MaxCat      int    `bson:"maxcat"`
	MaxProd     int    `bson:"maxprod"`
	MaxNews     int    `bson:"maxnews"`
	ShipFee     int    `bson:"shipfee"`
	FreeShip    int    `bson:"freeship"`
	Avatar      string `bson:"avatar"`
	FBPageId    string `bson:"fbpageid"`
	GHTKToken   string `bson:"ghtktoken"`
	GHTKWareID  string `bson:"ghtkwareid"`
	Tel         string `bson:"tel"`
	Address     string `bson:"address"`
	Province    string `bson:"province"`
	District    string `bson:"district"`
	Ward        string `bson:"ward"`

	Userdomain  bool   `bson:"userdomain"`
	Domain      string `bson:"domain"`
	Ftpusername string `bson:"ftpusername"`
	Ftppassword string `bson:"ftppassword"`

	Multilang   bool     `bson:"multilang"`
	Langs       []string `bson:"langs"`
	DefaultLang string   `bson:"defaultlang"`
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
