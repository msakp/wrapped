package wrapped

import (
	"fmt"
	"reflect"
	"strings"
)

// Wraps 's' struct to [Type] struct,
// panics if [Type] or 's' type is not a defined struct
func To[Type any](s any) (*Type){
	result := new(Type)
	
	// checks for result and s type being a defined struct
	CheckTypes(result, s)
		
	sValue := reflect.ValueOf(s).Elem()
	sFieldMap := fieldMap(sValue)
	fmt.Println(sFieldMap)


	return result
}


// check if wanted, and existing Type are both defined structs
func CheckTypes(wanted, existing any){
	dataKind := reflect.ValueOf(existing).Elem().Type().Kind()
	resultKind := reflect.ValueOf(wanted).Elem().Type().Kind()
	
	// [Type] and s value type check for being structs
	if resultKind.String() != "struct"{
		panic(fmt.Errorf("\nWrapped ERR:\n\nBad wrapping %T with %T: %T must be a *struct", existing, wanted, wanted))
	}
	if dataKind.String() != "struct"{
		fmt.Println(dataKind.String())
		panic(fmt.Errorf("\nWrapped ERR:\n\nBad wrapping %T with %T: %T must be a *struct", existing, wanted, existing))
	}
}

func fieldMap(v reflect.Value) map[string]reflect.Value{
	result := make(map[string]reflect.Value)
	for i := range v.NumField(){
		name := v.Type().Field(i).Name
		name = strings.ToLower(name)
		result[name] = v.Field(i)
	}
	return result
}
