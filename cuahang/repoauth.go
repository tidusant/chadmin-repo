package cuahang

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	"github.com/tidusant/chadmin-repo/models"

	"time"

	"gopkg.in/mgo.v2/bson"
)

/*for authentication aaa
*sdf
=============================================================================
*/
func GetLogin(session, userIP string) string {
	coluserlogin := db.C("addons_userlogin")
	var rs models.UserLogin
	coluserlogin.Find(bson.M{"session": session}).One(&rs)
	log.Debugf("user aut find:%v", rs)
	userid := rs.UserId.Hex()
	if userid == "" {
		return ""
	}
	if rs.ShopId == "" {
		rs.ShopId = GetShopDefault(rs.UserId.Hex())
		coluserlogin.UpsertId(rs.UserId, &rs)
	}
	return rs.UserId.Hex() + "[+]" + rs.ShopId
}
func UpdateShopLogin(session, ShopId string) bool {
	coluserlogin := db.C("addons_userlogin")
	var rs models.UserLogin
	coluserlogin.Find(bson.M{"session": session}).One(&rs)
	if rs.UserId.Hex() == "" {
		return false
	}
	//get shop id
	shop := GetShopById(rs.UserId.Hex(), ShopId)
	if rs.ID.Hex() == "" {
		return false
	}
	rs.ShopId = shop.ID.Hex()
	coluserlogin.UpsertId(rs.UserId, &rs)
	return true
}
func Login(user, pass, session, userIP string) string {
	hash := md5.Sum([]byte(pass))
	passmd5 := hex.EncodeToString(hash[:])
	coluser := db.C("addons_users")
	var result models.User
	coluser.Find(bson.M{"user": user, "password": passmd5}).One(&result)
	log.Debugf("user result %v", result)
	if result.Name != "" {
		coluserlogin := db.C("addons_userlogin")
		var userlogin models.UserLogin
		coluserlogin.Find(bson.M{"userid": result.ID}).One(&userlogin)
		if userlogin.UserId.Hex() == "" {
			userlogin.UserId = result.ID
		}
		userlogin.LastLogin = time.Now().UTC()
		userlogin.LoginIP = userIP
		userlogin.Session = session

		_, err := coluserlogin.UpsertId(userlogin.UserId, &userlogin)
		c3mcommon.CheckError("Upsert login", err)
		return result.Name
	}
	return ""

}
