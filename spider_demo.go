package main

import (
	"Golang/go_spider"
	"fmt"
)

func SpiderMain() {
	/*
		// Get力扣主页
		url := "https://www.leetcode.com"
		status, _ := go_spider.Fetch(url, map[string]string{"cookie": "csrftoken=GzJ58LxtEtFTSLPAGWCKrhQVEB4bFo6ebv9Vnrknpw5cdQPuoLRRraCfV48HJqjA; gr_user_id=4b8ffe66-09f9-4510-a6a8-cac5299d69f4; a2873925c34ecbd2_gr_last_sent_cs1=nervous-mestorfe6y; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1670051064,1670126428,1670401499,1670759225; _ga=GA1.1.278110060.1663154393; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1674779882,1674869675,1675907971,1676188108; _ga_PDVPZYN3CW=GS1.1.1676617700.325.1.1676617765.0.0.0; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjc1MTUzOSIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI5OTk2NTE4MDYzMGI3OTE2YjU0ZTI3ZjVlMTk3NTk4YjFjZmE3YmI0NDRjNDRiNTgxNDNmZTgwOTQ2MjA3YjgzIiwiaWQiOjI3NTE1MzksImVtYWlsIjoiMjFTMTUxMTU1QHN0dS5oaXQuZWR1LmNuIiwidXNlcm5hbWUiOiJuZXJ2b3VzLW1lc3RvcmZlNnkiLCJ1c2VyX3NsdWciOiJuZXJ2b3VzLW1lc3RvcmZlNnkiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL25lcnZvdXMtbWVzdG9yZmU2eS9hdmF0YXJfMTY2OTM4NzA5MS5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTY3NTU2NTk0NC42NDY5NjgsImV4cGlyZWRfdGltZV8iOjE2NzgxMjkyMDAsInZlcnNpb25fa2V5XyI6MiwibGF0ZXN0X3RpbWVzdGFtcF8iOjE2NzY2MTc3NjV9.n8L0GvppwdffObQlXCan0abVc-m7ygyCgs46QJYn0E4; a2873925c34ecbd2_gr_session_id_1512220a-7c91-4812-be1f-763a41484c43=true; a2873925c34ecbd2_gr_session_id=1512220a-7c91-4812-be1f-763a41484c43; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=1512220a-7c91-4812-be1f-763a41484c43; a2873925c34ecbd2_gr_cs1=nervous-mestorfe6y"})
		fmt.Println(status)
	*/

	/*
		// Get Gorm-Docs
		url := "https://gorm.io/zh_CN/docs/"
		_, html := go_spider.Fetch(url, map[string]string{"cookie": "_ga_YXBYDX14GJ=GS1.1.1676619317.1.0.1676619321.0.0.0; _ga=GA1.1.1598967530.1676619318"})
		//fmt.Println(html)

		// 解析Gorm-Docs的目录, 爬取每个链接
		go_spider.Parse(html)
	*/

	go_spider.InitDB()
	for i := 0; i < 10; i++ {
		fmt.Printf("正在爬取第%v页信息\n", i)
		go_spider.Douban(i * 25)
	}
}
