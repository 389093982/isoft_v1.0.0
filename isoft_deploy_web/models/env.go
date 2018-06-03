package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type EnvInfo struct {
	Id              int64       `json:"id"`
	EnvName         string    `json:"env_name"`
	EnvIp           string    `json:"env_ip"`
	EnvAccount      string    `json:"env_account"`
	EnvPasswd       string    `json:"env_passwd"`
	CreatedBy       string    `json:"created_by"`
	CreatedTime     time.Time `json:"created_time"`
	LastUpdatedBy   string    `json:"last_updated_by"`
	LastUpdatedTime time.Time `json:"last_updated_time"`
}

func CheckEnvInfoExists(condArr map[string]interface{}) (exists bool, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")

	if env_name, ok := condArr["env_name"]; ok {
		qs = qs.Filter("env_name", env_name)
	}
	if env_name, ok := condArr["env_ip"]; ok {
		qs = qs.Filter("env_ip", env_name)
	}
	counts, err := qs.Count()
	if err != nil{
		return false, err
	}
	return counts > 0, nil
}

func QueryEnvInfo(condArr map[string]interface{}, page int, offset int) (envInfos []EnvInfo, counts int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")
	var cond = orm.NewCondition()

	if _, ok := condArr["search"]; ok {
		subCond := orm.NewCondition()
		subCond = cond.And("env_name__contains", condArr["search"]).Or("env_ip__contains", condArr["search"])
		cond = cond.AndCond(subCond)
	}

	qs = qs.SetCond(cond)
	counts, _ = qs.Count()

	qs = qs.Limit(offset, (page-1)*offset)
	qs.All(&envInfos)
	return
}

func InsertOrUpdateEnvInfo(envInfo *EnvInfo) (id int64, err error) {
	o := orm.NewOrm()
	if envInfo.Id > 0 {
		id, err = o.Update(envInfo)
	} else {
		id, err = o.Insert(envInfo)
	}
	return
}

func QueryEnvInfoById(id int64) (envInfo EnvInfo, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("env_info")
	err = qs.Filter("id", id).One(&envInfo)
	return
}
