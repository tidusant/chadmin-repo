package cuahang

import (
	"crypto/md5"
	"encoding/hex"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	"github.com/tidusant/chadmin-repo/models"

	"time"

	"gopkg.in/mgo.v2/bson"
)

/*
Check login status by session and USERIP, return userId and shopid
*/

func GetUserInfo(UserId string) models.User {
	col := db.Collection("addons_users")
	var rs models.User
	col.FindOne(ctx, bson.M{"_id": bson.ObjectIdHex(UserId)}).Decode(&rs)
	return rs
}

//get user login by session and return current shop and user id
func GetLogin(session string) models.UserLogin {
	coluserlogin := db.Collection("addons_userlogin")
	var rs models.UserLogin
	coluserlogin.FindOne(ctx, bson.M{"session": session}).Decode(&rs)
	if rs.ShopId == "" {
		rs.ShopId = GetShopDefault(rs.UserId.Hex())
		filter := bson.D{{"userid", rs.UserId}}
		update := bson.D{{"$set", bson.M{"shopid": rs.ShopId}}}
		coluserlogin.UpdateOne(ctx, filter, update)
	}
	return rs
}
func UpdateShopLogin(session, ShopChangeId string) (shopchange models.Shop) {
	coluserlogin := db.Collection("addons_userlogin")
	var rs models.UserLogin
	coluserlogin.FindOne(ctx, bson.M{"session": session}).Decode(&rs)
	if rs.UserId.Hex() == "" {
		return shopchange
	}
	//get shop id

	shopchange = GetShopById(rs.UserId.Hex(), ShopChangeId)
	if shopchange.ID.Hex() == "" {
		return shopchange
	}
	rs.ShopId = shopchange.ID.Hex()

	filter := bson.D{{"userid", rs.UserId}}
	update := bson.D{{"$set", bson.M{"shopid": rs.ShopId}}}
	coluserlogin.UpdateOne(ctx, filter, update)
	return shopchange
}

//Login user and update session
func Login(user, pass, session, userIP string) bool {
	hash := md5.Sum([]byte(pass))
	passmd5 := hex.EncodeToString(hash[:])
	coluser := db.Collection("addons_users")
	var result models.User
	coluser.FindOne(ctx, bson.M{"user": user, "password": passmd5}).Decode(&result)
	log.Debugf("user result %v", result)
	if result.Name != "" {
		coluserlogin := db.Collection("addons_userlogin")
		var userlogin models.UserLogin
		coluserlogin.FindOne(ctx, bson.M{"userid": result.ID}).Decode(&userlogin)
		if userlogin.UserId.Hex() == "" {
			userlogin.UserId = result.ID
		}
		userlogin.LastLogin = time.Now().UTC()
		userlogin.LoginIP = userIP
		userlogin.Session = session

		opts := options.Update().SetUpsert(true)
		filter := bson.D{{"userid", userlogin.UserId}}
		update := bson.D{{"$set", userlogin}}

		_, err := coluserlogin.UpdateOne(ctx, filter, update, opts)
		c3mcommon.CheckError("Upsert login", err)
		return true
	}
	return false
}
func Logout(session string) string {

	col := db.Collection("addons_userlogin")
	col.DeleteOne(ctx, bson.M{"session": session})

	return ""
}
