package session

import (
	"time"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SaveWhook(whook models.Whook) string {

	col := db.C("addons_whook")
	whook.Created = int(time.Now().Unix())
	err := col.Insert(whook)
	c3mcommon.CheckError("whook insert", err)
	return "1"
}

func GetWhook() models.Whook {

	col := db.C("addons_whook")
	var bs models.Whook
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"status": 1, "modified": time.Now().Unix()}},
		ReturnNew: true,
	}
	_, err := col.Find(bson.M{"status": 0}).Apply(change, &bs)
	c3mcommon.CheckError("GetWhook script", err)
	return bs
}
