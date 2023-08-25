package stack

import "fmt"

/*
链栈的实现思路和顺序栈类似，顺序栈是将顺序表（数组）的一端做栈底，另一端做栈顶；
链栈也是如此，我们通常将链表的头部做栈顶，尾部做栈底，如图 1 所示：
*/

type LinkStackNode struct {
	data int
	next *LinkStackNode
}

type LinkStackList struct {
	head *LinkStackNode
}

// 将item插入到链表的head位置
func (l *LinkStackList) push(item *LinkStackNode) {
	if l.head == nil {
		l.head = item
	} else {
		item.next = l.head
		l.head = item
	}
}

func (l *LinkStackList) pop() {
	if l.head != nil {
		var temp = l.head.next
		l.head = nil
		l.head = temp
	} else {
		panic("link is Empty")
	}
}

func (l *LinkStackList) PrintNode() {
	if l.head == nil {
		fmt.Println("l is empty")
	} else {
		var temp = l.head
		for temp != nil {
			if temp.next == nil {
				fmt.Printf("%d\n", temp.data)
			} else {
				fmt.Printf("%d -> ", temp.data)
			}
			temp = temp.next
		}
	}
}

func LinkStackTest() {
	list := &LinkStackList{head: nil}
	for i := 0; i <= 5; i++ {
		var nodeOne = &LinkStackNode{data: i, next: nil}
		list.push(nodeOne)
	}
	list.PrintNode()
	list.pop()
	list.PrintNode()
	list.pop()
	list.PrintNode()
}
