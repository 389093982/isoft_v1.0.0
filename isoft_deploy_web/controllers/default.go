package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (m *MainController) Index()  {
	m.Data["Website"] = "beego.me"
	m.Data["Email"] = "astaxie@gmail.com"
	m.Layout = "layout/layout.html"
	m.TplName = "index.html"
}