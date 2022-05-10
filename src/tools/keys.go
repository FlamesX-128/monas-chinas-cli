package tools

import (
	"fmt"
	"os"
	"reflect"
)

func GetKeys(arr interface{}, em interface{}) (k []string) {
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(arr)

		for i := 0; i < s.Len(); i++ {
			k = append(k, s.Index(i).String())

		}

	case reflect.Map:
		s := reflect.ValueOf(arr)

		for _, key := range s.MapKeys() {
			k = append(k, key.String())

		}

	}

	if em != nil && len(k) == 0 {
		fmt.Printf("%s\n\n", em)

		os.Exit(1)
	}

	return
}
