package stream_test

import (
	"stream_test/stream"
	"strings"
)

// - Q1: 输入一个整数 int，字符串string。将这个字符串重复n遍返回
func Question3Sub1(str string, n int) string {
	r := stream.OfSlice(make([]int, n)).
		Map(func(a interface{}) interface{} { return str }).
		ReduceWith(&strings.Builder{}, func(a, b interface{}) interface{} {
			a.(*strings.Builder).WriteString(b.(string))
			return a
		})
	return r.(*strings.Builder).String()
}
