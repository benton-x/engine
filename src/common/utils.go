package common

import (
	"fmt"
	"strings"
	"reflect"
	"strconv"
)

func ToInt(v interface{}) int {
	switch v.(type) {
	case int:
		return v.(int)
	case int32:
		return int(v.(int32))
	case int64:
		return int(v.(int64))
	case string:
		val, _ := strconv.Atoi(v.(string))
		return int(val)
	case float32:
		return int(v.(float32))
	case float64:
		return int(v.(float64))
	case bool:
		val, ok := v.(bool)
		if ok && val {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func ToInt32(v interface{}) int32 {
	switch v.(type) {
	case int:
		return int32(v.(int))
	case int32:
		return v.(int32)
	case int64:
		return int32(v.(int64))
	case string:
		val, _ := strconv.Atoi(v.(string))
		return int32(val)
	case float32:
		return int32(v.(float32))
	case float64:
		return int32(v.(float64))
	case bool:
		val, ok := v.(bool)
		if ok && val {
			return 1
		}
		return 0
	default:
		return 0
	}
}

func ToInt64(v interface{}) int64 {
	switch v.(type) {
	case int:
		return int64(v.(int))
	case int32:
		return int64(v.(int32))
	case int64:
		return v.(int64)
	case string:
		val, _ := strconv.Atoi(v.(string))
		return int64(val)
	case float32:
		return int64(v.(float32))
	case float64:
		return int64(v.(float64))
	case bool:
		val, ok := v.(bool)
		if ok && val {
			return 1
		}
		return 0
	default:
		return 0
	}
}

//ToFloat32 : parse the input data into float64
func ToFloat32(v interface{}) float32 {
	switch v.(type) {
	case int:
		return float32(v.(int))
	case int64:
		return float32(v.(int64))
	case string:
		val, _ := strconv.ParseFloat(v.(string), 32)
		return float32(val)
	case float32:
		return v.(float32)
	default:
		return 0
	}
}

//ToFloat64 : parse the input data into float64
func ToFloat64(v interface{}) float64 {
	switch v.(type) {
	case int:
		return float64(v.(int))
	case int64:
		return float64(v.(int64))
	case string:
		val, _ := strconv.ParseFloat(v.(string), 64)
		return val
	case float64:
		return v.(float64)
	default:
		return 0
	}
}

// ConvertToFloat32 ...
func ConvertToFloat32(value interface{}) (float32, error) {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return float32(reflect.ValueOf(value).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float32(reflect.ValueOf(value).Uint()), nil
	case float32, float64:
		return float32(reflect.ValueOf(value).Float()), nil
	case string:
		v, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}
		return float32(v), nil
	default:
		return 0, fmt.Errorf("can not convert value: %v to float32", value)
	}
}

// ConvertToFloat64 ...
func ConvertToFloat64(value interface{}) (float64, error) {
	switch value.(type) {
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(value).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(value).Uint()), nil
	case float32, float64:
		return reflect.ValueOf(value).Float(), nil
	case string:
		v, err := strconv.ParseFloat(value.(string), 64)
		if err != nil {
			return 0, err
		}
		return v, nil
	default:
		return 0, fmt.Errorf("can not convert value: %v to float64", value)
	}
}

//ToBool : parse the input data into bool
func ToBool(v interface{}) bool {
	switch v.(type) {
	case bool:
		return v.(bool)
	case string:
		if strings.ToUpper(v.(string)) == "FALSE" {
			return false
		}
		if strings.ToUpper(v.(string)) == "TRUE" {
			return true
		}
		return false
	case int:
		return v.(int) > 0
	case int64:
		return v.(int64) > 0
	case float64:
		return v.(float64) > 0
	default:
		return false
	}
}

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v)
}
