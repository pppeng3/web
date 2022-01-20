package zhihu

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly/v2"
)

type ZhiHuModel struct {
	Title   string
	Content string
	Hot     string
}

const zhihuUrl = "https://www.zhihu.com/billboard"

func crawlZhihu(url string) {
	var zhihuList []ZhiHuModel
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          10,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 10 * time.Second,
	})
	c.OnRequest(func(r *colly.Request) {

	})
	c.OnError(func(_ *colly.Response, err error) {

	})
	c.OnResponse(func(r *colly.Response) {
		doc, err := htmlquery.Parse(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Fatal(err.Error())
		}
		nodes := htmlquery.Find(doc, `//div[@class="HotList-itemBody"]`)
		for _, node := range nodes {
			title := htmlquery.FindOne(node, `./div[@class="HotList-itemTitle"]/text()`)
			content := htmlquery.FindOne(node, `./div[@class="HotList-itemExcerpt"]/text()`)
			hot := htmlquery.FindOne(node, `./div[@class="HotList-itemMetrics"]/text()`)
			zhihuList = append(zhihuList, ZhiHuModel{
				Title: htmlquery.InnerText(title),
				Content: htmlquery.InnerText(content),
				Hot: htmlquery.InnerText(hot),
			})
		}
	})
	c.OnHTML("", func(e *colly.HTMLElement) {

	})
	c.OnScraped(func(r *colly.Response) {

	})
	c.Visit(zhihuUrl)
	c.Wait()
	fmt.Println(zhihuList)
}
