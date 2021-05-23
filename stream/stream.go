// Package stream 提供流式数据处理
package stream

import (
	"reflect"
)

// Iterator 迭代器
type Iterator func() (item interface{}, ok bool)

// Stream 流
type Stream struct {
	iterate func() Iterator
}

func (stream Stream) Iterator() Iterator {
	return stream.iterate()
}

type Iterable interface {
	Iterate() Iterator
}

// Of 构造流
func Of(source interface{}) Stream {
	src := reflect.ValueOf(source)
	switch src.Kind() {
	case reflect.Slice:
		return OfSlice(source)
	case reflect.String:
		return OfString(source.(string))
	}
	panic("unsupport source type")
}

// OfSlice 从 slice 构造流
func OfSlice(s interface{}) Stream {
	src := reflect.ValueOf(s)
	if src.Kind() != reflect.Slice {
		panic("s is not a slice")
	}
	len := src.Len()
	return Stream{
		iterate: func() Iterator {
			index := 0
			return func() (item interface{}, ok bool) {
				ok = index < len
				if ok {
					item = src.Index(index).Interface()
					index++
				}
				return
			}
		},
	}
}

// OfString 从 字符串 构造流
func OfString(v string) Stream {
	runes := []rune(v)
	len := len(runes)

	return Stream{
		iterate: func() Iterator {
			index := 0
			return func() (item interface{}, ok bool) {
				ok = index < len
				if ok {
					item = runes[index]
					index++
				}

				return
			}
		},
	}
}

type Predicate func(x interface{}) bool

// Filter 过滤流
func (stream Stream) Filter(p Predicate) Stream {
	return Stream{
		iterate: func() Iterator {
			return filter(p, stream.Iterator())
		},
	}
}

func filter(p Predicate, itr Iterator) Iterator {
	return func() (item interface{}, ok bool) {
		for {
			item, ok = itr()
			if !ok {
				return
			}
			if p(item) {
				return
			}
		}
	}
}

type Mapper func(a interface{}) interface{}

// Map 转换流
func (stream Stream) Map(mapper Mapper) Stream {
	return Stream{iterate: func() Iterator {
		return mapp(mapper, stream.Iterator())
	}}
}

func mapp(mapper Mapper, itr Iterator) Iterator {
	return func() (item interface{}, ok bool) {
		item, ok = itr()
		if ok {
			item = mapper(item)
		}
		return
	}
}
