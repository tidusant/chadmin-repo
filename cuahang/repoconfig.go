package cuahang

import (

	//	"c3m/log"

	//"strings"
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	"gopkg.in/mgo.v2/bson"
)

func GetTemplateLang(shopid, templatecode, lang string) []models.TemplateLang {
	col := db.C("addons_langs")
	var rs []models.TemplateLang
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode, "lang": lang}).All(&rs)
	c3mcommon.CheckError("get template langs", err)
	return rs
}

func SaveShopConfig(shop models.Shop) {
	col := db.C("addons_shops")
	//check  exist:
	cond := bson.M{"_id": shop.ID}
	change := bson.M{"$set": bson.M{"config": shop.Config}}
	err := col.Update(cond, change)

	c3mcommon.CheckError("SaveShopConfig :", err)

}
