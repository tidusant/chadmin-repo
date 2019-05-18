package cuahang

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	//	"c3m/log"

	//"strings"

	"gopkg.in/mgo.v2/bson"
)

func GetPosClient(userid, shopid, posname string) models.PosClient {

	col := db.C("addons_posclient")

	var rs models.PosClient
	err := col.Find(bson.M{"name": posname, "userid": userid, "shopid": shopid}).One(&rs)
	c3mcommon.CheckError("GetPosClient", err)
	return rs
}
func UpdateSyncPosClient(userid, shopid, posname string) {
	col := db.C("addons_posclient")

	cond := bson.M{"userid": userid, "shopid": shopid, "name": posname}
	change := bson.M{"issync": true}
	_, err := col.UpdateAll(cond, bson.M{"$set": change})
	c3mcommon.CheckError("UpdateSyncPosClient", err)

}
func SavePosClient(pos models.PosClient) {
	col := db.C("addons_posclient")
	//remove all old data
	col.RemoveAll(bson.M{})
	err := col.Insert(pos)
	c3mcommon.CheckError("SavePosClient", err)

}
func GetSyncProds(userid, shopid string) []models.Product {
	col := db.C("addons_products")
	var rs []models.Product

	err := col.Find(bson.M{"shopid": shopid, "issync": false}).All(&rs)
	c3mcommon.CheckError("GetSyncProds", err)
	return rs
}
func UpdateSyncProds(userid, shopid string, emplids []string) {
	col := db.C("addons_products")
	var arrEmplID []bson.ObjectId
	for _, eid := range emplids {
		arrEmplID = append(arrEmplID, bson.ObjectIdHex(eid))
	}
	cond := bson.M{"_id": bson.M{"$in": arrEmplID}}
	change := bson.M{"issync": true}
	_, err := col.UpdateAll(cond, bson.M{"$set": change})
	c3mcommon.CheckError("UpdateSyncProds", err)

}
func GetSyncCats(userid, shopid string) []models.ProdCat {
	col := db.C("addons_prodcats")
	var rs []models.ProdCat

	err := col.Find(bson.M{"shopid": shopid, "issync": false}).All(&rs)
	c3mcommon.CheckError("GetSyncCats", err)
	return rs
}
func UpdateSyncCats(userid, shopid string, emplids []string) {
	col := db.C("addons_prodcats")
	var arrEmplID []bson.ObjectId
	for _, eid := range emplids {
		arrEmplID = append(arrEmplID, bson.ObjectIdHex(eid))
	}
	cond := bson.M{"_id": bson.M{"$in": arrEmplID}}
	change := bson.M{"issync": true}
	_, err := col.UpdateAll(cond, bson.M{"$set": change})
	c3mcommon.CheckError("UpdateSyncCats", err)

}
func GetSyncEmployees(userid, shopid string) []models.Employee {
	col := db.C("addons_employee")
	var rs []models.Employee

	err := col.Find(bson.M{"shopid": shopid, "issync": false}).All(&rs)
	c3mcommon.CheckError("GetSyncEmployees", err)
	return rs
}
func UpdateSyncEmployees(userid, shopid string, emplids []string) {
	col := db.C("addons_employee")
	var arrEmplID []bson.ObjectId
	for _, eid := range emplids {
		arrEmplID = append(arrEmplID, bson.ObjectIdHex(eid))
	}
	cond := bson.M{"_id": bson.M{"$in": arrEmplID}}
	change := bson.M{"issync": true}
	_, err := col.UpdateAll(cond, bson.M{"$set": change})
	c3mcommon.CheckError("UpdateSyncEmployees", err)

}
