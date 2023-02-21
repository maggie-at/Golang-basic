package go_packages

import (
	"fmt"
	"math/rand"
	"time"
)

func Rand_() {
	// 每次运行的随机数值固定
	fmt.Println(rand.Int())

	// 设置随机种子
	rand.Seed(time.Now().UnixMicro())
	
	fmt.Println(rand.Intn(100))
}
