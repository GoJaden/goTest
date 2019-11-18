package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var nowTime string = "1568953571"
	t, _ := strconv.Atoi(nowTime)
	ti := time.Unix(int64(t), 0)
	fmt.Println(ti)
}
