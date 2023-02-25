package main

import "Golang/go_gorm"

func main() {
	// GO基础
	// fmt.Println("————main()————")
	// basic.CapTrack_()

	// GIN demo
	// GinMain()

	// GO并发
	// go_concurrent.AtomicOperations_()

	// GO爬虫
	// SpiderMain()

	// Go标准库
	// go_packages.Rand_()

	// go-sql
	// go_SQL.QueryMultiData_()

	// go_gorm
	db := go_gorm.InitDB_()
	//go_gorm.CreateTable_(db)
	//go_gorm.QueryData_(db)
	//go_gorm.InsertData_(db)
	//go_gorm.UpdateData_(db)
	//go_gorm.DeleteData_(db)

	//go_gorm.RawSql(db)

	//go_gorm.ManyToOne(db)
	go_gorm.ManyToMany(db)
}
