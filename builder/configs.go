package builder

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"
	"gopkg.in/mgo.v2/bson"
)

func GetConfig(shopid string) models.BuildConfig {
	col := db.C("configs")
	var bs models.BuildConfig

	err := col.Find(bson.M{"shopid": shopid}).One(&bs)
	c3mcommon.CheckError("GetConfigs", err)
	return bs

}
