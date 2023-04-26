package utils

import "reflect"

func F_InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				exists = true
				index = i
				return exists, index
			}
		}
	}

	return exists, index
}
