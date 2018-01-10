package cuahang

import (
	"c3m/apps/chadmin/models"
	"c3m/apps/common"
	"c3m/log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func GetOrdersReportByRange(shopid string, start time.Time, end time.Time) models.Campaign {
	col := db.C("addons_orders")
	var rs models.Campaign
	statusdetails := make(map[string]models.CampaignStatusDetail)
	cond := bson.M{"shopid": shopid}
	rpname := ""
	andcond := []bson.M{}
	if start.Year() > 1970 {
		andcond = append(andcond, bson.M{"created": bson.M{"$gt": start.Unix()}})
		rpname = start.Format("02/01/2006") + " - "
	} else {
		rpname = "Start - "
	}
	if end.Year() > 1970 {
		andcond = append(andcond, bson.M{"created": bson.M{"$lt": end.Unix()}})
		rpname += end.Format("02/01/2006")
	} else {
		rpname += "Now"
	}
	if len(andcond) > 0 {
		cond["$and"] = andcond
	}

	var ords []models.Order
	err := col.Find(cond).All(&ords)
	log.Debugf("start: %v\r\nend: %v\r\ncond: %v\r\nords: %v", start, end, cond, ords)
	stats := GetAllOrderStatus(shopid)
	statsmap := make(map[string]models.OrderStatus)
	for _, stat := range stats {
		statsmap[stat.ID.Hex()] = stat
	}
	if common.CheckError("GetOrdersReportByRange", err) {
		for _, ord := range ords {
			var detail models.CampaignStatusDetail
			if val, ok := statusdetails[ord.Status]; ok {
				detail = val
			}

			detail.Base += ord.BaseTotal
			detail.Total += ord.Total
			detail.PartnerShipFee += ord.PartnerShipFee
			detail.ShipFee += ord.ShipFee
			detail.Noo++
			statusdetails[ord.Status] = detail
			if statsmap[ord.Status].Finish {
				rs.Base += ord.BaseTotal
				rs.Total += ord.Total
				rs.PartnerShipFee += ord.PartnerShipFee
				rs.ShipFee += ord.ShipFee
				rs.Noo++
			}
		}
		rs.StatusDetail = statusdetails
	}
	rs.Name = rpname

	return rs
}
