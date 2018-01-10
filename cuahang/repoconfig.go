package cuahang

import (

	//	"c3m/log"

	//"strings"
	"c3m/apps/chadmin/models"
	"c3m/apps/common"

	"gopkg.in/mgo.v2/bson"
)

func GetTemplateConfigs(shopid, templatecode string) []models.TemplateConfig {
	col := db.C("addons_configs")
	var rs []models.TemplateConfig
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode}).All(&rs)
	common.CheckError("get template configs", err)
	return rs
}
func GetTemplateConfigByKey(shopid, templatecode, key string) models.TemplateConfig {
	col := db.C("addons_configs")
	var rs models.TemplateConfig
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode, "key": key}).One(&rs)
	common.CheckError("get template configs", err)
	return rs
}

func GetTemplateLang(shopid, templatecode, lang string) []models.TemplateLang {
	col := db.C("addons_langs")
	var rs []models.TemplateLang
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode, "lang": lang}).All(&rs)
	common.CheckError("get template langs", err)
	return rs
}

func SaveConfigs(newitem models.TemplateConfig) {
	col := db.C("addons_configs")

	if newitem.ID == "" {
		newitem.ID = bson.NewObjectId()
	}
	_, err := col.UpsertId(newitem.ID, &newitem)
	common.CheckError("save template configs", err)

}
