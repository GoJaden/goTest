package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goProject/gotest/basic/utils"
	"time"
)

type Data struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("+++", err)
		return
	}
	defer tx.Commit()
	//Query(tx)
	//insert(tx)
	queryIn(tx)
}

func Split(s string) string {
	return s
}
func queryIn(tx *sql.Tx) {
	placeHolders := utils.SqlParamsConverter([]interface{}{"323", "32", "jaden"})
	sql := fmt.Sprintf("select id,name,date from db where name in (%s)", placeHolders)
	stmt, err := tx.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	//ids := []int{1,3,11,12}
	rows, err := stmt.Query([]interface{}{"323", "32", "jaden"}...)
	if err != nil {
		fmt.Println("---", err)
		return
	}
	datas := make([]Data, 0)
	for rows.Next() {
		data := new(Data)
		rows.Scan(&data.Id, &data.Name, &data.Date)
		datas = append(datas, *data)
	}

	for _, val := range datas {
		fmt.Println("data:", val)
	}
}

func insert(tx *sql.Tx) {
	stmt, err := tx.Prepare("insert into db (name,date) values (?,?)")
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := stmt.Exec("jaden1", time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	id, _ := result.LastInsertId()
	fmt.Println("result:", id)
}

func Query(tx *sql.Tx) {
	rows, err := tx.Query("select id ,name,date from db")
	if err != nil {
		fmt.Println("---", err)
		return
	}
	datas := make([]Data, 0)
	for rows.Next() {
		data := new(Data)
		rows.Scan(&data.Id, &data.Name, &data.Date)
		datas = append(datas, *data)
	}

	for _, val := range datas {
		fmt.Println("data:", val)
	}
}
