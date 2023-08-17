package linear

import "fmt"

//import "fmt"

const maxSize = 10

type StaticNode struct {
	data int
	cur  int
}

type StaticList struct {
	Element [maxSize]StaticNode
	length  int
}

// 链表初始化
func (link *StaticList) InitList() {
	link.length = 0
	for i := 0; i < maxSize-1; i++ {
		link.Element[i].cur = i + 1
	}
	//目前静态链表为空，最后一个元素的下标为0
	link.Element[maxSize-1].cur = 0
}

// Malloc 返回备用链表的下标，表满返回0
func (link *StaticList) MallocList() int {
	//数组的第一个元素的cur存储备用链表的第一个元素的下标，表满的space[0].cur=0
	i := link.Element[0].cur
	if link.Element[0].cur != 0 {
		link.Element[0].cur = link.Element[i].cur
	}
	return i
}

// FreeSSL 将下标为k的节点回收到备用链表
func (link *StaticList) FreeList(k int) {
	link.Element[k].cur = link.Element[0].cur
	link.Element[0].cur = k
}

// 获取静态链表的长度
func (link *StaticList) ListLength() int {
	return link.length
}

// 输出

//显示链表结构
func (list *StaticList) traverse() {
	for _, v := range list.Element {
		fmt.Printf("%5d", v.cur)
	}
	fmt.Println()
	for _, v := range list.Element {
		fmt.Printf("%5d", v.data)
	}
	fmt.Println()
	for i, _ := range list.Element {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
}

// 判断是否为空
func (l *StaticList) ListEmpty() bool {
	if l.length > 0 {
		return false
	}
	return true
}

// 插入元素
func (l *StaticList) listInsert(value, index int) bool {

	if index < 1 || index > maxSize+1 {
		fmt.Println("插入位置错误")
		return false
	}
	// 获取备用链的位置，表满返回0
	j := l.MallocList()
	fmt.Println("....,", j)
	// k首先是最后一个元素的下标
	k := maxSize - 1
	if j != 0 {
		l.Element[j].data = value
		for m := 1; m < index; m++ {
			k = l.Element[k].cur
		}
		//循环完k是第i-1个元素的下标,把i-1的cur赋值给L[j].cur
		l.Element[j].cur = l.Element[k].cur
		l.Element[k].cur = j
		return true
	}

	return false
}

func StaticListTest() {
	link := new(StaticList)
	link.InitList()
	link.traverse()
	isEmpty := link.ListEmpty()
	fmt.Println(isEmpty)
	//link.listInsert(2, 1)
	//link.traverse()
}
