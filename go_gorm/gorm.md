### [GORM](https://gorm.io/)

```GO
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

#### 0. 建立数据库连接
```GO
import (
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

func InitDB_() *gorm.DB {
    dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8&parseTime=true"}, "")
    db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    return db
}
```


#### 1. 迁移schema(创建表) - `db.AutoMigrate(&Struct{}}`
```GO
type Product struct {
    // Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
    // It may be embedded into your model or you may build your own model without it 
	gorm.Model
    Code  string    // 字段首字母大写
    Price uint
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
```


#### 2. 查询数据
```GO
func QueryData_(db *gorm.DB) {
	// 主键查询
	var p1 Product
	db.First(&p1, 5)    // db.First(): 返回第一条记录, &p1用来接收查询结果
	fmt.Printf("p: %v\n", p1)

    // 主键批量查询: 切片
    var pList []Product
    ret := db.Find(&pList, []int{1, 2, 5})  // db.Find(): 返回所有匹配记录
    for idx, p := range pList {
        fmt.Println(idx, ": ", p)
    }
    fmt.Println(ret.RowsAffected, "rows found.")
	
    // 条件查询: First返回主键排序后的第一条; Last返回主键排序后的最后一条
    var p Product
    db.First(&p, "price=?", 0)
    fmt.Printf("p: %v\n", p)
    
    var p3 Product
    db.Last(&p3, "price=?", 0)
    fmt.Printf("p: %v\n", p3)
	
	// Where(query, args...) => = / <> / in / like / and / between
	//db.Where("Code <> ?", "A25").Find(&pList)
	//db.Where("Code in ?", []string{"K30", "M12"}).Find(&pList)
	//db.Where("Code like ?", "A%").Find(&pList)
	//db.Where("Code = ? AND price between ? AND ?", "A25", 0, 30).Find(&pList)
    // 配合Where(query, args...)定义任意查询
    var pList1 []Product
    db.Where("Code = ? AND price between ? AND ?", "A25", 0, 30).Find(&pList1)
    fmt.Println(pList1)
    
    // Not => select * from Table where NOT [clauses]
    var pList2 []Product
    db.Not("Code = ?", "A25").Find(&pList2)
    fmt.Println(pList2)

    // OR
    var pList3 []Product
    db.Where("Code = ?", "A25").Or("Price = ?", 25).Find(&pList3)
    fmt.Println(pList3)
    }

    // Select => 指定要检索的列
    var pList4 []Product
    db.Select("Code", "Price").Find(&pList4)
```


#### 3. 增/删/改
```GO
func InsertData_(db *gorm.DB) {
    // 1. 插入一条完整数据
    p := Product{Code: "A25", Price: 25}
    db.Create(&p)
    // 或者直接写成
    db.Create(&Product{Code: "D42", Price: 100})
    
    // 2. 插入部分字段
    pPart := Product{Code: "B20"}
    db.Select("Code").Create(&pPart)
    
    // 3. 批量插入, 用「结构体切片」或者「map切片」都可以
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
	// 方式一: db.Model(&m).Update("k", v) / .Updates(m)
    // 1. 更新单个字段
    db.Model(&product).Update("Price", 80)
    // 2. 使用 model结构体 或者 map 更新多个字段
    db.Model(&product).Updates(Product{Code: "A80", Price: 25})
    db.Model(&product).Updates(map[string]interface{}{"Code": "A25", "Price": 25})
    
	// 方式二: 修改model然后db.Save(&m)
    var product1 Product
    db.First(&product1, 3)
    product1.Code = "A13"
    product1.Price = 25
    db.Save(&product1)

    // 3. 批量更新
    db.Model(&Product{}).Where("Price=?", 0).Update("Price", 80)
    
    // 阻止全局更新 (WHERE conditions required), 至少要加一个.Where("1=1")
    fmt.Println(db.Model(&Product{}).Update("Code", "xxx").Error)
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
```


#### 4. 原生SQL
> `DB.Exec()`和`DB.Raw()`都是GORM框架中用于执行原生SQL语句的方法, 但它们有不同的用法。
> 
> - `DB.Exec()`
>   - 用于执行`SQL`语句并返回结果, 可以用来执行`INSERT`, `UPDATE`, `DELETE`等语句
>   - 执行成功返回一个`sql.Result`对象和一个`nil`作为错误，否则返回一个`非nil错误`
>   - 它也支持使用`?占位符`的方式进行参数绑定，比如: `db.Exec("UPDATE users SET name = ? WHERE id = ?", "Alice", 123)`
> - `DB.Raw()`
>   - 用于生成一个原生SQL**查询**语句，可以用于`SELECT`查询等操作
>   - 返回的是一个`*gorm.DB`对象, 可以通过该对象进行链式查询等操作, 或者配合`.Scan(&list)`保存查询结果
>   - 使用`?占位符`的方式进行参数绑定, 可以避免**SQL注入**
> 总的来说，`DB.Exec()`方法更适用于执行`INSERT`, `UPDATE`, `DELETE`等语句, 而`DB.Raw()`方法更适用于生成`SELECT`查询语句


##### SQL注入 => 参数化查询
> 用`?占位符`替代字符串拼接的方式防止SQL注入, 原理是采用了「预编译」的方法, 先将SQL语句中可被客户端控制的参数集进行编译, 生成对应的临时变量集, 再使用对应的设置方法, 为临时变量集里面的元素进行赋值
```GO
// 字符串拼接导致的「SQL注入」 => 解决方法: 参数化查询
var pList2 []PResult
var codeStr string = "xxx' or '1'='1"
db.Raw("select code, price from Products where code='" + codeStr + "'").Scan(&pList2)
fmt.Println(pList2)
```

```GO
func RawSql(db *gorm.DB) {
    // 1. db.Raw("sql stmt")
    // (1) .Row() / .Rows() => 逐行rows.Scan(&dest...)
    rows, _ := db.Raw("select code, price from Products where price > ?", 25).Rows()
    for rows.Next() {
        var pr PResult
        // rows.Scan(&dest1, ...): 逐个字段载入
        // rows.Scan(&pr.Code, &pr.Price)
        // fmt.Println(pr)
        
        // db.ScanRows(rows, &dest): 直接载入结构体
        db.ScanRows(rows, &pr)
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
```