package template

import (
	"c3m/apps/common"

	"c3m/apps/tempman/portal/models"
	"c3m/log"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init repo template")
	strErr := ""
	db, strErr = common.ConnectDB("chtemplate")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
}

func ActiveTemplate(userid, shopid string, template, oldtemplate models.Template) string {

	if userid == "" || shopid == "" {
		return ""
	}
	col := db.C("templates")

	cond := bson.M{"status": 1, "code": template.Code}
	change := bson.M{"$set": bson.M{"activedid": template.ActiveIDs}}
	col.Update(cond, change)

	cond = bson.M{"status": 1, "code": oldtemplate.Code}
	change = bson.M{"$set": bson.M{"activedid": oldtemplate.ActiveIDs}}
	col.Update(cond, change)

	return template.Code

}

func InstallTemplate(userid, shopid string, template models.Template) string {

	if userid == "" || shopid == "" {
		return ""
	}
	col := db.C("templates")

	cond := bson.M{"status": 1, "code": template.Code}
	change := bson.M{"$set": bson.M{"installedid": template.InstalledIDs}}
	err := col.Update(cond, change)
	if common.CheckError("install template", err) {
		return template.Code
	}
	return ""
}

func GetAllTemplates(userid, shopid string) []models.Template {
	col := db.C("templates")
	var rs []models.Template
	err := col.Find(bson.M{"status": 1, "installedid": bson.M{"$ne": shopid}}).Select(bson.M{"code": 1, "title": 1, "screenshot": 1}).All(&rs)
	common.CheckError("get all templates", err)
	return rs
}

func GetTemplateByCode(userid, shopid, code string) models.Template {
	var rs models.Template
	if userid == "" || shopid == "" {
		return rs
	}
	col := db.C("templates")
	err := col.Find(bson.M{"status": 1, "code": code}).One(&rs)
	common.CheckError("get template", err)
	return rs
}

func GetAllTemplatesInstalled(userid, shopid string) []models.Template {
	col := db.C("templates")
	var rs []models.Template
	err := col.Find(bson.M{"status": 1, "installedid": shopid}).Select(bson.M{"code": 1, "title": 1, "screenshot": 1}).All(&rs)
	common.CheckError("get all templates", err)
	return rs
}
