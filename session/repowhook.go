package session

import (
	"time"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"
)

func SaveWhook(whook models.Whook) string {

	col := db.C("addons_whook")
	whook.Created = int(time.Now().Unix())
	err := col.Insert(whook)
	c3mcommon.CheckError("whook insert", err)
	return "1"
}
