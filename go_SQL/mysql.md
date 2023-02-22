### go-mysql

#### 0. 安装
> Step 1: [MySQL安装](https://blog.csdn.net/guorenhao/article/details/124508441)
> 
> Step 2: [go get安装go-sql驱动](https://pkg.go.dev/github.com/go-sql-driver/mysql#section-readme)


#### 1. 建立连接

```GO
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	USERNAME = "alan"
	PASSWORD = "alantam."
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "golang"
)

func InitDB_() *sql.DB{
	dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8mb4&parseTime=True"}, "")

	// sql.Open(driverName, dataSourceName): 检查连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(5)

	// Ping: 建立连接
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established.")
	return db
}
```


#### 2. 增/删/改 - `Exec()`

##### (1) 插入数据
```GO
var db *sql.DB

func InsertData_() {
	db = InitDB_()
	sql := "insert into User(`username`, `password`) values(?, ?)"
	r, err := db.Exec(sql, "maggie", "24")
	if err != nil {
		panic(err)
	}

	id, err := r.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Insert successfully, the last record is %v.", id)
}
```

##### (2) 更新数据
```GO
func UpdateData_() {
	db = InitDB_()
	sql := "update User set `username`=?, `password`=? where id=?"
	ret, err := db.Exec(sql, `Maggie`, `24`, 2)
	if err != nil {
		panic(err)
	}
	affected, err := ret.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v records updated.", affected)
}
```

##### (3) 删除数据
```GO
func DeleteData_() {
	db = InitDB_()
	sql := "delete from User where id = ?"
	ret, err := db.Exec(sql, 2)
    if err != nil {
        panic(err)
    }
    affected, err := ret.RowsAffected()
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v records deleted.", affected)
}
```


#### 3. 查询 - Query

##### (1) 单行查询

```GO
func QuerySingleData_() {
	db = InitDB_()
	sql := "select * from User where id = ?"
	var u User
	err := db.QueryRow(sql, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("u: %v\n", u)
	}
}
```

##### (2) 多行查询
```GO
func QueryMultiData_() {
	db = InitDB_()
	sql := "select * from User"
	r, err := db.Query(sql)
	defer r.Close()
	if err != nil {
		panic(err)
	}
	var users []User
	for r.Next() {
		var user User
		fmt.Println(r.Scan(&user.id, &user.username, &user.password))
		users = append(users, user)
	}
	fmt.Println(users)
}
```