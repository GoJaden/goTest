package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	Id       int    `gorm:"Column:id"`
	Name     string `gorm:"Column:name"`
	Birthday *time.Time
}

func main() {

	currentTime := "\033[31m" + time.Unix(1589976794, 0).Format("2006-01-02 15:04:05") + "\033[0m"
	fmt.Println(currentTime)
	db, err := gorm.Open("mysql", fmt.Sprintf("root:root@(%v)/gorm?charset=utf8&parseTime=True", "127.0.0.1:3306"))
	if err != nil {
		fmt.Println("数据库开启失败", err)
		return
	}
	defer db.Close()
	var product Product
	db.Table("product").First(&product, "name=?", "电饭锅")

	fmt.Println(product)
}
