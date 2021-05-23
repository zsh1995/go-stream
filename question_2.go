package stream_test

import (
	"stream_test/stream"
	"unicode"
)

// - Q1: 计算一个 string 中小写字母的个数
func Question2Sub1(str string) int64 {
	// FromString(str).Where((x)=> x.IsLower()).Count()
	c := stream.OfString(str).Filter(func(x interface{}) bool {
		return unicode.IsLower(x.(rune))
	}).Count()
	return int64(c)
}

type tmp struct {
	a     string
	count int64
}

// - Q2: 找出 []string 中，包含小写字母最多的字符串
func Question2Sub2(list []string) string {
	// From(list).Map((x)=> [2]interface{}{x, FromString(x).Where(unicode.IsLower(x)).Count()}).MaxBy(x[1])[0]
	m, _ := stream.OfSlice(list).Map(func(a interface{}) interface{} {
		return tmp{a.(string), Question2Sub1(a.(string))}
	}).Sorted(func(a, b interface{}) int {
		ta := a.(tmp)
		tb := b.(tmp)
		return int(tb.count - ta.count)
	}).First().(tmp)
	return m.a
}
