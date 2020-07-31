package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	Id        int    `gorm:"Column:id"`
	Name      string `gorm:"Column:name"`
	StartTime *time.Time
}

func main() {

	currentTime := "\033[31m" + time.Unix(1589976794, 0).Format("2006-01-02 15:04:05") + "\033[0m"
	fmt.Println(currentTime)
	db, err := gorm.Open("mysql", fmt.Sprintf("web:123456@(%v)/test?charset=utf8&parseTime=True", "rm-bp1y6kuovzmbz1757.mysql.rds.aliyuncs.com:3306"))
	if err != nil {
		fmt.Println("数据库开启失败", err)
		return
	}
	defer db.Close()
	var product Product
	/*if err := db.Table("tt").First(&product, "name=?", "ff").Error;err != nil{

	}*/

	if err := db.Table("tt").Where("name=? and id=1", "jj").Scan(&product).Error; err != nil {
		fmt.Println(err)
	}

	fmt.Println(product)
}
