package basic

import (
	"encoding/json"
	"fmt"
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
	// 先将字符串转为[]byte数组, 然后传入map指针
	json.Unmarshal([]byte(jsonStr), &mp)
	fmt.Println(mp)

	// 2. 直接存入struct, 需要提前定义好struct
	var node JsonNode
	json.Unmarshal([]byte(jsonStr), &node)
	fmt.Println(node)
}
