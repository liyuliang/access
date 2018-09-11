package access

import (
	"reflect"
	"strconv"
	"strings"
)

func ToByte(obj interface{}) []byte {
	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()

	result := `{`
	for i := 0; i < structElem.NumField(); i++ {
		fieldName := structType.Field(i).Name
		structField := structElem.FieldByName(fieldName)
		fieldValue := structField.Interface()

		switch fieldValue.(type) {

		default:
			if strings.Contains(structField.Type().String(), "*") {
				result += `"` + fieldName + `":"` + ObjToStr(fieldValue) + `",`

			} else {
				result += `"` + fieldName + `":"` + ValToStr(fieldValue) + `",`
			}

		case []string:
			result += `"` + fieldName + `":[`
			result += ValToStr(fieldValue)
			result += "],"
		case map[string]string:
			result += `"` + fieldName + `":[`
			result += ValToStr(fieldValue)
			result += "],"
		}
	}

	result += `}`
	result = strings.Replace(result, ",]", "]", -1)
	result = strings.Replace(result, ",}", "}", -1)
	return []byte(result)
}

func ValToStr(val interface{}) (result string) {

	switch val.(type) {
	case string:
		result = val.(string)
	case int:
		result = strconv.Itoa(val.(int))
	case int32:
		result = strconv.FormatInt(int64(val.(int32)), 10)
	case int64:
		result = strconv.FormatInt(val.(int64), 10)
	case float32:
		result = strconv.FormatFloat(float64(val.(float32)), 'f', -1, 32)
	case float64:
		result = strconv.FormatFloat(val.(float64), 'f', -1, 64)
	case []string:
		for _, value := range val.([]string) {
			result += `"` + value + `",`
		}

	case map[string]string:
		field := val.(map[string]string)

		for key, value := range field {
			result += `{"` + key + `":"` + value + `"},`
		}
	}

	return result
}

func ObjToStr(obj interface{}) (result string) {

	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()
	result += "["
	for i := 0; i < structElem.NumField(); i++ {

		fieldName := structType.Field(i).Name
		fieldValue := structElem.FieldByName(fieldName).Interface()

		result += `{"` + fieldName + `":"` + ValToStr(fieldValue) + `"},`
	}
	result += "]"
	return result
}