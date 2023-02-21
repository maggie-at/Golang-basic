package go_spider

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func Parse(html string) {
	// 删除换行符
	html = strings.Replace(html, "\n", "", -1)
	// 获取sidebar(针对GORM-doc这个demo的目录)
	reSidebar := regexp.MustCompile(`<aside id="sidebar" role="navigation">(.*?)</aside>`) // 先定义正则表达式
	sidebar := reSidebar.FindString(html)                                                  // 然后在html字符串中应用正则表达式
	// 获取href (links)
	reLink := regexp.MustCompile(`href="(.*?)"`)
	links := reLink.FindAllString(sidebar, -1)
	//fmt.Println(links)

	baseUrl := "https://gorm.io/zh_CN/docs/"
	for _, v := range links {
		s := v[6 : len(v)-1]
		url := baseUrl + s // 目录条目的链接

		statusCode, body := Fetch(url, map[string]string{"cookie": "_ga_YXBYDX14GJ=GS1.1.1676619317.1.0.1676619321.0.0.0; _ga=GA1.1.1598967530.1676619318"})

		if statusCode == http.StatusOK {
			// 每个链接都用一个Goroutine去抓取
			go Parse2(body)
		}
	}
}

func Parse2(body string) {
	body = strings.Replace(body, "\n", "", -1)

	// 获取页面
	reContent := regexp.MustCompile(`<div class="article">(.*?)</div>`)
	content := reContent.FindString(body)
	// 获取标题
	reTitle := regexp.MustCompile(`<h1 class="article-title" itemprop="name">(.*?)</h1>`)
	title := reTitle.FindString(content)
	title = title[42 : len(title)-5]

	fmt.Println("title: ", title)

	// 按title保存整个content
	SaveFile(title, content)
}
