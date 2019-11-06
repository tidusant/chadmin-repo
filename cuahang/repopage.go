package cuahang

import (
	"encoding/json"
	"time"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	//	"c3m/log"

	//"strings"

	"gopkg.in/mgo.v2/bson"
)

//=========================cat function==================
func SavePage(newitem models.Page) string {

	col := db.C("pages")

	// if prod.Code {

	// 	err := col.Insert(prod)
	// 	c3mcommon.CheckError("product Insert", err)
	// } else {

	//if len(newitem.Langs) > 0 {
	if newitem.ID == "" {
		return ""
	}
	if newitem.Created.Equal(time.Time{}) {
		newitem.Created = time.Now()
	}
	newitem.Modified = time.Now()
	//slug
	//get all slug
	// slugs := GetAllSlug(newitem.UserID, newitem.ShopID)
	// mapslugs := make(map[string]string)
	// for i := 0; i < len(slugs); i++ {
	// 	mapslugs[slugs[i]] = slugs[i]
	// }
	// for lang, _ := range newitem.Langs {
	// 	if newitem.Langs[lang].Title != "" {
	// 		//newslug
	// 		// tb, _ := lzjs.DecompressFromBase64(newitem.Langs[lang].Title)
	// 		// newslug := inflect.Parameterize(string(tb))
	// 		// log.Debugf("title: %s", string(tb))
	// 		// log.Debugf("newslug: %s", newslug)
	// 		// newitem.Langs[lang].Slug = newslug

	// 		// //check slug duplicate
	// 		// i := 1
	// 		// for {
	// 		// 	if _, ok := mapslugs[newitem.Langs[lang].Slug]; ok {
	// 		// 		newitem.Langs[lang].Slug = newslug + strconv.Itoa(i)
	// 		// 		i++
	// 		// 	} else {
	// 		// 		mapslugs[newitem.Langs[lang].Slug] = newitem.Langs[lang].Slug
	// 		// 		break
	// 		// 	}
	// 		// }
	// 		//remove oldslug
	// 		// log.Debugf("page slug for lang %s,%v", lang, newitem.Langs[lang])
	// 		// newitem.Langs[lang].Slug = newitem.Code
	// 		// CreateSlug(newitem.Langs[lang].Slug, newitem.ShopID, "page")
	// 	} else {
	// 		delete(newitem.Langs, lang)
	// 	}
	// }

	_, err := col.UpsertId(newitem.ID, &newitem)
	c3mcommon.CheckError("news Update", err)
	// } else {
	// 	col.RemoveId(newitem.ID)
	// }

	//}
	// for lang, _ := range newitem.Langs {
	// 	newitem.Langs[lang].Content = ""
	// }
	langinfo, _ := json.Marshal(newitem.Langs)
	return "{\"Code\":\"" + newitem.Code + "\",\"Langs\":" + string(langinfo) + "}"
}
func GetAllPage(templatecode, shopid string) []models.Page {
	col := db.C("pages")
	var rs []models.Page
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode}).All(&rs)
	c3mcommon.CheckError("get all page", err)
	return rs
}
func GetAllPageCode(templatecode, shopid string) []string {
	col := db.C("pages")
	var rs []struct {
		Text string `bson:"code"`
	}
	var rt []string
	err := col.Find(bson.M{"shopid": shopid, "templatecode": templatecode}).Select(bson.M{"code": 1}).All(&rs)
	if c3mcommon.CheckError("get all page name", err) {
		for _, v := range rs {
			rt = append(rt, v.Text)
		}

	}
	return rt
}
func GetPageByCode(templatecode, shopid, code string) models.Page {
	col := db.C("pages")
	var rs models.Page
	cond := bson.M{"shopid": shopid, "code": code, "templatecode": templatecode}

	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("GetPageByCode shopid:"+shopid+" code:"+code+" templatecode:"+templatecode, err)
	return rs
}
func InsertPage(item models.Page) {
	col := db.C("pages")
	//check  exist:
	cond := bson.M{"shopid": item.ShopID, "templatecode": item.TemplateCode, "code": item.Code}
	var oldrs models.Page
	col.Find(cond).One(&oldrs)
	if oldrs.ID.Hex() != "" {
		//get map block of old page
		mapblocks := make(map[string]models.PageBlock)
		for _, block := range oldrs.Blocks {
			mapblocks[block.Name] = block
		}

		for i, block := range item.Blocks {
			//check block exist:
			_, ok := mapblocks[block.Name]
			if ok {
				//get map item of old block
				mapbitems := make(map[string]models.PageBlockItem)
				for _, bitem := range mapblocks[block.Name].Items {
					mapbitems[bitem.Key] = bitem
				}

				for bi, bitem := range block.Items {
					//check if exist: keep old value
					if _, ok := mapbitems[bitem.Key]; ok {
						item.Blocks[i].Items[bi] = mapbitems[bitem.Key]
					}
				}
			}
		}
		err := col.UpdateId(oldrs.ID, item)
		c3mcommon.CheckError("Update Page "+item.Code+" template:"+item.TemplateCode, err)
		return
	}
	err := col.Insert(item)
	c3mcommon.CheckError("Insert Page "+item.Code+" template:"+item.TemplateCode, err)

}

func RemoveOldTemplatePage(shop models.Shop, template models.Template) {
	//remove old config
	colcfg := db.C("pages")
	cond := bson.M{"shopid": shop.ID.Hex(), "templatecode": template.Code}
	_, err := colcfg.RemoveAll(cond)
	c3mcommon.CheckError("remove old template page,shopid:"+shop.ID.Hex()+",templatecode:"+template.Code, err)
}
func RemoveUnusedTemplatePage(shop models.Shop, template models.Template, installedPages []string) {
	//remove old config
	colcfg := db.C("pages")
	cond := bson.M{"shopid": shop.ID.Hex(), "templatecode": template.Code, "code": bson.M{"$nin": installedPages}}
	_, err := colcfg.RemoveAll(cond)
	c3mcommon.CheckError("remove old template page,shopid:"+shop.ID.Hex()+",templatecode:"+template.Code, err)
}
