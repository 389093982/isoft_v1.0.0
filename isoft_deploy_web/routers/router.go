package routers

import (
	"github.com/astaxie/beego"
	"isoft/isoft_deploy_web/controllers"
)

func init() {
	beego.Router("/index", &controllers.MainController{}, "get,post:Index")
	beego.Router("/env/list", &controllers.EnvController{}, "get:List;post:PostList")
	beego.Router("/env/edit", &controllers.EnvController{}, "post:PostEdit")
}
