package main

import (
	"fmt"
)

func main() {

	var emptyStringIds []string = make([]string, 0)

	fmt.Println(emptyStringIds)

	fmt.Println(emptyStringIds == nil)
	/*	var t *time.Timer

			t = time.NewTimer(time.Second*1)
		select {
		case <-t.C:
			fmt.Println("123")
			case <-t.C:
				fmt.Println("2222222")
		}*/

}
