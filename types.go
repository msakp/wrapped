package wrapped

import (
	"fmt"
	"reflect"
)

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

