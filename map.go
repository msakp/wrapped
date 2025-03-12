package wrapped

import (
	"reflect"
	"strings"
)


func fieldMap(s any) map[string]reflect.Value{
	v := reflect.ValueOf(s).Elem()
	result := make(map[string]reflect.Value)
	for i := range v.NumField(){
		name := v.Type().Field(i).Name
		name = strings.ToLower(name)
		result[name] = v.Field(i)
	}
	return result
}
