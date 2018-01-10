package session

import (
	"c3m/apps/chadmin/models"
	"c3m/apps/common"
	"c3m/log"

	"c3m/common/mystring"
	"os"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init reposession")
	strErr := ""
	db, strErr = common.ConnectDB("session")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}

}
func CreateSession() string {
	sex := mystring.RandString(20)
	col := db.C("sessions")
	err := col.Insert(bson.M{"uid": sex, "created": time.Now().Unix(), "expired": time.Now().Unix() + 30*60})
	if common.CheckError("Insert sessions", err) {
		return sex
	}
	return "-1"
}
func CheckSession(s string) bool {
	if s == "" {
		return false
	}
	col := db.C("sessions")
	var result models.Session
	err2 := col.Find(bson.M{"uid": s}).One(&result)

	if err2 != nil {
		log.Infof("Session not found uid '%s': %s\n", s, err2)
	} else {
		if result.Expired > time.Now().Unix() {
			//update session
			cond := bson.M{"_id": result.ID}
			change := bson.M{"$set": bson.M{"expired": time.Now().Unix() + 30*60}}
			col.Update(cond, change)
			return true
		} else {
			//remove session
			col.RemoveId(result.ID)
			log.Infof("Session expired: uid '%s'", s)
			return false
		}

	}

	return false
}
func CheckRequest(uri, useragent, referrer, remoteAddress, requestType string) bool {

	col := db.C("requestUrls")
	log.Printf("now: %d , check: %d", int(time.Now().Unix()), int(time.Now().Unix())-10)
	urlcount := -1
	var err error
	if requestType == "POST" {
		urlcount, err = col.Find(bson.M{"uri": uri}).Count()
	} else {
		urlcount, err = col.Find(bson.M{"uri": uri, "created": bson.M{"$gt": int(time.Now().Unix()) - 1}}).Count()
		if urlcount < 50 {
			log.Debugf("same url count %d", urlcount)
			urlcount = 0
		} else {
			log.Debugf("request limited %s", uri)
		}
	}

	if common.CheckError("checkRequest", err) {
		if urlcount == 0 {
			//check ip in 3 sec
			urlcount, err := col.Find(bson.M{"remoteAddress": remoteAddress, "created": bson.M{"$gt": int(time.Now().Unix()) - 1}}).Count()
			if urlcount < 500 {
				err = col.Insert(bson.M{"uri": uri, "created": int(time.Now().Unix()), "user-agent": useragent, "referer": referrer, "remoteAddress": remoteAddress})
				common.CheckError("checkRequest Insert", err)
				return true
			} else {
				log.Debugf("request ip limited %s", uri)
			}

		} else {
			log.Debugf("request limited 2 %s", uri)
		}
	}

	return false
}
