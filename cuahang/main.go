package cuahang

import (
	"c3m/apps/common"
	"c3m/log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init repo cuahang")
	strErr := ""
	db, strErr = common.ConnectDB("chadmin")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
}

//==============slug=================
func RemoveSlug(slug, shopid string) bool {
	col := db.C("addons_slugs")
	col.Remove(bson.M{"shopid": shopid, "slug": slug})
	return true
}
func CreateSlug(slug, shopid, object string) bool {
	col := db.C("addons_slugs")
	col.Insert(bson.M{"shopid": shopid, "slug": slug, "object": object})
	return true
}
func GetAllSlug(userid, shopid string) []string {
	col := db.C("addons_slugs")
	var rs []string
	cond := bson.M{"shopid": shopid}
	if userid != "594f665df54c58a2udfl54d3er" {
		cond["userid"] = userid
	}
	err := col.Find(cond).Select(bson.M{"slug": 1}).All(&rs)
	common.CheckError("getallslug", err)
	return rs
}
