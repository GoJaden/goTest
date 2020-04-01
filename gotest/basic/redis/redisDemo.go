package main

import (
	"common-lib/redis"
	"fmt"
	"time"
)

type User struct {
	name string `json:"name"`
	age  int    `json:"age"`
}

type User2 struct {
	name string `json:"name"`
	age  int    `json:"age"`
}
func main() {

	config := redis.Config{Address: "192.168.1.204:6379",
		Password:    "123456",
		MaxActive:   50,
		IdleTimeout: time.Duration(60) * time.Second}
	client := redis.New(config)
	//val, err := client.Get("fdfd;ljkdjkfoewj").Result()

	client.Set("test111","1122",0)

	client.HSet("name","jiang","123")
	client.HSet("name","jiang","111")
	client.HSet("name","jiang","222")

	result ,err := client.HGet("name","jiang").Result()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(result)
	/*client.Set("test111","1111",0)
	client.Set("test111","2222",0)*/

	/*fmt.Println(val, err)*/
	/*	for i:=0;i<100;i++{
		fmt.Println( utils.GetRandString(8))
	}*/
	//randomToken := utils.GetRandString(8)

	/*fmt.Println(randomToken)
	//检验token是否已经存在
	for thisValue :=client.Get(randomToken).Val(); thisValue != "";{
		//如果该token存在，生成下一个token
		randomToken = utils.GetRandString(8)
		thisValue = client.Get(randomToken).Val()
	}
	client.Set(randomToken,"random",-1)
	fmt.Println(randomToken)*/

	/*strs := make([]string,10)
	strs=append(strs,"123","321","abs")
	data ,_:=json.Marshal(strs)
	client.Set("user",string(data),time.Hour).Result()



	fmt.Println(client.Get("user").Result())*/
	//string存储
	/*err := client.Set("jaden","my name is jaden",time.Second*5).Err()
	if err != nil{
		panic(err)
	}
	//string获取
	res ,_:=client.Get("jaden").Result()
	println(res)
	//追加字符串
	client.Append("jaden","nancy wang")
	res1 ,_:=client.Get("jaden").Result()
	println(res1)*/
	//队列操作
	/*client.RPush("queue","1","2","3")

	result,_:=client.BLPop(time.Second,"queue").Result()
	for i := range result{
		println(i)
	}*/

	/*err:= client.BgSave().Err()
	fmt.Println(err)*/
	//fmt.Println(client)

}
