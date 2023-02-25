package go_gorm

import (
	"fmt"
	"gorm.io/gorm/logger"
	"strings"
	"time"

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
	// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
	// It may be embedded into your model or you may build your own model without it
	gorm.Model
	Code  string // 字段首字母大写
	Price uint
}

func InitDB_() *gorm.DB {
	dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8&parseTime=true"}, "")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

	// 主键批量查询 (要传入切片)
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

func InsertData_(db *gorm.DB) {
	// 1. 插入一条完整数据
	p := Product{Code: "A25", Price: 25}
	db.Create(&p)
	// 或者直接写成
	db.Create(&Product{Code: "D42", Price: 100})

	// 2. 插入部分字段
	pPart := Product{Code: "B20"}
	db.Select("Code").Create(&pPart)

	// 3. 批量插入 (要传入切片), 用「结构体切片」或者「map切片」都可以
	var pList = []Product{{Code: "K30", Price: 1999}, {Code: "M12", Price: 3999}}
	db.Create(&pList)

	// 尽量别用
	db.Model(&Product{}).Create(map[string]interface{}{
		"Code":  "K50",
		"Price": 1999,
	})
}

func UpdateData_(db *gorm.DB) {
	// 先把要更新的record查询出来
	var product Product
	db.First(&product, 1)
	// 1. 更新单个字段
	db.Model(&product).Update("Price", 80)
	// 2. 使用 结构体 或者 map 更新多个字段
	db.Model(&product).Updates(Product{Code: "A80", Price: 25})
	db.Model(&product).Updates(map[string]interface{}{"Code": "A25", "Price": 25})

	var product1 Product
	db.First(&product1, 3)
	product1.Code = "A13"
	product1.Price = 25
	db.Save(&product1)

	// 3. 批量更新
	ret := db.Model(&Product{}).Where("Price=?", 100).Update("Price", 80)
	fmt.Println(ret.RowsAffected, "rows updated.")

	// 阻止全局更新 (WHERE conditions required), 至少要加一个.Where("1=1")
	db.Model(&Product{}).Update("Code", "xxx") // WHERE conditions required
	db.Model(&Product{}).Where("1 = 1").Update("updated_at", time.Now())

	// 执行原生SQL语句
	db.Exec("Update Products set deleted_at = ?", nil)
}

func DeleteData_(db *gorm.DB) { // 软删除(将delete标识位设为非空)
	var product Product
	// 1. 根据主键删除
	db.Delete(&product, 2)

	// 2. 先查询后删除
	db.First(&product, 4)
	db.Delete(&product) // 删除已经删除的record => record not found

	// 3. 批量删除 (要传入切片)
	var pList []Product
	ret := db.Delete(pList, []int{5, 6})
	fmt.Println(ret.RowsAffected, "rows deleted.")

	ret1 := db.Where("Code like ?", "%2").Delete(pList)
	fmt.Println(ret1.RowsAffected, "rows deleted")

	ret2 := db.Delete(pList, "Code like ?", "%3")
	fmt.Println(ret2.RowsAffected, "rows deleted")
}
