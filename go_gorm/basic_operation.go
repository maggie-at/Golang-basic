package go_gorm

import (
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "alantam."
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "golang"
)

type Product struct {
	gorm.Model
	Code  string // 字段首字母大写
	Price uint
}

func InitDB_() *gorm.DB {
	dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8&parseTime=true"}, "")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

func CreateTable_(db *gorm.DB) {
	// 迁移schema(创建表)
	db.AutoMigrate(&Product{})
	/*
		+------------+-----------------+------+-----+---------+----------------+
		| Field      | Type            | Null | Key | Default | Extra          |
		+------------+-----------------+------+-----+---------+----------------+
		| id         | bigint unsigned | NO   | PRI | NULL    | auto_increment |
		| created_at | datetime(3)     | YES  |     | NULL    |                |
		| updated_at | datetime(3)     | YES  |     | NULL    |                |
		| deleted_at | datetime(3)     | YES  | MUL | NULL    |                |
		| code       | longtext        | YES  |     | NULL    |                |
		| price      | bigint unsigned | YES  |     | NULL    |                |
		+------------+-----------------+------+-----+---------+----------------+
	*/
}

func InsertData_(db *gorm.DB) {
	// 插入数据
	p := Product{Code: "A25", Price: 25}
	db.Create(&p)
	// 或者直接写成
	db.Create(&Product{Code: "D42", Price: 100})

	// 插入部分字段
	p_part := Product{Code: "B20"}
	db.Select("Code").Create(&p_part)

	// 批量插入, 用「结构体切片」或者「map切片」都可以
	var pList = []Product{{Code: "K30"}, {Code: "M12"}}
	db.Create(&pList)

	// 尽量别用
	db.Model(&Product{}).Create(map[string]interface{}{
		"Code":  "K50",
		"Price": 1999,
	})
}

func QueryData_(db *gorm.DB) {
	// 主键查询
	var p1 Product
	db.First(&p1, 5)
	fmt.Printf("p: %v\n", p1)

	// 条件查询: First返回主键排序后的第一条; Last返回主键排序后的最后一条
	var p2 Product
	db.First(&p2, "price=?", 0)
	fmt.Printf("p: %v\n", p2)

	var p3 Product
	db.Last(&p3, "price=?", 0)
	fmt.Printf("p: %v\n", p3)

	// 主键批量查询: 切片
	var pList []Product
	ret := db.Find(&pList, []int{1, 2, 5})
	for idx, p := range pList {
		fmt.Println(idx, ": ", p)
	}
	fmt.Println(ret.RowsAffected, "rows found.")

	// 配合Where(query, args...)定义任意查询
	var pList1 []Product
	db.Where("Code = ? AND price between ? AND ?", "A25", 0, 30).Find(&pList1)
	fmt.Println(pList1)

	// Not => select * from Table where NOT [clauses]
	var pList2 []Product
	db.Not("Code = ?", "A25").Find(&pList2)
	fmt.Println(pList2)

	// Or
	var pList3 []Product
	db.Where("Code = ?", "A25").Or("Price = ?", 25).Find(&pList3)
	fmt.Println(pList3)

	// Select => 指定要检索的列
	var pList4 []Product
	db.Select("Code", "Price").Find(&pList4)

	// .Group().Having().Rows()
	// .Joins()
	// .Distinct()
}

func UpdateData_(db *gorm.DB) {
	// 先把要更新的record查询出来
	var product Product
	db.First(&product, 1)
	// 更新单个字段
	db.Model(&product).Update("price", 80)
	// 使用 结构体 或者 map 更新多个字段
	db.Model(&product).Updates(Product{Code: "A80", Price: 25})
	db.Model(&product).Updates(map[string]interface{}{"Code": "A25", "Price": 25})
}

func DeleteData_(db *gorm.DB) {
	var product Product
	// 删除record
	db.Delete(&product, 2)
}
