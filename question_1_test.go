package stream_test_test

import (
	"stream_test"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptrOfInt64(v int64) *int64 {
	return &v
}

func ptrOfInt(v int) *int {
	return &v
}

func ptrOfString(v string) *string {
	return &v
}

func TestQuestion1Sub1(t *testing.T) {
	answer := stream_test.Question1Sub1([]*stream_test.Employee{
		{Age: ptrOfInt(23)},
		{Age: ptrOfInt(27)}})
	assert.EqualValues(t, answer, 50)
}

func TestQuestion1Sub2(t *testing.T) {
	answer := stream_test.Question1Sub2([]*stream_test.Employee{
		{Id: 11},
		{Id: 10},
		{Id: 9},
		{Id: 8},
		{Id: 7},
		{Id: 6},
		{Id: 5},
		{Id: 4},
		{Id: 3},
		{Id: 2},
		{Id: 1},
	})
	assert.EqualValues(t, answer[0].Id, 1)
	assert.EqualValues(t, answer[9].Id, 10)
}

func TestQuestion1Sub3(t *testing.T) {
	answer := stream_test.Question1Sub3([]*stream_test.Employee{
		{Phone: ptrOfString("123")},
		{Phone: nil}})
	assert.NotNil(t, answer[1].Phone)
}

func TestQuestion1Sub4(t *testing.T) {
	answer := stream_test.Question1Sub4([]*stream_test.Employee{
		{Id: 1, Age: ptrOfInt(10)},
		{Id: 2, Age: ptrOfInt(11)},
		{Id: 3, Age: ptrOfInt(10)},
		{Id: 4, Age: ptrOfInt(11)},
	})
	assert.Equal(t, answer[10], []int64{1, 3})
	assert.Equal(t, answer[11], []int64{2, 4})
}

func TestQuestion2Sub1(t *testing.T) {
	answer := stream_test.Question2Sub1(strings.Repeat("a", 10) + strings.Repeat("A", 21))
	assert.EqualValues(t, answer, 10)
}

func TestQuestion2Sub2(t *testing.T) {
	answer := stream_test.Question2Sub2([]string{
		strings.Repeat("a", 10) + strings.Repeat("A", 21),
		strings.Repeat("A", 21) + strings.Repeat("a", 10),
		strings.Repeat("A", 21) + strings.Repeat("a", 11),
		strings.Repeat("A", 21) + strings.Repeat("a", 2),
	})
	assert.Equal(t, answer, strings.Repeat("A", 21)+strings.Repeat("a", 11))
}

func TestQuestion3Sub1(t *testing.T) {
	answer := stream_test.Question3Sub1("123", 10)
	assert.Equal(t, answer, strings.Repeat("123", 10))
}
