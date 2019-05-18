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

func SaveConfigs(newitem models.TemplateConfig) {
	col := db.C("addons_configs")

	if newitem.ID == "" {
		newitem.ID = bson.NewObjectId()
	}
	_, err := col.UpsertId(newitem.ID, &newitem)
	c3mcommon.CheckError("save template configs", err)

}
