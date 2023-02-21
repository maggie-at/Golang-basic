package go_packages

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonNode struct {
	K1 string `json:"k1"`
	K2 string `json:"k2"`
	K3 int    `json:"k3"`
}

func JSON_() {
	var jsonStr string = `{"k1": "v111","k2": "v222","k3": 333}`
	fmt.Println(jsonStr)

	// 1. 存入map
	// GO中可以用空接口interface{}表示Object
	var mp map[string]interface{}
	// Unmarshal: 先将字符串转为[]byte数组, 然后传入map指针
	json.Unmarshal([]byte(jsonStr), &mp)
	fmt.Println(mp)

	// 2. 直接存入struct, 需要提前定义好struct
	var node JsonNode
	json.Unmarshal([]byte(jsonStr), &node)
	fmt.Println(node)

	// 3. io流
	// 从json文件中读取
	rf, _ := os.Open("go_packages/test.json")
	defer rf.Close()
	decoder := json.NewDecoder(rf)
	var mp1 map[string]interface{}
	decoder.Decode(&mp1)
	fmt.Println(mp1)

	// 写入json文件
	var newJson string = `{"a1": "b111","a2": "b222","a3": 333}`
	wf, _ := os.OpenFile("test.json", os.O_WRONLY, 0777)
	defer wf.Close()
	encoder := json.NewEncoder(wf)
	encoder.Encode(newJson)

}
