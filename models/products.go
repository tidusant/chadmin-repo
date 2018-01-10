package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	ID       bson.ObjectId           `bson:"_id,omitempty"`
	Code     string                  `bson:"code"`
	UserId   string                  `bson:"userid"`
	ShopId   string                  `bson:"shopid"`
	CatId    string                  `bson:"catid"`
	Langs    map[string]*ProductLang `bson:"langs"`
	Status   string                  `bson:"status"`
	Publish  bool                    `bson:"publish"`
	Created  time.Time               `bson:"created"`
	Modified time.Time               `bson:"modified"`
}

type ProductLang struct {
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	Price           int    `bson:"price"`
	BasePrice       int    `bson:"baseprice"`
	DiscountPrice   int    `bson:"discountprice"`
	PercentDiscount int    `bson:"percentdiscount"`
	Currency        string
	Description     string   `bson:"description"`
	Content         string   `bson:"content"`
	Avatar          string   `bson:"avatar"`
	Images          []string `bson:"images"`
	Viewed          int      `bson:"viewed"`
}

type ProdCatInfo struct {
	Slug        string `bson:"slug"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Content     string `bson:"content"`
	Avatar      string `bson:"avatar"`
}

type ProdCat struct {
	ID      bson.ObjectId           `bson:"_id,omitempty"`
	Code    string                  `bson:"code"`
	UserId  string                  `bson:"userid"`
	ShopId  string                  `bson:"shopid"`
	Created time.Time               `bson:"created"`
	Langs   map[string]*ProdCatInfo `bson:"langs"`
}
