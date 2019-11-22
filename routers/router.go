package routers

import (
	"FictionCrawler/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/crawler", &controllers.CrawlerController{})
}
