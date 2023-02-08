package routers

import (
	"api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/:id", &controllers.UserController{}, "get:FindUserById")
	beego.Router("/user", &controllers.UserController{}, "get:GetPaginateUser")
	beego.Router("/user/update/:id", &controllers.UserController{}, "post:UpdateUser")
	beego.Router("/user/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/crawl", &controllers.CrawlController{}, "get:GetData")
}
