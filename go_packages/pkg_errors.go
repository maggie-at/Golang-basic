package go_packages

import (
	"errors"
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// 实现了error接口的Error()方法 => 实现了error接口
func (err MyError) Error() string {
	return fmt.Sprintf("%v: %v", err.When, err.What)
}
func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}
func oops() error {
	return MyError{
		When: time.Now(),
		What: "Some error happened",
	}
}
func Errors_() {
	// 直接使用error接口
	s, err := check("string")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(s)
	}

	// 自定义MyError结构体实现error接口
	myErr := oops()
	if myErr != nil {
		fmt.Println(myErr.Error())
	}
}
