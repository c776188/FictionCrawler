package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// CrawlerController operations for Crawler
type CrawlerController struct {
	beego.Controller
}

type fictionCrawler struct {
	Id   int
	Name string
	Url  string
}

var visited = make(map[string]bool)

func (c *CrawlerController) Get() {
	c.TplName = "crawler/index.tpl"
}

// URLMapping ...
func (c *CrawlerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Crawler
// @Param	body		body 	models.Crawler	true		"body for Crawler content"
// @Success 201 {object} models.Crawler
// @Failure 403 body is empty
// @router / [post]
func (c *CrawlerController) Post() {
	pages, _ := c.GetInt("page")

	var result []fictionCrawler

	preUrl := "http://big5.quanben.io/c/xuanhuan_"
	proUrl := ".html"
	queue := make(chan int, 1)
	go func() {
		queue <- 1
	}()

	for page := 1; page <= pages; page++ {
		result = append(result, crawler(preUrl, page, proUrl, make(chan int, 1))...)
	}

	c.Data["json"] = &result
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Crawler by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Crawler
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CrawlerController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Crawler
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Crawler
// @Failure 403
// @router / [get]
func (c *CrawlerController) GetAll() {

}

// private funciton
func crawler(preUrl string, page int, proUrl string, queue chan int) []fictionCrawler {
	crawUrl := preUrl + strconv.Itoa(page) + proUrl
	visited[crawUrl] = true
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	// 取得網頁資料
	req, _ := http.NewRequest("GET", crawUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return nil
	}
	defer resp.Body.Close()

	// 解析
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	// 爬蟲
	var result []fictionCrawler

	var tmpNameList []string
	doc.Find("div[class=list2] span[itemprop=name]").Each(func(i int, selection *goquery.Selection) {
		logs.Debug("selection.Text(): ", selection.Text())
		tmpNameList = append(tmpNameList, selection.Text())
	})

	// 搜尋小說Url
	var tmpUrlList []string
	doc.Find("h3>a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if !ok {
			logs.Debug("error")
		}
		tmpUrlList = append(tmpUrlList, href)
	})

	for idx, _ := range tmpNameList {
		result = append(result, fictionCrawler{
			Id:   (page-1)*12 + idx + 1,
			Name: tmpNameList[idx],
			Url:  tmpUrlList[idx],
		})
	}

	return result
}
