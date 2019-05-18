package cuahang

import (
	"time"

	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/chadmin-repo/models"

	"gopkg.in/mgo.v2/bson"
)

func CountOrderByEmployee(employeeid, shopid string) int {
	col := db.C("addons_orders")
	rs := 0
	cond := bson.M{"shopid": shopid, "employeeid": employeeid}
	rs, err := col.Find(cond).Count()
	c3mcommon.CheckError("count order cus by employeeid", err)
	return rs
}
func GetAllEmployee(shopid string) []models.Employee {
	col := db.C("addons_employee")
	var rs []models.Employee
	cond := bson.M{"shopid": shopid}
	err := col.Find(cond).All(&rs)
	c3mcommon.CheckError("GetAllEmployee", err)
	return rs
}
func GetEmployeeByPhone(phone, shopid string) models.Employee {
	col := db.C("addons_employee")
	var rs models.Employee
	cond := bson.M{"shopid": shopid, "phone": phone}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("get employee by phone", err)
	return rs
}
func GetEmployeeById(employeeid, shopid string) models.Employee {
	col := db.C("addons_employee")
	var rs models.Employee
	cond := bson.M{"shopid": shopid, "id": employeeid}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("get employee by phone", err)
	return rs
}
func SaveEmployee(cus models.Employee) models.Employee {
	col := db.C("addons_employee")
	cus.Modified = time.Now().UTC()
	if cus.ID == "" {
		cus.ID = bson.NewObjectId()
		cus.Created = cus.Modified
	}
	_, err := col.UpsertId(cus.ID, &cus)
	if c3mcommon.CheckError("save employee ", err) {
		cus.Name = ""
		return cus
	}
	return cus
}

func RemoveEmployee(employeeid string) bool {
	col := db.C("addons_employee")
	err := col.RemoveId(bson.ObjectIdHex(employeeid))
	if c3mcommon.CheckError("save employee ", err) {
		return true
	}
	return false
}

func GetLoginEmployee(phone, password string) models.Employee {
	col := db.C("addons_employee")
	var rs models.Employee
	cond := bson.M{"phone": phone, "password": password}
	err := col.Find(cond).One(&rs)
	c3mcommon.CheckError("checkLoginEmployee", err)
	return rs
}
