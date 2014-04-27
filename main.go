package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"time"
)

func init() {
	beego.AddFuncMap("strtotime", strToTime)
}

func main() {
	beego.Run()
}

func strToTime(inttime int64) (formattime string) {
	formatetime := time.Unix(inttime, 0).Format("2006-01-02")
	return formatetime
}
