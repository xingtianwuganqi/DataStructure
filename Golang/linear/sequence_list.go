package linear

import (
	"errors"
)

const MaxSize int = 100

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

// InsertSequenceList InsertList 把value插入第i位置
func (seq *SequenceList) InsertSequenceList(value, i int) error {
	if i < 0 || i > seq.lenth {
		return errors.New("error insert location")
	}
	for j := seq.lenth - 1; j >= i; j-- {
		seq.data[j] = seq.data[j-1]
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
	for j := i; j < len(seq.data); j++ {
		seq.data[j] = seq.data[j+1]
	}
	seq.lenth--
	return nil
}
