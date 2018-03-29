package cuahang

import (
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
