package linear

import (
	"fmt"
	"strconv"
)

// ListNode 链表节点
type ListNode struct {
	data int
	next *ListNode
}

// LinkedList  定义链表结构
type LinkedList struct {
	head *ListNode
}

// CreateLinkedList 创建链表
func CreateLinkedList() *LinkedList {
	// 首先定义一个头指针
	return &LinkedList{
		head: nil,
	}
}

// AddNodes 添加几个节点
func (link *LinkedList) AddNodes(num int) {
	for i := 0; i < num; i++ {
		link.Insert(i)
	}
}

// Insert 添加元素
func (link *LinkedList) Insert(value int) {
	node := &ListNode{data: value, next: nil}
	if link.head == nil {
		link.head = node
		return
	}
	current := link.head
	for current.next != nil {
		current = current.next
	}
	current.next = node
}

// InsertIndex 插入节点
func (link *LinkedList) InsertIndex(value, index int) {
	if index == 1 {
		node := &ListNode{data: value, next: nil}
		next := link.head.next
		node.next = next
		link.head.next = node
	} else {
		current := link.head
		for i := 1; i < index; i++ {
			current = current.next
			if current.next == nil {
				fmt.Println("插入位置无效")
				return
			}
		}
		c := &ListNode{data: value, next: nil}
		c.next = current.next
		current.next = c
	}
}

// DeleteNode 删除节点
func (link *LinkedList) DeleteNode(value int) int {
	find := -1
	current := link.head
	for current != nil {
		// 找到要删除元素的前驱结点
		if current.next.data == value {
			find = 1
			break
		}
		current = current.next
	}
	if find == -1 {
		return find
	} else {
		current.next = current.next.next
		return 1
	}
}

// DeleteIndex 删除节点
func (link *LinkedList) DeleteIndex(index int) int {
	find := -1
	i := 0
	current := link.head
	for current != nil {
		if i+1 == index {
			find = 1
			break
		}
		// 找到删除节点的前驱节点
		current = current.next
		i++
	}
	current.next = current.next.next
	return find
}

// SelectNode 查询
func (link *LinkedList) SelectNode(value int) int {
	var i = 1
	current := link.head
	for current != nil {
		current = current.next
		if current.data == value {
			return i
		}
		i++
	}
	return -1
}

// ChangeNode 更改节点
func (link *LinkedList) ChangeNode(value, index int) int {
	if index < 1 {
		return -1
	}
	i := 1
	current := link.head
	for current != nil {
		current = current.next
		if i == index {
			current.data = value
			return 1
		}
		i++
	}
	return -1
}

// Print 打印元素
func (link *LinkedList) Print() {
	current := link.head
	for current != nil {
		fmt.Print(current.data, "->")
		current = current.next
	}
	fmt.Println("nil")
}

// IterationReverse 链表的反转 迭代反转
func (link *LinkedList) IterationReverse() {
	if link.head == nil || link.head.next == nil {
		return
	}

	var beg *ListNode
	mid := link.head
	end := link.head.next
	// 一直遍历
	for true {
		// 只需修改mid的指针的指向
		mid.next = beg
		if end == nil {
			break
		}

		beg = mid
		mid = end
		end = end.next
	}

	link.head = mid
}

// 两个链表相加
func addTwoNumbers(l1 *LinkedList, l2 *LinkedList) *LinkedList {
	l1.Print()
	l2.Print()
	var beg *ListNode
	mid := l1.head
	end := l1.head.next
	for true {
		mid.next = beg
		if end == nil {
			break
		}
		beg = mid
		mid = end
		end = end.next
	}
	l1.head = mid
	l1.Print()

	var beg2 *ListNode
	mid2 := l2.head
	end2 := l2.head.next
	for true {
		mid2.next = beg2
		if end2 == nil {
			break
		}
		beg2 = mid2
		mid2 = end2
		end2 = end2.next
	}
	l2.head = mid2
	l2.Print()

	// 将数字取出来
	current := l1.head
	i := 0
	total1 := ""
	for current != nil {
		if i == 0 && current.data == 0 {
			current = current.next
			i++
			continue
		}
		total1 = total1 + strconv.Itoa(current.data)
		current = current.next
		i++
	}
	fmt.Println("total1")
	fmt.Println(total1)

	current2 := l2.head
	i2 := 0
	total2 := ""
	for current2 != nil {
		if i2 == 0 && current2.data == 0 {
			current2 = current2.next
			i2++
			continue
		}
		total2 = total2 + strconv.Itoa(current2.data)
		current2 = current2.next
		i2++
	}
	fmt.Println("total2")
	fmt.Println(total2)

	total1Num, _ := strconv.Atoi(total1)
	total2Num, _ := strconv.Atoi(total2)
	total := total1Num + total2Num
	totalStr := strconv.Itoa(total)
	fmt.Println(totalStr)
	// 根据数字再生成链表
	link := &LinkedList{head: nil}
	//curData := link.head
	for i := 0; i < len(totalStr); i++ {
		numData, _ := strconv.Atoi(string(totalStr[i]))
		fmt.Println(numData)
		node := &ListNode{data: numData, next: nil}
		//if curData == nil {
		//	curData = node
		//	fmt.Println(curData)
		//	link.head = curData
		//} else {
		//	curData.next = node
		//	fmt.Println(curData)
		//}
		//curData = curData.next
		if link.head == nil {
			link.head = node
		} else {
			link.Insert(numData)
		}
	}
	fmt.Println(link)
	return link
}

func LinkTest() {
	link := CreateLinkedList()
	link.AddNodes(5)
	link.Print()
	link.InsertIndex(5, 1)
	link.Print()

	link.InsertIndex(5, 3)
	link.Print()

	finish := link.DeleteNode(1)
	if finish == 1 {
		fmt.Printf("删除元素%d\n", 1)
	}
	link.Print()
	selValue := link.SelectNode(4)
	fmt.Printf("查询到的元素的下标，%d\n", selValue)
	changeValue := link.ChangeNode(1, 2)
	link.Print()
	fmt.Printf("替换成功？%d\n", changeValue)
	link.DeleteIndex(2)
	link.Print()
	// 迭代法反转
	link.IterationReverse()
	link.Print()
}

func LinkIterationReverseTest() {
	l1 := CreateLinkedList()
	l1.AddNodes(4)

	l2 := CreateLinkedList()
	l2.AddNodes(5)

	l3 := addTwoNumbers(l1, l2)
	l3.Print()
}
