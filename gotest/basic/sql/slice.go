package utils

import (
	"common-lib/collection"
	"reflect"
)

type Selector func(element interface{}) interface{}

var SelfSelector = func(element interface{}) interface{} {
	return element
}

func Foreach(sliceData interface{}, f func(e interface{})) {
	if IsEmptyData(sliceData) || f == nil {
		return
	}
	sliceValues := reflect.ValueOf(sliceData)
	if sliceValues.Kind() != reflect.Slice {
		return
	}
	sliceLength := sliceValues.Len()
	for i := 0; i < sliceLength; i++ {
		f(sliceValues.Index(i).Interface())
	}
}

func FilterSame(data interface{}, selector func(element interface{}) interface{}, f func(element interface{})) {
	if IsEmptyData(data) || f == nil {
		return
	}
	hashSet := collection.NewHashSet()
	Foreach(data, func(e interface{}) {
		if selector == nil {
			selector = SelfSelector
		}
		hashSet.Add(selector(e))
	})
	hashSet.Foreach(f)
}
