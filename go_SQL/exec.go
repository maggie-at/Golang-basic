package go_SQL

import (
	"database/sql"
	"fmt"
)

type User struct {
	id       int
	username string
	password string
}

var db *sql.DB

func InsertData_() {
	db = InitDB_()
	sql := "insert into User(`username`, `password`) values(?, ?)"
	r, err := db.Exec(sql, "Maggie", "24")
	if err != nil {
		panic(err)
	}

	id, err := r.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Insert successfully, the last record is %v.", id)
}

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

func DeleteData_() {
	db = InitDB_()
	sql := "delete from User where id=?"
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
