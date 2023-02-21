package go_spider

import "os"

func SaveFile(title, content string) {
	err := os.WriteFile("./go_spider/pages/"+title+".html", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}
