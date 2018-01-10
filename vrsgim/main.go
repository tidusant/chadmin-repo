package cuahang

import (
	"c3m/apps/common"
	"c3m/log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func init() {
	log.Infof("init repo vrsgim")
	strErr := ""
	db, strErr = common.ConnectDB("vrsgim")
	if strErr != "" {
		log.Infof(strErr)
		os.Exit(1)
	}
}
