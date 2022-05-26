package utils

import (
	"encoding/json"
	"strconv"
)

func GetInterfaceToString(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func GetInterfaceToInt(value interface{}) int {
	var key int
	switch value.(type) {
	case uint:
		key = int(value.(uint))
		break
	case int8:
		key = int(value.(int8))
		break
	case uint8:
		key = int(value.(uint8))
		break
	case int16:
		key = int(value.(int16))
		break
	case uint16:
		key = int(value.(uint16))
		break
	case int32:
		key = int(value.(int32))
		break
	case uint32:
		key = int(value.(uint32))
		break
	case int64:
		key = int(value.(int64))
		break
	case uint64:
		key = int(value.(uint64))
		break
	case float32:
		key = int(value.(float32))
		break
	case float64:
		key = int(value.(float64))
		break
	case string:
		key, _ = strconv.Atoi(value.(string))
		break
	default:
		key = value.(int)
		break
	}
	return key
}
