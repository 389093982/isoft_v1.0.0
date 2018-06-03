package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"isoft/isoft/common"
	"encoding/json"
	"fmt"
	"isoft/isoft_deploy_web/models"
)

type EnvController struct {
	beego.Controller
}

func (this *EnvController) List()  {
	this.Layout = "layout/layout.html"
	this.TplName = "env/list.html"
}

func (this *EnvController) PostList()  {
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
