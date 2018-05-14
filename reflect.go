package gotool

import (
	"reflect"
)

func Inject(target interface{}, kv map[string]interface{}){
	for k, v := range kv{
		elem := reflect.ValueOf(target).Elem()
		field := elem.FieldByName(k)

		var val reflect.Value
		if reflect.ValueOf(v).Type() == field.Type() {
			val = reflect.ValueOf(v)
		}else{
			val = reflect.ValueOf(v).Convert(field.Type())
		}
		field.Set(val)
	}
}
