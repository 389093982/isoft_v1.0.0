package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common"
	"isoft/isoft_deploy_web/models"
	"time"
)

type EnvController struct {
	beego.Controller
}

func (this *EnvController) List() {
	this.Layout = "layout/layout.html"
	this.TplName = "env/list.html"
}

func (this *EnvController) PostList() {
	condArr := make(map[string]interface{})
	offset, _ := this.GetInt("offset", 10)            // 每页记录数
	current_page, _ := this.GetInt("current_page", 1) // 当前页
	search_text := this.GetString("search_text")

	if search_text != "" {
		condArr["search_text"] = search_text
	}

	envInfos, count, err := models.QueryEnvInfo(condArr, current_page, offset)
	paginator := pagination.SetPaginator(this.Ctx, offset, count)

	//初始化
	data := make(map[string]interface{}, 1)

	if err == nil {
		data["envInfos"] = envInfos
		data["paginator"] = common.Paginator(paginator.Page(), paginator.PerPageNums, paginator.Nums())
	}
	//序列化
	json_obj, err := json.Marshal(data)
	if err == nil {
		this.Data["json"] = string(json_obj)
	} else {
		fmt.Print(err.Error())
	}
	this.ServeJSON()
}

func (this *EnvController) PostEdit() {
	id, err:= this.GetInt64("id", -1)
	env_name := this.GetString("env_name")
	env_ip := this.GetString("env_ip")
	env_account := this.GetString("env_account")
	env_passwd := this.GetString("env_passwd")

	// 更新操作
	if err == nil && id > 0{
		envInfo, err := models.QueryEnvInfoById(id)
		if err != nil {
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
		}else{
			envInfo.EnvName = env_name
			envInfo.EnvIp = env_ip
			envInfo.EnvAccount = env_account
			envInfo.EnvPasswd = env_passwd
			_, err = models.InsertOrUpdateEnvInfo(&envInfo)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
			} else {
				this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
			}
		}
	}else{
		// 插入操作
		envInfo := models.EnvInfo{
			EnvName:         env_name,
			EnvIp:           env_ip,
			EnvAccount:      env_account,
			EnvPasswd:       env_passwd,
			CreatedBy:       "AutoInsert",
			CreatedTime:     time.Now(),
			LastUpdatedBy:   "AutoInsert",
			LastUpdatedTime: time.Now(),
		}
		// 根据联合条件 env_name 和 env_ip 判断是否存在
		condArr := make(map[string]interface{})
		condArr["env_name"] = env_name
		condArr["env_ip"] = env_ip
		exist, err := models.CheckEnvInfoExists(condArr)
		if exist == false{
			// 不存在则插入
			_, err = models.InsertOrUpdateEnvInfo(&envInfo)
			if err == nil {
				this.Data["json"] = &map[string]interface{}{"status": "SUCCESS"}
			} else {
				this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "保存失败！"}
			}
		}else{
			// 已经存在
			this.Data["json"] = &map[string]interface{}{"status": "ERROR", "errorMsg": "env_name 和 env_ip 已经存在！"}
		}
	}
	this.ServeJSON()
}
