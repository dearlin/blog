package admincontrollers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	// 判断是否为退出操作
	if this.Input().Get("exit") == "true" {
		/*this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")*/
		this.DelSession("ISLOGIN")
		this.Redirect("/", 302)
		return
	}

	this.TplNames = "admin/login.html"
}

func (this *LoginController) Post() {
	// 获取表单信息
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	//autoLogin := this.Input().Get("autoLogin") == "on"

	// 验证用户名及密码
	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		/*maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")*/
		this.SetSession("ISLOGIN", int(1))
		this.Redirect("/admin", 302)
		return
	}

	this.Redirect("/login", 302)
	return
}
