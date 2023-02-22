package go_spider

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

/* https://www.bilibili.com/video/BV1CR4y1g7wB?p=2&vd_source=e60de8162f155cdd464b9f11c355e633 */

const (
	USERNAME = "alan"
	PASSWORD = "alantam."
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "golang"
)

// 使用原生SQL语句
var DB *sql.DB

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	fmt.Println(path)
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(10)
	DB.SetMaxIdleConns(5)

	err := DB.Ping()
	if err != nil {
		fmt.Println("Open Database Failed")
		return
	}
	fmt.Println("Connect Successfully")
}
func Insert(movie MovieData) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("Begin error ", err)
		return false
	}
	stmt, err := tx.Prepare("insert into DoubanMovie (`Title`, `Director`, `Picture`, `Actor`, `Year`, `Score`, `Quote`) values (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare error ", err)
		return false
	}
	_, err = stmt.Exec(movie.Title, movie.Director, movie.Picture, movie.Actor, movie.Year, movie.Score, movie.Quote)
	if err != nil {
		fmt.Println("Execute err ", err)
		return false
	}
	_ = tx.Commit()
	return true
}

type MovieData struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Picture  string `json:"picture"`
	Actor    string `json:"actor"`
	Year     string `json:"year"`
	Score    string `json:"score"`
	Quote    string `json:"quote"`
}

func InfoParse(info string) (director, actor, year string) {
	// 定义正则表达式, 然后到info中去匹配各部分
	directorRe, _ := regexp.Compile(`导演: (.*)主演`)
	director = string(directorRe.Find([]byte(info)))

	actorRe, _ := regexp.Compile(`主演: (.*)`)
	actor = string(actorRe.Find([]byte(info)))

	yearRe, _ := regexp.Compile(`(\d+)`)
	year = string(yearRe.Find([]byte(info)))

	return director, actor, year
}
func Douban(page int) {
	// 1. 发送请求
	client := &http.Client{}
	url := "https://movie.douban.com/top250/?start=" + strconv.Itoa(page)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Request Error: ", err)
	}

	// 设置header和cookie, 为了防止网站监测到爬虫访问, 伪造成浏览器访问
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Referer", "https://movie.douban.com/chart")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Cookie", `ll="118282"; bid=ujCYZY6_inY; __utma=30149280.1441240252.1665111519.1665111519.1665111519.1; __utmz=30149280.1665111519.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; viewed="30351286_26772632_5377669_5286107"; gr_user_id=256643b2-98c8-48a5-8b96-1b719dd376e8; __gads=ID=84f144410487b4a0-22336759e6d6005f:T=1665111533:RT=1665111533:S=ALNI_MZVz5Fxi5Z7crGH_KN1h4R0ZsrtOQ; __gpi=UID=00000a1ff2542432:T=1665111533:RT=1665111533:S=ALNI_MbRqg6NyBywuyKvUGkIYzkwSFxQmw; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1676798241%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=6effe4795bd63dc2.1676640432.3.1676798241.1676728834.; __yadk_uid=TZBSpWnlJfbaQG2Qy8OVTB9xDn74f3ja; _vwo_uuid_v2=DC59D774B212D1A0B63FE410A613184A0|cee573fbbe10f23ce1d75388b372c276; _pk_ses.100001.4cf6=*; ap_v=0,6.0`)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Response Error: ", err)
	}

	// 2. 解析网页 - goquery
	docDetail, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Parse Failed", err)
	}

	// Step 3: 获取节点信息, 保存为结构体
	// 遍历每个<li></li>
	docDetail.Find("#content > div > div.article > ol > li").Each(func(i int, s *goquery.Selection) {
		var movie MovieData
		title := s.Find("div > div.info > div.hd > a > span:nth-child(1)").Text()
		img := s.Find("div > div.pic > a > img")
		imgPath, ok := img.Attr("src")
		info := s.Find("div > div.info > div.bd > p:nth-child(1)").Text()
		score := s.Find("div > div.info > div.bd > div > span.rating_num").Text()
		quote := s.Find("div > div.info > div.bd > p.quote > span").Text()
		if ok {
			fmt.Println("title: ", title)
			fmt.Println("img: ", imgPath)
			fmt.Println("score: ", score)
			fmt.Println("quote: ", quote)

			// 正则处理movie info
			director, actor, year := InfoParse(info)
			fmt.Println("director: ", director)
			fmt.Println("actor: ", actor)
			fmt.Println("year: ", year)

			// 保存为MovieData结构体
			movie.Title = title
			movie.Picture = imgPath
			movie.Score = score
			movie.Quote = quote
			movie.Director = director
			movie.Actor = actor
			movie.Year = year

			// Step 4: 保存信息 - 插入数据库
			if Insert(movie) {
				fmt.Println("插入成功")
			} else {
				fmt.Println("插入失败")
			}
		}
	})
}
