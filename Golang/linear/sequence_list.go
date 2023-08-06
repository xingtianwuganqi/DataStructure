package linear

import (
	"errors"
	"fmt"
)

const MaxSize int = 20

/*
通过数组实现顺序表
*/

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
func (seqSlice *SequenceList) CreateList(data []int) {
	if len(data) > MaxSize {
		fmt.Println("超过表的大小")
		return
	}
	for i := 0; i < len(data); i++ {
		seqSlice.data[i] = data[i]
		seqSlice.lenth++
	}
}

// RangeList 遍历
func (seqSlice *SequenceList) RangeList() {
	for i := 0; i < seqSlice.lenth; i++ {
		fmt.Printf("元素：%d,第%d个元素\n", seqSlice.data[i], i)
	}
}

// 添加到表的末尾
func (seqSlice *SequenceList) Append(value int) {
	if seqSlice.lenth < len(seqSlice.data) {
		seqSlice.data[seqSlice.lenth] = value
		seqSlice.lenth++
	} else {
		fmt.Println("表已满")
	}
}

// InsertSequenceList InsertList 把value插入第i位置
func (seqSlice *SequenceList) InsertSequenceList(value, i int) error {
	if i < 0 || i > seqSlice.lenth {
		return errors.New("error insert location")
	}
	if seqSlice.lenth >= len(seqSlice.data) {
		return errors.New("表已满")
	}
	for j := seqSlice.lenth - 1; j >= i; j-- {
		seqSlice.data[j+1] = seqSlice.data[j]
	}
	seqSlice.data[i] = value
	seqSlice.lenth++
	return nil
}

// DeleteItem 删除第i个数据
func (seqSlice *SequenceList) DeleteItem(i int) error {
	if i < 0 || i > seqSlice.lenth {
		return errors.New("error delete location")
	}
	for j := i; j < seqSlice.lenth; j++ {
		seqSlice.data[j] = seqSlice.data[j+1]
	}
	seqSlice.lenth--
	return nil
}

// GetValue 获取指定位置的元素
func (seqSlice *SequenceList) GetValue(index int) int {
	if index < 0 || index > seqSlice.lenth {
		return -1
	}
	return seqSlice.data[index]
}

// GetIndex 获取元素下标
func (seqSlice *SequenceList) GetIndex(value int) int {
	for i := 0; i < seqSlice.lenth; i++ {
		if seqSlice.data[i] == value {
			return i
		}
	}
	return -1
}

func (seqSlice *SequenceList) IsEmpty() bool {
	if seqSlice.lenth == 0 {
		return true
	} else {
		return false
	}
}

func (seqSlice *SequenceList) GetSize() int {
	return seqSlice.lenth
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

/*
	适用切片来实现顺序表
*/

type SeqSliceList struct {
	data []int
}

func CreateSeqSliceList() *SeqSliceList {
	return &SeqSliceList{}
}

// Append 添加元素
func (seqSlice *SeqSliceList) Append(value int) {
	seqSlice.data = append(seqSlice.data, value)
}

// InsertValue 插入元素
func (seqSlice *SeqSliceList) InsertValue(value, index int) {
	if index < 0 || index > len(seqSlice.data) {
		fmt.Println("插入位置错误")
		return
	}
	seqSlice.data = append(seqSlice.data[:index], append([]int{value}, seqSlice.data[index:]...)...)
}

// DeleteIndex 删除元素
func (seqSlice *SeqSliceList) DeleteIndex(index int) {
	if index < 0 || index >= len(seqSlice.data) {
		fmt.Println("删除位置错误")
		return
	}
	seqSlice.data = append(seqSlice.data[:index], seqSlice.data[index+1:]...)
}

// TraversingValue 遍历元素
func (seqSlice *SeqSliceList) TraversingValue() {
	for i := 0; i < len(seqSlice.data); i++ {
		var value = seqSlice.data[i]
		fmt.Printf("元素：%d,下标：%d\n", value, i)
	}
}

// GetSize 获取元素个数
func (seqSlice *SeqSliceList) GetSize() int {
	return len(seqSlice.data)
}

// GetIndex 获取下标
func (seqSlice *SeqSliceList) GetIndex(value int) int {
	for i := 0; i < len(seqSlice.data); i++ {
		var data = seqSlice.data[i]
		if value == data {
			fmt.Printf("元素：%d的下标是：%d\n", value, i)
			return i
		}
	}
	return -1
}

// GetValue 通过下标获取值
func (seqSlice *SeqSliceList) GetValue(index int) int {
	for i := 0; i < len(seqSlice.data); i++ {
		if i == index {
			return seqSlice.data[i]
		}
	}
	return -1
}

func SeqSliceListTest() {
	var arr = CreateSeqSliceList()
	arr.Append(0)
	arr.Append(1)
	arr.Append(2)
	arr.Append(3)
	arr.TraversingValue()
	fmt.Println("插入数据：")
	arr.InsertValue(5, 0)
	arr.InsertValue(6, 5)
	arr.InsertValue(7, 7)
	arr.TraversingValue()

	fmt.Println("删除数据")
	arr.DeleteIndex(3)
	arr.TraversingValue()
	fmt.Println("删除最后一个元素")
	arr.DeleteIndex(len(arr.data) - 1)
	arr.TraversingValue()
	value := arr.GetValue(2)
	fmt.Println(value)

	index := arr.GetIndex(3)
	fmt.Println(index)
}
