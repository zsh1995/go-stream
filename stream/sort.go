package stream

import "sort"

// Comparator 比较器，a > b = (>0) 1, a == b = 0, a < b = -1 (<0) 1
type Comparator func(a, b interface{}) int

// OrderedStream 排序流
type OrderedStream struct {
	Stream
	order Comparator
}

func newOrderedStream(itr Iterator, cmp Comparator) OrderedStream {
	return OrderedStream{
		Stream: Stream{
			iterate: func() Iterator {
				var r []interface{}
				for next, ok := itr(); ok; next, ok = itr() {
					r = append(r, next)
				}
				sort.Slice(r, func(i, j int) bool {
					return cmp(r[i], r[j]) < 0
				})
				lenght := len(r)
				index := 0
				return func() (item interface{}, ok bool) {
					ok = index < lenght
					if ok {
						item = r[index]
						index++
					}
					return
				}
			},
		},
		order: cmp,
	}
}

// Sorted 排序
func (stream Stream) Sorted(cmp Comparator) OrderedStream {
	return newOrderedStream(stream.Iterator(), cmp)
}

// Sorted 排序
func (stream OrderedStream) Sorted(cmp Comparator) OrderedStream {
	return newOrderedStream(stream.Iterator(), func(a, b interface{}) int {
		t := stream.order(a, b)
		if t != 0 {
			return t
		}
		return cmp(a, b)
	})
}
