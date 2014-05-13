package admincontrollers

import (
	"blog/models"
	//"fmt"
	"github.com/astaxie/beego"
	//"time"
	//"strconv"
	"github.com/beego/wetalk/modules/utils"
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
	/*var err error
	this.Data["Categories"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}*/
	/*var cs []*models.Category
	cs, _ = models.GetCategories()
	for _, v := range cs {
		fmt.Println(time.Unix(v.Createtime, 0).Format("2006-01-02"))
	}*/

	//cate := models.CateTree{}

	this.TplNames = "admin/categories.html"
}

func (this *CategoryController) Create() {
	/*var cateTree map[string]*models.Category

	var err error

	parentid := int64(0)
	cateTree, err = models.GetCategoryTree(parentid)
	if err != nil {
		beego.Error(err)
	}
	for _, v := range cateTree {
		fmt.Println(v)
		fmt.Println(1)
	}*/

	//cate := &models.CateTree{}
	models.GetCateTree(0)
	//fmt.Println(cate)
	//models.ShowNode(cate, "")
	models.Show()

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
	id := this.Input().Get("id")

	var err error
	this.Data["Cate"], err = models.GetCategory(id)

	if err != nil {
		beego.Error(err)
	}

	this.TplNames = "admin/category/edit.html"
}

func (this *CategoryController) Update() {
	category := models.Category{}

	if err := this.ParseForm(&category); err != nil {
		beego.Error(err)
	}

	err := models.ModifyCategory(&category)

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/admin/category/list", 302)

	return
}

func (this *CategoryController) Delete() {
	id := this.Input().Get("id")

	var err error

	err = models.DelCategory(id)

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/admin/category/list", 302)

	return
}

func (this *CategoryController) List() {
	/*var err error
	this.Data["Categories"], err = models.GetCategories()
	if err != nil {
		beego.Error(err)
	}*/

	limit := 2
	qs := models.GetCates()
	nums, _ := qs.Count()

	pager := this.SetPaginator(limit, nums)
	var categories []*models.Category
	qs.Limit(limit, pager.Offset()).All(&categories)
	this.Data["Categories"] = categories

	this.TplNames = "admin/category/categories.html"
}

func (this *CategoryController) SetPaginator(per int, nums int64) *utils.Paginator {
	p := utils.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["Paginator"] = p
	return p
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
