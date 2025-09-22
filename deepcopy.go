package deepcopy

import (
	"reflect"
)

// DeepCopy performs a deep copy of any value using reflection.
// It returns a new, independent copy of the original value.
func DeepCopy(src interface{}) (interface{}, error) {
	if src == nil {
		return nil, nil // correctly handles untyped nil
	}

	originalValue := reflect.ValueOf(src)

	// Special case: typed nil pointer (e.g., (*Person)(nil))
	if originalValue.Kind() == reflect.Ptr && originalValue.IsNil() {
		return nil, nil
	}

	copiedValue, err := deepCopy(originalValue)
	if err != nil {
		return nil, err
	}

	return copiedValue.Interface(), nil
}

// deepCopy is the recursive function that handles the core cloning logic.
func deepCopy(originalValue reflect.Value) (reflect.Value, error) {
	// If the value is invalid or a simple, non-reference type,
	// we can just return it. This handles primitives like int, string, etc.
	if !originalValue.IsValid() {
		return originalValue, nil
	}

	switch originalValue.Kind() {
	case reflect.Ptr:
		if originalValue.IsNil() {
			return reflect.Zero(originalValue.Type()), nil
		}
		newValue := reflect.New(originalValue.Elem().Type())
		copiedElem, err := deepCopy(originalValue.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		newValue.Elem().Set(copiedElem)
		return newValue, nil

	case reflect.Slice:
		if originalValue.IsNil() {
			return reflect.Zero(originalValue.Type()), nil
		}
		sliceType := originalValue.Type()
		newSlice := reflect.MakeSlice(sliceType, originalValue.Len(), originalValue.Cap())
		for i := 0; i < originalValue.Len(); i++ {
			elem, err := deepCopy(originalValue.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			newSlice.Index(i).Set(elem)
		}
		return newSlice, nil

	case reflect.Map:
		if originalValue.IsNil() {
			return reflect.Zero(originalValue.Type()), nil
		}
		mapType := originalValue.Type()
		newMap := reflect.MakeMap(mapType)
		iter := originalValue.MapRange()
		for iter.Next() {
			key, err := deepCopy(iter.Key())
			if err != nil {
				return reflect.Value{}, err
			}
			val, err := deepCopy(iter.Value())
			if err != nil {
				return reflect.Value{}, err
			}
			newMap.SetMapIndex(key, val)
		}
		return newMap, nil

	case reflect.Struct:
		newStruct := reflect.New(originalValue.Type()).Elem()
		for i := 0; i < originalValue.NumField(); i++ {
			field := originalValue.Field(i)
			if !newStruct.Field(i).CanSet() {
				continue
			}
			clonedField, err := deepCopy(field)
			if err != nil {
				return reflect.Value{}, err
			}
			newStruct.Field(i).Set(clonedField)
		}
		return newStruct, nil

	case reflect.Array:
		newArray := reflect.New(originalValue.Type()).Elem()
		for i := 0; i < originalValue.Len(); i++ {
			elem, err := deepCopy(originalValue.Index(i))
			if err != nil {
				return reflect.Value{}, err
			}
			newArray.Index(i).Set(elem)
		}
		return newArray, nil

	case reflect.Interface:
		if originalValue.IsNil() {
			return reflect.Zero(originalValue.Type()), nil
		}
		elem, err := deepCopy(originalValue.Elem())
		if err != nil {
			return reflect.Value{}, err
		}
		return elem.Convert(originalValue.Type()), nil

	case reflect.Chan:
		if originalValue.IsNil() {
			return reflect.Zero(originalValue.Type()), nil
		}
		// create a new channel with same type & buffer size
		newChan := reflect.MakeChan(originalValue.Type(), originalValue.Cap())
		return newChan, nil

	default:
		// For all other types (primitives, funcs, etc.), a shallow copy is enough.
		return originalValue, nil
	}
}
