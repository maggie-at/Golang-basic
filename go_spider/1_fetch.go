package go_spider

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string, params map[string]string) (int, string) {
	// 创建请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Request Error: ", err)
	}

	// 设置header和cookie, 为了防止网站监测到爬虫访问, 伪造成浏览器访问
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/110.0")
	req.Header.Add("Cookie", params["cookie"])
	//req.Header.Add("Cookie", "csrftoken=GzJ58LxtEtFTSLPAGWCKrhQVEB4bFo6ebv9Vnrknpw5cdQPuoLRRraCfV48HJqjA; gr_user_id=4b8ffe66-09f9-4510-a6a8-cac5299d69f4; a2873925c34ecbd2_gr_last_sent_cs1=nervous-mestorfe6y; Hm_lvt_fa218a3ff7179639febdb15e372f411c=1670051064,1670126428,1670401499,1670759225; _ga=GA1.1.278110060.1663154393; Hm_lvt_f0faad39bcf8471e3ab3ef70125152c3=1674779882,1674869675,1675907971,1676188108; _ga_PDVPZYN3CW=GS1.1.1676617700.325.1.1676617765.0.0.0; LEETCODE_SESSION=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJfYXV0aF91c2VyX2lkIjoiMjc1MTUzOSIsIl9hdXRoX3VzZXJfYmFja2VuZCI6ImF1dGhlbnRpY2F0aW9uLmF1dGhfYmFja2VuZHMuUGhvbmVBdXRoZW50aWNhdGlvbkJhY2tlbmQiLCJfYXV0aF91c2VyX2hhc2giOiI5OTk2NTE4MDYzMGI3OTE2YjU0ZTI3ZjVlMTk3NTk4YjFjZmE3YmI0NDRjNDRiNTgxNDNmZTgwOTQ2MjA3YjgzIiwiaWQiOjI3NTE1MzksImVtYWlsIjoiMjFTMTUxMTU1QHN0dS5oaXQuZWR1LmNuIiwidXNlcm5hbWUiOiJuZXJ2b3VzLW1lc3RvcmZlNnkiLCJ1c2VyX3NsdWciOiJuZXJ2b3VzLW1lc3RvcmZlNnkiLCJhdmF0YXIiOiJodHRwczovL2Fzc2V0cy5sZWV0Y29kZS5jbi9hbGl5dW4tbGMtdXBsb2FkL3VzZXJzL25lcnZvdXMtbWVzdG9yZmU2eS9hdmF0YXJfMTY2OTM4NzA5MS5wbmciLCJwaG9uZV92ZXJpZmllZCI6dHJ1ZSwiX3RpbWVzdGFtcCI6MTY3NTU2NTk0NC42NDY5NjgsImV4cGlyZWRfdGltZV8iOjE2NzgxMjkyMDAsInZlcnNpb25fa2V5XyI6MiwibGF0ZXN0X3RpbWVzdGFtcF8iOjE2NzY2MTc3NjV9.n8L0GvppwdffObQlXCan0abVc-m7ygyCgs46QJYn0E4; a2873925c34ecbd2_gr_session_id_1512220a-7c91-4812-be1f-763a41484c43=true; a2873925c34ecbd2_gr_session_id=1512220a-7c91-4812-be1f-763a41484c43; a2873925c34ecbd2_gr_last_sent_sid_with_cs1=1512220a-7c91-4812-be1f-763a41484c43; a2873925c34ecbd2_gr_cs1=nervous-mestorfe6y")

	// 发出请求, 接收response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http Get Error: ", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Http Status Code: ", resp.StatusCode)
		return resp.StatusCode, "请求错误"
	}
	defer resp.Body.Close()

	// 读取response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Error", err)
		return resp.StatusCode, "读取response失败"
	}
	return resp.StatusCode, string(body)
}
