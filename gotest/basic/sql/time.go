package utils

import (
	"fmt"
	"time"
)

func DateTimeing() {

	t, err := time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
	fmt.Println(err)
	fmt.Println(t.Unix())
}
