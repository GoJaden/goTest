package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	tx, err := db.Begin()
	defer func() {
		fmt.Println("start roback")
		if err := tx.Commit(); err != nil {
			fmt.Println("roback")
			//tx.Rollback()
		}
	}()
	type Date struct {
		id   int       `json:"id"`
		name string    `json:"name"`
		date time.Time `json:"date"`
	}
	stmt, _ := tx.Prepare("insert into db (name,date) values (?,?)")
	result, err := stmt.Exec("jaden123", time.Now())

	if err != nil {
		fmt.Println("报错内容是:", err)
		return
	}

	stmt1, _ := tx.Prepare("insert into db values (?,?,?)")
	result1, err := stmt1.Exec(137, "jaden123", time.Now())

	if err != nil {
		tx.Rollback()
		fmt.Println("报错内容是:", err)
		return
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result1.LastInsertId())
}
