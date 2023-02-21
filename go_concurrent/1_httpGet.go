package go_concurrent

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step 1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step 2: ", url)
	defer response.Body.Close()

	fmt.Println("Step 3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step 4: ", url, ": ", len(body))
}
func HttpGet() {
	go responseSize("https://www.baidu.com")
	go responseSize("https://www.jd.com")
	go responseSize("https://www.taobao.com")
	time.Sleep(time.Second * 3)
}
