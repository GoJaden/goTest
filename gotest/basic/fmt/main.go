package main

import (
	"fmt"
)

func main() {
	loghead := fmt.Sprint("loghead:", "userId=", 10086, "equipmentId=", 555, "type=", 1, ":")

	fmt.Println(loghead)
}
