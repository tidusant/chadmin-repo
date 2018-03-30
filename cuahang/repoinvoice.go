package cuahang

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	"github.com/tidusant/chadmin-repo/models"
	//	"c3m/log"

	//"strings"

	"gopkg.in/mgo.v2/bson"
)

func SaveInvoice(invc models.Invoice) models.Invoice {

	col := db.C("addons_invoice")

	if invc.ID == "" {
		invc.ID = bson.NewObjectId()
	}
	col.UpsertId(invc.ID, &invc)
	return invc
}

func GetInvoices(shopid string, imp bool) []models.Invoice {

	col := db.C("addons_invoice")
	log.Debugf("import:%v", imp)
	var rs []models.Invoice
	err := col.Find(bson.M{"shopid": shopid, "import": imp}).All(&rs)
	c3mcommon.CheckError("GetInvoices", err)
	return rs
}
