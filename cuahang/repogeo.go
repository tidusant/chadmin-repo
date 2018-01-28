package cuahang

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"
	"gopkg.in/mgo.v2/bson"
)

func GetCities() []models.City {
	col := db.C("geo_cities")
	var rs []models.City

	err := col.Find(bson.M{}).All(&rs)
	c3mcommon.CheckError("get cities", err)
	return rs
}
