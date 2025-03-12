package wrapped

import (
	"reflect"
	"strings"
)



func initStruct[Type any](zero Type, fieldMap map[string]reflect.Value) Type{
	zeroV := reflect.ValueOf(zero)
	zeroT := zeroV.Elem().Type()
	for i := range zeroV.Elem().NumField(){
		f := zeroV.Elem().Field(i)
		ft := zeroT.Field(i)
		name := strings.ToLower(ft.Name)
		if _, ok := fieldMap[name]; !ok{
			continue
		}
		f.Set(fieldMap[name])
	}
	return zeroV.Interface().(Type)
}


