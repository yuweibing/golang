package main

//只执行一次实例操作
import (
	"fmt"
	"sync"
)

type SliceNum []int

func NewSlice() SliceNum {
	return make(SliceNum, 0)
	// return make(SliceNum,0)
	// SliceNum = append(SliceNum[], 888)

}

func (s *SliceNum) Add(elem int) *SliceNum {
	*s = append(*s, elem)
	fmt.Println("add", elem)
	fmt.Println("add SliceNum end", s)
	return s
}

func main() {
	var once sync.Once
	s := NewSlice()
	s = append(s, 888)
	// 看源代码理解once的行为
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
	once.Do(func() {
		s.Add(16)
	})
}
