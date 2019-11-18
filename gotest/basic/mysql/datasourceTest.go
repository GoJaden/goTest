package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	var ()

	type Date struct {
		id   int       `json:"id"`
		name string    `json:"name"`
		date time.Time `json:"date"`
	}
	stmt, _ := db.Prepare("select id ,name,date from db where id=?")
	rows, err := stmt.Query(1111)

	if err != nil {
		fmt.Println("报错内容是:", err)
	}
	defer rows.Close()

	for rows.Next() {
		date := new(Date)
		err := rows.Scan(&date.id, &date.name, &date.date)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.id, date.name, date.date)
	}
	defer db.Close()
}
