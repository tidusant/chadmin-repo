package models

import (
	"gopkg.in/mgo.v2/bson"
)

type BuildScript struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Data     string        `bson:"data"`
	IsRemove bool          `bson:"isremove"`
	Status   int           `bson:"status"` //0: news, 1: building, 2: finish
	Created  int64         `bson:"created"`
	Modified int64         `bson:"modified"`
	Retry    int           `bson:"retry"`

	ObjectId     string `bson:"objectid"`
	Object       string `bson:"object"`
	ShopId       string `bson:"shopid"`
	TemplateCode string `bson:"templatecode`
	Domain       string `bson:"domain"`
}

type BuildConfig struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	ShopId      string        `bson:"shopid"`
	Domain      string        `bson:"domain"`
	Host        string        `bson:"host"` //0: news, 1: building, 2: finish
	Path        string        `bson:"path"`
	FTPUsername string        `bson:"ftpusername"`
	FTPPassword string        `bson:"ftppassword"`
}
