package cuahang

import (
	"time"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	//	"c3m/log"

	//"strings"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

func SaveNews(newitem *models.News) string {

	col := db.C("news")

	// if prod.Code {

	// 	err := col.Insert(prod)
	// 	c3mcommon.CheckError("product Insert", err)
	// } else {
	if len(newitem.Langs) > 0 {
		if newitem.ID == "" {
			newitem.ID = bson.NewObjectId()
		}
		newitem.Modified = time.Now()
		_, err := col.UpsertId(newitem.ID, &newitem)
		c3mcommon.CheckError("news Update", err)
	} else {
		col.RemoveId(newitem.ID)
	}
	//}

	//remove content in response
	for lang, _ := range newitem.Langs {
		newitem.Langs[lang].Content = ""

	}
	langinfo, _ := json.Marshal(newitem.Langs)
	return "{\"Langs\":" + string(langinfo) + "}"
}
func RemoveNews(item models.News) bool {
	col := db.C("news")
	err := col.RemoveId(item.ID)
	return c3mcommon.CheckError("RemoveNews "+item.ID.Hex(), err)
}
func GetAllNews(shopid string) []models.News {
	col := db.C("news")
	var rs []models.News

	err := col.Find(bson.M{"shopid": shopid}).Sort("-_id").All(&rs)
	c3mcommon.CheckError("get all news", err)
	return rs
}
func GetDemoNews() []models.News {
	col := db.C("addons_news")
	var rs []models.News
	shop := GetDemoShop()
	err := col.Find(bson.M{"shopid": shop.ID.Hex()}).All(&rs)
	c3mcommon.CheckError("get all demo news", err)
	return rs
}

func GetNewsByID(userid, shopid, id string) models.News {
	col := db.C("news")

	var rs models.News
	if bson.IsObjectIdHex(id) {
		cond := bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(id)}
		err := col.Find(cond).One(&rs)
		c3mcommon.CheckError("GetNewsByID", err)
	}

	return rs

}
func GetNewsByCatId(userid, shopid, catcode string) []models.Product {
	col := db.C("news")
	var rs []models.Product

	err := col.Find(bson.M{"userid": userid, "shopid": shopid, "catid": catcode}).All(&rs)
	c3mcommon.CheckError("getprod", err)

	return rs

}

//=========================cat function==================
func SaveNewsCat(cat *models.NewsCat) string {
	col := db.C("newscategories")
	if len(cat.Langs) > 0 {
		if cat.ID == "" {
			cat.ID = bson.NewObjectId()
		}
		col.UpsertId(cat.ID, &cat)
	} else {
		col.RemoveId(cat.ID)
	}
	langinfo, _ := json.Marshal(cat.Langs)
	return "{\"Langs\":" + string(langinfo) + "}"
}
func GetDemoNewsCats() []models.NewsCat {
	col := db.C("addons_newscats")
	shop := GetDemoShop()
	var rs []models.NewsCat
	err := col.Find(bson.M{"shopid": shop.ID.Hex()}).All(&rs)
	c3mcommon.CheckError("getcatprod", err)
	return rs
}
func GetAllNewsCats(shopid string) []models.NewsCat {
	col := db.C("newscategories")
	var rs []models.NewsCat
	cond := bson.M{"shopid": shopid}
	err := col.Find(cond).Sort("-created").All(&rs)
	c3mcommon.CheckError("getcat ", err)
	return rs
}

func GetSubCatsByID(shopid, code string) []models.NewsCat {
	col := db.C("newscategories")
	var rs []models.NewsCat
	err := col.Find(bson.M{"shopid": shopid, "parentid": code}).All(&rs)
	c3mcommon.CheckError("GetSubCatsByID", err)
	return rs
}
func GetNewsCatByID(shopid, catid string) models.NewsCat {
	col := db.C("newscategories")
	var rs models.NewsCat
	err := col.Find(bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(catid)}).One(&rs)
	c3mcommon.CheckError("GetNewsCatByID", err)
	return rs
}
