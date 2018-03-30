package lib

import (
	"reflect"
)

// Text 文本方法,如枚举文本,只要实现Text()方法即可
func Text(v interface{}) string {
	val := reflect.ValueOf(v)
	t := val.Type()
	// fmt.Printf("val:%v,type:%v \r\n", val, t)
	nums := val.NumMethod()
	// Text 方法
	for i := 0; i < nums; i++ {
		method := t.Method(i).Name
		// fmt.Printf("method:%v\r\n", method)
		if method == "Text" {
			// fmt.Printf("text method:%v\r\n", method)
			result := val.Method(i).Call(nil)
			return result[0].Interface().(string)
		}
	}
	// String 方法
	return val.String()
}
