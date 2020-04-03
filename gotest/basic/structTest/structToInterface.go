package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func main() {
	data := []string{"123", "3213", "fdafd", "在jog辅导费"}
	b, err := Encode(data)
	if err != nil {
		//错误处理
	}
	fmt.Println(string(b))
	/*if err := Decode(b, &to); err != nil {
		//错误处理
	}*/

}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
