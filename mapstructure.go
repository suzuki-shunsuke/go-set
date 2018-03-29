package set

import (
	"fmt"
	"reflect"
)

// MapstructureDecodeHookFromListToStrSet decodes []interface{} to StrSet.
// This function implements the mapstructure.DecodeHookFuncType interface.
// https://godoc.org/github.com/mitchellh/mapstructure#DecodeHookFuncType
func MapstructureDecodeHookFromListToStrSet(fromType reflect.Type, toType reflect.Type, from interface{}) (interface{}, error) {
	if reflect.TypeOf(StrSet{}) != toType {
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
