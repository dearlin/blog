package admincontrollers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Prepare() {
	isLogin := this.GetSession("ISLOGIN")

	if isLogin == nil {
		this.Redirect("/login", 302)
		return
	}
}

func (this *AdminController) Get() {
	this.TplNames = "admin/index.html"
}

func (this *AdminController) ShowMain() {
	this.TplNames = "admin/main.html"
}

func (this *AdminController) ShowTop() {
	this.TplNames = "admin/top.html"
}

func (this *AdminController) ShowMenu() {
	this.TplNames = "admin/menu.html"
}
