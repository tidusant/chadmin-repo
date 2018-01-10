package cuahang

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"
	//	"c3m/log"

	//"strings"
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

func SaveProd(prod models.Product) string {

	col := db.C("addons_products")

	// if prod.Code {

	// 	err := col.Insert(prod)
	// 	c3mcommon.CheckError("product Insert", err)
	// } else {
	if len(prod.Langs) > 0 {
		if prod.ID == "" {
			prod.ID = bson.NewObjectId()
		}
		_, err := col.UpsertId(prod.ID, &prod)
		c3mcommon.CheckError("product Update", err)
	} else {
		col.RemoveId(prod.ID)
	}
	//}
	langinfo, _ := json.Marshal(prod.Langs)
	return "{\"Code\":\"" + prod.Code + "\",\"Langs\":" + string(langinfo) + "}"
}
func GetAllProds(userid, shopid string) []models.Product {
	col := db.C("addons_products")
	var rs []models.Product
	shop := GetShopById(userid, shopid)
	err := col.Find(bson.M{"shopid": shop.ID.Hex()}).All(&rs)
	c3mcommon.CheckError("getprod", err)
	return rs
}
func GetDemoProds() []models.Product {
	col := db.C("addons_products")
	var rs []models.Product
	shop := GetDemoShop()
	err := col.Find(bson.M{"shopid": shop.ID.Hex()}).All(&rs)
	c3mcommon.CheckError("get demo prod", err)
	return rs
}
func GetProdBySlug(userid, shopid, slug string) models.Product {
	col := db.C("addons_products")
	var rs models.Product
	cond := bson.M{"shopid": shopid, "slug": slug}
	if userid != "594f665df545c58a2builder" {
		cond["userid"] = userid
	}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("getprod", err)
	return rs
}
func GetProdByCode(userid, shopid, code string) models.Product {
	col := db.C("addons_products")
	var rs models.Product
	cond := bson.M{"shopid": shopid, "code": code}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("getprod", err)
	return rs
}

func GetProdsByCatId(userid, shopid, catcode string) []models.Product {
	col := db.C("addons_products")
	var rs []models.Product
	cond := bson.M{"shopid": shopid, "catid": catcode}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("getprod", err)

	return rs

}

//=========================cat function==================
func SaveCat(cat models.ProdCat) string {
	col := db.C("addons_prodcats")
	if len(cat.Langs) > 0 {
		if cat.ID == "" {
			cat.ID = bson.NewObjectId()
			//save slug
		} else {
			//update slug
		}

		col.UpsertId(cat.ID, cat)
	} else {
		col.RemoveId(cat.ID)
	}
	langinfo, _ := json.Marshal(cat.Langs)
	return "{\"Code\":\"" + cat.Code + "\",\"Langs\":" + string(langinfo) + "}"
}
func GetAllCats(userid, shopid string) []models.ProdCat {
	col := db.C("addons_prodcats")
	var rs []models.ProdCat
	cond := bson.M{"shopid": shopid}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("getcatprod", err)
	return rs
}

func GetDemoProdCats() []models.ProdCat {
	col := db.C("addons_prodcats")
	shop := GetDemoShop()
	var rs []models.ProdCat
	err := col.Find(bson.M{"shopid": shop.ID.Hex()}).All(&rs)
	c3mcommon.CheckError("getcatprod", err)
	return rs
}
func GetCatByCode(userid, shopid, code string) models.ProdCat {
	col := db.C("addons_prodcats")
	var rs models.ProdCat
	cond := bson.M{"shopid": shopid, "code": code}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("getcatbycode", err)
	return rs
}
