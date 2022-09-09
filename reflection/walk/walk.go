// Package walk walks.
package walk

import "reflect"

// Walk walks.
func Walk(x interface{}, function func(string)) {
	val := getValue(x)
	var noOfValues int
	var getField func(int) reflect.Value

	switch val.Kind() { //nolint
	case reflect.Slice, reflect.Array:
		noOfValues = val.Len()
		getField = val.Index
	case reflect.Struct:
		noOfValues = val.NumField()
		getField = val.Field
	case reflect.String:
		function(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), function)
		}
	}

	for i := 0; i < noOfValues; i++ {
		Walk(getField(i).Interface(), function)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
