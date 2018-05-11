package gotool

import (
	"reflect"
	"fmt"
)

func Inject(target interface{}, kv map[string]interface{}) error{
	var success = true

	for k, v := range kv{
		elem := reflect.ValueOf(target).Elem()
		field := elem.FieldByName(k)
		if reflect.ValueOf(v).Type() == field.Type() {
			field.Set(reflect.ValueOf(v))
		}else{
			success = false
		}
	}

	if !success{
		return fmt.Errorf("some field(s) can not inject")
	}

	return nil
}
