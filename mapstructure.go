package set

import (
	"fmt"
	"reflect"
)

var strSetType = reflect.TypeOf(StrSet{})

// MapstructureDecodeHookFromListToStrSet decodes []interface{} to StrSet.
// This function implements the mapstructure.DecodeHookFuncType interface.
// https://godoc.org/github.com/mitchellh/mapstructure#DecodeHookFuncType
func MapstructureDecodeHookFromListToStrSet(fromType, toType reflect.Type, from interface{}) (interface{}, error) {
	if strSetType != toType {
		return from, nil
	}
	if arr, ok := from.([]interface{}); ok {
		s := NewStrSet()
		for _, a := range arr {
			if k, ok := a.(string); ok {
				s.Add(k)
				continue
			}
			return from, fmt.Errorf("type assertion to string is failure: %v", a)
		}
		return s, nil
	}
	return from, nil
}
