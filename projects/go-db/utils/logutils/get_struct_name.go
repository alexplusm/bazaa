package logutils

import "reflect"

func GetStructName(value interface{}) string {
	if t := reflect.TypeOf(value); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name() + ":"
	} else {
		return t.Name() + ":"
	}
}
