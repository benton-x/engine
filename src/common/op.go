package common

import (
	"context"
	"strconv"
)

// OpType ...
type OpType string

const (
	OpTypeIg   OpType = "ig"
	OpTypeEq   OpType = "eq"
	OpTypeNoEq OpType = "neq"
	OpTypeGt   OpType = "gt"
	OpTypeLt   OpType = "lt"
)

// Calc ...
func Calc(ctx context.Context, opType OpType, condVal interface{}, realVal interface{}) bool {
	switch opType {
	case OpTypeEq:
		return Equal(condVal, realVal)
	case OpTypeNoEq:
		return NoEqual(condVal, realVal)
	case OpTypeGt:
		return More(condVal, realVal)
	case OpTypeLt:
		return Less(condVal, realVal)
	case OpTypeIg:
		return true
	}

	return false
}

// Equal ...
func Equal(condVal, realVal interface{}) bool {
	switch condVal.(type) {
	case int:
		return condVal.(int) == ToInt(realVal)
	case int32:
		return condVal.(int32) == ToInt32(realVal)
	case int64:
		return condVal.(int64) == ToInt64(realVal)
	case string:
		return condVal.(string) == ToString(realVal)
	case bool:
		return condVal.(bool) == ToBool(realVal)
	}
	return false
}

// NoEqual ...
func NoEqual(condVal, realVal interface{}) bool {
	return !Equal(condVal, realVal)
}

// More ...
func More(condVal, reqVal interface{}) bool {
	switch reqVal.(type) {
	case string:
		return ToFloat64(reqVal) > ToFloat64(condVal)
	case int:
		return reqVal.(int) > ToInt(condVal)
	case int64:
		return reqVal.(int64) > ToInt64(condVal)
	case float64:
		return reqVal.(float64) > ToFloat64(condVal)
	default:
	}
	return false
}

// Less ...
func Less(condVal, reqVal interface{}) bool {
	switch reqVal.(type) {
	case string:
		val, err := strconv.ParseFloat(reqVal.(string), 64)
		if err != nil {
			return false
		}
		return val < ToFloat64(condVal)
	case int:
		return reqVal.(int) < ToInt(condVal)
	case int64:
		return reqVal.(int64) < ToInt64(condVal)
	case float64:
		return reqVal.(float64) < ToFloat64(condVal)
	default:
	}
	return false
}
