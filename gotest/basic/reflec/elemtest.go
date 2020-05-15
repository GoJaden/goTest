package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	e := TestEvent{
		Timestamp: time.Unix(1231421123, 223),
	}
	ToOnWriteEvent(&e)

}

func ToOnWriteEvent(event interface{}) (int, error) {
	/*eventType := reflect.TypeOf(event).Elem()
	 */
	baseField := reflect.Indirect(reflect.ValueOf(event)).FieldByName("Timestamp")
	if baseField.IsValid() {
		b := baseField.Interface().(time.Time)
		fmt.Println("is ", b)
	} else {
		b := time.Now()
		fmt.Println("is not", b)
	}

	return 1, nil
}

type TestEvent struct {
	Timestamp time.Time
}
type Person1 struct {
	Name string
	Age  int
}

func TestElem(i interface{}) {

	ei := reflect.TypeOf(i).Elem()
	fmt.Println("ei:", ei)
	ve := reflect.ValueOf(i).Elem()
	rv := ve.Interface()

	fmt.Println("has", rv)

	fmt.Println(rv)
}
