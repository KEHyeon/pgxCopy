package utils

import "reflect"

func GetValuesFromStruct(i interface{}, tagName string) []interface{} {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	var values []interface{}

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)

		// JSON 태그 추출
		tag := field.Tag.Get(tagName)
		if tag != "" {
			value := val.Field(i).Interface()
			values = append(values, value)
		}
	}

	return values
}

// GetValuesFromMap extracts values from a map[string]interface{} based on the order
// of columns provided in the orderedColumns slice, and returns them in the same order.
func GetValuesFromMap(inputMap map[string]interface{}, orderedColumns []string) []interface{} {
	// Create a slice with pre-allocated capacity based on the orderedColumns length.
	values := make([]interface{}, 0, len(orderedColumns))

	// Iterate over the orderedColumns and extract the corresponding values from the map.
	for _, column := range orderedColumns {
		if value, exists := inputMap[column]; exists {
			// Directly append the value into the pre-allocated slice.
			values = append(values, value)
		} else {
			// Optionally handle the case where a column is missing in the map.
			values = append(values, nil)
		}
	}

	return values
}
