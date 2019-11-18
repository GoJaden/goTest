package utils

import "reflect"

func IsEmptyData(data interface{}) bool {
	if data == nil {
		return true
	}
	if value := reflect.ValueOf(data); value.Kind() == reflect.Slice {
		return value.Len() == 0
	}
	return false
}
