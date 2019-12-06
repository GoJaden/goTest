package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("errrr")

	err = errors.Wrap(err, "附加信息")

	fmt.Printf("%+v", err)
}
