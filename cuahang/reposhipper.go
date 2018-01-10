package cuahang

import (
	"c3m/apps/chadmin/models"
	"c3m/apps/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetAllShipper(shopid string) []models.Shipper {
	col := db.C("addons_shippers")
	var rs []models.Shipper
	cond := bson.M{"shopid": shopid}
	err := col.Find(cond).All(&rs)
	common.CheckError("GetAllShipper", err)
	return rs
}
func GetShipperByID(itemid, shopid string) models.Shipper {
	col := db.C("addons_shippers")
	var rs models.Shipper
	cond := bson.M{"shopid": shopid, "_id": bson.ObjectIdHex(itemid)}
	err := col.Find(cond).One(&rs)
	common.CheckError("GetShipperByID", err)
	return rs
}

func GetDefaultShipper(shopid string) models.Shipper {
	col := db.C("addons_shippers")
	var rs models.Shipper
	cond := bson.M{"shopid": shopid, "default": true}
	err := col.Find(cond).One(&rs)
	common.CheckError("GetDefaultShipper", err)
	return rs
}

func SaveShipper(shipper models.Shipper) models.Shipper {
	col := db.C("addons_shippers")
	if shipper.ID == "" {
		shipper.ID = bson.NewObjectId()
		shipper.Created = time.Now().UTC()
	}

	shipper.Modified = shipper.Created
	_, err := col.UpsertId(shipper.ID, shipper)
	common.CheckError("SaveShipper", err)
	return shipper
}

func GetCountOrderByShipper(shipper models.Shipper) int {
	col := db.C("addons_orders")

	cond := bson.M{"shopid": shipper.ShopId, "shipperid": shipper.ID.Hex()}
	n, err := col.Find(cond).Count()
	common.CheckError("GetCountOrderByShipper", err)
	return n
}

func DeleteShipper(shipper models.Shipper) bool {
	col := db.C("addons_shippers")

	cond := bson.M{"shopid": shipper.ShopId, "_id": shipper.ID}
	err := col.Remove(cond)
	return common.CheckError("DeleteShipper", err)

}

func UnSetShipperDefault(shopid string) {
	col := db.C("addons_shippers")

	cond := bson.M{"shopid": shopid, "default": true}
	change := bson.M{"$set": bson.M{"default": false}}
	err := col.Update(cond, change)
	common.CheckError("UnSetShipperDefault", err)

}
