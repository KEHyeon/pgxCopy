package utils

import (
	"reflect"
	"sort"
)

func GetColumnFromStruct(i interface{}, tagName string) []string {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var columns []string // 초기화할 때 크기를 미리 지정하지 않고, append로 추가

	// 구조체 필드 순회
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// JSON 태그 추출
		column := field.Tag.Get(tagName)

		// JSON 태그가 비어있지 않으면 배열에 추가
		if column != "" {
			columns = append(columns, column)
		}
	}
	return columns
}

func GetColumnFromMap(inputMap map[string]interface{}) []string {
	// Initialize a slice with the expected capacity based on the number of keys in the map.
	keys := make([]string, 0, len(inputMap))

	// Iterate over the map and append each key to the slice.
	for key := range inputMap {
		keys = append(keys, key)
	}

	// Sort the keys alphabetically.
	sort.Strings(keys)

	// Return the sorted slice of keys.
	return keys
}
