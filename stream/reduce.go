package stream

// Combinator 结合器
type Combinator func(old, current interface{}) interface{}

// Reduce reduce 流，初始值为流的第一个值
func (stream Stream) Reduce(combinator Combinator) interface{} {
	itr := stream.Iterator()
	result, any := itr()
	if !any {
		return nil
	}
	for current, ok := itr(); ok; current, ok = itr() {
		result = combinator(result, current)
	}
	return result
}

// ReduceWith 指定初始值 reduce 流
func (stream Stream) ReduceWith(initVal interface{}, combinator Combinator) interface{} {
	itr := stream.Iterator()
	result := initVal

	for current, ok := itr(); ok; current, ok = itr() {
		result = combinator(result, current)
	}

	return result
}
