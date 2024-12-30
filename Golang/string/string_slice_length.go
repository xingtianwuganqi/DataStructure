package string

import "fmt"

/*
2.串的堆区存储结构
*/
type Slice struct {
	data []rune
}

func NewStringSlice(s string) *Slice {
	runes := []rune(s)
	return &Slice{
		data: runes,
	}
}

func (s *Slice) string() string {
	return string(s.data)
}

// 插入
func (s *Slice) insertStr(insertStr string) *Slice {
	insertRunes := []rune(insertStr)
	s.data = append(s.data, insertRunes...)
	return s
}

func StringSliceTest() {
	data := NewStringSlice("hello")
	data = data.insertStr(" world")
	fmt.Println(data.string())
}
