package utils

import (
	"net/url"
	"reflect"
)

func EncodeQuery(data interface{}) string {
	params := url.Values{}
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段
			name := []byte(t.Field(i).Name)
			val := v.Field(i).Interface().(string)
			name[0] ^= ' '
			if len(val) != 0 {
				params.Set(string(name), val)
			}
		}
	}
	return params.Encode()
}
