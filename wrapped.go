package wrapped


// Wraps 's' struct to [Type] struct,
// panics if [Type] or 's' type is not a defined struct
func To[Type any](s any) (*Type){
	result := new(Type)
	
	// checks for result and s type being a defined struct
	CheckTypes(result, s)
		
	sFieldMap := fieldMap(s)

	newV := initStruct(result, sFieldMap)
	return newV
}

















