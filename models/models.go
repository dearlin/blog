package models

import (
	//"github.com/astaxie/beego"
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"time"
)

// category
type Category struct {
	Id         int64
	Parentid   int64
	Catename   string
	Articlenum int64
	Createtime int64
	Updatetime int64
}

func init() {
	orm.RegisterModel(new(Category))
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8")
}

func GetCatetory() (category *Category, error) {
	o:=orm.NewOrm()
	
}

func AddCategory(category *Category) error {
	o := orm.NewOrm()
	o.Using("default")

	/*parentId, err := strconv.Atoi(params["parentid"])
	articleNum, err := strconv.Atoi(params["articlenum"])*/

	category.Createtime = int64(time.Now().Unix())

	_, err := o.Insert(category)
	if err != nil {
		return err
	}

	return nil
}

func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
