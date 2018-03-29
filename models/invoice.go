package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Invoice struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	ShopId      string        `bson:"shopid"`
	UserId      string        `bson:"userid"`
	Description string        `bson:"description"`
	Images      []string      `bson:"images"`
	Items       []InvoiceItem `bson:"items"`
	Created     int64         `bson:"created"`
	Modified    int64         `bson:"modified"`
	Total       int           `bson:"total"`
}

type InvoiceItem struct {
	ProductName  string `json:"prodname"`
	ProductCode  string `json:"prodcode"`
	PropertyName string `json:"propname"`
	PropertyCode string `json:"propcode"`
	Stock        int    `json:"propstock"`
	BasePrice    int    `json:"propbaseprice"`
}
