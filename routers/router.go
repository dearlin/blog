package routers

import (
	"blog/controllers"
	"blog/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// admin
	beego.Router("/login", &admincontrollers.LoginController{})
	beego.Router("/admin", &admincontrollers.AdminController{})

	beego.Router("/admin/main", &admincontrollers.AdminController{}, "get:ShowMain")
	beego.Router("/admin/top", &admincontrollers.AdminController{}, "get:ShowTop")
	beego.Router("/admin/menu", &admincontrollers.AdminController{}, "get:ShowMenu")

	// category
	beego.Router("/admin/category", &admincontrollers.CategoryController{})
	beego.Router("/admin/category/create", &admincontrollers.CategoryController{}, "get:Create")
	beego.Router("/admin/category/store", &admincontrollers.CategoryController{}, "post:Store")
	beego.Router("/admin/category/edit", &admincontrollers.CategoryController{}, "get:Edit")
	beego.Router("/admin/category/update", &admincontrollers.CategoryController{}, "post:Update")
	beego.Router("/admin/category/list", &admincontrollers.CategoryController{}, "get:List")
	beego.Router("/admin/category/delete", &admincontrollers.CategoryController{}, "*:Delete")
}
