package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "web:123456@tcp(rm-bp1y6kuovzmbz1757.mysql.rds.aliyuncs.com:3306)/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	var ()

	type Date struct {
		id      int       `json:"id"`
		name    string    `json:"name"`
		date    time.Time `json:"start_time"`
		endTime string    `json:"end_time"`
	}
	stmt, _ := db.Prepare("select id ,name,start_time,end_time from tt where id=?")
	rows, err := stmt.Query(1)

	if err != nil {
		fmt.Println("报错内容是:", err)
	}
	defer rows.Close()

	for rows.Next() {
		date := new(Date)
		err := rows.Scan(&date.id, &date.name, &date.date, &date.endTime)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(date.id, date.name, date.date, date.endTime)
	}
	defer db.Close()
}
