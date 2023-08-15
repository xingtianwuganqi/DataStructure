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

func (link *StaticList) InitList() {
	link.length = 0
	for i := 0; i < maxSize; i++ {
		link.Element[i] = StaticNode{data: -1, cur: i}
	}
	link.Element[maxSize-1].cur = 0
}

// 输出
func (l *StaticList) Echo() {
	start := l.Element[maxSize-1].cur
	index := start
	for index != 0 {
		fmt.Print(l.Element[index].data, l.Element[index].cur, "-> ")
		index = l.Element[index].cur
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
func (l *StaticList) listInsert(value, index int) {
	if l.length == maxSize {
		fmt.Println("静态链表已满")
		return
	}

	if index < 1 || index > maxSize+1 {
		fmt.Println("插入位置错误")
		return
	}

}

func StaticListTest() {
	link := new(StaticList)
	link.InitList()
	link.Echo()
	isEmpty := link.ListEmpty()
	fmt.Println(isEmpty)
}
