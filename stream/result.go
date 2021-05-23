package stream

import "reflect"

// First 返回流中第一个
func (stream Stream) First() interface{} {
	v, _ := stream.Iterator()()
	return v
}

// FirstN 返回流中前 N 个
func (stream Stream) FirstN(n int, ptrOfSlice interface{}) {
	res := reflect.ValueOf(ptrOfSlice)
	slice := reflect.Indirect(res)

	cap := slice.Cap()
	res.Elem().Set(slice.Slice(0, cap))

	next := stream.Iterator()
	index := 0
	for item, ok := next(); ok && index < n; item, ok = next() {
		slice = reflect.Append(slice, reflect.ValueOf(item))
		index++
	}

	res.Elem().Set(slice.Slice(0, index))
}

// Count 返回流中数目
func (stream Stream) Count() (cnt int) {
	itr := stream.Iterator()
	for {
		_, ok := itr()
		if !ok {
			return
		}
		cnt++
	}
}

// ToSlice 将流收集为 slice
func (stream Stream) ToSlice(ptrOfSlice interface{}) {
	res := reflect.ValueOf(ptrOfSlice)
	slice := reflect.Indirect(res)
	cap := slice.Cap()
	res.Elem().Set(slice.Slice(0, cap))

	next := stream.Iterator()
	index := 0
	for item, ok := next(); ok; item, ok = next() {
		slice = reflect.Append(slice, reflect.ValueOf(item))
		index++
	}
	res.Elem().Set(slice.Slice(0, index))
}
