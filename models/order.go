package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	ShopId         string        `bson:"shopid"`
	CampaignId     string        `bson:"campaignid"`
	ShipperId      string        `bson:"shipperid"`
	ShipmentCode   string        `bson:"shipmentcode"`
	Name           string
	Phone          string `bson:"phone"`
	OrderCount     int
	C              string `bson:"c"`
	City           string
	District       string
	Ward           string
	Address        string
	Note           string `bson:"note"`
	CusNote        string
	Email          string
	Status         string `bson:"status"`
	L              string `bson:"l"`
	Total          int    `bson:"total"`
	BaseTotal      int    `bson:"basetotal"`
	PartnerShipFee int    `bson:"partnershipfee"`
	ShipFee        int    `bson:"shipfee"`

	Items    []OrderItem `bson:"items"`
	Created  int64       `bson:"created"`
	Modified int64       `bson:"modified"`
}

type OrderItem struct {
	Code      string `bson:"code"`
	Title     string `bson:"title"`
	Avatar    string `bson:"avatar"`
	BasePrice int    `bson:"baseprice"`
	Price     int    `bson:"price"`
	Num       int    `bson:"num"`
}

type OrderStatus struct {
	ID            bson.ObjectId       `bson:"_id,omitempty"`
	Title         string              `bson:"title"`
	Default       bool                `bson:"default"`
	Finish        bool                `bson:"finish"`
	UserId        string              `bson:"userid"`
	ShopId        string              `bson:"shopid"`
	Created       time.Time           `bson:"created"`
	Modified      time.Time           `bson:"modified"`
	Color         string              `bson:"color"`
	PartnerStatus map[string][]string `bson:partnerstatus`
}
