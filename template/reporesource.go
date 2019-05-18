package template

import (
	"encoding/json"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	//	"c3m/log"

	//"strings"

	"gopkg.in/mgo.v2/bson"
)

//=========================cat function==================
func SaveResource(newitem models.Resource) string {

	col := db.C("resources")

	// if prod.Code {

	// 	err := col.Insert(prod)
	// 	c3mcommon.CheckError("product Insert", err)
	// } else {

	//if len(newitem.Langs) > 0 {
	if newitem.ID == "" {
		return ""
	}

	_, err := col.UpsertId(newitem.ID, &newitem)
	c3mcommon.CheckError("news Update", err)
	// } else {
	// 	col.RemoveId(newitem.ID)
	// }

	//}
	// for lang, _ := range newitem.Langs {
	// 	newitem.Langs[lang].Content = ""
	// }
	langinfo, _ := json.Marshal(newitem.Value)
	return "{\"Code\":\"" + newitem.Key + "\",\"Value\":" + string(langinfo) + "}"
}
func GetAllResource(templatecode, shopid string) []models.Resource {
	col := db.C("resources")
	var rs []models.Resource
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode}).All(&rs)
	c3mcommon.CheckError("GetAllResource", err)
	return rs
}
func GetResourceByKey(templatecode, shopid, key string) models.Resource {
	col := db.C("resources")
	var rs models.Resource
	cond := bson.M{"shopid": shopid, "key": key, "templatecode": templatecode}

	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetResourceByKey", err)
	return rs
}
