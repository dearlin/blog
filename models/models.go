package models

import (
	//"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
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
	Haschild   int64
}

//var CategoryTree map[string]*Category

type CateTree struct {
	id       int64
	cate     Category
	children []*CateTree
}

var (
	nodeTable = map[int64]*CateTree{}
	root      *CateTree
)

func init() {
	orm.RegisterModel(new(Category))
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/blog?charset=utf8")
}

func GetCategory(id string) (*Category, error) {
	cid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()

	category := new(Category)

	qs := o.QueryTable("category")
	err = qs.Filter("id", cid).One(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func AddCategory(category *Category) error {
	o := orm.NewOrm()

	/*parentId, err := strconv.Atoi(params["parentid"])
	articleNum, err := strconv.Atoi(params["articlenum"])*/

	category.Createtime = int64(time.Now().Unix())

	_, err := o.Insert(category)

	return err
}

func ModifyCategory(cate *Category) error {
	o := orm.NewOrm()

	_, err := o.QueryTable("Category").Filter("id", cate.Id).Update(orm.Params{
		"catename":   cate.Catename,
		"parentid":   cate.Parentid,
		"updatetime": int64(time.Now().Unix()),
	})

	return err
}

func DelCategory(id string) error {
	o := orm.NewOrm()

	cid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return err
	}

	_, err = o.QueryTable("category").Filter("id", cid).Delete()

	return err
}

func ShowNode(categ *CateTree, prefix string) {
	if prefix == "" {
		fmt.Printf("%v\n\n", categ.cate.Catename)
	} else {
		fmt.Printf("%v %v\n\n", prefix, categ.cate.Catename)
	}
	for _, n := range categ.children {
		ShowNode(n, prefix+"--")
	}
}

func Show() {
	if root == nil {
		//	fmt.Printf("show: root node not found\n")
		return
	}
	//fmt.Printf("RESULT:\n")
	ShowNode(root, "")
}

func AddChildTree(id int64, cate Category, parentId int64) {
	node := &CateTree{id: id, cate: cate, children: []*CateTree{}}

	//node := &Node{name: name, children: []*Node{}}

	if parentId == 0 {
		root = node
	} else {
		parent, ok := nodeTable[parentId]

		if !ok {
			return
		}

		parent.children = append(parent.children, node)
	}

	nodeTable[id] = node
}

func GetCateTree(parentid int64) {
	cates, _ := GetCategoriesOfParentId(parentid)

	for _, v := range cates {
		AddChildTree(v.Id, *v, v.Parentid)

		if v.Haschild == 1 {
			GetCateTree(v.Id)
		}
	}
}

func GetCategoriesOfParentId(parentid int64) ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	var err error
	_, err = o.QueryTable("category").Filter("parentid", parentid).All(&cates)

	return cates, err
}

func GetCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")

	_, err := qs.All(&cates)

	return cates, err
}

/*func GetCategoryTree(parentid int64, catesTree *catesTree) {
	o := orm.NewOrm()
	//catesTree := make(map[string]*Category, 0)
	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.Filter("parentid", parentid).All(&cates)

	for k, v := range cates {
		catesTree[k]["cate"] = v
		//catesTree[k]["child"] = make(map[string]*Category)
		if v.Haschild == 1 {
			catesTree[k]["child"] = GetCategoryTree(v.Parentid, catesTree)
		}
		if v.Haschild == 1 {
			GetCategoryTree(v.Category.Id)
		}
	}
}
*/
