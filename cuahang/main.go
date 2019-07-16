package cuahang

import (
	"encoding/json"
	"os"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	"github.com/tidusant/chadmin-repo/models"

	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init repo cuahang")
	strErr := ""
	db, strErr = c3mcommon.ConnectDB("chadmin")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
}

func CreateBuild(object, objectid, data string, usex models.UserSession) string {
	request := "createbuild|" + usex.Session
	var bs models.BuildScript
	bs.Object = object
	bs.ObjectId = objectid
	bs.Data = data
	bs.ShopConfigs = usex.Shop.Config
	bs.TemplateCode = usex.Shop.Theme
	bs.ShopId = usex.Shop.ID.Hex()
	b, _ := json.Marshal(bs)
	c3mcommon.RequestBuildServiceAsync(request, "POST", string(b))

	return ""
}

func CreateCommonDataBuild(usex models.UserSession) string {

	var bs models.BuildScript
	bs.ShopConfigs = usex.Shop.Config
	bs.TemplateCode = usex.Shop.Theme
	bs.ShopId = usex.Shop.ID.Hex()
	bs.Object = "common"
	var common models.CommonData
	//get allpage
	common.Pages = GetAllPage(usex.Shop.Theme, usex.Shop.ID.Hex())
	//remove page content
	for i, _ := range common.Pages {
		common.Pages[i].Blocks = nil
		common.Pages[i].Seo = ""
	}
	//get all product

	//get all productcat

	//get all news

	//get all newscat

	//get all albums

	//get all images
	b, _ := json.Marshal(common)
	bs.Data = string(b)
	request := "createbuild|" + usex.Session
	b, _ = json.Marshal(bs)
	c3mcommon.RequestBuildServiceAsync(request, "POST", string(b))

	return ""
}
