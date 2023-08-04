package linear

import (
	"errors"
	"fmt"
)

const MaxSize int = 20

type SequenceList struct {
	data  []int
	lenth int
}

// NewSqeList  声明并初始化顺序表，为顺序表分配内存空间；
func NewSqeList() *SequenceList {
	if MaxSize == 0 {
		return nil
	}
	return &SequenceList{data: make([]int, MaxSize, MaxSize), lenth: 0}
}

// CreateList 线性表赋值
func (seq *SequenceList) CreateList(data []int) {
	if len(data) > MaxSize {
		fmt.Println("超过表的大小")
		return
	}
	for i := 0; i < len(data); i++ {
		seq.data[i] = data[i]
		seq.lenth++
	}
}

// RangeList 遍历
func (seq *SequenceList) RangeList() {
	for i := 0; i < seq.lenth; i++ {
		fmt.Printf("元素：%d,第%d个元素\n", seq.data[i], i)
	}
}

// 添加到表的末尾
func (seq *SequenceList) Append(value int) {
	if seq.lenth < len(seq.data) {
		seq.data[seq.lenth] = value
		seq.lenth++
	} else {
		fmt.Println("表已满")
	}
}

// InsertSequenceList InsertList 把value插入第i位置
func (seq *SequenceList) InsertSequenceList(value, i int) error {
	if i < 0 || i > seq.lenth {
		return errors.New("error insert location")
	}
	if seq.lenth >= len(seq.data) {
		return errors.New("表已满")
	}
	for j := seq.lenth - 1; j >= i; j-- {
		seq.data[j+1] = seq.data[j]
	}
	seq.data[i] = value
	seq.lenth++
	return nil
}

// DeleteItem 删除第i个数据
func (seq *SequenceList) DeleteItem(i int) error {
	if i < 0 || i > seq.lenth {
		return errors.New("error delete location")
	}
	for j := i; j < seq.lenth; j++ {
		seq.data[j] = seq.data[j+1]
	}
	seq.lenth--
	return nil
}

// GetValue 获取指定位置的元素
func (seq *SequenceList) GetValue(index int) int {
	if index < 0 || index > seq.lenth {
		return -1
	}
	return seq.data[index]
}

// GetIndex 获取元素下标
func (seq *SequenceList) GetIndex(value int) int {
	for i := 0; i < seq.lenth; i++ {
		if seq.data[i] == value {
			return i
		}
	}
	return -1
}

func (seq *SequenceList) IsEmpty() bool {
	if seq.lenth == 0 {
		return true
	} else {
		return false
	}
}

func (seq *SequenceList) GetSize() int {
	return seq.lenth
}

func Text() {
	var arr = NewSqeList()
	arrData := []int{3, 4, 5, 6}
	arr.CreateList(arrData)
	arr.RangeList()

	fmt.Println("----")

	arr.InsertSequenceList(2, 0)
	arr.InsertSequenceList(5, 5)
	arr.Append(7)
	arr.RangeList()
	size := arr.GetSize()
	fmt.Printf("表元素的长度%d\n", size)
	arr.DeleteItem(2)
	arr.RangeList()

	vale := arr.GetValue(5)
	fmt.Printf("下标元素为5的元素：%d\n", vale)

	index := arr.GetIndex(6)
	fmt.Printf("元素6的下标：%d\n", index)
}
