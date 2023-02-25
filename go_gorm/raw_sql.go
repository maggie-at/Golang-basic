package go_gorm

import (
	"fmt"
	"gorm.io/gorm"
)

// PResult 一个缩略版的Product结构体
type PResult struct {
	Code  string // 字段首字母大写
	Price uint
}

func RawSql(db *gorm.DB) {
	// 1. db.Raw("sql stmt")
	// (1) .Row() / .Rows() => 逐行rows.Scan(&dest...)
	rows, _ := db.Raw("select code, price from Products where price > ?", 25).Rows()
	for rows.Next() {
		var pr PResult
		rows.Scan(&pr.Code, &pr.Price)
		fmt.Println(pr)
	}

	// (2) .Scan(&s) => 直接放到某个类型参数中
	// Scan(&s): Scanning results into a struct works similarly to the way we use Find
	var pList PResult // PResult 一个缩略版的Product结构体, 也可以是任何类型, 或者map
	db.Raw("select code, price from Products where code=?", "A25").Scan(&pList)
	fmt.Println(pList)
	var pList1 []PResult
	db.Raw("SELECT code, price FROM Products WHERE code = @code OR price > @price", map[string]interface{}{"code": "xxx' or '1'='1", "price": 25}).Scan(&pList1)
	fmt.Println(pList1)

	var codeList []string
	db.Raw("select code from Products where price=?", 25).Scan(&codeList)
	fmt.Println(codeList)

	// 聚合函数
	var avgPrice float32
	db.Raw("select avg(price) from Products").Scan(&avgPrice)
	fmt.Println("Average Price:", avgPrice)

	// 字符串拼接导致的「SQL注入」 => 解决方法: 参数化查询
	var pList2 []PResult
	var codeStr string = "xxx' or '1'='1"
	db.Raw("select code, price from Products where code='" + codeStr + "'").Scan(&pList2)
	fmt.Println(pList2)

	// 参数化查询可以避免SQL注入 => select code, price from Products where code='xxx\' or \'1\'=\'1'
	var pList3 []PResult
	db.Raw("select code, price from Products where code=?", "A25").Scan(&pList3)
	fmt.Println(pList3)

	// 2. db.Exec("sql stmt")
	db.Exec("Update Products set deleted_at = ?", nil)
}
