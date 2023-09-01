package linear

import "fmt"

type doubleCircularNode struct {
	prior *doubleCircularNode
	data  int
	next  *doubleCircularNode
}

func initDoubleCircularList(head *doubleCircularNode) *doubleCircularNode {
	var temp = head
	for i := 2; i <= 3; i++ {
		current := &doubleCircularNode{prior: nil, data: i, next: nil}
		temp.next = current
		current.prior = temp
		temp = current
	}
	head.prior = temp
	temp.next = head
	return head
}

func displayDoubleCircular(head *doubleCircularNode) {
	var temp = head
	for temp.next != head {
		if temp.next == nil {
			fmt.Printf("%d\n", temp.data)
		} else {
			fmt.Printf("%d<->", temp.data)
		}
		temp = temp.next
	}
	// 输出最后一个元素
	fmt.Printf("%d\n", temp.data)
}

func DoubleCircularTest() {
	var head = &doubleCircularNode{
		prior: nil,
		data:  1,
		next:  nil,
	}
	list := initDoubleCircularList(head)
	displayDoubleCircular(list)
}
