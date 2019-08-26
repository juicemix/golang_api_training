package common

import "reflect"

func ValidateStructValues(i interface{}) string {
	ret := ""

	switch v := i.(type) {
	case int, int16, int32, int64, int8:
		if v == 0 {
			return "0"
		} else {
			return "1"
		}
	case float32, float64:
		if v == 0.0 {
			return "0"
		} else {
			return "1"
		}
	case string:
		if v == "" {
			return "0"
		} else {
			return "1"
		}
	default:
		val := reflect.ValueOf(i)

		for a := 0; a < reflect.Indirect(val).NumField(); a++ {
			result := ValidateStructValues(reflect.Indirect(val).Field(a).Interface())
			if result != "1" && result != "" {
				ret = ret + reflect.Indirect(val).Type().Field(a).Name + ","
			}
		}

		return ret
	}
}
