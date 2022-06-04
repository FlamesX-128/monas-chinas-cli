package tools

import (
	"log"
	"reflect"
)

func GetKeys[T string, Y any](m map[T]Y) (res []T) {
	for k := range m {
		res = append(res, k)

	}

	return
}

func Guard[T any](val T, err error) T {
	if err != nil {
		log.Panicln(err.Error())
	}

	return val
}

func FindFields[T any, Y any](arr []Y, field string) (res []T) {
	for _, e := range arr {
		res = append(res,
			reflect.ValueOf(e).FieldByName(field).Interface().(T),
		)
	}

	return
}

func FindIndex(arr []string, target string) (res int) {
	for i, e := range arr {
		if e == target {

			return i
		}

	}

	return -1
}
