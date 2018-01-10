package cuahang

import (
	"c3m/apps/chadmin/models"
	"c3m/apps/common"

	"gopkg.in/mgo.v2/bson"
)

/*for upload
=============================================================================
*/
func SaveImages(images []models.CHImage) string {
	col := db.C("files")

	if len(images) == 0 {
		return "0"
	}
	imgfiles := make([]interface{}, len(images))
	for i, image := range images {
		imgfiles[i] = image
	}
	bulk := col.Bulk()
	bulk.Unordered()
	bulk.Insert(imgfiles...)
	_, bulkErr := bulk.Run()
	common.CheckError("insert bulk", bulkErr)

	return "1"
}

func ImageCount(shopid string) int {
	col := db.C("files")
	count := -1
	count, err := col.Find(bson.M{"shopid": shopid, "appname": "chadmin"}).Count()
	common.CheckError("image count error", err)

	return count
}

func RemoveImage(shopid, userid, filename string) bool {
	col := db.C("files")
	cond := bson.M{"filename": filename}
	var image models.CHImage
	err := col.Find(cond).One(&image)
	if image.Shopid == shopid && image.Filename == filename {
		if userid == "1" || image.Uid == userid {
			err = col.Remove(bson.M{"filename": filename, "shopid": shopid})
			if common.CheckError("remove image", err) {
				return true
			}
		}
	}

	return false
}

func GetImages(shopid, userid, album string) []models.CHImage {
	col := db.C("files")
	var rs []models.CHImage
	cond := bson.M{}
	if userid == "1" {
		cond = bson.M{"shopid": shopid, "albumname": album, "appname": "chadmin"}

	} else {
		cond = bson.M{"shopid": shopid, "albumname": album, "appname": "chadmin"}
	}
	err := col.Find(cond).All(&rs)
	common.CheckError("error get images shopid:"+shopid+" userid:"+userid+" album:"+album, err)
	return rs
}
