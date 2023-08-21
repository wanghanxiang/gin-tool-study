package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

func StrToInt(val string) int {
	v1, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return int(v1)
}

// ToString 各种类型转string
// 整数转换为10进制的字符串
func ToString(v interface{}) string {
	t := reflect.TypeOf(v)
	var s string
	switch t.Kind() {
	case reflect.Int:
		s = strconv.FormatInt(int64(v.(int)), 10)
	case reflect.Int64:
		s = strconv.FormatInt(int64(v.(int64)), 10)
	case reflect.Int16:
		s = strconv.FormatInt(int64(v.(int16)), 10)
	case reflect.Int8:
		s = strconv.FormatInt(int64(v.(int8)), 10)
	case reflect.Uint:
		s = strconv.FormatUint(uint64(v.(uint)), 10)
	case reflect.Uint64:
		s = strconv.FormatUint(v.(uint64), 10)
	case reflect.Uint16:
		s = strconv.FormatUint(uint64(v.(uint16)), 10)
	case reflect.Uint8:
		s = strconv.FormatUint(uint64(v.(uint8)), 10)
	case reflect.Bool:
		s = strconv.FormatBool(v.(bool))
	case reflect.Float32:
		// 默认以(-ddd.dddd, no exponent)格式转化浮点数
		s = strconv.FormatFloat(float64(v.(float32)), 'f', -1, 64)
	case reflect.Float64:
		s = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case reflect.Map, reflect.Struct, reflect.Slice:
		s = ToJson(v)
	default:
		fmt.Printf("type %s is not support, use fmt.Sprintf instead", t.Kind())
	}
	return s
}

// ToJson 转成json字符串
func ToJson(m interface{}) string {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// TypeOf 获取变量类型
func TypeOf(data interface{}) string {
	if data == nil {
		return "nil"
	}
	return reflect.TypeOf(data).Kind().String()
}
