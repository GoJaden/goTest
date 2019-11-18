package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	type Date struct {
		id   *int       `json:"id"`
		name *string    `json:"name"`
		date *time.Time `json:"date"`
	}
	handler := new(UserDao)
	DBQuery(db, "select id,name from db where id=?", []interface{}{1111}, handler)

}

func DBQuery(db *sql.DB, sqlStr string, arg []interface{}, handler DBHandler) error {
	var err error

	stmt, err := db.Prepare(sqlStr)
	if err != nil {

		return errors.New("fd")
	}
	defer stmt.Close()
	rows, err := stmt.Query(arg...)
	if err != nil {
		return errors.New("fd")
	}

	defer rows.Close()
	return deal(rows, handler)

}

func deal(rows *sql.Rows, handler DBHandler) error {
	return handler.HandlerResults(rows)
}

type UserDao struct {
}

var (
	id   int
	name string
)

func (dao UserDao) HandlerResults(rows *sql.Rows) error {
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return errors.New("error")
		}
		fmt.Println("数据:", id, name)
	}
	return nil
}

type DBHandler interface {
	HandlerResults(rows *sql.Rows) error
}
