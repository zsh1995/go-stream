package stream_test

import (
	"stream_test/stream"

	"github.com/Pallinder/go-randomdata"
)

// - Q1: 输入 employees，返回 年龄 >22岁 的所有员工，年龄总和
func Question1Sub1(employees []*Employee) int64 {
	// Where((x)=> x.age > 22).Reduce((a, b)=> return a.age + b.age, 0)
	c := stream.OfSlice(employees).Filter(func(x interface{}) bool { return (*x.(*Employee).Age) > 22 }).
		ReduceWith(0, func(a, b interface{}) interface{} {
			return a.(int) + (*b.(*Employee).Age)
		})
	return int64(c.(int))
}

// - Q2: - 输入 employees，返回 id 最小的十个员工，按 id 升序排序
func Question1Sub2(employees []*Employee) (result []*Employee) {
	// OrderBy((a, b) => a.id < b.id).First(10)
	stream.OfSlice(employees).Sorted(func(a, b interface{}) int {
		return int(a.(*Employee).Id - b.(*Employee).Id)
	}).FirstN(10, &result)
	return
}

// - Q3: - 输入 employees，对于没有手机号为0的数据，随机填写一个
func Question1Sub3(employees []*Employee) (result []*Employee) {
	// Where((x)=> x.phoneNumber==0).ForEach(x=>x.phone=random)
	stream.OfSlice(employees).Map(func(a interface{}) interface{} {
		ea := a.(*Employee)
		cp := *ea
		if ea.Phone == nil {
			phone := randomdata.RandStringRunes(10)
			cp.Phone = &phone
		}
		return &cp
	}).ToSlice(&result)
	return
}

// - Q4: - 输入 employees ，返回一个map[int][]int，其中 key 为 员工年龄 Age，value 为该年龄段员工ID
func Question1Sub4(employees []*Employee) (resp map[int][]int64) {
	// Reduce((a, b)=> a[b.age].apend(b.id), map[int][]int)
	resp = make(map[int][]int64)
	stream.OfSlice(employees).ReduceWith(resp, func(a, b interface{}) interface{} {
		result := a.(map[int][]int64)
		cur := b.(*Employee)
		result[*cur.Age] = append(result[*cur.Age], cur.Id)
		return result
	})
	return
}
