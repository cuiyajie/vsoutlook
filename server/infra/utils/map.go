package utils

import (
	"encoding/json"
	"reflect"
)

type Map = map[string]any

func ShrinkMap(m Map) {
	for k, v := range m {
		if reflect.ValueOf(v).IsZero() {
			delete(m, k)
		}
	}
}

func StructToMap(obj any) (result map[string]any) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	json.Unmarshal(data, &result)
	return
}

func MapToStruct(m Map, s interface{}) {
	for k, v := range m {
		if reflect.TypeOf(s).Elem().Kind() == reflect.Struct {
			field := reflect.ValueOf(s).Elem().FieldByName(k)
			if field.IsValid() && field.CanSet() {
				field.Set(reflect.ValueOf(v))
			}
		}
	}
}

func MapKeys[T any](m map[string]T) []string {
	var result = make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func MapKeysNoneZero[T any](m map[string]T) []string {
	var result = make([]string, 0, len(m))
	for k := range m {
		if len(k) > 0 {
			result = append(result, k)
		}
	}
	return result
}
