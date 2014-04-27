package admincontrollers

import (
	"blog/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Prepare() {
	isLogin := this.GetSession("ISLOGIN")

	if isLogin == nil {
		this.Redirect("/login", 302)
		return
	}
}

func (this *CategoryController) Index() {
	this.TplNames = "admin/archivers.html"
}

func (this *CategoryController) Get() {
	var err error
	this.Data["Categories"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}
	var cs []*models.Category
	cs, _ = models.GetCategories()
	for _, v := range cs {
		fmt.Println(time.Unix(v.Createtime, 0).Format("2006-01-02"))
	}

	this.TplNames = "admin/categories.html"
}

func (this *CategoryController) Create() {
	this.TplNames = "admin/category/add.html"
}

func (this *CategoryController) Store() {
	category := models.Category{}

	if err := this.ParseForm(&category); err != nil {
		beego.Error(err)
	}

	err := models.AddCategory(&category)
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "admin/category/add.html"
}

func (this *CategoryController) Edit() {

}

func (this *CategoryController) Update() {

}

func (this *CategoryController) Delete() {

}

func (this *CategoryController) List() {
	var err error
	this.Data["Categories"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}
	this.TplNames = "admin/category/categories.html"
}

/*func (this *CategoryController) Post() {
	m := models.Category{}

	if err := this.ParseForm(&m); err != nil {
		//handle error
		beego.Error(err)
	}

	fmt.Println(m)
	m := make(map[string]string)
	m["parentid"] = "0"
	m["catename"] = "dearlin"
	m["atriclenum"] = "1"

	err := models.AddCategory(m)
	if err != nil {
		beego.Error(err)
	}
}*/
