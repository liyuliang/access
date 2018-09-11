package access

import (
	"reflect"
)

func ToByte(obj interface{}) (data []byte) {
	targetObj := reflect.ValueOf(obj)
	structElem := targetObj.Elem()
	structType := structElem.Type()

	for i := 0; i < structElem.NumField(); i++ {
		fieldName := structType.Field(i).Name
		structField := structElem.FieldByName(fieldName)

		switch structField.Type().String() {

		case "string":
			fieldValue := structField.String()
			println(fieldValue)
		case "int":
			fieldValue := structField.Int()
			println(fieldValue)
			break
		case "int32":
			fieldValue := structField.Interface().(int32)
			println(fieldValue)
			break
		case "int64":
			fieldValue := structField.Interface().(int64)
			println(fieldValue)
			break
		case "float32":
			fieldValue := structField.Interface().(float32)
			println(fieldValue)
			break
		case "float64":
			fieldValue := structField.Interface().(float64)
			println(float64(fieldValue))
			break
		case "[]string":
			fieldValue := structField.Interface().([]string)
			for key, value := range fieldValue {
				println(key, value)
			}
			break
		case "map[string]string":
			fieldValue := structField.Interface().(map[string]string)
			for key, value := range fieldValue {
				println(key, value)
			}
			break
		}
	}
	return
}
