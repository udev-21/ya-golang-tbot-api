package utils

import (
	"reflect"
)

func NotNil(values ...interface{}) bool {
	for _, value := range values {
		v := reflect.ValueOf(value)
		if value == nil || (v.Kind() == reflect.Ptr && v.IsNil()) {
			return false
		}
	}
	return true
}
