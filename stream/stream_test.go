// Package stream 提供流式数据处理
package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	t.Run("simple stream", func(t *testing.T) {
		stream := OfSlice([]int{4, 1, 3, 6, 2})
		t.Run("count", func(t *testing.T) {
			c := stream.Count()
			assert.Equal(t, c, 5)
		})
		t.Run("filter", func(t *testing.T) {
			c := stream.Filter(func(x interface{}) bool { return x.(int) >= 2 }).Count()
			assert.Equal(t, c, 4)
		})
		t.Run("to slice", func(t *testing.T) {
			var result []int
			stream.Filter(func(x interface{}) bool { return x.(int) >= 2 }).ToSlice(&result)
			assert.Equal(t, result, []int{4, 3, 6, 2})
		})
		t.Run("reduce", func(t *testing.T) {
			sum := stream.Filter(func(x interface{}) bool { return x.(int) >= 6 }).Reduce(func(a, b interface{}) interface{} {
				return a.(int) + b.(int)
			})
			assert.Equal(t, sum, 6)
		})
		t.Run("first", func(t *testing.T) {
			var result []int
			stream.FirstN(1, &result)
			assert.Equal(t, result, []int{4})
		})
		t.Run("reduce with value", func(t *testing.T) {
			sum := stream.ReduceWith(10, func(a, b interface{}) interface{} {
				return a.(int) + b.(int)
			})
			assert.Equal(t, sum, 10+16)
		})
	})
	t.Run("sorted stream", func(t *testing.T) {
		stream := OfSlice([]int{4, 1, 3, 6, 2})
		t.Run("sort", func(t *testing.T) {
			var result []int
			stream.Sorted(func(a, b interface{}) int { return a.(int) - b.(int) }).ToSlice(&result)
			assert.Equal(t, result, []int{1, 2, 3, 4, 6})
		})

	})
}

func BenchmarkStream(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := Of([]int{3, 3, 5, 2, 34})
		var result []int
		s.ToSlice(&result)
	}
}
// BenchmarkStream-8   	 1131432	      1029 ns/op	     600 B/op	      22 allocs/op
// BenchmarkStream-8   	 1210950	      1099 ns/op	     568 B/op	      21 allocs/op
