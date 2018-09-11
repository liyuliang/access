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

		switch structField.Type().String() {

		case "string":
			fieldValue := structField.String()
			result += `"` + fieldName + `":"` + fieldValue + `",`
		case "int":
			fieldValue := structField.Interface().(int)
			result += `"` + fieldName + `":"` + strconv.Itoa(fieldValue) + `",`
			break
		case "int32":
			fieldValue := structField.Interface().(int32)
			result += `"` + fieldName + `":"` + strconv.FormatInt(int64(fieldValue), 10) + `",`
			break
		case "int64":
			fieldValue := structField.Interface().(int64)
			result += `"` + fieldName + `":"` + strconv.FormatInt(fieldValue, 10) + `",`
			break
		case "float32":
			fieldValue := structField.Interface().(float32)
			result += `"` + fieldName + `":"` + strconv.FormatFloat(float64(fieldValue), 'f', -1, 32) + `",`
			break
		case "float64":
			fieldValue := structField.Interface().(float64)
			result += `"` + fieldName + `":"` + strconv.FormatFloat(fieldValue, 'f', -1, 64) + `",`
			break
		case "[]string":
			fieldValue := structField.Interface().([]string)
			result += `"` + fieldName + `":[`
			for _, value := range fieldValue {
				result += `"` + value + `",`
			}
			result += "],"
			break
		case "map[string]string":
			fieldValue := structField.Interface().(map[string]string)
			result += `"` + fieldName + `":[`
			for key, value := range fieldValue {
				result += `{"` + key + `":"` + value + `"},`
			}
			result += "],"
			break
		}
	}

	result += `}`
	result = strings.Replace(result,",]","]",-1)
	result = strings.Replace(result,",}","}",-1)
	return []byte(result)
}
