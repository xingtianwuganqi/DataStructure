package linear

import "fmt"

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

// ChangeNode 改节点
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
	fmt.Printf("查询到的元素，%d\n", selValue)
	changeValue := link.ChangeNode(1, 2)
	link.Print()
	fmt.Printf("替换成功？%d", changeValue)
}
